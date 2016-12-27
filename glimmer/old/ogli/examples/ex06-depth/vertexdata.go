package main

import "github.com/PieterD/glimmer/gli"

var vertexExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 3,
}

var colorExtent = gli.Extent{
	Type:       gli.GlFloat,
	Components: 4,
	Start:      2 * 18 * 3 * 4,
}

var wedgeObject1 = gli.Object{
	Mode:      gli.Triangles,
	Vertices:  36,
	IndexType: gli.GlUShort,
}

var wedgeObject2 = gli.Object{
	Mode:      gli.Triangles,
	Vertices:  36,
	IndexType: gli.GlUShort,
	IndexBase: 18,
}

const (
	RIGHT_EXTENT  float32 = 0.8
	LEFT_EXTENT   float32 = -RIGHT_EXTENT
	TOP_EXTENT    float32 = 0.2
	MIDDLE_EXTENT float32 = 0.0
	BOTTOM_EXTENT float32 = -TOP_EXTENT
	FRONT_EXTENT  float32 = -1.25
	REAR_EXTENT   float32 = -1.75
)

var vertexData = []float32{
	// Object 1 positions
	LEFT_EXTENT, TOP_EXTENT, REAR_EXTENT,
	LEFT_EXTENT, MIDDLE_EXTENT, FRONT_EXTENT,
	RIGHT_EXTENT, MIDDLE_EXTENT, FRONT_EXTENT,
	RIGHT_EXTENT, TOP_EXTENT, REAR_EXTENT,

	LEFT_EXTENT, BOTTOM_EXTENT, REAR_EXTENT,
	LEFT_EXTENT, MIDDLE_EXTENT, FRONT_EXTENT,
	RIGHT_EXTENT, MIDDLE_EXTENT, FRONT_EXTENT,
	RIGHT_EXTENT, BOTTOM_EXTENT, REAR_EXTENT,

	LEFT_EXTENT, TOP_EXTENT, REAR_EXTENT,
	LEFT_EXTENT, MIDDLE_EXTENT, FRONT_EXTENT,
	LEFT_EXTENT, BOTTOM_EXTENT, REAR_EXTENT,

	RIGHT_EXTENT, TOP_EXTENT, REAR_EXTENT,
	RIGHT_EXTENT, MIDDLE_EXTENT, FRONT_EXTENT,
	RIGHT_EXTENT, BOTTOM_EXTENT, REAR_EXTENT,

	LEFT_EXTENT, BOTTOM_EXTENT, REAR_EXTENT,
	LEFT_EXTENT, TOP_EXTENT, REAR_EXTENT,
	RIGHT_EXTENT, TOP_EXTENT, REAR_EXTENT,
	RIGHT_EXTENT, BOTTOM_EXTENT, REAR_EXTENT,

	// Object 2 positions
	TOP_EXTENT, RIGHT_EXTENT, REAR_EXTENT,
	MIDDLE_EXTENT, RIGHT_EXTENT, FRONT_EXTENT,
	MIDDLE_EXTENT, LEFT_EXTENT, FRONT_EXTENT,
	TOP_EXTENT, LEFT_EXTENT, REAR_EXTENT,

	BOTTOM_EXTENT, RIGHT_EXTENT, REAR_EXTENT,
	MIDDLE_EXTENT, RIGHT_EXTENT, FRONT_EXTENT,
	MIDDLE_EXTENT, LEFT_EXTENT, FRONT_EXTENT,
	BOTTOM_EXTENT, LEFT_EXTENT, REAR_EXTENT,

	TOP_EXTENT, RIGHT_EXTENT, REAR_EXTENT,
	MIDDLE_EXTENT, RIGHT_EXTENT, FRONT_EXTENT,
	BOTTOM_EXTENT, RIGHT_EXTENT, REAR_EXTENT,

	TOP_EXTENT, LEFT_EXTENT, REAR_EXTENT,
	MIDDLE_EXTENT, LEFT_EXTENT, FRONT_EXTENT,
	BOTTOM_EXTENT, LEFT_EXTENT, REAR_EXTENT,

	BOTTOM_EXTENT, RIGHT_EXTENT, REAR_EXTENT,
	TOP_EXTENT, RIGHT_EXTENT, REAR_EXTENT,
	TOP_EXTENT, LEFT_EXTENT, REAR_EXTENT,
	BOTTOM_EXTENT, LEFT_EXTENT, REAR_EXTENT,

	//Object 1 colors
	0.75, 0.75, 1.0, 1.0,
	0.75, 0.75, 1.0, 1.0,
	0.75, 0.75, 1.0, 1.0,
	0.75, 0.75, 1.0, 1.0,

	0.0, 0.5, 0.0, 1.0,
	0.0, 0.5, 0.0, 1.0,
	0.0, 0.5, 0.0, 1.0,
	0.0, 0.5, 0.0, 1.0,

	1.0, 0.0, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,

	0.8, 0.8, 0.8, 1.0,
	0.8, 0.8, 0.8, 1.0,
	0.8, 0.8, 0.8, 1.0,

	0.5, 0.5, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,

	//Object 2 colors
	1.0, 0.0, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,

	0.5, 0.5, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,
	0.5, 0.5, 0.0, 1.0,

	0.0, 0.5, 0.0, 1.0,
	0.0, 0.5, 0.0, 1.0,
	0.0, 0.5, 0.0, 1.0,

	0.75, 0.75, 1.0, 1.0,
	0.75, 0.75, 1.0, 1.0,
	0.75, 0.75, 1.0, 1.0,

	0.8, 0.8, 0.8, 1.0,
	0.8, 0.8, 0.8, 1.0,
	0.8, 0.8, 0.8, 1.0,
	0.8, 0.8, 0.8, 1.0,
}

var indexData = []uint16{
	0, 2, 1,
	3, 2, 0,

	4, 5, 6,
	6, 7, 4,

	8, 9, 10,
	11, 13, 12,

	14, 16, 15,
	17, 16, 14,
}