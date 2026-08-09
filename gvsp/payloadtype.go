package gvsp

import "fmt"

// GVSP payload type IDs. These match GigE Vision / GenTL PayloadType
// constants: leader packet offset 0, 32-bit field (GenCP Table: CONTENT_TYPE
// carried in the leader's payload-type-specific header).
const (
	PayloadTypeImage          = 0x00000001 // standard image data
	PayloadTypeRawData        = 0x00000002
	PayloadTypeFile           = 0x00000003
	PayloadTypeChunkData      = 0x00000004 // chunk data attached before image (GenTL v1.2/v1.4)
	PayloadTypeJPEG           = 0x00000005
	PayloadTypeJPEG2000       = 0x00000006
	PayloadTypeH264           = 0x00000007
	PayloadTypeChunkOnly      = 0x00000008 // chunk-only, no image part (GenTL v1.4)
	PayloadTypeDeviceSpecific = 0x00000009
	PayloadTypeMultiPart      = 0x0000000A // multi-part (GenTL v1.5)
	PayloadTypeGenDC          = 0x0000000B // GenDC container (GenTL v1.6 / GenDC 1.1)

	// Some vendors extend the payload type field with the high bit set while
	// still using the lower byte as the GenTL id. Accepted aliases below.
	payloadTypeAliasGenDC     = 0x80000008 // vendor encoding observed in the wild
	payloadTypeAliasMultiPart = 0x80000007
	payloadTypeAliasChunk     = 0x80000009
)

// IsPayloadTypeImage reports whether a payload type denotes a plain image.
func IsPayloadTypeImage(t uint32) bool { return t == PayloadTypeImage }

// IsPayloadTypeGenDC accepts both the spec id and vendor aliases.
func IsPayloadTypeGenDC(t uint32) bool { return t == PayloadTypeGenDC || t == payloadTypeAliasGenDC }

// IsPayloadTypeMultiPart accepts both the spec id and vendor aliases.
func IsPayloadTypeMultiPart(t uint32) bool {
	return t == PayloadTypeMultiPart || t == payloadTypeAliasMultiPart
}

// IsPayloadTypeChunk accepts chunk-data and chunk-only forms.
func IsPayloadTypeChunk(t uint32) bool {
	return t == PayloadTypeChunkData || t == PayloadTypeChunkOnly || t == payloadTypeAliasChunk
}

// PayloadTypeName returns a human-readable payload type name.
func PayloadTypeName(t uint32) string {
	switch t {
	case PayloadTypeImage:
		return "IMAGE"
	case PayloadTypeRawData:
		return "RAW_DATA"
	case PayloadTypeFile:
		return "FILE"
	case PayloadTypeChunkData, payloadTypeAliasChunk:
		return "CHUNK_DATA"
	case PayloadTypeJPEG:
		return "JPEG"
	case PayloadTypeJPEG2000:
		return "JPEG2000"
	case PayloadTypeH264:
		return "H264"
	case PayloadTypeChunkOnly:
		return "CHUNK_ONLY"
	case PayloadTypeDeviceSpecific:
		return "DEVICE_SPECIFIC"
	case PayloadTypeMultiPart, payloadTypeAliasMultiPart:
		return "MULTI_PART"
	case PayloadTypeGenDC, payloadTypeAliasGenDC:
		return "GENDC"
	default:
		return fmt.Sprintf("0x%08x", t)
	}
}
