package main

import "github.com/PieterD/glimmer/gli"

var vertexData = []float32{
	-0.75, +0.75,
	-0.25, +0.75,
	+0.25, +0.75,
	+0.75, +0.75,

	-0.75, +0.25,
	-0.25, +0.25,
	+0.25, +0.25,
	+0.75, +0.25,

	-0.75, -0.25,
	-0.25, -0.25,
	+0.25, -0.25,
	+0.75, -0.25,

	-0.75, -0.75,
	-0.25, -0.75,
	+0.25, -0.75,
	+0.75, -0.75,
}

var vertexExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 2,
}

var vertexObject = gli.Object{
	Mode:     gli.Points,
	Vertices: 16,
}
