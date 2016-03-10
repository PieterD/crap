package gli

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type VertexArrayObject interface {
	Id() uint32
	Delete()
	Enable(attr ProgramAttribute, pointer DataPointer)
	Disable(attr ProgramAttribute)
}

type iVertexArrayObject struct {
	id uint32
}

func (context iContext) CreateVertexArrayObject() VertexArrayObject {
	var id uint32
	gl.GenVertexArrays(1, &id)
	return iVertexArrayObject{id: id}
}

func (context iContext) BindVertexArrayObject(vao VertexArrayObject) {
	gl.BindVertexArray(vao.Id())
}

func (context iContext) UnbindVertexArrayObject() {
	gl.BindVertexArray(0)
}

func (vao iVertexArrayObject) Id() uint32 {
	return vao.id
}

func (vao iVertexArrayObject) Delete() {
	gl.DeleteVertexArrays(1, &vao.id)
}

func (vao iVertexArrayObject) Enable(attr ProgramAttribute, pointer DataPointer) {
	BindBuffer(ArrayBuffer, pointer.Buffer)
	BindVertexArrayObject(vao)
	gl.EnableVertexAttribArray(uint32(attr.Index))
	gl.VertexAttribPointer(uint32(attr.Index), int32(pointer.Components), uint32(pointer.Type), pointer.Normalize, int32(pointer.Stride), unsafe.Pointer(uintptr(pointer.Start)))
	UnbindVertexArrayObject()
	UnbindBuffer(ArrayBuffer)
}

func (vao iVertexArrayObject) Disable(attr ProgramAttribute) {
	BindVertexArrayObject(vao)
	gl.DisableVertexAttribArray(uint32(attr.Index))
	UnbindVertexArrayObject()
}
