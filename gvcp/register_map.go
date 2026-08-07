package gvcp

// Bootstrap register addresses.
//
// GenCP 1.3.1 Technology-Agnostic BRM occupies 0x0000–0x024F (reserved from
// 0x0250). GigE Vision uses a different ABRM/SBRM layout in the same low
// address space (First URL at 0x0200 overlaps GenCP ImplementationEndianness
// at 0x020C). GigE stream/CCP registers live higher (0x0938, 0x0A00, 0x0D00…).

// GenCP Technology-Agnostic Bootstrap Register Map (ABRM).
const (
	AbrmGenCPVersion                 = 0x0000 // 4: major/minor
	AbrmManufacturerName             = 0x0004 // 64
	AbrmModelName                    = 0x0044 // 64
	AbrmFamilyName                   = 0x0084 // 64 (CM)
	AbrmDeviceVersion                = 0x00C4 // 64
	AbrmManufacturerInfo             = 0x0104 // 64
	AbrmSerialNumber                 = 0x0144 // 64
	AbrmUserDefinedName              = 0x0184 // 64 (CM)
	AbrmDeviceCapability             = 0x01C4 // 8
	AbrmMaximumDeviceResponseTime    = 0x01CC // 4, ms
	AbrmManifestTableAddress         = 0x01D0 // 8
	AbrmSBRMAddress                  = 0x01D8 // 8 (CM)
	AbrmDeviceConfiguration          = 0x01E0 // 8
	AbrmHeartbeatTimeout             = 0x01E8 // 4, ms (CM)
	AbrmMessageChannelID             = 0x01EC // 4 (CM)
	AbrmTimestamp                    = 0x01F0 // 8 (CM)
	AbrmTimestampLatch               = 0x01F8 // 4 (CM)
	AbrmTimestampIncrement           = 0x01FC // 8 (CM)
	AbrmAccessPrivilege              = 0x0204 // 4 (CM)
	AbrmProtocolEndiannessDeprecated = 0x0208 // 4, do not use
	AbrmImplementationEndianness     = 0x020C // 4 (CM): 0=BE, 0xFFFFFFFF=LE
	AbrmDeviceSoftwareInterfaceVer   = 0x0210 // 64 (CM)
	AbrmReserved                     = 0x0250 // start of reserved space

	// DeviceCapability bit: Endianness Register Supported (bit 10).
	AbrmCapEndiannessRegister = 1 << 10

	EndiannessBig    = uint32(0)
	EndiannessLittle = uint32(0xFFFFFFFF)
)

// GigE Vision ABRM / SBRM (technology-specific layout used by this library).
const (
	gevVersion              = 0x0000
	gevDeviceMode           = 0x0004
	gevMACHigh              = 0x0008
	gevMACLow               = 0x000C
	gevCurrentIP            = 0x0024
	gevCurrentSubnet        = 0x0034
	gevCurrentGateway       = 0x0044
	gevManufacturerName     = 0x0048 // 32
	gevModelName            = 0x0068 // 32
	gevDeviceVersion        = 0x0088 // 32
	gevManufacturerInfo     = 0x00A8 // 48
	gevSerialNumber         = 0x00D8 // 16
	gevUserDefinedName      = 0x00E8 // 16
	gvbsXMLURL0             = 0x0200 // First URL, 512 bytes
	gvbsXMLURLSize          = 512
	gevSecondURL            = 0x0400 // 512
	gevNumNetworkInterfaces = 0x0600
	gevPersistentIP         = 0x064C
	gevPersistentSubnet     = 0x065C
	gevPersistentGateway    = 0x066C
	gevLinkSpeed            = 0x0670
	gevNumMessageChannels   = 0x0900
	gevNumStreamChannels    = 0x0904
	gevNumActionSignals     = 0x092C
	gevActionDeviceKey      = 0x0934
	gvbsHeartbeatTO         = 0x0938
	gevTimestampTickHigh    = 0x093C
	gevTimestampTickLow     = 0x0940
	gevTimestampControl     = 0x0944
	gevTimestampValueHigh   = 0x0948
	gevTimestampValueLow    = 0x094C
	gevDiscoveryACKDelay    = 0x0950
	gvbsCCP                 = 0x0A00
	gvbsCCPControl          = 1 << 1
	gevPrimaryAppPort       = 0x0A04
	gevPrimaryAppIP         = 0x0A14
	gevMessageChannelPort   = 0x0B00

	Stream0Port        = 0x0D00
	Stream0PacketSize  = 0x0D04
	Stream0PacketDelay = 0x0D08
	Stream0IP          = 0x0D18
	Stream0SourcePort  = 0x0D1C
	Stream0Capability  = 0x0D20
	Stream0Config      = 0x0D24
)
