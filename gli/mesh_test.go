package gli

import "testing"

func TestBasicObjectFromStructSlice(t *testing.T) {
	mb := MeshBuilder("mesh-name")
	pos := mb.Attribute("position", FmFloat, Float4.Full(), false)
	col := mb.Attribute("color", FmByte, true)
	pos.Add(0.75, 0.75, 0.0, .0)
	col.Add(255, 0, 0, 255)

	pos.Add(0.75, -0.75, 0.0, 1.0)
	col.Add(0, 255, 0, 255)

	pos.Add(-0.75, -0.75, 0.0, 1.0)
	col.Add(0, 0, 255, 255)

	mb.Done()
}
