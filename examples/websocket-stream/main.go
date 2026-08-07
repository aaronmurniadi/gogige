// Live camera preview over WebSocket, with a browser page at /.
//
//	go run . -ip 192.168.1.10
//
// Open http://127.0.0.1:8080/ — JPEG frames on /ws; send text "freeze"/"resume".
package main

import (
	"context"
	"embed"
	"flag"
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
	ip := flag.String("ip", "127.0.0.1", "camera IP")
	addr := flag.String("addr", "127.0.0.1:8080", "HTTP listen address")
	flag.Parse()

	ctx := context.Background()
	dev, err := gige.Open(ctx, *ip)
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

	live := gige.NewLive(dev, gige.WithSink(hub))
	live.Start(ctx)
	defer live.Stop()

	log.Printf("open http://%s/  (ws://%s/ws)", *addr, *addr)
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

var _ gige.FrameSink = (*hub)(nil)
