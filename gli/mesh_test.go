package gli

import "testing"

type testMesh struct {
	Position [4]float32
	Color    [3]byte
	Integer  byte `meshattr:"integer"`
	Noattr   int  `meshattr:"-"`
}

var testMeshData = []testMesh{
	testMesh{
		Position: [4]float32{1.0, 2.0, 3.0, 1.0},
		Color:    [3]byte{255, 0, 0},
		Integer:  1,
	},
	testMesh{
		Position: [4]float32{3.0, 2.0, 1.0, 1.0},
		Color:    [3]byte{0, 255, 0},
		Integer:  2,
	},
}

var testMeshIndex = []uint32{0, 1, 1, 0}

func TestMesh(t *testing.T) {
	mesh, err := NewMeshBuilder(testMesh{}).
		Format("integer", FmUShort.Full(1)).
		Index(IndexShort).
		Mode(DrawTriangles).
		Build()
	if err != nil {
		t.Fatalf("MeshBuilder.Build failed: %v", err)
	}

	instance := mesh.Instance()
	_, err = instance.Object(testMeshData, testMeshIndex)
	if err != nil {
		t.Fatalf("MeshInstance.Object failed: %v", err)
	}
}
