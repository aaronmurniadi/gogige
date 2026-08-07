package gvsp

// PacketRange is an inclusive GVSP packet_id span for PACKETRESEND_CMD.
type PacketRange struct {
	First, Last uint32
}

// ResendFunc requests retransmission of GVSP packets over the GVCP control channel.
// first/last are inclusive packet_ids; extended matches the GVSP EI header mode for the block.
type ResendFunc func(blockID uint64, first, last uint32, extended bool)

// MissingPayloadRanges returns inclusive gaps in [from, to] for a frame under reassembly.
// Packets with id < nextContiguous are already in the contiguous buffer; have holds OOO payloads.
func MissingPayloadRanges(nextContiguous uint32, have map[uint32][]byte, from, to uint32) []PacketRange {
	if to < from {
		return nil
	}
	var out []PacketRange
	var gapStart uint32
	inGap := false
	for id := from; id <= to; id++ {
		ok := id < nextContiguous
		if !ok && have != nil {
			_, ok = have[id]
		}
		if !ok {
			if !inGap {
				gapStart = id
				inGap = true
			}
			continue
		}
		if inGap {
			out = append(out, PacketRange{First: gapStart, Last: id - 1})
			inGap = false
		}
	}
	if inGap {
		out = append(out, PacketRange{First: gapStart, Last: to})
	}
	return out
}

func (fb *frameBuild) missingPayloadRanges(from, to uint32) []PacketRange {
	if fb == nil {
		return nil
	}
	return MissingPayloadRanges(fb.nextPkt, fb.parts, from, to)
}
