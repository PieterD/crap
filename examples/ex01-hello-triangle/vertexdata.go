package main

import "github.com/PieterD/glimmer/gli"

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	0.0, 1.0, 0.0, 1.0,
	0.0, 0.0, 1.0, 1.0,
}

var vertices = 3
var vertexPosition, vertexColor gli.Extent

func init() {
	gli.ExtentBuild(4).Ext(4, &vertexPosition).Seq(vertices).Ext(4, &vertexColor).Seq(vertices)
}
