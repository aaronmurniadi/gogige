// Smoke exercises as much of the current gogige stack as practical on one camera:
// discovery, CCP, ImplementationEndianness, heartbeat, FirstURL/XML, high-level Open,
// stream MTU/SCPS (via logger), grab + BSCF Sample, pause/resume, optional Live OnSample.
//
//	go run .                              # discover first camera
//	go run . -ip 192.168.1.108 -frames 3 -out ./out
package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/aaronmurniadi/gogige"
	"github.com/aaronmurniadi/gogige/genapi"
	"github.com/aaronmurniadi/gogige/gvcp"
	"github.com/aaronmurniadi/gogige/gvsp"
)

func main() {
	ipFlag := flag.String("ip", "", "camera IP (empty = first discovery hit)")
	frames := flag.Int("frames", 3, "frames to grab after stream open")
	outDir := flag.String("out", "", "if set, write frame-0.jpg here")
	hold := flag.Duration("hold", 2*time.Second, "how long to hold CCP with heartbeat before streaming")
	liveFor := flag.Duration("live", 0, "if >0, also run Live+OnSample for this duration")
	flag.Parse()

	ip, err := resolveIP(*ipFlag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== camera", ip, "===")

	if err := phase1Control(ip, *hold); err != nil {
		log.Fatal(err)
	}
	if err := phase2Stream(ip, *frames, *outDir, *liveFor); err != nil {
		log.Fatal(err)
	}
	fmt.Println("SMOKE OK")
}

func phase1Control(ip string, hold time.Duration) error {
	fmt.Println("--- Phase 1: GVCP control ---")
	g, err := gvcp.DialGVCP(ip, 3*time.Second)
	if err != nil {
		return err
	}
	defer g.Close()

	if err := g.TakeControl(); err != nil {
		return fmt.Errorf("TakeControl: %w", err)
	}
	defer g.LeaveControl()

	order := g.DeviceByteOrder()
	fmt.Printf("DeviceByteOrder=%v (big=%v)\n", order, order == binary.BigEndian)

	overlap, err := g.ReadReg(gvcp.AbrmImplementationEndianness)
	if err != nil {
		return err
	}
	fmt.Printf("ReadReg 0x020C (GenCP ImplEndian / GigE FirstURL overlap)=0x%08x\n", overlap)

	for _, a := range []struct {
		name string
		addr uint32
	}{
		{"GEV Version", 0x0000},
		{"CCP", 0x0a00},
		{"HeartbeatTO", 0x0938},
		{"Stream0PacketSize", gvcp.Stream0PacketSize},
	} {
		v, err := g.ReadReg(a.addr)
		if err != nil {
			fmt.Printf("  %-18s @0x%04x ERR %v\n", a.name, a.addr, err)
			continue
		}
		fmt.Printf("  %-18s @0x%04x = 0x%08x\n", a.name, a.addr, v)
	}

	url, err := g.FirstURL()
	if err != nil {
		return err
	}
	fmt.Printf("FirstURL=%q\n", url)
	fmt.Printf("HeartbeatTimeout()=%s\n", g.HeartbeatTimeout())

	hb := g.StartHeartbeat()
	defer hb.Stop()
	fmt.Printf("StartHeartbeat: holding CCP for %s…\n", hold)
	time.Sleep(hold)
	if err := g.PulseHeartbeat(); err != nil {
		return fmt.Errorf("PulseHeartbeat: %w", err)
	}
	fmt.Println("PulseHeartbeat: ok")

	xmlData, err := genapi.FetchXML(g)
	if err != nil {
		return fmt.Errorf("FetchXML: %w", err)
	}
	nm, err := genapi.ParseNodeMap(xmlData, g)
	if err != nil {
		return fmt.Errorf("ParseNodeMap: %w", err)
	}
	fmt.Printf("FetchXML=%d bytes Has(Width)=%v Has(AcquisitionStart)=%v\n",
		len(xmlData), nm.Has("Width"), nm.Has("AcquisitionStart"))

	camIP := net.ParseIP(ip)
	mtu := gvsp.PathMTU(camIP)
	fmt.Printf("PathMTU(%s)=%d → PacketSizeForMTU=%d\n", ip, mtu, gvsp.PacketSizeForMTU(mtu))
	return nil
}

func phase2Stream(ip string, n int, outDir string, liveFor time.Duration) error {
	fmt.Println("--- Phase 2: Open + stream ---")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dev, err := gige.Open(ctx, ip, gige.WithLogger(printLogger{}), gige.WithTimeout(3*time.Second))
	if err != nil {
		return fmt.Errorf("Open: %w", err)
	}
	defer dev.Close()

	for _, name := range []string{"Width", "Height", "PixelFormat", "AcquisitionStart", "AcquisitionStop"} {
		fmt.Printf("  Features.Has(%s)=%v\n", name, dev.Features().Has(name))
	}

	grabber, err := dev.StartGrabber(ctx)
	if err != nil {
		return fmt.Errorf("StartGrabber: %w", err)
	}
	defer grabber.Close()

	var firstJPEG []byte
	for i := 0; i < n; i++ {
		gctx, gcancel := context.WithTimeout(ctx, 5*time.Second)
		sample, err := grabber.Grab(gctx)
		gcancel()
		if err != nil {
			return fmt.Errorf("Grab[%d]: %w", i, err)
		}
		fmt.Printf("frame[%d] jpeg=%dB %dx%d packs=%d L=%.2f W=%.2f H=%.2f stable=%v\n",
			i, len(sample.JPEG), sample.Width, sample.Height, sample.PackCount,
			sample.Length, sample.WidthMm, sample.HeightMm, sample.Stable)
		if i == 0 {
			firstJPEG = append([]byte(nil), sample.JPEG...)
		}
	}

	fmt.Println("Pause 1s…")
	if err := grabber.Pause(ctx); err != nil {
		return fmt.Errorf("Pause: %w", err)
	}
	time.Sleep(time.Second)
	if err := grabber.Resume(ctx); err != nil {
		return fmt.Errorf("Resume: %w", err)
	}
	gctx, gcancel := context.WithTimeout(ctx, 5*time.Second)
	sample, err := grabber.Grab(gctx)
	gcancel()
	if err != nil {
		return fmt.Errorf("Grab after resume: %w", err)
	}
	fmt.Printf("after resume: jpeg=%dB packs=%d\n", len(sample.JPEG), sample.PackCount)

	if outDir != "" && len(firstJPEG) > 0 {
		if err := os.MkdirAll(outDir, 0o755); err != nil {
			return err
		}
		path := filepath.Join(outDir, "frame-0.jpg")
		if err := os.WriteFile(path, firstJPEG, 0o644); err != nil {
			return err
		}
		fmt.Println("wrote", path)
	}

	_ = grabber.Close()

	if liveFor > 0 {
		fmt.Printf("Live OnSample for %s…\n", liveFor)
		var count int
		live := gige.NewLive(dev, gige.WithOnSample(func(s gige.Sample) {
			count++
			if count <= 3 || count%10 == 0 {
				fmt.Printf("  live[%d] packs=%d L=%.2f\n", count, s.PackCount, s.Length)
			}
		}))
		lctx, lcancel := context.WithTimeout(context.Background(), liveFor)
		live.Start(lctx)
		<-lctx.Done()
		live.Stop()
		lcancel()
		fmt.Printf("Live samples=%d latestJPEG=%dB\n", count, len(live.LatestJPEG()))
	}
	return nil
}

type printLogger struct{}

func (printLogger) Debug(msg string, kv ...any) { logLine("DEBUG", msg, kv...) }
func (printLogger) Info(msg string, kv ...any)  { logLine("INFO", msg, kv...) }
func (printLogger) Warn(msg string, kv ...any)  { logLine("WARN", msg, kv...) }
func (printLogger) Error(msg string, kv ...any) { logLine("ERROR", msg, kv...) }

func logLine(level, msg string, kv ...any) {
	fmt.Printf("[%s] %s", level, msg)
	for i := 0; i+1 < len(kv); i += 2 {
		fmt.Printf(" %v=%v", kv[i], kv[i+1])
	}
	fmt.Println()
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
	d := devs[0]
	fmt.Printf("discovered %s %s sn=%s @ %s\n", d.Manufacturer, d.Model, d.Serial, d.IP)
	return d.IP, nil
}
