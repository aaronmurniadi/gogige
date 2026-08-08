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

	devs, err := gvcp.Discover(context.Background(), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	var dev gvcp.DiscoveredDevice
	if *ip != "" {
		dev = gvcp.DiscoveredDevice{IP: *ip}
		for _, d := range devs {
			if d.IP == *ip {
				dev = d
				break
			}
		}
	} else {
		if len(devs) == 0 {
			log.Fatal("no cameras found; pass -ip")
		}
		dev = devs[0]
		fmt.Printf("discovered %s (%s)\n", dev.Model, dev.IP)
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
		vendor, serial := fileToken(dev.Manufacturer), fileToken(dev.Serial)
		if vendor == "" {
			vendor = "unknown"
		}
		if serial == "" {
			serial = "unknown"
		}
		outName = vendor + "-" + serial + "-genicam.xml"
	}
	outPath := filepath.Join(*dir, outName)
	if err := os.WriteFile(outPath, xmlData, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %s (%d bytes)\n", outPath, len(xmlData))
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
