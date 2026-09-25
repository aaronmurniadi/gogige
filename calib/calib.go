// Package calib converts between camera-frame 3D coordinates (millimetres)
// and image pixel coordinates using pinhole intrinsics.
//
// The math mirrors the vendor SDK's stereoConvetPoint3dToDepth /
// stereoConvetDepthToPoint3d (libStereoCamera.so, Huaray/Dahua 5000 series):
// pure pinhole projection without distortion, rescaled from the calibration
// resolution to the actual output resolution by the width ratio only:
//
//	scale = intrinsicImgWidth / imgW
//	u     = (x*K[fx]/z + K[cx]) / scale
//	v     = (y*K[fy]/z + K[cy]) / scale
//
// Points are expected in the target camera's own frame: origin on the optical
// axis, X right, Y down, Z forward along the optical axis (millimetres).
package calib

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
)

// Indices into the row-major 3x3 intrinsic matrix
// { fx, 0, cx, 0, fy, cy, 0, 0, 1 }.
const (
	FX = 0
	FY = 4
	CX = 2
	CY = 5
)

// CamCalib holds the pinhole intrinsics of one image type.
type CamCalib struct {
	// IntrinsicImgWidth/Height is the resolution the intrinsics were
	// calibrated at. Projections onto other resolutions are rescaled by the
	// width ratio, matching the vendor SDK.
	IntrinsicImgWidth  int
	IntrinsicImgHeight int

	// K is the row-major 3x3 intrinsic matrix { fx, 0, cx, 0, fy, cy, 0, 0, 1 }.
	K [9]float64
}

// ProjectPoint3D projects a camera-frame point (mm) to pixel coordinates on an
// image of the given size. It returns NaN pixels for z <= 0 or imgW <= 0.
func (c CamCalib) ProjectPoint3D(x, y, z float64, imgW, imgH int) (u, v float64) {
	if z <= 0 || imgW <= 0 {
		return math.NaN(), math.NaN()
	}
	scale := float64(c.IntrinsicImgWidth) / float64(imgW)
	return (x*c.K[FX]/z + c.K[CX]) / scale, (y*c.K[FY]/z + c.K[CY]) / scale
}

// DeprojectPixel back-projects a pixel plus depth (mm) to a camera-frame
// point. It is the exact inverse of ProjectPoint3D.
func (c CamCalib) DeprojectPixel(u, v, z float64, imgW, imgH int) (x, y float64) {
	if z <= 0 || imgW <= 0 {
		return math.NaN(), math.NaN()
	}
	scale := float64(c.IntrinsicImgWidth) / float64(imgW)
	return (u*scale - c.K[CX]) / c.K[FX] * z, (v*scale - c.K[CY]) / c.K[FY] * z
}

// CameraCalibJSON is the CameraCalib object of a vendor calibration export:
// every field CalibFile::struct2Json serializes from MvSstereoCalibrateResult
// (libStereoCamera.so), which StereoCameraViewer saves as CalibData.json and
// vendor tools as "IPC4.94 Camera Calibration.json".
//
// Fields are declared in ASCII order because the vendor serializer stores keys
// in a std::map and therefore writes them sorted ("Q" first, then the lower-case
// keys); marshaling this struct reproduces that layout. rightCamValidRoi is
// deliberately absent — the vendor never serializes it.
type CameraCalibJSON struct {
	Q                               []float64 `json:"Q"`
	WorkDistance                    float64   `json:"WorkDistance,omitempty"`
	AveEpipolarError                float64   `json:"aveEpipolarError"`
	ColorCamDistortion              []float64 `json:"colorCamDistortion"`
	ColorCamImgHeight               int       `json:"colorCamImgHeight"`
	ColorCamImgWidth                int       `json:"colorCamImgWidth"`
	ColorCamIntrinsic               []float64 `json:"colorCamIntrinsic"`
	ColorCamRmsError                float64   `json:"colorCamRmsError"`
	LeftCamDistortion               []float64 `json:"leftCamDistortion"`
	LeftCamImgHeight                int       `json:"leftCamImgHeight"`
	LeftCamImgWidth                 int       `json:"leftCamImgWidth"`
	LeftCamIntrinsic                []float64 `json:"leftCamIntrinsic"`
	LeftCamValidRoi                 []int32   `json:"leftCamValidRoi"`
	LeftP                           []float64 `json:"leftP"`
	LeftRectifyR                    []float64 `json:"leftRectifyR"`
	LeftToColorCamRmsError          float64   `json:"leftToColorCamRmsError"`
	LeftToRightExtrinsic            []float64 `json:"leftToRightExtrinsic"`
	RectLeftCamToColorCamExtrinsic  []float64 `json:"rectLeftCamToColorCamExtrinsic"`
	RectRightCamToColorCamExtrinsic []float64 `json:"rectRightCamToColorCamExtrinsic"`
	RightCamDistortion              []float64 `json:"rightCamDistortion"`
	RightCamImgHeight               int       `json:"rightCamImgHeight"`
	RightCamImgWidth                int       `json:"rightCamImgWidth"`
	RightCamIntrinsic               []float64 `json:"rightCamIntrinsic"`
	RightP                          []float64 `json:"rightP"`
	RightRectifyR                   []float64 `json:"rightRectifyR"`
	RightToColorCamRmsError         float64   `json:"rightToColorCamRmsError"`
	StereoRmsError                  float64   `json:"stereoRmsError"`
}

// VendorCalibJSON mirrors the calibration export written by the vendor tools
// ("IPC4.94 Camera Calibration.json" style): a top-level CameraCalib object
// with per-camera intrinsics and resolutions.
type VendorCalibJSON struct {
	CameraCalib CameraCalibJSON `json:"CameraCalib"`
}

// Color returns the color-camera intrinsics from a vendor calibration export.
func (v VendorCalibJSON) Color() (CamCalib, error) {
	return fromVendor(v.CameraCalib.ColorCamIntrinsic, v.CameraCalib.ColorCamImgWidth, v.CameraCalib.ColorCamImgHeight)
}

// Left returns the raw left (IR) camera intrinsics from a vendor export.
func (v VendorCalibJSON) Left() (CamCalib, error) {
	return fromVendor(v.CameraCalib.LeftCamIntrinsic, v.CameraCalib.LeftCamImgWidth, v.CameraCalib.LeftCamImgHeight)
}

// RectifiedLeft returns the rectified-left projection P as cam calib. Volume
// results are computed on rectified images, so pack centres are most often in
// this frame; P has the form { fx', 0, cx', 0, fy', cy', ... }.
func (v VendorCalibJSON) RectifiedLeft() (CamCalib, error) {
	p := v.CameraCalib.LeftP
	if len(p) < 7 || v.CameraCalib.LeftCamImgWidth <= 0 {
		return CamCalib{}, errors.New("calib: leftP missing")
	}
	return CamCalib{
		IntrinsicImgWidth:  v.CameraCalib.LeftCamImgWidth,
		IntrinsicImgHeight: v.CameraCalib.LeftCamImgHeight,
		K:                  [9]float64{p[0], 0, p[2], 0, p[5], p[6], 0, 0, 1},
	}, nil
}

func fromVendor(k []float64, w, h int) (CamCalib, error) {
	if len(k) < 6 || w <= 0 || h <= 0 {
		return CamCalib{}, errors.New("calib: incomplete vendor calib")
	}
	return CamCalib{
		IntrinsicImgWidth:  w,
		IntrinsicImgHeight: h,
		K:                  [9]float64{k[0], k[1], k[2], k[3], k[4], k[5], 0, 0, 1},
	}, nil
}

// LoadVendorFile reads a vendor calibration export from path.
func LoadVendorFile(path string) (VendorCalibJSON, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return VendorCalibJSON{}, fmt.Errorf("calib: read %s: %w", path, err)
	}
	var v VendorCalibJSON
	if err := json.Unmarshal(b, &v); err != nil {
		return VendorCalibJSON{}, fmt.Errorf("calib: parse %s: %w", path, err)
	}
	return v, nil
}
