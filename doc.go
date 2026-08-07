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
//   - gvcp — GigE Vision Control Protocol
//   - genapi — GenICam GenApi XML / node map
//   - gvsp — GigE Vision Streaming Protocol
//
// Device discovery (broadcast) lives on this root package: Discover.
// Preview sinks are app-owned (see examples/websocket-stream).
package gogige // import "github.com/aaronmurniadi/gogige"
