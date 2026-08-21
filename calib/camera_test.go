package calib

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakePort serves one memory bank the way the camera firmware does.
type fakePort struct {
	bank    []byte
	writes  map[uint32]uint32
	failCRC bool
}

func (f *fakePort) ReadReg(addr uint32) (uint32, error) {
	switch addr {
	case regMemSize:
		return uint32(len(f.bank)), nil
	case regMemCRC:
		if f.failCRC {
			return 0xDEADBEEF, nil
		}
		return crc32.ChecksumIEEE(f.bank), nil
	}
	return 0, errors.New("unexpected read")
}

func (f *fakePort) WriteReg(addr, value uint32) error {
	if f.writes == nil {
		f.writes = map[uint32]uint32{}
	}
	f.writes[addr] = value
	return nil
}

func (f *fakePort) ReadMem(addr uint32, n int) ([]byte, error) {
	off := int(addr) - regMemData
	return f.bank[off : off+n], nil
}

// buildStereoBlob encodes a MvSstereoCalibrateResult using values from the
// real DS5131MG30CE calibration export.
func buildStereoBlob(t *testing.T) []byte {
	t.Helper()
	b := make([]byte, 0x640)
	putI := func(off uint32, v int32) { binary.LittleEndian.PutUint32(b[off:], uint32(v)) }
	putF := func(off uint32, v float64) { binary.LittleEndian.PutUint64(b[off:], math.Float64bits(v)) }
	putArr := func(off uint32, vs ...float64) {
		for i, v := range vs {
			putF(off+uint32(8*i), v)
		}
	}

	putI(0x000, 1280) // leftCamImgWidth
	putI(0x004, 960)  // leftCamImgHeight
	putArr(0x008, 1063.397805476623, 0, 638.60711689473806, 0, 1063.2707377290042, 482.97497012149216, 0, 0, 1)
	putArr(0x050, -0.11460978866207634, 0.11781795232413562, 0.0006051239508532486, -0.0003967324)
	putI(0x0B0, 1280)
	putI(0x0B4, 960)
	putArr(0x0B8, 1065.568799802541, 0, 635.9802119194284, 0, 1065.251323993112, 458.8036134395502, 0, 0, 1)
	putArr(0x160, 0.9999817124927021, -0.0057118828383232595, -0.0019872278692052584, -80.1823356)
	putArr(0x270, 1073.4602683226065, 0, 627.19732666015625, 0, 0, 1073.4602683226065, 471.2895393371582, 0, 0, 0, 1, 0)
	putArr(0x330, 1, 0, 0, -627.1973266601562, 0, 1, 0, -471.2895393371582, 0, 0, 0, 1073.4602683226065)
	putI(0x3B0, 0) // leftCamValidRoi x
	putI(0x3B4, 0)
	putI(0x3B8, 1280)
	putI(0x3BC, 960)
	putI(0x3C0, 2) // rightCamValidRoi (never serialized by vendor JSON)
	putI(0x3C4, 3)
	putI(0x3C8, 4)
	putI(0x3CC, 5)
	putF(0x3D0, 0.10674393575477786) // stereoRmsError
	putF(0x3D8, 0.11541362242265181) // aveEpipolarError
	putI(0x3E0, 1280)                // colorCamImgWidth
	putI(0x3E4, 1024)                // colorCamImgHeight
	putArr(0x3E8, 996.45574642371707, 0, 627.18293542287472, 0, 996.15876619015796, 523.28861070003722, 0, 0, 1)
	putArr(0x590, 0.0835428386926651) // leftToColorCamRmsError
	putArr(0x598, 0.08280157297849655)
	putArr(0x5A0, 0.0936705445956812) // colorCamRmsError
	return b
}

func TestReadStereoCalib(t *testing.T) {
	p := &fakePort{bank: buildStereoBlob(t)}
	s, err := ReadStereoCalib(p)
	require.NoError(t, err)

	assert.Equal(t, map[uint32]uint32{regMemType: memTypeStereo, regMemAck: 0}, p.writes)

	assert.Equal(t, 1280, s.LeftImgW)
	assert.Equal(t, 960, s.LeftImgH)
	assert.InDelta(t, 1063.397805476623, s.LeftK[FX], 1e-9)
	assert.InDelta(t, 482.97497012149216, s.LeftK[CY], 1e-9)
	assert.InDelta(t, -80.1823356, s.LeftToRightExtrinsic[3], 1e-9)
	assert.InDelta(t, 471.2895393371582, s.LeftP[6], 1e-9)
	assert.InDelta(t, -627.1973266601562, s.Q[3], 1e-9)
	assert.Equal(t, [4]int32{0, 0, 1280, 960}, s.LeftValidRoi)
	assert.Equal(t, [4]int32{2, 3, 4, 5}, s.RightValidRoi)
	assert.InDelta(t, 0.10674393575477786, s.StereoRmsError, 1e-12)
	assert.InDelta(t, 0.11541362242265181, s.AveEpipolarError, 1e-12)

	color, err := s.Color()
	require.NoError(t, err)
	assert.Equal(t, 1280, color.IntrinsicImgWidth)
	assert.Equal(t, 1024, color.IntrinsicImgHeight)
	assert.InDelta(t, 996.45574642371707, color.K[FX], 1e-9)
	assert.InDelta(t, 523.28861070003722, color.K[CY], 1e-9)

	left, err := s.Left()
	require.NoError(t, err)
	assert.InDelta(t, 1063.2707377290042, left.K[FY], 1e-9)

	rect, err := s.RectifiedLeft()
	require.NoError(t, err)
	assert.InDelta(t, 1073.4602683226065, rect.K[FX], 1e-9)
	assert.InDelta(t, 627.19732666015625, rect.K[CX], 1e-9)
}

func TestReadStereoCalibErrors(t *testing.T) {
	p := &fakePort{}
	_, err := ReadStereoCalib(p)
	assert.ErrorIs(t, err, ErrNoCalib)
	assert.Equal(t, uint32(0), p.writes[regMemAck], "empty bank must still ack")

	p = &fakePort{bank: buildStereoBlob(t), failCRC: true}
	_, err = ReadStereoCalib(p)
	assert.ErrorContains(t, err, "crc mismatch")

	short := make([]byte, 0x100)
	p = &fakePort{bank: short}
	_, err = ReadStereoCalib(p)
	assert.ErrorContains(t, err, "short")
}

func TestReadCalibTypes(t *testing.T) {
	p := &fakePort{bank: []byte("calibration_pd;volume;")}
	names, err := ReadCalibTypes(p)
	require.NoError(t, err)
	assert.Equal(t, []string{"calibration_pd", "volume"}, names)
}
