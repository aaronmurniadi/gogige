// Live camera preview over WebSocket, with a browser page at /.
//
//	go run .                         # discover first camera, color stream
//	go run . -component depth            # depth preview
//	go run . -ip 192.168.1.108 -component mono
//
// Open http://127.0.0.1:8080/ — JPEG frames on /ws; send text "freeze"/"resume".
package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/aaronmurniadi/gogige"
	"github.com/gorilla/websocket"
)

//go:embed index.html
var content embed.FS

func main() {
	ip := flag.String("ip", "", "camera IP (empty = first GigE discovery hit)")
	addr := flag.String("addr", "127.0.0.1:8080", "HTTP listen address")
	component := flag.String("component", "color", "component: color|depth|mono")
	flag.Parse()

	kind, err := gogige.ParseComponent(*component)
	if err != nil {
		log.Fatal(err)
	}

	cameraIP, err := resolveIP(*ip)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	dev, err := gogige.Open(ctx, cameraIP)
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	hub := newHub()
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", hub.serveWS)
	static, err := fs.Sub(content, ".")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/", http.FileServer(http.FS(static)))

	live := gogige.NewLive(dev, gogige.WithSink(hub), gogige.WithLiveComponent(kind))
	live.Start(ctx)
	defer live.Stop()

	log.Printf("streaming %s — open http://%s/  (ws://%s/ws)", kind, *addr, *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}

type hub struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]chan []byte
	frozen  bool
	up      websocket.Upgrader
}

func newHub() *hub {
	return &hub{
		clients: make(map[*websocket.Conn]chan []byte),
		up:      websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }},
	}
}

func (h *hub) SendJPEG(jpeg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.frozen || len(jpeg) == 0 {
		return
	}
	payload := append([]byte(nil), jpeg...)
	for _, ch := range h.clients {
		select {
		case ch <- payload:
		default:
		}
	}
}

func (h *hub) Freeze() {
	h.mu.Lock()
	h.frozen = true
	h.mu.Unlock()
}

func (h *hub) Resume() {
	h.mu.Lock()
	h.frozen = false
	h.mu.Unlock()
}

func (h *hub) serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := h.up.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	ch := make(chan []byte, 1)
	h.mu.Lock()
	h.clients[conn] = ch
	h.mu.Unlock()

	go func() {
		defer conn.Close()
		for jpeg := range ch {
			_ = conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
			if err := conn.WriteMessage(websocket.BinaryMessage, jpeg); err != nil {
				return
			}
		}
	}()

	for {
		mt, data, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if mt == websocket.TextMessage {
			switch string(data) {
			case `{"type":"freeze"}`, "freeze":
				h.Freeze()
			case `{"type":"resume"}`, "resume":
				h.Resume()
			}
		}
	}

	h.mu.Lock()
	delete(h.clients, conn)
	close(ch)
	h.mu.Unlock()
}

var _ gogige.FrameSink = (*hub)(nil)

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
	log.Printf("discovered %s", devs[0].IP)
	return devs[0].IP, nil
}
