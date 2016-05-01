package main

import "github.com/PieterD/glimmer/gli"

type MeshType struct {
	Position [4]float32
	Color    [4]float32
}

var meshData = []MeshType{
	MeshType{
		Position: [4]float32{0.75, 0.75, 0.0, 1.0},
		Color:    [4]float32{1.0, 0.0, 0.0, 1.0},
	},
	MeshType{
		Position: [4]float32{0.75, -0.75, 0.0, 1.0},
		Color:    [4]float32{0.0, 1.0, 0.0, 1.0},
	},
	MeshType{
		Position: [4]float32{-0.75, -0.75, 0.0, 1.0},
		Color:    [4]float32{0.0, 0.0, 1.0, 1.0},
	},
}

func DefineMyMesh() *gli.Mesh {
	mb, err := gli.NewMeshBuilder(MeshType{})
	if err != nil {
		panic(err)
	}
	mb.Attribute("Position")
	mb.Attribute("Color")
	mb.Interleave("Position", "Color")
	mb.Mode(gli.DrawTriangles)
	mesh, err := mb.Build()
	if err != nil {
		panic(err)
	}
	return mesh
}

func LoadMyMesh(ctx *gli.Context, mesh *gli.Mesh, data []MeshType) {

}
