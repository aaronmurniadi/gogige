// Stream frames through the Phase 4 channel API.
//
//	go run .                  # discover first camera
//	go run . -ip 192.168.1.108
//	go run . -ip 192.168.1.108 -frames 10
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
	frames := flag.Int("frames", 0, "stop after this many frames (0 = until interrupted)")
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

	cam, err := gogige.OpenDevice(ctx, cameraIP, gogige.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	stream, err := cam.StartStream(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	n := 0
	for frame := range stream.Frames() {
		fmt.Printf("frame %d  id=%d  %dx%d  fmt=0x%08x  %dB\n",
			n+1, frame.ID, frame.Width, frame.Height, frame.PixelFormat, len(frame.Data))
		frame.Release()
		n++
		if *frames > 0 && n >= *frames {
			stream.Stop()
		}
		if ctx.Err() != nil {
			break
		}
	}
	fmt.Printf("received %d frames\n", n)
}
