package calib

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"math"
	"strings"

	"github.com/aaronmurniadi/gogige/gvcp"
)

// Camera-memory bank protocol used by the vendor SDK
// (libStereoCamera.so CalibDataManager::Impl::readData): select a bank,
// read its length and CRC32, pull the bytes through the data window in
// <=512-byte chunks, then acknowledge.
const (
	regMemType = 0xE0000000 // W: memory-bank selector
	regMemSize = 0xE0000004 // R: byte length of bank contents
	regMemCRC  = 0xE0000008 // R: IEEE CRC32 of bank contents
	regMemAck  = 0xE000000C // W: finish/acknowledge transfer
	regMemData = 0xE0000100 // R: bulk data window

	memTypeCalibNames = 0x20000 // semicolon-separated calib type names
	memTypeStereo     = 0x20001 // MvSstereoCalibrateResult (1600 B)
)

// RegisterPort is the register-level transport needed to read camera
// memory banks. *gvcp.GVCP implements it.
type RegisterPort interface {
	ReadReg(addr uint32) (uint32, error)
	WriteReg(addr, value uint32) error
	ReadMem(addr uint32, n int) ([]byte, error)
}

var _ RegisterPort = (*gvcp.GVCP)(nil)

// ErrNoCalib is returned when the camera reports an empty calibration bank
// (device shipped without on-camera calibration).
var ErrNoCalib = errors.New("calib: camera calibration bank is empty")

// StereoCalib mirrors the vendor MvSstereoCalibrateResult POD (1600 bytes,
// little-endian) stored in camera memory bank 0x20001. Field offsets were
// recovered from CalibFile::struct2Json and cross-checked against a real
// DS5131MG30CE calibration export.
type StereoCalib struct {
	LeftImgW, LeftImgH        int
	LeftK                     [9]float64  // leftCamIntrinsic @0x008
	LeftDistortion            [12]float64 // leftCamDistortion @0x050
	RightImgW, RightImgH      int
	RightK                    [9]float64  // rightCamIntrinsic @0x0B8
	RightDistortion           [12]float64 // rightCamDistortion @0x100
	LeftToRightExtrinsic      [16]float64 // @0x160
	LeftRectifyR              [9]float64  // @0x1E0
	RightRectifyR             [9]float64  // @0x228
	LeftP                     [12]float64 // @0x270
	RightP                    [12]float64 // @0x2D0
	Q                         [16]float64 // disparity-to-depth matrix @0x330
	LeftValidRoi              [4]int32    // @0x3B0
	RightValidRoi             [4]int32    // @0x3C0 (never serialized by the vendor JSON)
	StereoRmsError            float64     // @0x3D0
	AveEpipolarError          float64     // @0x3D8
	ColorImgW, ColorImgH      int
	ColorK                    [9]float64  // colorCamIntrinsic @0x3E8
	ColorDistortion           [12]float64 // colorCamDistortion @0x430
	RectLeftToColorExtrinsic  [16]float64 // @0x490
	RectRightToColorExtrinsic [16]float64 // @0x510
	LeftToColorRmsError       float64     // @0x590
	RightToColorRmsError      float64     // @0x598
	ColorRmsError             float64     // @0x5A0
}

// stereoCalibMinLen is the end of the last field (0x5A8); cameras store the
// full 1600-byte (0x640) struct with trailing reserved space.
const stereoCalibMinLen = 0x5A8

// ReadStereoCalib downloads the stereo/color calibration from the camera's
// memory bank 0x20001 over GVCP, verifying the vendor CRC32.
func ReadStereoCalib(p RegisterPort) (*StereoCalib, error) {
	b, err := readMemBank(p, memTypeStereo)
	if err != nil {
		return nil, err
	}
	return parseStereoCalib(b)
}

// ReadCalibTypes returns the calibration type names the camera reports in
// bank 0x20000 (semicolon-separated, e.g. "calibration_pd").
func ReadCalibTypes(p RegisterPort) ([]string, error) {
	b, err := readMemBank(p, memTypeCalibNames)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, n := range strings.Split(string(b), ";") {
		if n = strings.TrimSpace(n); n != "" {
			names = append(names, n)
		}
	}
	return names, nil
}

// readMemBank replicates the vendor readData flow for one memory bank.
func readMemBank(p RegisterPort, memType uint32) ([]byte, error) {
	if err := p.WriteReg(regMemType, memType); err != nil {
		return nil, fmt.Errorf("calib: select bank %#x: %w", memType, err)
	}
	size, err := p.ReadReg(regMemSize)
	if err != nil {
		return nil, fmt.Errorf("calib: read size: %w", err)
	}
	if size == 0 {
		_ = p.WriteReg(regMemAck, 0)
		return nil, ErrNoCalib
	}
	want, err := p.ReadReg(regMemCRC)
	if err != nil {
		return nil, fmt.Errorf("calib: read crc: %w", err)
	}
	data, err := p.ReadMem(regMemData, int(size))
	if err != nil {
		return nil, fmt.Errorf("calib: read data: %w", err)
	}
	if crc32.ChecksumIEEE(data) != want {
		return nil, fmt.Errorf("calib: crc mismatch (got %#08x want %#08x)", crc32.ChecksumIEEE(data), want)
	}
	if err := p.WriteReg(regMemAck, 0); err != nil {
		return nil, fmt.Errorf("calib: ack: %w", err)
	}
	return data, nil
}

var le = binary.LittleEndian

func f64(b []byte, off int) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b[off:]))
}

func f64s(b []byte, off, n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		out[i] = f64(b, off+i*8)
	}
	return out
}

func parseStereoCalib(b []byte) (*StereoCalib, error) {
	if len(b) < stereoCalibMinLen {
		return nil, fmt.Errorf("calib: stereo blob short (%d < %d)", len(b), stereoCalibMinLen)
	}
	s := &StereoCalib{
		LeftImgW:             int(le.Uint32(b[0x000:])),
		LeftImgH:             int(le.Uint32(b[0x004:])),
		RightImgW:            int(le.Uint32(b[0x0B0:])),
		RightImgH:            int(le.Uint32(b[0x0B4:])),
		ColorImgW:            int(le.Uint32(b[0x3E0:])),
		ColorImgH:            int(le.Uint32(b[0x3E4:])),
		StereoRmsError:       f64(b, 0x3D0),
		AveEpipolarError:     f64(b, 0x3D8),
		LeftToColorRmsError:  f64(b, 0x590),
		RightToColorRmsError: f64(b, 0x598),
		ColorRmsError:        f64(b, 0x5A0),
	}
	copy(s.LeftK[:], f64s(b, 0x008, 9))
	copy(s.LeftDistortion[:], f64s(b, 0x050, 12))
	copy(s.RightK[:], f64s(b, 0x0B8, 9))
	copy(s.RightDistortion[:], f64s(b, 0x100, 12))
	copy(s.LeftToRightExtrinsic[:], f64s(b, 0x160, 16))
	copy(s.LeftRectifyR[:], f64s(b, 0x1E0, 9))
	copy(s.RightRectifyR[:], f64s(b, 0x228, 9))
	copy(s.LeftP[:], f64s(b, 0x270, 12))
	copy(s.RightP[:], f64s(b, 0x2D0, 12))
	copy(s.Q[:], f64s(b, 0x330, 16))
	for i := range s.LeftValidRoi {
		s.LeftValidRoi[i] = int32(le.Uint32(b[0x3B0+4*i:]))
		s.RightValidRoi[i] = int32(le.Uint32(b[0x3C0+4*i:]))
	}
	copy(s.ColorK[:], f64s(b, 0x3E8, 9))
	copy(s.ColorDistortion[:], f64s(b, 0x430, 12))
	copy(s.RectLeftToColorExtrinsic[:], f64s(b, 0x490, 16))
	copy(s.RectRightToColorExtrinsic[:], f64s(b, 0x510, 16))
	return s, nil
}

// Color returns the color-camera intrinsics.
func (s *StereoCalib) Color() (CamCalib, error) {
	return fromVendor(s.ColorK[:], s.ColorImgW, s.ColorImgH)
}

// Left returns the raw left (IR) camera intrinsics.
func (s *StereoCalib) Left() (CamCalib, error) {
	return fromVendor(s.LeftK[:], s.LeftImgW, s.LeftImgH)
}

// RectifiedLeft returns the rectified-left projection P as cam calib. Volume
// results are computed on rectified images, so pack centres are most often in
// this frame.
func (s *StereoCalib) RectifiedLeft() (CamCalib, error) {
	if s.LeftImgW <= 0 || s.LeftImgH <= 0 {
		return CamCalib{}, errors.New("calib: leftP missing")
	}
	return CamCalib{
		IntrinsicImgWidth:  s.LeftImgW,
		IntrinsicImgHeight: s.LeftImgH,
		K:                  [9]float64{s.LeftP[0], 0, s.LeftP[2], 0, s.LeftP[5], s.LeftP[6], 0, 0, 1},
	}, nil
}
