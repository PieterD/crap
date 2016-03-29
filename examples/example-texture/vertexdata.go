package main

import "github.com/PieterD/glimmer/gli"

var vertexExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 3,
	Stride:     5 * 4,
}

var textureExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 2,
	Stride:     5 * 4,
	Start:      3 * 4,
}

var triangleObject = gli.Object{
	Mode:     gli.Triangles,
	Vertices: 3,
}

var vertexData = []float32{
	0.75, 0.75, 0.0,
	1, 0,
	0.75, -0.75, 0.0,
	1, 1,
	-0.75, -0.75, 0.0,
	0, 1,
}
