// Package gogige is a pure-Go GigE Vision client.
//
// Happy path (Phase 4):
//
//	cam, err := gogige.OpenDevice(ctx, "192.168.1.10")
//	defer cam.Close()
//	_ = cam.SetInteger("Width", 1920)
//	_ = cam.SetEnum("PixelFormat", "Mono8")
//	stream, err := cam.StartStream(ctx)
//	defer stream.Stop()
//	for frame := range stream.Frames() {
//		// use frame.Data
//		frame.Release()
//	}
//
// Sample/JPEG path (Huaray BSCF): Device.StartGrabber → Grabber.Grab.
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
