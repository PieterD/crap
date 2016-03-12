package main

import "github.com/PieterD/crap/glimmer/gli"

var vertexExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 4,
}

var triangleObject = gli.Object{
	Mode:     gli.Triangles,
	Vertices: 3,
}

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
}
