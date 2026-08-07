// Grab one color frame and write it as JPEG into -dir.
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
	dir := flag.String("dir", ".", "directory to write the JPEG into")
	name := flag.String("name", "", "output filename (default: frame-<unix>.jpg)")
	flag.Parse()

	cameraIP, err := resolveIP(*ip)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(*dir, 0o755); err != nil {
		log.Fatal(err)
	}
	outName := *name
	if outName == "" {
		outName = fmt.Sprintf("frame-%d.jpg", time.Now().Unix())
	}
	outPath := filepath.Join(*dir, outName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jpeg, err := gige.GrabJPEG(ctx, cameraIP)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(outPath, jpeg, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(outPath)
}

func resolveIP(ip string) (string, error) {
	if ip != "" {
		return ip, nil
	}
	devs, err := gige.Discover(context.Background(), 2*time.Second)
	if err != nil {
		return "", err
	}
	if len(devs) == 0 {
		return "", fmt.Errorf("no cameras found; pass -ip")
	}
	fmt.Printf("discovered %s (%s)\n", deviceLabel(devs[0]), devs[0].IP)
	return devs[0].IP, nil
}

func deviceLabel(d gige.DeviceInfo) string {
	s := d.Manufacturer
	if d.Model != "" {
		if s != "" {
			s += "-"
		}
		s += d.Model
	}
	if d.Serial != "" {
		if s != "" {
			s += "-"
		}
		s += d.Serial
	}
	if s == "" {
		return d.IP
	}
	return s
}
