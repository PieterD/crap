package gli

import "testing"

type testMesh struct {
	Position [4]float32
	Color    [3]byte
	Integer  uint32 `meshattr:"integer"`
}

func TestMesh(t *testing.T) {
	mb, err := MeshBuilder(testMesh{})
	mb.Attribute("Position", FmFloat)
	mb.Attribute("Color", FmUByte)
	mb.Attribute("integer", FmUShort)
	mb.Interleave("Position", "Color")
	mb.Interleave("integer")
}
