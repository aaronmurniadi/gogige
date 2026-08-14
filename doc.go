// Package gogige is a pure-Go GigE Vision client.
//
// Controlled access: OpenDevice returns a *Camera. Use Camera.SetInteger /
// SetEnum / SetFloat / SetBoolean / SetString to write features and their
// matching getter (Integer / Enum / Float / Boolean / String) to read them.
//
//	cam, err := gogige.OpenDevice(ctx, "192.168.1.10")
//	defer cam.Close()
//	_ = cam.SetInteger("Width", 1920)
//	_ = cam.SetEnum("PixelFormat", "Mono8")
//	_ = cam.SetBoolean("AcquisitionStart", true)
//
// Live frames (Phase 4): Camera.StartStream → stream.Frames() of pooled frames;
// Release() each frame back to the buffer pool.
//
//	stream, err := cam.StartStream(ctx)
//	defer stream.Stop()
//	for frame := range stream.Frames() {
//		// use frame.Data
//		frame.Release()
//	}
//
// One-shot samples (Huaray BSCF): Camera.GrabSample / GrabAllSamples / GrabJPEG
// grab a single frame directly from the Camera. For continuous preview use a
// Device (Open → StartGrabber → Grabber.Grab) or live.NewLive.
//
// dev, _ := gogige.Open(ctx, ip)
// g, _ := dev.StartGrabber(ctx)
// sample, _ := g.Grab(ctx) // sample.JPEG + mm measurements
//
// Protocol packages:
//   - gvcp — GigE Vision Control Protocol (GenCP)
//   - genapi — GenICam GenApi XML / node map
//   - gvsp — GigE Vision Streaming Protocol
//   - gentl — GenTL constants (no CGO)
//
// References
//
//   - GigE Vision for Realtime MV (2010)
//   - GenICam GenCP Standard v1.3.1
//   - GenICam Standard v2.1.1
//   - GenICam GenApi Standard v2.1.1
//   - GenICam GenTL Standard v1.6
//   - GenICam GenDC Standard v1.1
//   - GenICam SFNC v2.7
//   - https://www.emva.org/standards-technical-documents/
package gogige // import "github.com/aaronmurniadi/gogige"
