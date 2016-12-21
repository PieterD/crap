package vision

import "image"

type ShadowCastMap interface {
	SetVisible(p image.Point)
	IsTransparent(p image.Point) bool
}

func ShadowCast(m ShadowCastMap, r Radius, source image.Point) {
	vt := visionTransformer{source: source}
	for i := range quads {
		shadowCastOctant(1, 1.0, 0.0, m, r, vt.quad(i).norm)
		shadowCastOctant(1, 1.0, 0.0, m, r, vt.quad(i).swap)
	}
}

func shadowCastOctant(col int, startSlope, endSlope float64, m ShadowCastMap, r Radius, trans visionTransform) {
	wall := false
	for x := col; startSlope > endSlope && !wall && r.In(x, 0); x++ {
		approxStart := int(startSlope*float64(x)) + 1
		for y := approxStart; y >= 0; y-- {
			if !r.In(x, y) {
				continue
			}
			hiSlope := (float64(y) - 0.5) / (float64(x) - 0.5)
			if hiSlope > startSlope {
				continue
			}
			loSlope := (float64(y) + 0.5) / (float64(x) - 0.5)
			if loSlope < endSlope {
				break
			}
			pos := trans(image.Point{X: x, Y: y})
			m.SetVisible(pos)
			if !m.IsTransparent(pos) {
				if !wall {
					wall = true
					shadowCastOctant(x+1, startSlope, loSlope, m, r, trans)
				}
			} else {
				if wall {
					wall = false
					startSlope = (float64(y) + 0.5) / (float64(x) + 0.5)
				}
			}
		}
	}
}
