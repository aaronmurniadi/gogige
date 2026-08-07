package gvsp

import (
	"fmt"
	"net"
	"syscall"
)

const (
	// DefaultRecvBufSize is the target SO_RCVBUF for GVSP (bytes).
	DefaultRecvBufSize = 16 << 20 // 16 MiB
	// MinRecvBufSize is the GigE streaming floor; warn when the kernel grants less.
	MinRecvBufSize = 8 << 20 // 8 MiB
	// DefaultMTU used when the path interface cannot be resolved.
	DefaultMTU = 1500
)

// PathMTU returns the MTU of a local interface that can reach dst (same subnet).
// Falls back to DefaultMTU when no matching interface is found.
func PathMTU(dst net.IP) int {
	if dst == nil {
		return DefaultMTU
	}
	dst = dst.To4()
	if dst == nil {
		return DefaultMTU
	}
	ifaces, err := net.Interfaces()
	if err != nil {
		return DefaultMTU
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			ipnet, ok := a.(*net.IPNet)
			if !ok || ipnet.IP.To4() == nil {
				continue
			}
			if ipnet.Contains(dst) && iface.MTU > 0 {
				return iface.MTU
			}
		}
	}
	return DefaultMTU
}

// PacketSizeForMTU maps link MTU to GevSCPSPacketSize (IP packet size including
// IP/UDP/GVSP headers). Clamped to a practical GigE range.
func PacketSizeForMTU(mtu int) int {
	if mtu < 576 {
		return DefaultMTU
	}
	if mtu > 16384 {
		return 16384
	}
	return mtu
}

// SetRecvBuffer requests SO_RCVBUF size (bytes) on the GVSP socket.
// Linux doubles the value for kernel bookkeeping; the returned size is the
// effective receive buffer (half of what Getsockopt reports on Linux).
func (s *Stream) SetRecvBuffer(bytes int) (got int, err error) {
	if s == nil || s.conn == nil {
		return 0, fmt.Errorf("gige: gvsp socket not open")
	}
	if bytes <= 0 {
		bytes = DefaultRecvBufSize
	}
	raw, err := s.conn.SyscallConn()
	if err != nil {
		return 0, fmt.Errorf("gige: gvsp syscall conn: %w", err)
	}
	var setErr, getErr error
	var kernelGot int
	cerr := raw.Control(func(fd uintptr) {
		setErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, bytes)
		kernelGot, getErr = syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF)
	})
	if cerr != nil {
		return 0, fmt.Errorf("gige: gvsp rcvbuf control: %w", cerr)
	}
	if setErr != nil {
		return 0, fmt.Errorf("gige: set SO_RCVBUF: %w", setErr)
	}
	if getErr != nil {
		return 0, fmt.Errorf("gige: get SO_RCVBUF: %w", getErr)
	}
	// Linux returns 2× the user request; report the application-visible size.
	got = kernelGot / 2
	if got <= 0 {
		got = kernelGot
	}
	s.mu.Lock()
	s.rcvBuf = got
	s.mu.Unlock()
	return got, nil
}

// RecvBuffer returns the last observed SO_RCVBUF size (0 if unset).
func (s *Stream) RecvBuffer() int {
	if s == nil {
		return 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.rcvBuf
}
