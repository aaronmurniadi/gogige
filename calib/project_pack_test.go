package calib

import (
	"testing"

	"github.com/aaronmurniadi/gogige/gvsp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// identity orientation: box axes align with world axes, so the projected
// bounding box is simply the axis-aligned extent of (±L/2, ±W/2, ±H/2) around
// the projected centre.
func TestProjectPackIdentity(t *testing.T) {
	c := testCalib() // 1280x1024, fx=fy=1000, cx=640, cy=512
	pack := gvsp.PackDet{
		CenterX: 0, CenterY: 0, CenterZ: 1000,
		LengthMm: 100, WidthMm: 80, HeightMm: 60,
	}
	box, ok := c.ProjectPack(pack, 1280, 1024)
	require.True(t, ok)
	// centre projects to principal point (640,512); half-extents at z=1000 are
	// (50,40,30) -> +-50 x, +-40 y.
	assert.InDelta(t, 640-50, box.MinX, 2)
	assert.InDelta(t, 640+50, box.MaxX, 2)
	assert.InDelta(t, 512-40, box.MinY, 2)
	assert.InDelta(t, 512+40, box.MaxY, 2)
	assert.False(t, box.Empty())
}

func TestProjectPackPerspective(t *testing.T) {
	c := testCalib()
	pack := gvsp.PackDet{
		CenterX: 0, CenterY: 0, CenterZ: 2000,
		LengthMm: 100, WidthMm: 100, HeightMm: 100,
	}
	box, ok := c.ProjectPack(pack, 1280, 1024)
	require.True(t, ok)
	// At double the depth, projected half-extent halves (~±25 from centre),
	// plus the z-extent (±30) widens the range slightly on the near edges.
	assert.InDelta(t, 640-26, box.MinX, 2)
	assert.InDelta(t, 640+26, box.MaxX, 2)
	assert.InDelta(t, 512-26, box.MinY, 2)
	assert.InDelta(t, 512+26, box.MaxY, 2)
}

func TestProjectPackBehindCamera(t *testing.T) {
	c := testCalib()
	pack := gvsp.PackDet{
		CenterX: 0, CenterY: 0, CenterZ: -100,
		LengthMm: 100, WidthMm: 80, HeightMm: 60,
	}
	_, ok := c.ProjectPack(pack, 1280, 1024)
	assert.False(t, ok)
}

func TestOverlayBoxesIndexAlignment(t *testing.T) {
	c := testCalib()
	packs := []gvsp.PackDet{
		{CenterX: 0, CenterY: 0, CenterZ: 1000, LengthMm: 100, WidthMm: 100, HeightMm: 100},
		{CenterX: 0, CenterY: 0, CenterZ: -100, LengthMm: 100, WidthMm: 100, HeightMm: 100},
	}
	boxes := OverlayBoxes(packs, 1280, 1024, c)
	require.Len(t, boxes, 2)
	assert.False(t, boxes[0].Empty())
	assert.True(t, boxes[1].Empty()) // behind camera
}
