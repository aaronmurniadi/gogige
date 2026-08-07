// gogige-stream captures N frames (JPEG + BSCF measurements) from a GigE camera.
//
//	go run ./cmd/gogige-stream -n 5 -dir ./out
//	go run ./cmd/gogige-stream -ip 192.168.1.108 -n 1
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	ipFlag := flag.String("ip", "", "camera IP (empty = first discovery hit)")
	n := flag.Int("n", 3, "number of frames to capture")
	dir := flag.String("dir", ".", "directory for frame-*.jpg")
	timeout := flag.Duration("timeout", 5*time.Second, "per-frame grab timeout")
	flag.Parse()

	ip, err := resolveIP(*ipFlag)
	if err != nil {
		fatal(err)
	}
	if err := os.MkdirAll(*dir, 0o755); err != nil {
		fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*n+2)*(*timeout))
	defer cancel()

	dev, err := gige.Open(ctx, ip, gige.WithLogger(stdioLog{}), gige.WithTimeout(*timeout))
	if err != nil {
		fatal(err)
	}
	defer dev.Close()

	g, err := dev.StartGrabber(ctx)
	if err != nil {
		fatal(err)
	}
	defer g.Close()

	for i := 0; i < *n; i++ {
		gctx, gcancel := context.WithTimeout(ctx, *timeout)
		sample, err := g.Grab(gctx)
		gcancel()
		if err != nil {
			fatal(fmt.Errorf("frame %d: %w", i, err))
		}
		path := filepath.Join(*dir, fmt.Sprintf("frame-%04d.jpg", i))
		if err := os.WriteFile(path, sample.JPEG, 0o644); err != nil {
			fatal(err)
		}
		fmt.Printf("%s  %dx%d  packs=%d  L=%.3f W=%.3f H=%.3f stable=%v\n",
			path, sample.Width, sample.Height, sample.PackCount,
			sample.Length, sample.WidthMm, sample.HeightMm, sample.Stable)
	}
}

type stdioLog struct{}

func (stdioLog) Debug(msg string, kv ...any) { logKV("debug", msg, kv...) }
func (stdioLog) Info(msg string, kv ...any)  { logKV("info", msg, kv...) }
func (stdioLog) Warn(msg string, kv ...any)  { logKV("warn", msg, kv...) }
func (stdioLog) Error(msg string, kv ...any) { logKV("error", msg, kv...) }

func logKV(level, msg string, kv ...any) {
	fmt.Fprintf(os.Stderr, "%s: %s", level, msg)
	for i := 0; i+1 < len(kv); i += 2 {
		fmt.Fprintf(os.Stderr, " %v=%v", kv[i], kv[i+1])
	}
	fmt.Fprintln(os.Stderr)
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
	fmt.Fprintf(os.Stderr, "discovered %s %s @ %s\n", devs[0].Manufacturer, devs[0].Model, devs[0].IP)
	return devs[0].IP, nil
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "gogige-stream: %v\n", err)
	os.Exit(1)
}
