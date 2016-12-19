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
		return p
	}
}

func (vt visionTransformer) identity(p image.Point) image.Point {
	return p.Add(vt.source)
}

func (vt visionTransformer) swap(p image.Point) image.Point {
	return image.Point{X: p.Y, Y: p.X}
}

func (vt visionTransformer) invx(p image.Point) image.Point {
	return image.Point{X: -p.X, Y: p.Y}
}

func (vt visionTransformer) invy(p image.Point) image.Point {
	return image.Point{X: p.X, Y: -p.Y}
}

type ShadowCast struct {
	m Map
}

func NewShadowCast(m Map) *ShadowCast {
	return &ShadowCast{
		m: m,
	}
}

func (v *ShadowCast) Vision(source image.Point) {
	vt := visionTransformer{source: source}
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.identity))
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.swap, vt.identity))
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.invx, vt.identity))
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.swap, vt.invx, vt.identity))
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.invy, vt.identity))
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.swap, vt.invy, vt.identity))
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.invy, vt.invx, vt.identity))
	v.visionOctant(1, 1.0, 0.0, vt.compose(vt.swap, vt.invy, vt.invx, vt.identity))
}

func (v *ShadowCast) visionOctant(col int, startSlope, endSlope float64, trans visionTransform) {
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
			v.m.MakeVisible(pos)
			if !v.m.IsTransparent(pos) {
				if !wall {
					wall = true
					v.visionOctant(x+1, startSlope, loSlope, trans)
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
