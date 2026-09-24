// Draw detected-object bounding boxes on a color frame, Aerial-View style.
//
// The camera reports pack detections as 3D world boxes (BSCF PackDet). With a
// color calibration, gogige projects those 3D boxes onto the image plane and
// attaches pixel-space Box2D rectangles to Sample.Overlay; this example streams
// frames until one contains at least one detected object, then draws green box
// outlines around each pack on the color JPEG.
//
//	go run . -ip 192.168.1.108 -calib ./calib.json -out boxes.png
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/aaronmurniadi/gogige"
	"github.com/aaronmurniadi/gogige/calib"
)

func main() {
	ip := flag.String("ip", "192.168.1.108", "camera IP")
	calibPath := flag.String("calib", "calib.json", "path to vendor camera calibration JSON")
	out := flag.String("out", "boxes.png", "output PNG path")
	maxTries := flag.Int("tries", 200, "max frames to wait for a detection")
	flag.Parse()

	cam, err := gogige.Open(context.Background(), *ip, gogige.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	vendor, err := calib.LoadVendorFile(*calibPath)
	if err != nil {
		log.Fatalf("calib: %v", err)
	}
	colorCalib, err := vendor.Color()
	if err != nil {
		log.Fatalf("calib color: %v", err)
	}
	fmt.Printf("color calib fx=%.2f fy=%.2f cx=%.2f cy=%.2f\n",
		colorCalib.K[calib.FX], colorCalib.K[calib.FY], colorCalib.K[calib.CX], colorCalib.K[calib.CY])

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, err := cam.StartGrabber(ctx, gogige.WithOverlay(true, colorCalib))
	if err != nil {
		log.Fatalf("start grabber: %v", err)
	}
	defer g.Close()

	for i := 0; i < *maxTries; i++ {
		samples, err := g.GrabAll(ctx)
		if err != nil {
			log.Fatalf("grab: %v", err)
		}
		// prefer the color component's frame; fall back to any sample with packs
		for _, s := range samples {
			if s.Component != gogige.ComponentColor && len(samples) > 1 {
				continue
			}
			if len(s.Packs) > 0 {
				drawBoxes(s, *out)
				return
			}
		}
		// color component had no packs yet; try again
	}
	log.Fatalf("no detections after %d frames", *maxTries)
}

func drawBoxes(s gogige.Sample, out string) {
	img, err := jpeg.Decode(bytes.NewReader(s.JPEG))
	if err != nil {
		log.Fatalf("jpeg decode: %v", err)
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
	green := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	packs := 0
	for i, b := range s.Overlay {
		if b.Empty() {
			continue
		}
		drawBox(rgba, green, b)
		fmt.Printf("pack %d box L=%.0f T=%.0f R=%.0f B=%.0f\n", i, b.MinX, b.MinY, b.MaxX, b.MaxY)
		packs++
	}
	if packs == 0 {
		log.Fatal("sample had packs but no overlay boxes; check calibration")
	}
	f, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, rgba); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote", out, "with", packs, "boxes")
}

// drawBox draws a green rectangle outline.
func drawBox(img *image.RGBA, c color.RGBA, b gogige.Box2D) {
	const t = 1
	x0, y0 := int(b.MinX), int(b.MinY)
	x1, y1 := int(b.MaxX), int(b.MaxY)
	rects := []image.Rectangle{
		image.Rect(x0, y0, x0+t, y1), // left
		image.Rect(x1-t, y0, x1, y1), // right
		image.Rect(x0, y0, x1, y0+t), // top
		image.Rect(x0, y1-t, x1, y1), // bottom
	}
	for _, r := range rects {
		draw.Draw(img, r, &image.Uniform{c}, image.Point{}, draw.Src)
	}
}
