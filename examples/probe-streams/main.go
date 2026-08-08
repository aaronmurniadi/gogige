// Probe-streams reports what a camera can stream beyond the default JPEG color
// path: GenICam stream features (pixel formats, payload type, chunks,
// components) plus one grabbed frame's actual BSCF/raw components and their
// JPEG encodability.
//
//	go run .                  # discover first camera
//	go run . -ip 192.168.1.108
//	go run . -ip 192.168.1.108 -fmt Mono16   # set PixelFormat before grab
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aaronmurniadi/gogige"
)

// streamFeatures are probed for presence/kind/entries when available.
var streamFeatures = []string{
	"PixelFormat",
	"PayloadType",
	"StreamPayloadType",
	"ComponentSelector",
	"ComponentEnable",
	"ComponentIDValue",
	"ChunkSelector",
	"ChunkModeActive",
	"GevStreamChannelCount",
	"GevSCPSPacketSize",
	"GevSCPSBigEndian",
	"GevSCPD",
	"GevSCPDirection",
	"GevSCCExtendedChunkData",
	"GevSCCFGExtendedChunkData",
	"GevSCCUnconditionalStreaming",
}

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	fmtName := flag.String("fmt", "", "optional PixelFormat to set before grab (e.g. Mono16)")
	timeout := flag.Duration("timeout", 3*time.Second, "frame grab timeout")
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
		log.Printf("discovered %s @ %s", devs[0].Model, cameraIP)
	}

	cam, err := gogige.Connect(cameraIP)
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	nm := cam.NodeMap()
	fmt.Printf("camera %s\n", cameraIP)
	fmt.Println("-- GenICam stream features --")
	probeFeatures(nm)

	if *fmtName != "" {
		if !nm.Has("PixelFormat") {
			log.Fatalf("camera has no PixelFormat feature")
		}
		if err := cam.SetStringFeature("PixelFormat", *fmtName); err != nil {
			log.Fatalf("set PixelFormat=%s: %v", *fmtName, err)
		}
		fmt.Printf("set PixelFormat=%s\n", *fmtName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout+3*time.Second)
	defer cancel()
	s := gogige.NewFromCamera(cam)
	if err := s.Open(cameraIP); err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	gctx, gcancel := context.WithTimeout(ctx, *timeout)
	samples, err := s.GrabComponents(gctx)
	gcancel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-- grabbed frame --")
	reportComponents(samples)
}

func probeFeatures(nm *gogige.NodeMap) {
	for _, name := range streamFeatures {
		if !nm.Has(name) {
			fmt.Printf("  %-24s absent\n", name)
			continue
		}
		switch nm.Kind(name) {
		case "Enumeration":
			entries, err := nm.EnumEntries(name)
			if err != nil {
				fmt.Printf("  %-24s enum   (error: %v)\n", name, err)
				continue
			}
			fmt.Printf("  %-24s enum   entries=%v current=%s\n", name, entries, currentEnum(nm, name))
		default:
			if v, err := nm.ReadInteger(name); err == nil {
				fmt.Printf("  %-24s %-8s = %d\n", name, nm.Kind(name), v)
			} else {
				fmt.Printf("  %-24s %s\n", name, nm.Kind(name))
			}
		}
	}
}

func currentEnum(nm *gogige.NodeMap, name string) string {
	cur, err := nm.CurrentEnum(name)
	if err != nil {
		return fmt.Sprintf("(read error: %v)", err)
	}
	if cur == "" {
		return "(no matching entry)"
	}
	return cur
}

func reportComponents(samples []gogige.Sample) {
	if len(samples) == 0 {
		fmt.Println("  (no components)")
		return
	}
	if len(samples) == 1 && samples[0].Component == gogige.ComponentUnknown {
		s := samples[0]
		fmt.Printf("  non-BSCF raw payload\n")
		fmt.Printf("    %-6s %dx%d  fmt=0x%08x %-12s raw=%dB  %s\n",
			"-", s.Width, s.Height, s.PixelFormat, pixelFormatName(s.PixelFormat),
			len(s.RawColor), jpegProbe(s))
		return
	}
	for _, s := range samples {
		fmt.Printf("  %-6s %dx%d  fmt=0x%08x %-12s raw=%dB  %s\n",
			s.Component, s.Width, s.Height, s.PixelFormat, pixelFormatName(s.PixelFormat),
			len(s.RawColor), jpegProbe(s))
	}
}

// jpegProbe tries to JPEG-encode a sample and reports the result.
func jpegProbe(s gogige.Sample) string {
	if s.Width <= 0 || s.Height <= 0 || len(s.RawColor) == 0 {
		return "empty"
	}
	jpeg, err := gogige.EncodeJPEG(s.RawColor, s.Width, s.Height, s.PixelFormat, 60)
	if err != nil {
		return fmt.Sprintf("jpeg: %v", err)
	}
	return fmt.Sprintf("jpeg=%dB", len(jpeg))
}

func pixelFormatName(fmtID uint32) string {
	if name, ok := pixelFormatNames[fmtID]; ok {
		return name
	}
	return ""
}

var pixelFormatNames = map[uint32]string{
	0x01080001: "Mono8",
	0x01080009: "BayerRG8",
	0x01080008: "BayerGR8",
	0x0108000b: "BayerGB8",
	0x0108000a: "BayerBG8",
	0x01080046: "Coord3D_A8",
	0x01100007: "Mono16",
	0x0110000f: "BayerRG16",
	0x01280047: "Coord3D_A16",
	0x01400044: "Coord3D_A32f",
	0x02100032: "YUV422_8",
	0x02100010: "YUV411_8",
	0x02100034: "YUV444_8",
	0x02180014: "RGB8",
	0x02180015: "BGR8",
	0x02300018: "RGB10",
	0x02300019: "BGR10",
	0x0230001c: "RGB12",
	0x0230001d: "BGR12",
	0x0234002e: "RGB10V1Packed",
	0x02360022: "RGB12V1Packed",
	0x0280001e: "RGB8_Planar",
	0x010c003b: "Coord3D_C16",
}
