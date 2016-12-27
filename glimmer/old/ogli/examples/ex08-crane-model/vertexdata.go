package main

import "github.com/PieterD/glimmer/gli"

var vertexExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 3,
}

var colorExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 4,
	Start:      6 * 4 * 3 * 4,
}

var rectObject = gli.Object{
	Mode:      gli.Triangles,
	Vertices:  6 * 2 * 3,
	IndexType: gli.GlUShort,
}

var vertexData = []float32{
	//Front
	+1.0, +1.0, +1.0,
	+1.0, -1.0, +1.0,
	-1.0, -1.0, +1.0,
	-1.0, +1.0, +1.0,

	//Top
	+1.0, +1.0, +1.0,
	-1.0, +1.0, +1.0,
	-1.0, +1.0, -1.0,
	+1.0, +1.0, -1.0,

	//Let
	+1.0, +1.0, +1.0,
	+1.0, +1.0, -1.0,
	+1.0, -1.0, -1.0,
	+1.0, -1.0, +1.0,

	//Back
	+1.0, +1.0, -1.0,
	-1.0, +1.0, -1.0,
	-1.0, -1.0, -1.0,
	+1.0, -1.0, -1.0,

	//Bottom
	+1.0, -1.0, +1.0,
	+1.0, -1.0, -1.0,
	-1.0, -1.0, -1.0,
	-1.0, -1.0, +1.0,

	//Right
	-1.0, +1.0, +1.0,
	-1.0, -1.0, +1.0,
	-1.0, -1.0, -1.0,
	-1.0, +1.0, -1.0,

	//Colors
	0, 1, 0, 1,
	0, 1, 0, 1,
	0, 1, 0, 1,
	0, 1, 0, 1,

	0, 0, 1, 1,
	0, 0, 1, 1,
	0, 0, 1, 1,
	0, 0, 1, 1,

	1, 0, 0, 1,
	1, 0, 0, 1,
	1, 0, 0, 1,
	1, 0, 0, 1,

	1, 1, 0, 1,
	1, 1, 0, 1,
	1, 1, 0, 1,
	1, 1, 0, 1,

	0, 1, 1, 1,
	0, 1, 1, 1,
	0, 1, 1, 1,
	0, 1, 1, 1,

	1, 0, 1, 1,
	1, 0, 1, 1,
	1, 0, 1, 1,
	1, 0, 1, 1,
}

var indexData = []uint16{
	0, 1, 2,
	2, 3, 0,

	4, 5, 6,
	6, 7, 4,

	8, 9, 10,
	10, 11, 8,

	12, 13, 14,
	14, 15, 12,

	16, 17, 18,
	18, 19, 16,

	20, 21, 22,
	22, 23, 20,
}