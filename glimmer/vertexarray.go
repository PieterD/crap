package glimmer

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type vertexArray struct {
	id uint32
}

func createVertexArray() *vertexArray {
	va := new(vertexArray)
	gl.GenVertexArrays(1, &va.id)
	return va
}

func (va *vertexArray) bind() {
	gl.BindVertexArray(va.id)
}

func (va *vertexArray) delete() {
	gl.DeleteVertexArrays(1, &va.id)
}

func (va *vertexArray) enable(index uint32, pointer *ArrayPointer) {
	pointer.buffer.Bind(gl.ARRAY_BUFFER)
	va.bind()
	gl.EnableVertexAttribArray(index)
	gl.VertexAttribPointer(index, int32(pointer.datasize), pointer.buffer.datatype, pointer.normalize, int32(pointer.stride), unsafe.Pointer(uintptr(pointer.start)))
	pointer.buffer.Unbind()
}
