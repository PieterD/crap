package main

import "github.com/PieterD/glimmer/gli"

var vertexExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 4,
}

var triangleObject = gli.Object{
	Mode:     gli.Triangles,
	Vertices: 3,
}

var vertexData = []float32{
	0.25, 0.25, 0.0, 1.0,
	0.25, -0.25, 0.0, 1.0,
	-0.25, -0.25, 0.0, 1.0,
}
