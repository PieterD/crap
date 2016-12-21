package vision

import "image"

type visionTransform func(image.Point) image.Point

type visionTransformer struct {
	source image.Point
	mul    image.Point
}

var quads = []image.Point{
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
}

func (vt visionTransformer) quad(i int) visionTransformer {
	vt.mul = quads[i]
	return vt
}

func (vt visionTransformer) norm(p image.Point) image.Point {
	x := p.X * vt.mul.X
	y := p.Y * vt.mul.Y
	return image.Point{X: x + vt.source.X, Y: y + vt.source.Y}
}

func (vt visionTransformer) swap(p image.Point) image.Point {
	x := p.X * vt.mul.X
	y := p.Y * vt.mul.Y
	return image.Point{X: y + vt.source.X, Y: x + vt.source.Y}
}
