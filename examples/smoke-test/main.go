// Smoke-test exercises the whole current gogige stack against one camera:
// GigE discovery, raw GVCP control channel (CCP + heartbeat), GenICam XML and
// GenApi node-map introspection, high-level Camera feature get/set, MTU/SCPS
// negotiation, streaming (Session grab, GrabAll, component switch, pause/
// resume), the Phase-4 channel API (StartStream/Frames/Release), and optional
// Live preview.
//
// All phases run sequentially on real hardware; every phase reports what it
// did. Writes are non-destructive: integer/enum nodes are re-written with the
// currently-read value (a no-op that still round-trips through the device map).
//
//	go run .                               # discover first camera
//	go run . -ip 192.168.1.108 -out ./out  # keep JPEGs + full checks
//	go run . -live 2s -channel             # also run Live 2s + the channel API
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
	"github.com/aaronmurniadi/gogige/live"
)

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first discovery hit)")
	frames := flag.Int("frames", 3, "frames to grab after stream open")
	outDir := flag.String("out", "", "if set, write <component>.jpg captures here")
	hold := flag.Duration("hold", 2*time.Second, "how long to hold CCP with heartbeat before streaming")
	liveFor := flag.Duration("live", 0, "if >0, also run Live+OnSample for this duration")
	channel := flag.Bool("channel", false, "exercise the Phase-4 StartStream/Frames() API")
	component := flag.String("component", "color", "BSCF component for Grab (color|depth|mono)")
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
		d := devs[0]
		deviceIP = d.IP
		fmt.Printf("discovered %s %s sn=%s mac=%s user=%q @ %s\n",
			d.Manufacturer, d.Model, d.Serial, d.MAC, d.UserName, d.IP)
		if len(devs) > 1 {
			fmt.Printf("  (%d total; using first)\n", len(devs))
		}
	}
	comp, err := gogige.ParseComponent(*component)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=== camera %s (component %s) ===\n", deviceIP, comp)

	if err := phase1Control(deviceIP, *hold); err != nil {
		log.Fatal(err)
	}
	if err := phase2GenApi(deviceIP); err != nil {
		log.Fatal(err)
	}
	if err := phase3Stream(deviceIP, *frames, *outDir, comp); err != nil {
		log.Fatal(err)
	}
	if *channel {
		if err := phase4Channel(deviceIP); err != nil {
			log.Fatal(err)
		}
	}
	if *liveFor > 0 {
		if err := phase5Live(deviceIP, *liveFor); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("SMOKE OK")
}

// phase1Control exercises the raw GVCP control channel: CCP access privilege,
// byte order, register map reads, heartbeat maintenance, FirstURL and the
// GenApi XML node-map build.
func phase1Control(ip string, hold time.Duration) error {
	fmt.Println("--- Phase 1: GVCP control + GenApi XML ---")
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
		return fmt.Errorf("ReadReg 0x020C: %w", err)
	}
	fmt.Printf("ReadReg 0x020C (GenCP ImplEndian / GigE FirstURL overlap)=0x%08x\n", overlap)

	for _, a := range []struct {
		name string
		addr uint32
	}{
		{"GEV Version", 0x0000},
		{"ManifestTable", gvcp.AbrmManifestTableAddress},
		{"AccessPrivilege", gvcp.AbrmAccessPrivilege},
		{"HeartbeatTO", gvcp.AbrmHeartbeatTimeout},
		{"Stream0PacketSize", gvcp.Stream0PacketSize},
		{"Stream0Port", gvcp.Stream0Port},
	} {
		v, rerr := g.ReadReg(a.addr)
		if rerr != nil {
			fmt.Printf("  %-20s @0x%04x ERR %v\n", a.name, a.addr, rerr)
			continue
		}
		fmt.Printf("  %-20s @0x%04x = 0x%08x\n", a.name, a.addr, v)
	}

	url, err := g.FirstURL()
	if err != nil {
		return fmt.Errorf("FirstURL: %w", err)
	}
	fmt.Printf("FirstURL=%q\nHeartbeatTimeout()=%s\n", url, g.HeartbeatTimeout())

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
	fmt.Printf("FetchXML parsed %d bytes; node introspection:\n", len(xmlData))
	probeFeatures(nm, []string{
		"Width", "Height", "PixelFormat",
		"DeviceModelName", "DeviceUserID",
		"AcquisitionStart", "AcquisitionStartOrStop",
		"GevSCPSPacketSize", "GevStreamChannelCount",
	})

	camIP := net.ParseIP(ip)
	mtu := gvsp.PathMTU(camIP)
	fmt.Printf("PathMTU(%s)=%d → PacketSizeForMTU=%d\n", ip, mtu, gvsp.PacketSizeForMTU(mtu))
	return nil
}

// phase2GenApi exercises the high-level Camera + NodeMap feature surface:
// presence, kind, enum entries/current, integer reads and a non-destructive
// set (re-writes the current value to prove the write path).
func phase2GenApi(ip string) error {
	fmt.Println("--- Phase 2: feature get/set via OpenDevice ---")
	cam, err := gogige.OpenDevice(context.Background(), ip,
		gogige.WithLogger(printLogger{}), gogige.WithTimeout(3*time.Second))
	if err != nil {
		return fmt.Errorf("OpenDevice: %w", err)
	}
	defer cam.Close()

	for _, name := range []string{"Width", "Height", "PixelFormat", "DeviceUserID", "AcquisitionStart"} {
		fmt.Printf("  cam.Has(%s)=%v Kind=%q\n", name, cam.Has(name), cam.NodeMap().Kind(name))
	}

	// Read + re-write current value (a no-op write) to prove setters round-trip.
	for _, name := range []string{"Width", "Height"} {
		if !cam.Has(name) {
			continue
		}
		cur, err := cam.NodeMap().ReadInteger(name)
		if err != nil {
			fmt.Printf("  read %s: %v\n", name, err)
			continue
		}
		if err := cam.SetInteger(name, cur); err != nil {
			fmt.Printf("  set %s=%d: %v\n", name, cur, err)
			continue
		}
		fmt.Printf("  SetInteger(%s=%d): ok\n", name, cur)
	}
	if ent, err := cam.NodeMap().EnumEntries("PixelFormat"); err == nil {
		cur, _ := cam.NodeMap().CurrentEnum("PixelFormat")
		fmt.Printf("  PixelFormat enum entries=%d current=%s\n", len(ent), cur)
		if cur != "" {
			if err := cam.SetEnum("PixelFormat", cur); err != nil {
				fmt.Printf("  SetEnum(PixelFormat=%s): %v\n", cur, err)
			} else {
				fmt.Printf("  SetEnum(PixelFormat=%s): ok\n", cur)
			}
		}
	}
	return nil
}

// probeFeature kinds supported by the NodeMap (get only; no writes).
func probeFeatures(nm *gogige.NodeMap, names []string) {
	for _, name := range names {
		if !nm.Has(name) {
			fmt.Printf("  %-24s absent\n", name)
			continue
		}
		kind := nm.Kind(name)
		switch kind {
		case "Enumeration":
			entries, err := nm.EnumEntries(name)
			if err != nil {
				fmt.Printf("  %-24s enum   (err %v)\n", name, err)
				continue
			}
			cur, _ := nm.CurrentEnum(name)
			fmt.Printf("  %-24s enum   entries=%d current=%q\n", name, len(entries), cur)
			for _, e := range entries {
				fmt.Printf("      - %s\n", e)
			}
		case "Integer", "IntReg", "MaskedIntReg", "SwissKnife", "Converter":
			v, err := nm.ReadInteger(name)
			if err != nil {
				fmt.Printf("  %-24s %-12s (read %v)\n", name, kind, err)
				continue
			}
			min, hasMin, _ := nm.GetMin(name)
			max, hasMax, _ := nm.GetMax(name)
			inc, hasInc, _ := nm.GetInc(name)
			fmt.Printf("  %-24s %-12s = %d (min=%d max=%d inc=%d; has=%v/%v/%v)\n",
				name, kind, v, min, max, inc, hasMin, hasMax, hasInc)
		case "Boolean":
			if b, err := nm.ReadBoolean(name); err == nil {
				fmt.Printf("  %-24s bool    = %v\n", name, b)
			} else {
				fmt.Printf("  %-24s bool    (read %v)\n", name, err)
			}
		case "Float", "FloatReg":
			// No ReadFloat yet; report kind + constraints when available.
			if v, ok, _ := nm.GetMax(name); ok {
				fmt.Printf("  %-24s %-10s max=%g (via int)\n", name, kind, float64(v))
			} else {
				fmt.Printf("  %-24s %s (no read)\n", name, kind)
			}
		case "Command":
			fmt.Printf("  %-24s command\n", name)
		case "String", "StringReg":
			fmt.Printf("  %-24s %s (no read)\n", name, kind)
		default:
			fmt.Printf("  %-24s %s\n", name, kind)
		}
	}
}

// phase3Stream opens a high-level device, grabs frames via the Session,
// enumerates all components with GrabAll, and exercises pause/resume.
func phase3Stream(deviceIP string, n int, outDir string, comp gogige.Component) error {
	fmt.Println("--- Phase 3: Open + stream Session ---")
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	dev, err := gogige.Open(ctx, deviceIP, gogige.WithLogger(printLogger{}), gogige.WithTimeout(3*time.Second))
	if err != nil {
		return fmt.Errorf("Open: %w", err)
	}
	defer dev.Close()

	for _, name := range []string{"Width", "Height", "PixelFormat", "AcquisitionStart", "AcquisitionStop"} {
		fmt.Printf("  Features.Has(%s)=%v\n", name, dev.Features().Has(name))
	}

	grabber, err := dev.StartGrabber(ctx, gogige.GrabComponent(comp))
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
		fmt.Printf("  frame[%d] %-6s jpeg=%dB %dx%d packs=%d L=%.2f W=%.2f H=%.2f stable=%v\n",
			i, sample.Component, len(sample.JPEG), sample.Width, sample.Height, sample.PackCount,
			sample.Length, sample.WidthMm, sample.HeightMm, sample.Stable)
		if i == 0 {
			firstJPEG = append([]byte(nil), sample.JPEG...)
		}
	}

	// GrabAll enumerates every BSCF component available on the current frame.
	gctx, gcancel := context.WithTimeout(ctx, 5*time.Second)
	all, err := grabber.GrabAll(gctx)
	gcancel()
	if err != nil {
		return fmt.Errorf("GrabAll: %w", err)
	}
	fmt.Printf("  grab all: %d components\n", len(all))
	for _, s := range all {
		fmt.Printf("    %-8s %dx%d pixfmt=0x%08x jpeg=%dB\n",
			s.Component, s.Width, s.Height, s.PixelFormat, len(s.JPEG))
	}

	// Component switch (only when a different one exists) + pause/resume.
	if other, ok := otherComponent(comp); ok {
		grabber.SetComponent(other)
		fmt.Printf("  SetComponent(%s) applied; pause 1s…\n", other)
		if err := grabber.Pause(ctx); err != nil {
			return fmt.Errorf("Pause: %w", err)
		}
		time.Sleep(time.Second)
		if err := grabber.Resume(ctx); err != nil {
			return fmt.Errorf("Resume: %w", err)
		}
	}
	gctx, gcancel = context.WithTimeout(ctx, 5*time.Second)
	sample, err := grabber.Grab(gctx)
	gcancel()
	if err != nil {
		return fmt.Errorf("Grab after resume: %w", err)
	}
	fmt.Printf("  after resume: %-6s jpeg=%dB packs=%d\n", sample.Component, len(sample.JPEG), sample.PackCount)

	if outDir != "" && len(firstJPEG) > 0 {
		if err := os.MkdirAll(outDir, 0o755); err != nil {
			return err
		}
		path := filepath.Join(outDir, "frame-0-"+comp.String()+".jpg")
		if err := os.WriteFile(path, firstJPEG, 0o644); err != nil {
			return err
		}
		fmt.Println("wrote", path)
	}

	_ = grabber.Close()
	return nil
}

func otherComponent(c gogige.Component) (gogige.Component, bool) {
	for _, cand := range []gogige.Component{gogige.ComponentDepth, gogige.ComponentMono, gogige.ComponentColor} {
		if cand != c {
			return cand, true
		}
	}
	return 0, false
}

// phase4Channel exercises the Phase-4 channel stream API: StartStream, the
// Frames() channel of pooled frames, and frame.Release() back to the pool.
func phase4Channel(ip string) error {
	fmt.Println("--- Phase 4: StartStream channel API ---")
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	cam, err := gogige.OpenDevice(ctx, ip, gogige.WithLogger(printLogger{}))
	if err != nil {
		return fmt.Errorf("OpenDevice: %w", err)
	}
	defer cam.Close()

	st, err := cam.StartStream(ctx)
	if err != nil {
		return fmt.Errorf("StartStream: %w", err)
	}
	defer st.Stop()

	var got int
	end := time.After(5 * time.Second)
	for {
		select {
		case f, ok := <-st.Frames():
			if !ok {
				fmt.Printf("  frames channel closed; got %d\n", got)
				return nil
			}
			got++
			if got <= 3 {
				fmt.Printf("  ch[%d] id=0x%x len=%d pixfmt=0x%x\n",
					got, f.ID, len(f.Data), f.PixelFormat)
			}
			f.Release()
		case <-end:
			fmt.Printf("  channel drained; frames=%d\n", got)
			return nil
		}
	}
}

// phase5Live runs the Live preview loop with OnSample + LatestSample.
func phase5Live(ip string, d time.Duration) error {
	fmt.Println("--- Phase 5: Live preview ---")
	dev, err := gogige.Open(context.Background(), ip, gogige.WithLogger(printLogger{}))
	if err != nil {
		return fmt.Errorf("Open: %w", err)
	}
	defer dev.Close()

	var count int
	live := live.NewLive(dev, live.WithOnSample(func(s gogige.Sample) {
		count++
		if count <= 3 || count%10 == 0 {
			fmt.Printf("  live[%d] %-6s packs=%d L=%.2f\n", count, s.Component, s.PackCount, s.Length)
		}
	}))
	lctx, lcancel := context.WithTimeout(context.Background(), d)
	live.Start(lctx)
	<-lctx.Done()
	live.Stop()
	lcancel()
	latest := live.LatestSample()
	fmt.Printf("  live samples=%d latest packs=%d L=%.2f\n", count, latest.PackCount, latest.Length)
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
		if str, ok := kv[i+1].(string); ok {
			fmt.Printf(" %s=%q", kv[i], str)
		} else {
			fmt.Printf(" %s=%v", kv[i], kv[i+1])
		}
	}
	fmt.Println()
}
