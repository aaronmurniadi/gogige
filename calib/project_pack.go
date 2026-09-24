package calib

import (
	"math"

	"github.com/aaronmurniadi/gogige/gvsp"
)

// ProjectPack projects a 3D pack's oriented bounding box onto the color image
// plane, returning the axis-aligned pixel bounding box (sub-pixel). The pack is
// treated as a box centred at (CenterX, CenterY, CenterZ) with half-extents
// LengthMm/WidthMm/HeightMm along the pack's local axes Orientation[0..2]
// (row 0 = length axis, row 1 = width axis, row 2 = height / normal). Corners
// behind the camera (z <= 0) or unprojectable yield NaN and are excluded; if
// no corner projects to the image, ok is false.
//
// This is the data consumers need to draw the green bounding-box overlay; the
// library only supplies coordinates, drawing is left to the caller.
func (c CamCalib) ProjectPack(pack gvsp.PackDet, imgW, imgH int) (gvsp.Box2D, bool) {
	cx, cy, cz := float64(pack.CenterX), float64(pack.CenterY), float64(pack.CenterZ)
	h := [3]float64{float64(pack.LengthMm) / 2, float64(pack.WidthMm) / 2, float64(pack.HeightMm) / 2}
	axis := [3][3]float64{}
	for r := 0; r < 3; r++ {
		norm := 0.0
		for k := 0; k < 3; k++ {
			axis[r][k] = float64(pack.Orientation[r][k])
			norm += axis[r][k] * axis[r][k]
		}
		if norm == 0 { // no orientation supplied: assume world-axis aligned box
			axis[r][r] = 1
		}
	}
	var minU, minV, maxU, maxV float64
	set := false
	for sx := -1.0; sx <= 1; sx += 2 {
		for sy := -1.0; sy <= 1; sy += 2 {
			for sz := -1.0; sz <= 1; sz += 2 {
				x := cx + sx*h[0]*axis[0][0] + sy*h[1]*axis[1][0] + sz*h[2]*axis[2][0]
				y := cy + sx*h[0]*axis[0][1] + sy*h[1]*axis[1][1] + sz*h[2]*axis[2][1]
				z := cz + sx*h[0]*axis[0][2] + sy*h[1]*axis[1][2] + sz*h[2]*axis[2][2]
				u, v := c.ProjectPoint3D(x, y, z, imgW, imgH)
				if math.IsNaN(u) || math.IsNaN(v) {
					continue
				}
				if !set {
					minU, minV, maxU, maxV = u, v, u, v
					set = true
					continue
				}
				if u < minU {
					minU = u
				}
				if u > maxU {
					maxU = u
				}
				if v < minV {
					minV = v
				}
				if v > maxV {
					maxV = v
				}
			}
		}
	}
	if !set {
		return gvsp.Box2D{}, false
	}
	return gvsp.Box2D{MinX: float32(minU), MinY: float32(minV), MaxX: float32(maxU), MaxY: float32(maxV)}, true
}

// OverlayBoxes returns one pixel bounding box per pack, in Packs order.
// Packs that cannot be projected produce an empty Box2D so indices line up
// with Sample.Packs / Sample.Overlay.
func OverlayBoxes(packs []gvsp.PackDet, imgW, imgH int, c CamCalib) []gvsp.Box2D {
	out := make([]gvsp.Box2D, len(packs))
	for i := range packs {
		if box, ok := c.ProjectPack(packs[i], imgW, imgH); ok {
			out[i] = box
		}
	}
	return out
}
