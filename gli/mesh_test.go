package gli

import "testing"

type testMesh struct {
	Position [4]float32
	Color    [3]byte
	Integer  byte `meshattr:"integer"`
}

func TestMesh(t *testing.T) {
	mb, err := NewMeshBuilder(testMesh{})
	if err != nil {
		t.Fatalf("MeshBuilder failed: %v", err)
	}
	mb.Attribute("Position")
	mb.Attribute("Color")
	mb.Attribute("integer").Format(FmUShort.Full(1))
	mb.Interleave("Position", "Color")
	mb.Interleave("integer")
	_, err = mb.Build()
	if err != nil {
		t.Fatalf("MeshBuilder.Build failed: %v", err)
	}
}
