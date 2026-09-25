package calib

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewVendorCalibJSON converts a calibration downloaded from the camera into
// the JSON document the vendor tools write (CalibFile::struct2Json in
// libStereoCamera.so): StereoCameraViewer saves it as CalibData.json, vendor
// tools as "IPC4.94 Camera Calibration.json".
//
// Everything in MvSstereoCalibrateResult that the vendor serializes is mapped;
// rightCamValidRoi is skipped for the same reason the vendor skips it. The
// VolumeCalib section (WorkDistance, MeasureRect, ...) lives in a separate
// camera memory bank and is not part of this document, so WorkDistance stays
// zero and is omitted when marshaling.
func NewVendorCalibJSON(s *StereoCalib) VendorCalibJSON {
	return VendorCalibJSON{CameraCalib: CameraCalibJSON{
		Q:                               s.Q[:],
		AveEpipolarError:                s.AveEpipolarError,
		ColorCamDistortion:              s.ColorDistortion[:],
		ColorCamImgHeight:               s.ColorImgH,
		ColorCamImgWidth:                s.ColorImgW,
		ColorCamIntrinsic:               s.ColorK[:],
		ColorCamRmsError:                s.ColorRmsError,
		LeftCamDistortion:               s.LeftDistortion[:],
		LeftCamImgHeight:                s.LeftImgH,
		LeftCamImgWidth:                 s.LeftImgW,
		LeftCamIntrinsic:                s.LeftK[:],
		LeftCamValidRoi:                 s.LeftValidRoi[:],
		LeftP:                           s.LeftP[:],
		LeftRectifyR:                    s.LeftRectifyR[:],
		LeftToColorCamRmsError:          s.LeftToColorRmsError,
		LeftToRightExtrinsic:            s.LeftToRightExtrinsic[:],
		RectLeftCamToColorCamExtrinsic:  s.RectLeftToColorExtrinsic[:],
		RectRightCamToColorCamExtrinsic: s.RectRightToColorExtrinsic[:],
		RightCamDistortion:              s.RightDistortion[:],
		RightCamImgHeight:               s.RightImgH,
		RightCamImgWidth:                s.RightImgW,
		RightCamIntrinsic:               s.RightK[:],
		RightP:                          s.RightP[:],
		RightRectifyR:                   s.RightRectifyR[:],
		RightToColorCamRmsError:         s.RightToColorRmsError,
		StereoRmsError:                  s.StereoRmsError,
	}}
}

// SaveVendorFile writes v to path as vendor calibration JSON (two-space
// indented), readable back with LoadVendorFile.
func SaveVendorFile(path string, v VendorCalibJSON) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("calib: encode: %w", err)
	}
	if err := os.WriteFile(path, append(b, '\n'), 0o644); err != nil {
		return fmt.Errorf("calib: write %s: %w", path, err)
	}
	return nil
}

// WriteCalibFile is the export path StereoCameraViewer's "save calibration"
// performs: download the camera's current stereo/color calibration from memory
// bank 0x20001 (CRC-verified) and write it to path in the vendor JSON format,
// e.g. "calib.json". It returns the calibration it wrote.
func WriteCalibFile(p RegisterPort, path string) (*StereoCalib, error) {
	s, err := ReadStereoCalib(p)
	if err != nil {
		return nil, err
	}
	if err := SaveVendorFile(path, NewVendorCalibJSON(s)); err != nil {
		return nil, err
	}
	return s, nil
}
