package gli

import "testing"

type testMesh struct {
	Position [4]float32
	Color    [3]byte
	Integer  uint32 `meshattr:"integer"`
}

func TestMesh(t *testing.T) {
	mb, err := MeshBuilder(testMesh{})
	mb.Attribute("Position")
	mb.Attribute("Color")
	mb.Attribute("integer").Format(FmUShort.Full(1))
	mb.Interleave("Position", "Color")
	mb.Interleave("integer")
}
