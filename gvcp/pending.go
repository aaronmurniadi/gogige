package gvcp

import (
	"encoding/binary"
	"time"
)

// pendingAckTimeout parses GenCP PENDING_ACK SCD (Table 12):
//
//	offset 0: reserved (2)
//	offset 2: temporary_timeout (2), milliseconds from when the ack was sent
//
// Falls back to defaultTO when the field is missing or zero.
func pendingAckTimeout(scd []byte, defaultTO time.Duration) time.Duration {
	if len(scd) < 4 {
		return defaultTO
	}
	ms := binary.BigEndian.Uint16(scd[2:])
	if ms == 0 {
		return defaultTO
	}
	return time.Duration(ms) * time.Millisecond
}
