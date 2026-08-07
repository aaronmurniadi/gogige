// Features exercises GenICam Has / ApplyControlPair / Execute via Camera.Connect.
//
//	go run . -has Width -has AcquisitionStart
//	go run . -set DeviceUserID=line1
//	go run . -ip 192.168.1.108 -exec AcquisitionStop
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	ipFlag := flag.String("ip", "", "camera IP (empty = first discovery hit)")
	var hasNames, sets, execs multiFlag
	flag.Var(&hasNames, "has", "feature name to probe with Has (repeatable)")
	flag.Var(&sets, "set", "Name=value via ApplyControlPair (repeatable)")
	flag.Var(&execs, "exec", "command feature to Execute (repeatable)")
	flag.Parse()

	if len(hasNames) == 0 && len(sets) == 0 && len(execs) == 0 {
		fmt.Fprintln(os.Stderr, "usage: pass at least one of -has / -set / -exec")
		flag.PrintDefaults()
		os.Exit(2)
	}

	ip, err := resolveIP(*ipFlag)
	if err != nil {
		log.Fatal(err)
	}

	cam, err := gige.Connect(ip)
	if err != nil {
		log.Fatal(err)
	}
	defer cam.Close()

	for _, name := range hasNames {
		fmt.Printf("Has(%s)=%v\n", name, cam.Has(name))
	}
	for _, pair := range sets {
		if err := gige.ApplyControlPair(cam, pair); err != nil {
			log.Fatalf("set %s: %v", pair, err)
		}
		fmt.Printf("set %s: ok\n", pair)
	}
	for _, name := range execs {
		if err := cam.ExecuteCommand(name); err != nil {
			log.Fatalf("exec %s: %v", name, err)
		}
		fmt.Printf("exec %s: ok\n", name)
	}
}

type multiFlag []string

func (m *multiFlag) String() string { return fmt.Sprint([]string(*m)) }
func (m *multiFlag) Set(v string) error {
	*m = append(*m, v)
	return nil
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
	fmt.Printf("discovered %s @ %s\n", devs[0].Model, devs[0].IP)
	return devs[0].IP, nil
}
