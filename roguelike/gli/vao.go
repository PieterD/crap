package gli

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type VAO struct {
	id uint32
}

func (vao *VAO) Id() uint32 {
	return vao.id
}

func (vao *VAO) Use() {
	gl.BindVertexArray(vao.id)
}

func (vao *VAO) Delete() {
	gl.DeleteVertexArrays(1, &vao.id)
}

func NewVAO() (*VAO, error) {
	var id uint32
	gl.GenVertexArrays(1, &id)
	return &VAO{
		id: id,
	}, nil
}

/*
func (vao *VAO) Enable(uint32 index, buffer *Buffer, elements int, stride int, offset int) {
	gl.BindVertexArray(vao.id)
	defer gl.BindVertexArray(0)
	gl.EnableVertexAttribArray(index)
	gl.VertexAttribPointer(index, elements, buffer.data.typ, false, stride*buffer.data.siz, gl.PtrOffset(offset*buffer.data.siz))
}
*/

func (vao *VAO) Enable(index uint32, elements int, buffer *Buffer) vertexAttribOption {
	return vertexAttribOption{
		vao:        vao,
		index:      index,
		elements:   elements,
		buffer:     buffer,
		stride:     0,
		offset:     gl.PtrOffset(0),
		normalized: false,
	}
}

type vertexAttribOption struct {
	vao        *VAO
	index      uint32
	elements   int
	buffer     *Buffer
	stride     int
	offset     unsafe.Pointer
	normalized bool
}

func (opt vertexAttribOption) Normalized(normalized bool) vertexAttribOption {
	opt.normalized = normalized
	return opt
}

func (opt vertexAttribOption) Stride(stride int) vertexAttribOption {
	opt.stride = stride * opt.buffer.data.siz
	return opt
}

func (opt vertexAttribOption) Offset(offset int) vertexAttribOption {
	opt.offset = gl.PtrOffset(offset * opt.buffer.data.siz)
	return opt
}

func (opt vertexAttribOption) Done() {
	vao := opt.vao
	gl.BindVertexArray(vao.id)
	defer gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, opt.buffer.id)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.EnableVertexAttribArray(opt.index)
	gl.VertexAttribPointer(
		opt.index,
		int32(opt.elements),
		opt.buffer.data.typ,
		opt.normalized,
		int32(opt.stride),
		opt.offset)
}
