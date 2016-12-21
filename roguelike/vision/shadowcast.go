package vision

import "image"

type ShadowCastMap interface {
	SetVisible(p image.Point)
	IsTransparent(p image.Point) bool
}

func ShadowCast(m ShadowCastMap, r Radius, source image.Point) {
	vt := visionTransformer{source: source}
	for i := range quads {
		shadowCastOctant(1, one(), zero(), m, r, vt.quad(i).norm)
		shadowCastOctant(1, one(), zero(), m, r, vt.quad(i).swap)
	}
}

func shadowCastOctant(col int, startSlope, endSlope fraction, m ShadowCastMap, r Radius, trans visionTransform) {
	wall := false
	for x := col; endSlope.less(startSlope) && !wall && r.In(x, 0); x++ {
		approxStart := startSlope.mul(x).whole() + 1
		for y := approxStart; y >= 0; y-- {
			if !r.In(x, y) {
				continue
			}
			hiSlope := newFraction(y*2-1, x*2-1)
			if startSlope.less(hiSlope) {
				continue
			}
			loSlope := newFraction(y*2+1, x*2-1)
			if loSlope.less(endSlope) {
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
					startSlope = newFraction(y*2+1, x*2+1)
				}
			}
		}
	}
}
