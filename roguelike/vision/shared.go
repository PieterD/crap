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
