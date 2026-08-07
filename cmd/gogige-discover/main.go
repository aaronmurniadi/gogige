// gogige-discover broadcasts GigE Vision DISCOVERY_CMD and prints peers.
//
//	go run ./cmd/gogige-discover
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	timeout := flag.Duration("timeout", 2*time.Second, "how long to wait for DISCOVERY_ACK")
	flag.Parse()

	devs, err := gogige.Discover(context.Background(), *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "discover: %v\n", err)
		os.Exit(1)
	}
	if len(devs) == 0 {
		fmt.Fprintln(os.Stderr, "no devices found")
		os.Exit(1)
	}
	for _, d := range devs {
		label := d.Manufacturer
		if d.Model != "" {
			if label != "" {
				label += "-"
			}
			label += d.Model
		}
		if d.Serial != "" {
			if label != "" {
				label += "-"
			}
			label += d.Serial
		}
		if label == "" {
			label = d.IP
		}
		fmt.Printf("%s (%s)", label, d.IP)
		if d.MAC != "" {
			fmt.Printf(" mac=%s", d.MAC)
		}
		if d.UserName != "" {
			fmt.Printf(" user=%s", d.UserName)
		}
		fmt.Println()
	}
}
