// Grab-components discovers BSCF/SFNC components in one frame and writes a JPEG for each.
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
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	dir := flag.String("dir", ".", "directory for <component>.jpg")
	timeout := flag.Duration("timeout", 5*time.Second, "grab timeout")
	flag.Parse()

	deviceIP := *ip
	if deviceIP == "" {
		devs, err := gogige.Discover(context.Background(), 2*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		if len(devs) == 0 {
			log.Fatal("no cameras found; pass -ip")
		}
		deviceIP = devs[0].IP
		fmt.Printf("discovered %s %s @ %s\n", devs[0].Manufacturer, devs[0].Model, deviceIP)
	}

	if err := os.MkdirAll(*dir, 0o755); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout+3*time.Second)
	defer cancel()

	dev, err := gogige.Open(ctx, deviceIP, gogige.WithTimeout(*timeout))
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	g, err := dev.StartGrabber(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	gctx, gcancel := context.WithTimeout(ctx, *timeout)
	samples, err := g.GrabAll(gctx)
	gcancel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("available components (%d):\n", len(samples))
	for _, s := range samples {
		fmt.Printf("  %-8s  %dx%d  pixfmt=0x%08x  jpeg=%dB\n",
			s.Component, s.Width, s.Height, s.PixelFormat, len(s.JPEG))
	}

	for _, s := range samples {
		compName := s.Component.String()
		if compName == "" || compName == "unknown" {
			compName = "component"
		}
		path := filepath.Join(*dir, compName+".jpg")
		if err := os.WriteFile(path, s.JPEG, 0o644); err != nil {
			log.Fatal(err)
		}
		fmt.Println("wrote", path)
	}
}
