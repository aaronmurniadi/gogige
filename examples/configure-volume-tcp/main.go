// Apply the Huaray volume-camera TCP transfer control preset.
//
//	go run .                  # discover first camera
//	go run . -ip 192.168.1.108 -tcp 3100
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	tcp := flag.Int64("tcp", 3100, "TCP result port")
	flag.Parse()

	cameraIP, err := resolveIP(*ip)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dev, err := gogige.Open(ctx, cameraIP)
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	if err := configureVolumeTCP(dev, *tcp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("configured %s tcp=%d\n", cameraIP, *tcp)
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

func configureVolumeTCP(dev gogige.Device, tcpPort int64) error {
	if tcpPort <= 0 {
		tcpPort = 3100
	}
	f := dev.Features()
	for _, pair := range []string{"CommEnable=false"} {
		if err := applyFeaturePair(f, pair); err != nil {
			return err
		}
	}
	time.Sleep(200 * time.Millisecond)
	controls := []string{
		"TransferEnable=true",
		"TransferWorkMode=TCPServer",
		"TCPPort=" + strconv.FormatInt(tcpPort, 10),
		"TriggerEnable=true",
		"TriggerCmdType=STATIC_CMD",
		"TriggerType=FreeRun",
		"TriggerSource=Software",
		"CommSelector=TCP",
		"CommEnable=true",
		"CommDataTailer=DataTailer_LF",
		"TransferCmdType=TCP_CMD_NONE",
		"FilterDuplicateDataEnable=false",
		"ImageStoreEnable=false",
		"ImageStoreFilter=AllVolumes",
		"ResultMode=NotOnEdgeVolumes",
	}
	for _, pair := range controls {
		if err := applyFeaturePair(f, pair); err != nil {
			return err
		}
	}
	return nil
}

func applyFeaturePair(f gogige.Features, pair string) error {
	eq := strings.IndexByte(pair, '=')
	if eq <= 0 || eq == len(pair)-1 {
		return fmt.Errorf("bad control pair %q", pair)
	}
	feature := pair[:eq]
	val := pair[eq+1:]
	switch strings.ToLower(val) {
	case "true":
		return f.SetBool(feature, true)
	case "false":
		return f.SetBool(feature, false)
	default:
		if i, perr := strconv.ParseInt(val, 10, 64); perr == nil && strconv.FormatInt(i, 10) == val {
			return f.SetInt(feature, i)
		}
		if fl, perr := strconv.ParseFloat(val, 64); perr == nil && strings.ContainsAny(val, ".eE") {
			return f.SetFloat(feature, fl)
		}
		return f.SetString(feature, val)
	}
}
