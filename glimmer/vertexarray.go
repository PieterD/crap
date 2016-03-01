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

func (va *vertexArray) Bind() {
	gl.BindVertexArray(va.id)
}

func (va *vertexArray) Unbind() {
	gl.BindVertexArray(0)
}

func (va *vertexArray) Delete() {
	gl.DeleteVertexArrays(1, &va.id)
}

func (va *vertexArray) Enable(index uint32, pointer *ArrayPointer) {
	if pointer.buffer.target != ArrayBuffer {
		return
	}
	va.Bind()
	pointer.buffer.bind()
	gl.EnableVertexAttribArray(index)
	gl.VertexAttribPointer(index, int32(pointer.datasize), uint32(pointer.buffer.datatype), pointer.normalize, int32(pointer.stride), unsafe.Pointer(uintptr(pointer.start)))
	pointer.buffer.unbind()
	va.Unbind()
}
