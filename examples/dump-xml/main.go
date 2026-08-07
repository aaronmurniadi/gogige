// Discover the camera GenICam URL, fetch the feature XML, and write it to -dir.
//
//	go run .                  # discover first camera
//	go run . -ip 192.168.1.108 -dir ./out
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/aaronmurniadi/gogige/genapi"
	"github.com/aaronmurniadi/gogige/gvcp"
)

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	dir := flag.String("dir", ".", "directory to write the XML into")
	name := flag.String("name", "", "output filename (default: <vendor>-<serial>-genicam.xml)")
	timeout := flag.Duration("timeout", 5*time.Second, "GVCP timeout")
	flag.Parse()

	dev, err := resolveDevice(*ip)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(*dir, 0o755); err != nil {
		log.Fatal(err)
	}

	g, err := gvcp.DialGVCP(dev.IP, *timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	if err := g.TakeControl(); err != nil {
		log.Fatal(err)
	}
	defer func() { _ = g.LeaveControl() }()

	url, err := g.FirstURL()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GenICam URL:", url)

	xmlData, err := genapi.FetchXML(g)
	if err != nil {
		log.Fatal(err)
	}

	outName := *name
	if outName == "" {
		outName = defaultXMLName(dev)
	}
	outPath := filepath.Join(*dir, outName)
	if err := os.WriteFile(outPath, xmlData, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %s (%d bytes)\n", outPath, len(xmlData))
}

func defaultXMLName(d gvcp.DiscoveredDevice) string {
	vendor := fileToken(d.Manufacturer)
	serial := fileToken(d.Serial)
	if vendor == "" {
		vendor = "unknown"
	}
	if serial == "" {
		serial = "unknown"
	}
	return vendor + "-" + serial + "-genicam.xml"
}

// fileToken keeps letters/digits; other runs become a single '-'.
func fileToken(s string) string {
	var b strings.Builder
	dash := false
	for _, r := range strings.TrimSpace(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			dash = false
			continue
		}
		if !dash && b.Len() > 0 {
			b.WriteByte('-')
			dash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

func resolveDevice(ip string) (gvcp.DiscoveredDevice, error) {
	devs, err := gvcp.Discover(context.Background(), 2*time.Second)
	if err != nil {
		return gvcp.DiscoveredDevice{}, err
	}
	if ip != "" {
		for _, d := range devs {
			if d.IP == ip {
				return d, nil
			}
		}
		return gvcp.DiscoveredDevice{IP: ip}, nil
	}
	if len(devs) == 0 {
		return gvcp.DiscoveredDevice{}, fmt.Errorf("no cameras found; pass -ip")
	}
	fmt.Printf("discovered %s (%s)\n", devs[0].Model, devs[0].IP)
	return devs[0], nil
}
