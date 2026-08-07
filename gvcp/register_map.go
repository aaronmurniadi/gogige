package gvcp

// Bootstrap / stream-channel register addresses (GigE Vision ABRM/SBRM subset).
const (
	gvbsXMLURL0     = 0x00000200
	gvbsXMLURLSize  = 512
	gvbsCCP         = 0x00000a00
	gvbsCCPControl  = 1 << 1
	gvbsHeartbeatTO = 0x00000938

	Stream0Port        = 0x00000d00
	Stream0PacketSize  = 0x00000d04
	Stream0PacketDelay = 0x00000d08
	Stream0IP          = 0x00000d18
)
