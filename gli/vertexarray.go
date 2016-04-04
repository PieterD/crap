package gli

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type VertexArrayObject interface {
	Id() uint32
	Delete()
	Enable(attr ProgramAttribute, buffer Buffer, extent Extent)
	Disable(attr ProgramAttribute)
	Elements(buffer Buffer)
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

func (vao iVertexArrayObject) Enable(attr ProgramAttribute, buffer Buffer, extent Extent) {
	if !attr.Valid() {
		panic(fmt.Errorf("VertexArrayObject.Enable: invalid attribute %#v", attr))
	}
	BindBuffer(ArrayBuffer, buffer)
	BindVertexArrayObject(vao)
	gl.EnableVertexAttribArray(uint32(attr.Index))
	switch attr.Type {
	case GlByte, GlUByte, GlShort, GlUShort, GlInt, GlUInt, GlIntV2, GlIntV3, GlIntV4, GlUIntV2, GlUIntV3, GlUIntV4:
		gl.VertexAttribIPointer(uint32(attr.Index), int32(extent.Components), uint32(extent.Type), int32(extent.Stride), unsafe.Pointer(uintptr(extent.Start)))
	case GlDouble, GlDoubleV2, GlDoubleV3, GlDoubleV4, GlDoubleMat2, GlDoubleMat2x3, GlDoubleMat2x4, GlDoubleMat3x2, GlDoubleMat4x2, GlDoubleMat4x3, GlDoubleMat4:
		gl.VertexAttribLPointer(uint32(attr.Index), int32(extent.Components), uint32(extent.Type), int32(extent.Stride), unsafe.Pointer(uintptr(extent.Start)))
	default:
		gl.VertexAttribPointer(uint32(attr.Index), int32(extent.Components), uint32(extent.Type), extent.Normalize, int32(extent.Stride), unsafe.Pointer(uintptr(extent.Start)))
	}
	UnbindVertexArrayObject()
	UnbindBuffer(ArrayBuffer)
}

func (vao iVertexArrayObject) Disable(attr ProgramAttribute) {
	BindVertexArrayObject(vao)
	gl.DisableVertexAttribArray(uint32(attr.Index))
	UnbindVertexArrayObject()
}

func (vao iVertexArrayObject) Elements(buffer Buffer) {
	BindVertexArrayObject(vao)
	BindBuffer(ElementArrayBuffer, buffer)
	UnbindVertexArrayObject()
}

type ElementInstance struct {
	Vao   VertexArrayObject
	First int
	Count int
	Type  DataType
}
