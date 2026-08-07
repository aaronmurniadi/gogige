// Grab one color frame and write it as JPEG into -dir.
//
//	go run . -ip 192.168.1.10 -dir ./out
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aaronmurniadi/gogige"
)

func main() {
	ip := flag.String("ip", "192.168.1.10", "camera IP")
	dir := flag.String("dir", ".", "directory to write the JPEG into")
	name := flag.String("name", "", "output filename (default: frame-<unix>.jpg)")
	flag.Parse()

	if err := os.MkdirAll(*dir, 0o755); err != nil {
		log.Fatal(err)
	}
	outName := *name
	if outName == "" {
		outName = fmt.Sprintf("frame-%d.jpg", time.Now().Unix())
	}
	outPath := filepath.Join(*dir, outName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jpeg, err := gige.GrabJPEG(ctx, *ip)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(outPath, jpeg, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(outPath)
}
