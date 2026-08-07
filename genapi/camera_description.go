package genapi

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/aaronmurniadi/gogige/gvcp"
)

// URLReader can read the GenICam XML URL from bootstrap and device memory.
type URLReader interface {
	gvcp.Port
	FirstURL() (string, error)
}

// FetchXML loads GenICam XML described by FirstURL (Local: or http://).
func FetchXML(g URLReader) ([]byte, error) {
	url, err := g.FirstURL()
	if err != nil {
		return nil, err
	}
	return fetchDeviceXML(g, url)
}

func fetchDeviceXML(port gvcp.Port, url string) ([]byte, error) {
	if url == "" {
		return nil, fmt.Errorf("gige: empty GenICam URL")
	}
	// local:filename.zip;addr;length — scheme case-insensitive; addr/len are hex (AIA / Aravis).
	lower := strings.ToLower(url)
	if strings.HasPrefix(lower, "local:") {
		rest := url[len("local:"):]
		name, addr, length, err := parseLocalGenICamURL(rest)
		if err != nil {
			return nil, fmt.Errorf("gige: bad Local URL %q: %w", url, err)
		}
		raw, err := port.ReadMem(uint32(addr), int(length))
		if err != nil {
			return nil, err
		}
		return maybeUnzipXML(raw, name)
	}
	if strings.HasPrefix(lower, "http://") || strings.HasPrefix(lower, "https://") {
		resp, err := http.Get(url) //nolint:gosec // camera-local GenICam URL
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("gige: fetch XML HTTP %d", resp.StatusCode)
		}
		return io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	}
	return nil, fmt.Errorf("gige: unsupported GenICam URL scheme %q", url)
}

// parseLocalGenICamURL parses "path;hexAddr;hexLength" (optional 0x prefix).
func parseLocalGenICamURL(rest string) (name string, addr, length uint64, err error) {
	parts := strings.Split(rest, ";")
	if len(parts) < 3 {
		return "", 0, 0, fmt.Errorf("need path;addr;length")
	}
	name = strings.TrimSpace(parts[0])
	name = strings.TrimPrefix(name, "///")
	name = strings.TrimPrefix(name, "//")
	addr, err = parseHexU64(parts[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("addr %q: %w", parts[1], err)
	}
	length, err = parseHexU64(parts[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("length %q: %w", parts[2], err)
	}
	if length == 0 {
		return "", 0, 0, fmt.Errorf("zero length")
	}
	return name, addr, length, nil
}

func parseHexU64(s string) (uint64, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(strings.ToLower(s), "0x")
	return strconv.ParseUint(s, 16, 64)
}

func unzipFirstXML(raw []byte) ([]byte, error) {
	zr, err := zip.NewReader(bytes.NewReader(raw), int64(len(raw)))
	if err != nil {
		return nil, fmt.Errorf("gige: unzip GenICam: %w", err)
	}
	var fallback *zip.File
	for _, f := range zr.File {
		name := strings.ToLower(f.Name)
		if strings.HasSuffix(name, ".xml") {
			return readZipFile(f)
		}
		if fallback == nil && !f.FileInfo().IsDir() {
			fallback = f
		}
	}
	if fallback != nil {
		return readZipFile(fallback)
	}
	return nil, errors.New("gige: zip has no XML")
}

func readZipFile(f *zip.File) ([]byte, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(io.LimitReader(rc, 8<<20))
}

func maybeUnzipXML(raw []byte, name string) ([]byte, error) {
	lower := strings.ToLower(name)
	if strings.HasSuffix(lower, ".zip") || (len(raw) >= 2 && raw[0] == 'P' && raw[1] == 'K') {
		return unzipFirstXML(raw)
	}
	return raw, nil
}
