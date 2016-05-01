package gli

import (
	"bytes"
	"testing"
)

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
		Color:    [3]byte{0, 255, 128},
		Integer:  2,
	},
}

var testMeshIndexOne = []uint32{0, 1, 1, 0}
var testMeshIndexTwo = []uint32{1, 0, 0}

func TestMesh(t *testing.T) {
	mesh, err := NewMeshBuilder(testMesh{}).
		Format("integer", FmUShort.Full(1)).
		Interleave("integer").
		Index(IndexShort).
		Mode(DrawTriangles).
		Build()
	if err != nil {
		t.Fatalf("MeshBuilder.Build failed: %v", err)
	}

	instance := mesh.Instance()
	object1, err := instance.Object(testMeshData, testMeshIndexOne)
	if err != nil {
		t.Fatalf("MeshInstance.Object failed: %v", err)
	}
	object2, err := instance.Object(testMeshData, testMeshIndexTwo)
	if err != nil {
		t.Fatalf("MeshInstance.Object failed: %v", err)
	}
	if object1.indexbase != 0 || object1.vertexbase != 0 || object1.indexnum != 4 || object1.vertexnum != 2 {
		t.Fatalf("Unexpected values in object1: %#v", object1)
	}
	if object2.indexbase != 4 || object2.vertexbase != 2 || object2.indexnum != 3 || object2.vertexnum != 2 {
		t.Fatalf("Unexpected values in object2: %#v", object2)
	}
	if len(instance.buffers) != 2 {
		t.Fatalf("Expected 2 buffers, got %d", len(instance.buffers))
	}
	exp := []byte{0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0}
	got := instance.indexbuf.buffer.Bytes()
	if !bytes.Equal(exp, got) {
		t.Fatalf("Expected index buffer %#v, got %#v", exp, got)
	}
	exp = []byte{0, 1, 0, 2, 0, 1, 0, 2}
	got = instance.buffers[1].buffer.Bytes()
	if !bytes.Equal(exp, got) {
		t.Fatalf("Expected integer buffer %#v, got %#v", exp, got)
	}
	w := bytes.NewBuffer(nil)
	mw := NewMeshWriter(w)
	mw.PutFloat32(1.0)
	mw.PutFloat32(2.0)
	mw.PutFloat32(3.0)
	mw.PutFloat32(1.0)
	mw.PutUint8(255)
	mw.PutUint8(0)
	mw.PutUint8(0)
	mw.PutUint8(0)
	mw.PutFloat32(3.0)
	mw.PutFloat32(2.0)
	mw.PutFloat32(1.0)
	mw.PutFloat32(1.0)
	mw.PutUint8(0)
	mw.PutUint8(255)
	mw.PutUint8(128)
	oexp := w.Bytes()
	exp = append(oexp, 0) // padding
	exp = append(exp, oexp...)
	got = instance.buffers[0].buffer.Bytes()
	if !bytes.Equal(exp, got) {
		t.Fatalf("Expected position and color buffer %#v, got %#v", exp, got)
	}
}
