package gli

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type VertexArrayObject struct {
	id uint32
}

func (context iContext) CreateVertexArrayObject() VertexArrayObject {
	var id uint32
	gl.GenVertexArrays(1, &id)
	return VertexArrayObject{id: id}
}

func (context iContext) BindVertexArrayObject(vao VertexArrayObject) {
	gl.BindVertexArray(vao.id)
}

func (context iContext) UnbindVertexArrayObject() {
	gl.BindVertexArray(0)
}

func (vao VertexArrayObject) Delete() {
	gl.DeleteVertexArrays(1, &vao.id)
}

func (vao VertexArrayObject) EnableAttrib(index uint32) {
	CreateContext().BindVertexArrayObject(vao)
	gl.EnableVertexAttribArray(index)
	CreateContext().UnbindVertexArrayObject()
}

func (vao VertexArrayObject) AttribPointer(index uint32, datasize int32, datatype DataType, normalize bool, stride int32, start int32) {
	CreateContext().BindVertexArrayObject(vao)
	gl.VertexAttribPointer(index, datasize, uint32(datatype), normalize, stride, unsafe.Pointer(uintptr(start)))
	CreateContext().UnbindVertexArrayObject()
}
