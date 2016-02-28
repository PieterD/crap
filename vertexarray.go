package glimmer

import "github.com/go-gl/gl/v3.3-core/gl"

type VertexArray struct {
	id uint32
}

func CreateVertexArray() *VertexArray {
	va := new(VertexArray)
	gl.GenVertexArrays(1, &va.id)
	return va
}

func (va *VertexArray) Bind() {
	gl.BindVertexArray(va.id)
}

func (va *VertexArray) Delete() {
	gl.DeleteVertexArrays(1, &va.id)
}

// TODO: Specify pointer (perhaps from ArrayBuffer?)
func (va *VertexArray) Enable(index uint32, ab *ArrayBuffer) {
	ab.Bind()
	va.Bind()
	gl.EnableVertexAttribArray(index)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 0, nil)
}
