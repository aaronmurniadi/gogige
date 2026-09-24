package gogige

import "testing"

func TestSFNCFeatureConstants(t *testing.T) {
	// All exported feature-name constants must be non-empty so that
	// Camera.Integer / Camera.String / etc. resolve a real GenApi node name.
	consts := []string{
		FeatureDeviceVendorName,
		FeatureDeviceModelName,
		FeatureDeviceSerialNumber,
		FeatureDeviceUserID,
		FeatureDeviceTLType,
		FeatureDeviceSFNCVersionMajor,
		FeatureDeviceLinkHeartbeatTimeout,
		FeatureDeviceStreamChannelCount,
		FeatureGevSCPSPacketSize,
		FeatureGevSCPD,
		FeatureGevSCDA,
		FeatureGevSCPHostPort,
		FeatureGevCCP,
		FeatureGevGVSPExtendedIDMode,
		FeatureGevHeartbeatTimeout,
	}
	for _, c := range consts {
		if c == "" {
			t.Fatal("empty SFNC feature constant")
		}
	}
}
