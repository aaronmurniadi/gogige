package calib

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testCalib() CamCalib {
	return CamCalib{
		IntrinsicImgWidth:  1280,
		IntrinsicImgHeight: 1024,
		K:                  [9]float64{1000, 0, 640, 0, 1000, 512, 0, 0, 1},
	}
}

func TestProjectPoint3DHandComputed(t *testing.T) {
	c := testCalib()

	// Optical axis centre maps to the principal point.
	u, v := c.ProjectPoint3D(0, 0, 1000, 1280, 1024)
	assert.InDelta(t, 640, u, 1e-9)
	assert.InDelta(t, 512, v, 1e-9)

	// 1 px per mm at z = 1000 with fx = 1000.
	u, v = c.ProjectPoint3D(-250, 125, 1000, 1280, 1024)
	assert.InDelta(t, 640-250, u, 1e-9)
	assert.InDelta(t, 512+125, v, 1e-9)

	// Doubling depth halves the pixel offset (perspective).
	u, _ = c.ProjectPoint3D(200, 0, 2000, 1280, 1024)
	assert.InDelta(t, 640+100, u, 1e-9)
}

func TestProjectRescalesToOutputResolution(t *testing.T) {
	c := testCalib() // calibrated at 1280 wide

	// Half-size output: same point must land at half the offset from the
	// scaled principal point.
	u, v := c.ProjectPoint3D(-250, 125, 1000, 640, 512)
	assert.InDelta(t, 320-125, u, 1e-9)
	assert.InDelta(t, 256+62.5, v, 1e-9)
}

func TestDeprojectIsInverseOfProject(t *testing.T) {
	c := testCalib()

	for _, tc := range []struct{ x, y, z float64 }{
		{-57.74, 62.09, 1376.18},
		{300, -150, 800},
		{0, 0, 1300},
	} {
		u, v := c.ProjectPoint3D(tc.x, tc.y, tc.z, 1280, 1024)
		x, y := c.DeprojectPixel(u, v, tc.z, 1280, 1024)
		assert.InDelta(t, tc.x, x, 1e-9, "x round-trip")
		assert.InDelta(t, tc.y, y, 1e-9, "y round-trip")
	}
}

func TestInvalidInputsYieldNaN(t *testing.T) {
	c := testCalib()

	u, v := c.ProjectPoint3D(10, 10, 0, 1280, 1024)
	assert.True(t, math.IsNaN(u) && math.IsNaN(v))

	u, v = c.ProjectPoint3D(10, 10, -5, 1280, 1024)
	assert.True(t, math.IsNaN(u) && math.IsNaN(v))

	x, y := c.DeprojectPixel(100, 100, 1000, 0, 1024)
	assert.True(t, math.IsNaN(x) && math.IsNaN(y))
}

// goldenVendorJSON is the calibration export of a DS5131MG30CE
// ("IPC4.94 Camera Calibration.json" format).
const goldenVendorJSON = `{"CameraCalib": {"Q": [1], "aveEpipolarError": 0.115,
"colorCamDistortion": [-0.115, 0.138, 0.0007, 0.0012, -0.031, 0, 0, 0, 0, 0, 0, 0],
"colorCamImgHeight": 1024, "colorCamImgWidth": 1280,
"colorCamIntrinsic": [996.45574642371707, 0, 627.18293542287472, 0, 996.15876619015796, 523.28861070003722, 0, 0, 1],
"leftCamImgHeight": 960, "leftCamImgWidth": 1280,
"leftCamIntrinsic": [1063.397805476623, 0, 638.60711689473806, 0, 1063.2707377290042, 482.97497012149216, 0, 0, 1],
"leftP": [1073.4602683226065, 0, 627.19732666015625, 0, 0, 1073.4602683226065, 471.2895393371582, 0, 0, 0, 1, 0],
"WorkDistance": 1300}}`

func TestLoadVendorFileAndAccessors(t *testing.T) {
	path := filepath.Join(t.TempDir(), "calib.json")
	require.NoError(t, os.WriteFile(path, []byte(goldenVendorJSON), 0o644))

	v, err := LoadVendorFile(path)
	require.NoError(t, err)
	assert.InDelta(t, 1300, v.CameraCalib.WorkDistance, 1e-9)

	color, err := v.Color()
	require.NoError(t, err)
	assert.Equal(t, 1280, color.IntrinsicImgWidth)
	assert.Equal(t, 1024, color.IntrinsicImgHeight)
	assert.InDelta(t, 996.45574642371707, color.K[FX], 1e-9)
	assert.InDelta(t, 996.15876619015796, color.K[FY], 1e-9)
	assert.InDelta(t, 627.18293542287472, color.K[CX], 1e-9)
	assert.InDelta(t, 523.28861070003722, color.K[CY], 1e-9)

	left, err := v.Left()
	require.NoError(t, err)
	assert.Equal(t, 960, left.IntrinsicImgHeight)
	assert.InDelta(t, 1063.397805476623, left.K[FX], 1e-9)

	rect, err := v.RectifiedLeft()
	require.NoError(t, err)
	assert.InDelta(t, 1073.4602683226065, rect.K[FX], 1e-9)
	assert.InDelta(t, 1073.4602683226065, rect.K[FY], 1e-9)
	assert.InDelta(t, 627.19732666015625, rect.K[CX], 1e-9)
	assert.InDelta(t, 471.2895393371582, rect.K[CY], 1e-9)
}

func TestLoadVendorFileErrors(t *testing.T) {
	_, err := LoadVendorFile(filepath.Join(t.TempDir(), "missing.json"))
	assert.Error(t, err)

	path := filepath.Join(t.TempDir(), "bad.json")
	require.NoError(t, os.WriteFile(path, []byte("not json"), 0o644))
	_, err = LoadVendorFile(path)
	assert.Error(t, err)

	var empty VendorCalibJSON
	_, err = empty.Color()
	assert.Error(t, err)
	_, err = empty.RectifiedLeft()
	assert.Error(t, err)
}

// TestDS5131SampleProjection pins the projection of a real BSCF sample
// (pack centre measured live at (-38.39, 47.69, 1387.39) mm) using the
// exported color intrinsics, so frame-convention regressions are visible.
func TestDS5131SampleProjection(t *testing.T) {
	var vendor VendorCalibJSON
	require.NoError(t, json.Unmarshal([]byte(goldenVendorJSON), &vendor))
	color, err := vendor.Color()
	require.NoError(t, err)

	u, vv := color.ProjectPoint3D(-38.39, 47.69, 1387.39, 1280, 1024)
	assert.InDelta(t, 599.6103, u, 1e-3)  // ~40 px left of centre
	assert.InDelta(t, 557.5305, vv, 1e-3) // ~45 px below centre
}
