// Read the Huaray volume-camera ImageStoreEnable, toggle it, then restore it.
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
		deviceName := devs[0].Manufacturer + " " + devs[0].Model + " " + devs[0].Serial + " " + devs[0].IP
		fmt.Printf("%s\n", deviceName)
	}

	cam, err := gogige.Connect(deviceIP)
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	const feature = "ImageStoreEnable"

	cur, err := cam.BooleanFeature(feature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s current=%v\n", feature, cur)

	if err := gogige.ApplyControlPair(cam, fmt.Sprintf("%s=%v", feature, !cur)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s toggled to %v\n", feature, !cur)

	if err := gogige.ApplyControlPair(cam, fmt.Sprintf("%s=%v", feature, cur)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s restored to %v\n", feature, cur)
}
