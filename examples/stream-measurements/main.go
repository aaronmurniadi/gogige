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
	flag.Parse()

	cameraIP := *ip
	if cameraIP == "" {
		devs, err := gogige.Discover(context.Background(), 2*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		if len(devs) == 0 {
			log.Fatal("no cameras found; pass -ip")
		}
		cameraIP = devs[0].IP
		fmt.Printf("discovered %s\n", cameraIP)
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

	for {
		grabCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		samples, err := g.GrabAll(grabCtx)
		cancel()
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("grab: %v", err)
			continue
		}
		for _, sample := range samples {
			fmt.Printf("Sample.PackCount=%d\n", sample.PackCount)
			for j, pack := range sample.Packs {
				fmt.Printf("  Pack[%d]: LengthMm=%f, WidthMm=%f, HeightMm=%f, center.x=%f, center.y=%f, center.z=%f\n", j, pack.LengthMm, pack.WidthMm, pack.HeightMm, pack.CenterX, pack.CenterY, pack.CenterZ)
			}
		}
	}
}
