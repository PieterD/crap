package vision

import (
	"image"
	"sync"
)

type ShadowCastMap interface {
	SetVisible(p image.Point)
	IsTransparent(p image.Point) bool
}

func ShadowCast(m ShadowCastMap, r Radius, source image.Point) {
	vt := visionTransformer{source: source}
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose())
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose(swap))
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose(invx))
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose(swap, invx))
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose(invy))
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose(swap, invy))
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose(invy, invx))
	shadowCastOctant(1, 1.0, 0.0, m, r, vt.compose(swap, invy, invx))
}

func ShadowCastPar(m ShadowCastMap, r Radius, source image.Point) {
	vt := visionTransformer{source: source}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go shadowCastPar(wg, m, r, vt.compose())
	go shadowCastPar(wg, m, r, vt.compose(invx))
	go shadowCastPar(wg, m, r, vt.compose(invy))
	go shadowCastPar(wg, m, r, vt.compose(invy, invx))
	wg.Wait()
	wg.Add(4)
	go shadowCastPar(wg, m, r, vt.compose(swap))
	go shadowCastPar(wg, m, r, vt.compose(swap, invx))
	go shadowCastPar(wg, m, r, vt.compose(swap, invy))
	go shadowCastPar(wg, m, r, vt.compose(swap, invy, invx))
	wg.Wait()
}

func shadowCastPar(wg *sync.WaitGroup, m ShadowCastMap, r Radius, trans visionTransform) {
	defer wg.Done()
	shadowCastOctant(1, 1.0, 0.0, m, r, trans)
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
