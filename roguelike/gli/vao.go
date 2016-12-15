package gli

import (
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

func (vao *VAO) Enable(elements int, buffer *Buffer, attrib Attrib, opts ...VAOOption) {
	opt := vaoOption{
		stride:    0,
		offset:    0,
		normalize: false,
	}
	for _, o := range opts {
		o(&opt)
	}
	gl.BindVertexArray(vao.id)
	defer gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.id)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.EnableVertexAttribArray(attrib.Location())
	gl.VertexAttribPointer(
		attrib.Location(),
		int32(elements),
		buffer.data.typ,
		opt.normalize,
		int32(opt.stride*buffer.data.siz),
		gl.PtrOffset(opt.offset*buffer.data.siz))
}

type vaoOption struct {
	stride    int
	offset    int
	normalize bool
}

type VAOOption func(opt *vaoOption)

func VAONormalize() VAOOption {
	return func(opt *vaoOption) {
		opt.normalize = true
	}
}

func VAOStride(stride int) VAOOption {
	return func(opt *vaoOption) {
		opt.stride = stride
	}
}

func VAOOffset(offset int) VAOOption {
	return func(opt *vaoOption) {
		opt.offset = offset
	}
}
