// Stream volume measurements in a loop (no image output).
//
//	go run . -ip 192.168.1.10
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
	ip := flag.String("ip", "192.168.1.10", "camera IP")
	validOnly := flag.Bool("valid", false, "print only single-pack readings with positive dimensions")
	stable := flag.Bool("stable", false, "with -valid, also require Stable")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	dev, err := gige.Open(ctx, *ip)
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
func okMeasurement(s gige.Sample, requireStable bool) bool {
	if s.PackCount != 1 || s.Length <= 0 || s.WidthMm <= 0 || s.HeightMm <= 0 {
		return false
	}
	return !requireStable || s.Stable
}
