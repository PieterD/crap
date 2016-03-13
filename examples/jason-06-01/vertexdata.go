package main

import "github.com/PieterD/crap/glimmer/gli"

var vertexExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 3,
}

var colorExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 4,
	Start:      8 * 3 * 4,
}

var starObject = gli.Object{
	Mode:      gli.Triangles,
	Vertices:  24,
	IndexType: gli.GlUShort,
}

var vertexData = []float32{
	+1.0, +1.0, +1.0,
	-1.0, -1.0, +1.0,
	-1.0, +1.0, -1.0,
	+1.0, -1.0, -1.0,

	-1.0, -1.0, -1.0,
	+1.0, +1.0, -1.0,
	+1.0, -1.0, +1.0,
	-1.0, +1.0, +1.0,

	0.0, 1.0, 0.0, 1.0,
	0.0, 0.0, 1.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,

	0.0, 1.0, 0.0, 1.0,
	0.0, 0.0, 1.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,
}

var indexData = []uint16{
	0, 1, 2,
	1, 0, 3,
	2, 3, 0,
	3, 2, 1,

	5, 4, 6,
	4, 5, 7,
	7, 6, 4,
	6, 7, 5,
}
