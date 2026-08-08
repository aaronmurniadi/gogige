// Package gentl provides GenTL (GenICam Transport Layer) definitions.
//
// GenTL Version 1.6
// Reference: GenICam_GenTL_1_6.pdf
package gentl

// Error codes (GC_ERROR from GenTL.h)
const (
	ErrSuccess           = 0
	ErrError             = -1001
	ErrNotInitialized    = -1002
	ErrNotImplemented    = -1003
	ErrResourceInUse     = -1004
	ErrAccessDenied      = -1005
	ErrInvalidHandle     = -1006
	ErrInvalidID         = -1007
	ErrNoData            = -1008
	ErrInvalidParameter  = -1009
	ErrIO                = -1010
	ErrTimeout           = -1011
	ErrAbort             = -1012 // GenTL v1.1
	ErrInvalidBuffer     = -1013 // GenTL v1.1
	ErrNotAvailable      = -1014 // GenTL v1.2
	ErrInvalidAddress    = -1015 // GenTL v1.3
	ErrBufferTooSmall    = -1016 // GenTL v1.4
	ErrInvalidIndex      = -1017 // GenTL v1.4
	ErrParsingChunkData  = -1018 // GenTL v1.4
	ErrInvalidValue      = -1019 // GenTL v1.4
	ErrResourceExhausted = -1020 // GenTL v1.4
	ErrOutOfMemory       = -1021 // GenTL v1.4
	ErrBusy              = -1022 // GenTL v1.5
	ErrAmbiguous         = -1023 // GenTL v1.6
	ErrCustomID          = -10000
)

// GenTL header version constants
const (
	GenTLMajorVersion    = 1
	GenTLMinorVersion    = 6
	GenTLSubMinorVersion = 0
)

// Transport layer types
const (
	TLTypeMixed    = "Mixed"
	TLTypeCustom   = "Custom"
	TLTypeGEV      = "GEV"  // GigE Vision
	TLTypeCL       = "CL"   // Camera Link
	TLTypeIIDC     = "IIDC" // IIDC 1394
	TLTypeUVC      = "UVC"  // USB Video Class
	TLTypeCXP      = "CXP"  // CoaXPress
	TLTypeCLHS     = "CLHS" // Camera Link HS
	TLTypeU3V      = "U3V"  // USB3 Vision
	TLTypeEthernet = "Ethernet"
	TLTypePCI      = "PCI"
)

// GenTL module names
const (
	TLSystemModuleName       = "TLSystem"
	TLInterfaceModuleName    = "TLInterface"
	TLDeviceModuleName       = "TLDevice"
	TLDataStreamModuleName   = "TLDataStream"
	TLBufferModuleName       = "TLBuffer"
	TLRemoteDeviceModuleName = "Device"
)

// Info data types (INFO_DATATYPE_LIST)
const (
	InfoDataTypeUnknown    = 0
	InfoDataTypeString     = 1
	InfoDataTypeStringList = 2
	InfoDataTypeInt16      = 3
	InfoDataTypeUInt16     = 4
	InfoDataTypeInt32      = 5
	InfoDataTypeUInt32     = 6
	InfoDataTypeInt64      = 7
	InfoDataTypeUInt64     = 8
	InfoDataTypeFloat64    = 9
	InfoDataTypePtr        = 10
	InfoDataTypeBool8      = 11
	InfoDataTypeSizeT      = 12
	InfoDataTypeBuffer     = 13
	InfoDataTypePtrDiff    = 14
)

// Character encodings (TL_CHAR_ENCODING_LIST)
const (
	TLCharEncodingASCII = 0
	TLCharEncodingUTF8  = 1
)

// System module info commands (TL_INFO_CMD_LIST)
const (
	TLInfoID            = 0
	TLInfoVendor        = 1
	TLInfoModel         = 2
	TLInfoVersion       = 3
	TLInfoTLType        = 4
	TLInfoName          = 5
	TLInfoPathname      = 6
	TLInfoDisplayName   = 7
	TLInfoCharEncoding  = 8
	TLInfoGenTLVerMajor = 9
	TLInfoGenTLVerMinor = 10
)

// Interface info commands (INTERFACE_INFO_CMD_LIST)
const (
	InterfaceInfoID          = 0
	InterfaceInfoDisplayName = 1
	InterfaceInfoTLType      = 2
)

// Device access flags (DEVICE_ACCESS_FLAGS_LIST)
const (
	DeviceAccessUnknown   = 0
	DeviceAccessNone      = 1
	DeviceAccessReadOnly  = 2
	DeviceAccessControl   = 3
	DeviceAccessExclusive = 4
)

// Device access status (DEVICE_ACCESS_STATUS_LIST)
const (
	DeviceAccessStatusUnknown       = 0
	DeviceAccessStatusReadWrite     = 1
	DeviceAccessStatusReadOnly      = 2
	DeviceAccessStatusNoAccess      = 3
	DeviceAccessStatusBusy          = 4 // GenTL v1.5
	DeviceAccessStatusOpenReadWrite = 5 // GenTL v1.5
	DeviceAccessStatusOpenReadOnly  = 6 // GenTL v1.5
)

// Device info commands (DEVICE_INFO_CMD_LIST)
const (
	DeviceInfoID                 = 0
	DeviceInfoVendor             = 1
	DeviceInfoModel              = 2
	DeviceInfoTLType             = 3
	DeviceInfoDisplayName        = 4
	DeviceInfoAccessStatus       = 5
	DeviceInfoUserDefinedName    = 6 // GenTL v1.4
	DeviceInfoSerialNumber       = 7 // GenTL v1.4
	DeviceInfoVersion            = 8 // GenTL v1.4
	DeviceInfoTimestampFrequency = 9 // GenTL v1.4
)

// Acquisition stop flags (ACQ_STOP_FLAGS_LIST)
const (
	AcqStopFlagsDefault = 0
	AcqStopFlagsKill    = 1
)

// Acquisition start flags (ACQ_START_FLAGS_LIST)
const (
	AcqStartFlagsDefault = 0
)

// Acquisition queue types (ACQ_QUEUE_TYPE_LIST)
const (
	AcqQueueInputToOutput   = 0
	AcqQueueOutputDiscard   = 1
	AcqQueueAllToInput      = 2
	AcqQueueUnqueuedToInput = 3
	AcqQueueAllDiscard      = 4
)

// Stream info commands (STREAM_INFO_CMD_LIST)
const (
	StreamInfoID                      = 0
	StreamInfoNumDelivered            = 1
	StreamInfoNumUnderrun             = 2
	StreamInfoNumAnnounced            = 3
	StreamInfoNumQueued               = 4
	StreamInfoNumAwaitDelivery        = 5
	StreamInfoNumStarted              = 6
	StreamInfoPayloadSize             = 7
	StreamInfoIsGrabbing              = 8
	StreamInfoDefinesPayloadSize      = 9
	StreamInfoTLType                  = 10
	StreamInfoNumChunksMax            = 11 // GenTL v1.3
	StreamInfoBufAnnounceMin          = 12 // GenTL v1.3
	StreamInfoBufAlignment            = 13 // GenTL v1.3
	StreamInfoFlowTable               = 14 // GenTL v1.6
	StreamInfoGenDCPrefetchDescriptor = 15 // GenTL v1.6
)

// Buffer info commands (BUFFER_INFO_CMD_LIST)
const (
	BufferInfoBase                      = 0
	BufferInfoSize                      = 1
	BufferInfoUserPtr                   = 2
	BufferInfoTimestamp                 = 3
	BufferInfoNewData                   = 4
	BufferInfoIsQueued                  = 5
	BufferInfoIsAcquiring               = 6
	BufferInfoIsIncomplete              = 7
	BufferInfoTLType                    = 8
	BufferInfoSizeFilled                = 9
	BufferInfoWidth                     = 10 // GenTL v1.2
	BufferInfoHeight                    = 11 // GenTL v1.2
	BufferInfoXOffset                   = 12 // GenTL v1.2
	BufferInfoYOffset                   = 13 // GenTL v1.2
	BufferInfoXPadding                  = 14 // GenTL v1.2
	BufferInfoYPadding                  = 15 // GenTL v1.2
	BufferInfoFrameID                   = 16 // GenTL v1.2
	BufferInfoImagePresent              = 17 // GenTL v1.2
	BufferInfoImageOffset               = 18 // GenTL v1.2
	BufferInfoPayloadType               = 19 // GenTL v1.2
	BufferInfoPixelFormat               = 20 // GenTL v1.2
	BufferInfoPixelFormatNamespace      = 21 // GenTL v1.2
	BufferInfoDeliveredImageHeight      = 22 // GenTL v1.2
	BufferInfoDeliveredChunkPayloadSize = 23 // GenTL v1.2
	BufferInfoChunkLayoutID             = 24 // GenTL v1.2
	BufferInfoFilename                  = 25 // GenTL v1.2
	BufferInfoPixelEndianness           = 26 // GenTL v1.4
	BufferInfoDataSize                  = 27 // GenTL v1.4
	BufferInfoTimestampNS               = 28 // GenTL v1.4
	BufferInfoDataLargerThanBuffer      = 29 // GenTL v1.4
	BufferInfoContainsChunkData         = 30 // GenTL v1.4
	BufferInfoIsComposite               = 31 // GenTL v1.6
)

// Buffer part info commands (BUFFER_PART_INFO_CMD_LIST)
const (
	BufferPartInfoBase                 = 0
	BufferPartInfoDataSize             = 1
	BufferPartInfoDataType             = 2
	BufferPartInfoDataFormat           = 3
	BufferPartInfoDataFormatNamespace  = 4
	BufferPartInfoWidth                = 5
	BufferPartInfoHeight               = 6
	BufferPartInfoXOffset              = 7
	BufferPartInfoYOffset              = 8
	BufferPartInfoXPadding             = 9
	BufferPartInfoSourceID             = 10
	BufferPartInfoDeliveredImageHeight = 11
	BufferPartInfoRegionID             = 12 // GenTL v1.6
	BufferPartInfoDataPurposeID        = 13 // GenTL v1.6
)

// Payload type IDs (PAYLOADTYPE_INFO_IDS)
const (
	PayloadTypeUnknown        = 0
	PayloadTypeImage          = 1
	PayloadTypeRawData        = 2
	PayloadTypeFile           = 3
	PayloadTypeChunkData      = 4 // Deprecated in GenTL 1.5
	PayloadTypeJPEG           = 5
	PayloadTypeJPEG2000       = 6
	PayloadTypeH264           = 7
	PayloadTypeChunkOnly      = 8
	PayloadTypeDeviceSpecific = 9
	PayloadTypeMultiPart      = 10 // GenTL v1.5
	PayloadTypeGenDC          = 11 // GenTL v1.6
)

// Pixel format namespace IDs (PIXELFORMAT_NAMESPACE_IDS)
const (
	PixelFormatNamespaceUnknown   = 0
	PixelFormatNamespaceGEV       = 1
	PixelFormatNamespaceIIDC      = 2
	PixelFormatNamespacePFNC16Bit = 3
	PixelFormatNamespacePFNC32Bit = 4
)

// Pixel endianness (PIXELENDIANNESS_IDS)
const (
	PixelEndiannessUnknown = 0
	PixelEndiannessLittle  = 1
	PixelEndiannessBig     = 2
)

// Part data type IDs (PARTDATATYPE_IDS)
const (
	PartDataTypeUnknown           = 0
	PartDataType2DImage           = 1
	PartDataType2DPlaneBiplanar   = 2
	PartDataType2DPlaneTriplanar  = 3
	PartDataType2DPlaneQuadplanar = 4
	PartDataType3DImage           = 5
	PartDataType3DPlaneBiplanar   = 6
	PartDataType3DPlaneTriplanar  = 7
	PartDataType3DPlaneQuadplanar = 8
	PartDataTypeConfidenceMap     = 9
	PartDataTypeJPEG              = 10
	PartDataTypeJPEG2000          = 11
)

// Port info commands (PORT_INFO_CMD_LIST)
const (
	PortInfoID           = 0
	PortInfoVendor       = 1
	PortInfoModel        = 2
	PortInfoTLType       = 3
	PortInfoModule       = 4
	PortInfoLittleEndian = 5
	PortInfoBigEndian    = 6
	PortInfoAccessRead   = 7
	PortInfoAccessWrite  = 8
	PortInfoAccessNA     = 9
	PortInfoAccessNI     = 10
	PortInfoVersion      = 11
	PortInfoPortName     = 12
)

// URL scheme IDs (URL_SCHEME_IDS)
const (
	URLSchemeLocal = 0
	URLSchemeHTTP  = 1
	URLSchemeFile  = 2
)

// URL info commands (URL_INFO_CMD_LIST)
const (
	URLInfoURL                 = 0
	URLInfoSchemaVerMajor      = 1
	URLInfoSchemaVerMinor      = 2
	URLInfoFileVerMajor        = 3
	URLInfoFileVerMinor        = 4
	URLInfoFileVerSubminor     = 5
	URLInfoFileSHA1Hash        = 6
	URLInfoFileRegisterAddress = 7
	URLInfoFileSize            = 8
	URLInfoScheme              = 9
	URLInfoFilename            = 10
)

// Event types (EVENT_TYPE_LIST)
const (
	EventTypeError             = 0
	EventTypeNewBuffer         = 1
	EventTypeFeatureInvalidate = 2
	EventTypeFeatureChange     = 3
	EventTypeRemoteDevice      = 4
	EventTypeModule            = 5 // GenTL v1.4
)

// Event info commands (EVENT_INFO_CMD_LIST)
const (
	EventInfoEventType   = 0
	EventInfoNumInQueue  = 1
	EventInfoNumFired    = 2
	EventInfoSizeMax     = 3
	EventInfoDataSizeMax = 4
)

// Event data info commands (EVENT_DATA_INFO_CMD_LIST)
const (
	EventDataInfoID    = 0
	EventDataInfoValue = 1
	EventDataInfoNumID = 2
)

// Flow info commands (FLOW_INFO_CMD_LIST)
const (
	FlowInfoSize = 0
)

// Segment info commands (SEGMENT_INFO_CMD_LIST)
const (
	SegmentInfoBase         = 0
	SegmentInfoSize         = 1
	SegmentInfoIsIncomplete = 2
	SegmentInfoSizeFilled   = 3
	SegmentInfoDataSize     = 4
)

// Invalid handle and infinite timeout values
const (
	InvalidHandle = 0
	Infinite      = 0xFFFFFFFFFFFFFFFF
)

// Version returns the GenTL standard version supported.
func Version() (major, minor, subMinor int) {
	return GenTLMajorVersion, GenTLMinorVersion, GenTLSubMinorVersion
}
