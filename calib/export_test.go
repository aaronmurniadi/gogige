package calib

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// goldenPath is a real DS5131MG30CE export in the exact document the vendor
// serializer produces (StereoCameraViewer "save calibration" / IPC4.94 tool).
const goldenPath = "testdata/ds5131_calib.json"

// goldenStereoCalib reads the vendor export and maps it back onto the
// in-memory camera structure, so NewVendorCalibJSON output can be compared
// against the document the camera itself was exported from.
func goldenStereoCalib(t *testing.T) (*StereoCalib, CameraCalibJSON) {
	t.Helper()
	b, err := os.ReadFile(goldenPath)
	require.NoError(t, err)
	var doc VendorCalibJSON
	require.NoError(t, json.Unmarshal(b, &doc))
	g := doc.CameraCalib

	arr9 := func(v []float64) [9]float64 {
		var a [9]float64
		copy(a[:], v)
		return a
	}
	arr12 := func(v []float64) [12]float64 {
		var a [12]float64
		copy(a[:], v)
		return a
	}
	arr16 := func(v []float64) [16]float64 {
		var a [16]float64
		copy(a[:], v)
		return a
	}

	return &StereoCalib{
		LeftImgW:                  g.LeftCamImgWidth,
		LeftImgH:                  g.LeftCamImgHeight,
		LeftK:                     arr9(g.LeftCamIntrinsic),
		LeftDistortion:            arr12(g.LeftCamDistortion),
		RightImgW:                 g.RightCamImgWidth,
		RightImgH:                 g.RightCamImgHeight,
		RightK:                    arr9(g.RightCamIntrinsic),
		RightDistortion:           arr12(g.RightCamDistortion),
		LeftToRightExtrinsic:      arr16(g.LeftToRightExtrinsic),
		LeftRectifyR:              arr9(g.LeftRectifyR),
		RightRectifyR:             arr9(g.RightRectifyR),
		LeftP:                     arr12(g.LeftP),
		RightP:                    arr12(g.RightP),
		Q:                         arr16(g.Q),
		LeftValidRoi:              [4]int32{g.LeftCamValidRoi[0], g.LeftCamValidRoi[1], g.LeftCamValidRoi[2], g.LeftCamValidRoi[3]},
		StereoRmsError:            g.StereoRmsError,
		AveEpipolarError:          g.AveEpipolarError,
		ColorImgW:                 g.ColorCamImgWidth,
		ColorImgH:                 g.ColorCamImgHeight,
		ColorK:                    arr9(g.ColorCamIntrinsic),
		ColorDistortion:           arr12(g.ColorCamDistortion),
		RectLeftToColorExtrinsic:  arr16(g.RectLeftCamToColorCamExtrinsic),
		RectRightToColorExtrinsic: arr16(g.RectRightCamToColorCamExtrinsic),
		LeftToColorRmsError:       g.LeftToColorCamRmsError,
		RightToColorRmsError:      g.RightToColorCamRmsError,
		ColorRmsError:             g.ColorCamRmsError,
	}, g
}

// TestNewVendorCalibJSONMatchesVendorExport proves the export writes exactly
// the document the vendor tools write: every key of the real export survives
// the conversion, nothing extra appears, and keys stay in vendor (ASCII-sorted)
// order.
func TestNewVendorCalibJSONMatchesVendorExport(t *testing.T) {
	s, want := goldenStereoCalib(t)

	out, err := json.Marshal(NewVendorCalibJSON(s))
	require.NoError(t, err)

	var got VendorCalibJSON
	require.NoError(t, json.Unmarshal(out, &got))
	require.Equal(t, want, got.CameraCalib, "vendor CameraCalib must round-trip unchanged")

	text := string(out)
	idx := func(k string) int { return bytes.Index(out, []byte(`"`+k+`":`)) }
	assert.NotContains(t, text, "rightCamValidRoi", "vendor never serializes the right ROI")
	assert.NotContains(t, text, "WorkDistance", "WorkDistance lives in VolumeCalib, not the camera bank")
	assert.Less(t, idx("Q"), idx("aveEpipolarError"), "Q sorts first in the vendor document")
	assert.Less(t, idx("leftCamValidRoi"), idx("leftP"))
	assert.Less(t, idx("leftToRightExtrinsic"), idx("rectLeftCamToColorCamExtrinsic"))
	assert.Less(t, idx("rightP"), idx("stereoRmsError"))
}

// TestSaveVendorFileRoundTrip loads a vendor export, rewrites it with
// SaveVendorFile, and checks the document is preserved.
func TestSaveVendorFileRoundTrip(t *testing.T) {
	v, err := LoadVendorFile(goldenPath)
	require.NoError(t, err)

	path := filepath.Join(t.TempDir(), "calib.json")
	require.NoError(t, SaveVendorFile(path, v))

	back, err := LoadVendorFile(path)
	require.NoError(t, err)
	require.Equal(t, v, back)

	b, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Contains(t, string(b), "\n  \"CameraCalib\": {", "output is indented for humans")
	assert.True(t, bytes.HasSuffix(b, []byte("\n")))
	assert.True(t, json.Valid(b))
}

func TestSaveVendorFileError(t *testing.T) {
	err := SaveVendorFile(filepath.Join(t.TempDir(), "missing", "calib.json"), VendorCalibJSON{})
	assert.ErrorContains(t, err, "calib: write")
}

// TestWriteCalibFile covers the whole path the goal asks for: pull the current
// calibration out of the camera and land it in a calib.json file.
func TestWriteCalibFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "calib.json")
	p := &fakePort{bank: buildStereoBlob(t)}

	s, err := WriteCalibFile(p, path)
	require.NoError(t, err)
	assert.Equal(t, 1280, s.LeftImgW)
	assert.Equal(t, uint32(memTypeStereo), p.writes[regMemType], "stereo bank selected")

	v, err := LoadVendorFile(path)
	require.NoError(t, err)
	assert.Equal(t, 1280, v.CameraCalib.LeftCamImgWidth)
	assert.Equal(t, 960, v.CameraCalib.LeftCamImgHeight)
	assert.InDelta(t, 1063.397805476623, v.CameraCalib.LeftCamIntrinsic[FX], 1e-9)
	assert.InDelta(t, 996.45574642371707, v.CameraCalib.ColorCamIntrinsic[FX], 1e-9)
	assert.InDelta(t, -627.1973266601562, v.CameraCalib.Q[3], 1e-9)
	assert.Equal(t, []int32{0, 0, 1280, 960}, v.CameraCalib.LeftCamValidRoi)
	assert.NotContains(t, mustJSON(t, v), "rightCamValidRoi")
	assert.Zero(t, v.CameraCalib.WorkDistance, "volume bank is not part of the camera calibration")

	// the written file feeds the existing consumers directly
	color, err := v.Color()
	require.NoError(t, err)
	assert.Equal(t, 1280, color.IntrinsicImgWidth)
	rect, err := v.RectifiedLeft()
	require.NoError(t, err)
	assert.InDelta(t, 1073.4602683226065, rect.K[FX], 1e-9)
}

func TestWriteCalibFileEmptyBank(t *testing.T) {
	path := filepath.Join(t.TempDir(), "calib.json")
	_, err := WriteCalibFile(&fakePort{}, path)
	assert.ErrorIs(t, err, ErrNoCalib)
	_, statErr := os.Stat(path)
	assert.True(t, os.IsNotExist(statErr), "no file may be written when the bank is empty")
}

func mustJSON(t *testing.T, v VendorCalibJSON) string {
	t.Helper()
	b, err := json.Marshal(v)
	require.NoError(t, err)
	return string(b)
}

// TestVendorKeyOrderPinned keeps the exported key order locked to the vendor
// document (a std::map, so ASCII-sorted).
func TestVendorKeyOrderPinned(t *testing.T) {
	b, err := os.ReadFile(goldenPath)
	require.NoError(t, err)

	m := regexp.MustCompile(`"CameraCalib": \{"([^"]+)"`).FindSubmatch(b)
	require.NotNil(t, m, "golden CameraCalib object not found")
	assert.Equal(t, "Q", string(m[1]))
}
