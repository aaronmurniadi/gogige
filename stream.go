package gige

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aaronmurniadi/gogige/gvcp"
	"github.com/aaronmurniadi/gogige/gvsp"
	"github.com/aaronmurniadi/gogige/internal/color"
)

// Session is a persistent GVSP stream session. It implements Grabber.
type Session struct {
	cam        *Camera
	stream     *gvsp.Stream
	opened     atomic.Bool
	mu         sync.Mutex
	ownCam     bool // if true, Close() closes cam
	ip         string
	packetSize int

	hb *gvcp.Heartbeat
}

// NewSession returns an empty GVSP session (Connects on Open if no camera set).
func NewSession() *Session { return &Session{} }

// NewFromCamera returns a Session that streams using cam without taking ownership.
func NewFromCamera(cam *Camera) *Session {
	return &Session{cam: cam}
}

// Open connects (if needed), starts GVSP, and begins acquisition.
func (s *Session) Open(ip string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.opened.Load() {
		return nil
	}
	s.ip = ip
	if s.cam == nil {
		cam, err := Connect(ip)
		if err != nil {
			return err
		}
		s.cam = cam
		s.ownCam = true
	}
	stream, err := gvsp.ListenStream(0)
	if err != nil {
		return err
	}
	s.stream = stream
	if g := s.cam.GVCP(); g != nil {
		stream.SetResender(func(blockID uint64, first, last uint32, extended bool) {
			_ = g.RequestResend(0, blockID, first, last, extended)
		})
	}
	if got := stream.RecvBuffer(); got > 0 && got < gvsp.MinRecvBufSize {
		s.cam.Logger().Warn("gvsp SO_RCVBUF below recommended",
			"got", got, "min", gvsp.MinRecvBufSize, "want", gvsp.DefaultRecvBufSize)
	}

	local := s.cam.GVCP().LocalAddr()
	if local == nil || local.IP == nil {
		stream.Close()
		s.stream = nil
		return errors.New("gige: no local address for stream destination")
	}
	destIP := local.IP
	if destIP.IsUnspecified() {
		destIP = firstIPv4()
	}
	camIP := net.ParseIP(ip)
	if camIP == nil {
		camIP = net.ParseIP(s.cam.IP)
	}
	mtu := gvsp.PathMTU(camIP)
	want := gvsp.PacketSizeForMTU(mtu)
	if err := gvcp.StartAcquisition(s.cam.GVCP(), s.cam, destIP, stream.Port(), want); err != nil {
		stream.Close()
		s.stream = nil
		return err
	}
	s.packetSize = want
	if reg, err := s.cam.GVCP().ReadReg(gvcp.Stream0PacketSize); err == nil {
		s.packetSize = int(reg & 0xffff)
	}
	if s.packetSize < want {
		s.cam.Logger().Warn("GevSCPSPacketSize clamped by device",
			"want", want, "got", s.packetSize, "path_mtu", mtu)
	}
	s.cam.Logger().Info("gvsp stream setup",
		"mtu", mtu, "scps", s.packetSize, "rcvbuf", stream.RecvBuffer())
	s.startHeartbeatLocked()
	s.opened.Store(true)
	return nil
}

func firstIPv4() net.IP {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, a := range addrs {
			var ip net.IP
			switch v := a.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip4 := ip.To4(); ip4 != nil {
				return ip4
			}
		}
	}
	return net.IPv4(127, 0, 0, 1)
}

// Close stops acquisition and releases stream resources.
// The Camera is closed only when the Session owns it (not when Device owns it).
func (s *Session) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.opened.Store(false)
	s.stopHeartbeatLocked()
	if s.cam != nil {
		_ = gvcp.StopAcquisition(s.cam)
	}
	if s.stream != nil {
		_ = s.stream.Close()
		s.stream = nil
	}
	if s.ownCam && s.cam != nil {
		s.cam.Close()
		s.cam = nil
	}
	return nil
}

// PauseStreaming stops image transfer but keeps CCP, heartbeat, and GVSP socket.
func (s *Session) PauseStreaming() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.opened.Load() {
		return nil
	}
	if s.cam != nil {
		_ = gvcp.StopAcquisition(s.cam)
		if g := s.cam.GVCP(); g != nil {
			_ = g.WriteReg(gvcp.Stream0Port, 0)
		}
	}
	return nil
}

// ResumeStreaming re-programs the stream channel and starts acquisition again.
func (s *Session) ResumeStreaming() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.opened.Load() {
		return errors.New("gige: stream not open")
	}
	if s.cam == nil || s.stream == nil {
		return errors.New("gige: stream not open")
	}
	local := s.cam.GVCP().LocalAddr()
	if local == nil || local.IP == nil {
		return errors.New("gige: no local address for stream destination")
	}
	destIP := local.IP
	if destIP.IsUnspecified() {
		destIP = firstIPv4()
	}
	ps := s.packetSize
	if ps <= 0 {
		ps = gvsp.PacketSizeForMTU(gvsp.PathMTU(net.ParseIP(s.ip)))
		s.packetSize = ps
	}
	return gvcp.StartAcquisition(s.cam.GVCP(), s.cam, destIP, s.stream.Port(), ps)
}

// Pause implements Grabber.
func (s *Session) Pause(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.PauseStreaming()
}

// Resume implements Grabber.
func (s *Session) Resume(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.ResumeStreaming()
}

func (s *Session) startHeartbeatLocked() {
	s.stopHeartbeatLocked()
	if s.cam == nil || s.cam.GVCP() == nil {
		return
	}
	s.hb = s.cam.GVCP().StartHeartbeat()
}

func (s *Session) stopHeartbeatLocked() {
	if s.hb == nil {
		return
	}
	s.hb.Stop()
	s.hb = nil
}

// Opened reports whether the session is open.
func (s *Session) Opened() bool { return s.opened.Load() }

// Grab implements Grabber: receives one GVSP frame, parses BSCF, returns JPEG Sample.
func (s *Session) Grab(ctx context.Context) (Sample, error) {
	sample := Sample{PackCount: -1}
	if err := ctx.Err(); err != nil {
		return sample, err
	}
	timeout := time.Second
	if dl, ok := ctx.Deadline(); ok {
		timeout = time.Until(dl)
		if timeout <= 0 {
			return sample, context.DeadlineExceeded
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.opened.Load() || s.stream == nil {
		return sample, errors.New("gige: stream not open")
	}
	frame, err := s.stream.Recv(timeout)
	if err != nil {
		return sample, err
	}
	defer frame.Release()
	sample, err = SampleFromBSCF(frame.Data)
	if err != nil {
		// Copy off the pooled buffer before Release; JPEG path needs ownership.
		raw := append([]byte(nil), frame.Data...)
		sample.RawColor = raw
		sample.Width = int(frame.Width)
		sample.Height = int(frame.Height)
		sample.PixelFormat = frame.PixelFormat
		if sample.Width == 0 || sample.Height == 0 {
			return sample, fmt.Errorf("gige: grab: %w", err)
		}
	}
	jpeg, jerr := color.EncodeJPEG(sample.RawColor, sample.Width, sample.Height, sample.PixelFormat, 60)
	if jerr != nil {
		return sample, jerr
	}
	sample.JPEG = jpeg
	return sample, nil
}
