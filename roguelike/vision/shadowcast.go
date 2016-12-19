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
	v.visionOctant(1, 1, 1, 0.0, vt.compose(vt.identity))
	/*
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.identity))
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.swap, vt.identity))
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.invx, vt.identity))
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.swap, vt.invx, vt.identity))
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.invy, vt.identity))
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.swap, vt.invy, vt.identity))
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.invy, vt.invx, vt.identity))
		atlas.visionOctant(1, 0.0, 1.0, vt.compose(vt.swap, vt.invy, vt.invx, vt.identity))
	*/
}

func (v *ShadowCast) visionOctant(col, row int, startSlope, endSlope float64, trans visionTransform) {
	//fmt.Printf("vision col=%d row=%d startSlope=%f endSlope=%f\n", col, row, startSlope, endSlope)
	//defer fmt.Printf("return\n")
	wall := false
	for x := col; startSlope > endSlope && !wall; x++ {
		//fmt.Printf(" x=%d\n", x)
		for y := x; y >= 0; y-- {
			//fmt.Printf("  startSlope=%f endSlope=%f\n", startSlope, endSlope)
			hiSlope := (float64(y) - 0.5) / (float64(x) + 0.5)
			loSlope := (float64(y) + 0.5) / (float64(x) - 0.5)
			//fmt.Printf("  y=%d lo=%f hi=%f\n", y, loSlope, hiSlope)
			if hiSlope > startSlope {
				//fmt.Printf("   continue\n")
				continue
			}
			if loSlope < endSlope {
				//fmt.Printf("   break\n")
				break
			}
			pos := trans(image.Point{X: x, Y: y})
			v.m.MakeVisible(pos)
			//fmt.Printf("   wall = %t\n", wall)
			if !v.m.IsTransparent(pos) {
				if !wall {
					//fmt.Printf("   new wall\n")
					wall = true
					v.visionOctant(x+1, y+1, startSlope, loSlope, trans)
				}
			} else {
				if wall {
					//fmt.Printf("   end wall\n")
					wall = false
					startSlope = (float64(y) + 0.5) / (float64(x) + 0.5)
				}
			}
		}
	}
}

/*
func (v *ShadowCast) visionOctant(col int, minSlope, maxSlope float64, trans visionTransform) {
	fmt.Printf("vision %d\n", col)
	wall := false
	x := col
	for maxSlope > minSlope && !wall {
		fmt.Printf("  %f %f\n", minSlope, maxSlope)
		fStart := maxSlope * (float64(x) + 0.5)
		fEnd := minSlope * (float64(x) + 0.5)
		yStart := int(fStart + 0.5)
		yEnd := int(fEnd + 0.5)
		yRem := fStart - float64(yStart)
		fmt.Printf("  yStart=%3d yEnd=%3d rem=%f\n", yStart, yEnd, yRem)
		for y := yStart; y >= yEnd; y-- {
			fmt.Printf("    x=%3d y=%3d\n", x, y)
			pos := trans(image.Point{X: x, Y: y})
			v.m.MakeVisible(pos)
			if !v.m.IsTransparent(pos) {
				fmt.Printf("      blocker\n")
				if !wall {
					wall = true
					newSlope := float64(y*2+1) / float64(x*2-1)
					if newSlope <= 1.0 {
						fmt.Printf("        next\n")
						v.visionOctant(col+1, newSlope, maxSlope, trans)
						fmt.Printf("        end\n")
					}
				}
				maxSlope = float64(y*2-1) / float64(x*2+1)
			} else {
				if wall {
					wall = false
				}
			}
		}
		x++
	}
}
*/
