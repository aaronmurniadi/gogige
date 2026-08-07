// Package gige is a pure-Go GigE Vision client.
//
// Happy path (root package):
//
//	dev, err := gige.Open(ctx, "192.168.1.10")
//	defer dev.Close()
//	g, err := dev.StartGrabber(ctx)
//	defer g.Close()
//	sample, err := g.Grab(ctx)
//
// Protocol packages:
//   - gvcp — GigE Vision Control Protocol
//   - genapi — GenICam GenApi XML / node map
//   - gvsp — GigE Vision Streaming Protocol
//
// Device discovery (broadcast) lives on this root package: Discover.
// Preview sinks are app-owned (see examples/websocket-stream).
package gige // import "github.com/aaronmurniadi/gogige"
