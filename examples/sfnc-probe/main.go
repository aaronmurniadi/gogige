// Sfnc-probe opens a GigE Vision camera and reports its well-known SFNC
// identity and streaming features using the typed Camera getters and the
// GVCP MaximumDeviceResponseTime / DeviceLinkHeartbeatTimeout wiring.
//
//	go run .                  # discover first camera
//	go run . -ip 192.168.1.108
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	timeout := flag.Duration("timeout", 3*time.Second, "GVCP timeout")
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

	cam, err := gogige.OpenDevice(context.Background(), cameraIP, gogige.WithTimeout(*timeout))
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	g := cam.GVCP()

	// --- Device identity (typed getters) ---
	fmt.Println("-- device identity --")
	printStr(cam.DeviceVendorName, "DeviceVendorName")
	printStr(cam.DeviceModelName, "DeviceModelName")
	printStr(cam.DeviceSerialNumber, "DeviceSerialNumber")
	printStr(cam.DeviceUserID, "DeviceUserID")

	// --- Device info via generic API ---
	fmt.Println("\n-- device info --")
	printInt(cam, gogige.FeatureDeviceSFNCVersionMajor)
	printStrGeneric(cam, gogige.FeatureDeviceTLType)

	// --- Streaming features (typed getters) ---
	fmt.Println("\n-- streaming features --")
	printInt(cam, gogige.FeatureGevSCPSPacketSize)
	printInt(cam, gogige.FeatureGevSCPD)
	printInt(cam, gogige.FeatureDeviceStreamChannelCount)
	printInt(cam, gogige.FeatureGevCCP)
	printInt(cam, gogige.FeatureGevGVSPExtendedIDMode)

	// --- MDRT ---
	fmt.Println("\n-- GVCP timing --")
	if mdrt, err := g.MaximumDeviceResponseTime(); err != nil {
		fmt.Printf("  MaximumDeviceResponseTime  error (%v)\n", err)
	} else {
		fmt.Printf("  MaximumDeviceResponseTime  %d ms\n", mdrt.Milliseconds())
	}

	// --- Heartbeat timeout ---
	printInt(cam, gogige.FeatureGevHeartbeatTimeout)
	if hb, err := cam.Integer(gogige.FeatureDeviceLinkHeartbeatTimeout); err == nil {
		fmt.Printf("  DeviceLinkHeartbeatTimeout %d ms (SFNC 2.7)\n", hb)
	} else {
		fmt.Println("  DeviceLinkHeartbeatTimeout absent (will use GigE SBRM register)")
	}
}

func printStr(fn func() (string, error), label string) {
	v, err := fn()
	if err != nil {
		fmt.Printf("  %-28s error (%v)\n", label, err)
	} else {
		fmt.Printf("  %-28s %s\n", label, v)
	}
}

func printInt(cam *gogige.Camera, name string) {
	if !cam.Has(name) {
		fmt.Printf("  %-28s absent\n", name)
		return
	}
	v, err := cam.Integer(name)
	if err != nil {
		fmt.Printf("  %-28s error (%v)\n", name, err)
	} else {
		fmt.Printf("  %-28s %d\n", name, v)
	}
}

func printStrGeneric(cam *gogige.Camera, name string) {
	if !cam.Has(name) {
		fmt.Printf("  %-28s absent\n", name)
		return
	}
	v, err := cam.String(name)
	if err != nil {
		fmt.Printf("  %-28s error (%v)\n", name, err)
	} else {
		fmt.Printf("  %-28s %s\n", name, v)
	}
}
