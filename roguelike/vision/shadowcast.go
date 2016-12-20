package vision

import "image"

type visionTransform func(image.Point) image.Point

type visionTransformer struct {
	source image.Point
}

func (vt visionTransformer) compose(ts ...visionTransform) visionTransform {
	return func(p image.Point) image.Point {
		for _, t := range ts {
			p = t(p)
		}
		return p.Add(vt.source)
	}
}

func swap(p image.Point) image.Point {
	return image.Point{X: p.Y, Y: p.X}
}

func invx(p image.Point) image.Point {
	return image.Point{X: -p.X, Y: p.Y}
}

func invy(p image.Point) image.Point {
	return image.Point{X: p.X, Y: -p.Y}
}

func ShadowCast(m Map, source image.Point) {
	vt := visionTransformer{source: source}
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose())
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose(swap))
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose(invx))
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose(swap, invx))
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose(invy))
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose(swap, invy))
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose(invy, invx))
	shadowCastOctant(1, 1.0, 0.0, m, vt.compose(swap, invy, invx))
}

func shadowCastOctant(col int, startSlope, endSlope float64, m Map, trans visionTransform) {
	wall := false
	for x := col; startSlope > endSlope && !wall; x++ {
		approxStart := int(startSlope*float64(x)) + 1
		for y := approxStart; y >= 0; y-- {
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
					shadowCastOctant(x+1, startSlope, loSlope, m, trans)
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
