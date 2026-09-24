package gogige

// Well-known SFNC (Standard Feature Naming Convention) feature names for
// Device* and Gev* streaming-essential features. Use these constants with
// Camera.Integer / Camera.String / Camera.SetInteger / etc. to avoid magic
// strings. Features declared in the camera's GenICam XML are also accessible
// by their raw name via the generic accessors.

// Device identity features (SFNC §3).
const (
	FeatureDeviceVendorName       = "DeviceVendorName"
	FeatureDeviceModelName        = "DeviceModelName"
	FeatureDeviceSerialNumber     = "DeviceSerialNumber"
	FeatureDeviceUserID           = "DeviceUserID"
	FeatureDeviceTLType           = "DeviceTLType"
	FeatureDeviceSFNCVersionMajor = "DeviceSFNCVersionMajor"
)

// Device streaming / link features (SFNC §3.34–§3.46).
const (
	FeatureDeviceLinkHeartbeatTimeout = "DeviceLinkHeartbeatTimeout"
	FeatureDeviceStreamChannelCount   = "DeviceStreamChannelCount"
)

// Gev streaming features (GigE Vision / SFNC §27.4).
const (
	FeatureGevSCPSPacketSize     = "GevSCPSPacketSize"
	FeatureGevSCPD               = "GevSCPD"
	FeatureGevSCDA               = "GevSCDA"
	FeatureGevSCPHostPort        = "GevSCPHostPort"
	FeatureGevCCP                = "GevCCP"
	FeatureGevGVSPExtendedIDMode = "GevGVSPExtendedIDMode"
)

// Gev heartbeat feature (deprecated in SFNC 2.7, replaced by
// DeviceLinkHeartbeatTimeout, but still widely present in GigE Vision cameras).
const FeatureGevHeartbeatTimeout = "GevHeartbeatTimeout"
