// Discover the camera GenICam URL, fetch the feature XML, and write it to -dir.
//
//	go run . -ip 192.168.1.10 -dir ./out
package main

import (
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
	ip := flag.String("ip", "192.168.1.10", "camera IP")
	dir := flag.String("dir", ".", "directory to write the XML into")
	name := flag.String("name", "", "output filename (default: <ip>-genapi.xml)")
	timeout := flag.Duration("timeout", 5*time.Second, "GVCP timeout")
	flag.Parse()

	if err := os.MkdirAll(*dir, 0o755); err != nil {
		log.Fatal(err)
	}

	g, err := gvcp.DialGVCP(*ip, *timeout)
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
		safeIP := strings.ReplaceAll(*ip, ".", "-")
		outName = safeIP + "-genapi.xml"
	}
	outPath := filepath.Join(*dir, outName)
	if err := os.WriteFile(outPath, xmlData, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %s (%d bytes)\n", outPath, len(xmlData))
}
