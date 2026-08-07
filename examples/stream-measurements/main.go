// Stream volume measurements in a loop (no image output).
//
//	go run .                  # discover first camera
//	go run . -ip 192.168.1.108
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	validOnly := flag.Bool("valid", false, "print only single-pack readings with positive dimensions")
	stable := flag.Bool("stable", false, "with -valid, also require Stable")
	flag.Parse()

	cameraIP, err := resolveIP(*ip)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	dev, err := gogige.Open(ctx, cameraIP)
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	g, err := dev.StartGrabber(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	fmt.Println("packs\tlength_mm\twidth_mm\theight_mm\tstable")
	for {
		grabCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		sample, err := g.Grab(grabCtx)
		cancel()
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("grab: %v", err)
			continue
		}
		if *validOnly && !okMeasurement(sample, *stable) {
			continue
		}
		fmt.Printf("%d\t%.3f\t%.3f\t%.3f\t%v\n",
			sample.PackCount, sample.Length, sample.WidthMm, sample.HeightMm, sample.Stable)
	}
}

// Caller-owned filter — the library does not validate measurements.
func okMeasurement(s gogige.Sample, requireStable bool) bool {
	if s.PackCount != 1 || s.Length <= 0 || s.WidthMm <= 0 || s.HeightMm <= 0 {
		return false
	}
	return !requireStable || s.Stable
}

func resolveIP(ip string) (string, error) {
	if ip != "" {
		return ip, nil
	}
	devs, err := gogige.Discover(context.Background(), 2*time.Second)
	if err != nil {
		return "", err
	}
	if len(devs) == 0 {
		return "", fmt.Errorf("no cameras found; pass -ip")
	}
	fmt.Printf("discovered %s\n", devs[0].IP)
	return devs[0].IP, nil
}
