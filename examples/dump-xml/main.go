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

	"github.com/aaronmurniadi/gogige/genapi"
	"github.com/aaronmurniadi/gogige/gvcp"
)

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	dir := flag.String("dir", ".", "directory to write the XML into")
	name := flag.String("name", "", "output filename (default: <ip>-genapi.xml)")
	timeout := flag.Duration("timeout", 5*time.Second, "GVCP timeout")
	flag.Parse()

	cameraIP, err := resolveIP(*ip)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(*dir, 0o755); err != nil {
		log.Fatal(err)
	}

	g, err := gvcp.DialGVCP(cameraIP, *timeout)
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
		safeIP := strings.ReplaceAll(cameraIP, ".", "-")
		outName = safeIP + "-genapi.xml"
	}
	outPath := filepath.Join(*dir, outName)
	if err := os.WriteFile(outPath, xmlData, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %s (%d bytes)\n", outPath, len(xmlData))
}

func resolveIP(ip string) (string, error) {
	if ip != "" {
		return ip, nil
	}
	devs, err := gvcp.Discover(context.Background(), 2*time.Second)
	if err != nil {
		return "", err
	}
	if len(devs) == 0 {
		return "", fmt.Errorf("no cameras found; pass -ip")
	}
	fmt.Printf("discovered %s (%s)\n", devs[0].Model, devs[0].IP)
	return devs[0].IP, nil
}
