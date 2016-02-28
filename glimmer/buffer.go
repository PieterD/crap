package glimmer

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Buffer struct {
	id uint32
}

func CreateBuffer() *Buffer {
	buffer := new(Buffer)
	gl.GenBuffers(1, &buffer.id)
	return buffer
}

func (buffer *Buffer) Delete() {
	gl.DeleteBuffers(1, &buffer.id)
}

func (buffer *Buffer) Array() *ArrayBuffer {
	return &ArrayBuffer{buffer: buffer}
}

type ArrayBuffer struct {
	buffer *Buffer
}

func (ab *ArrayBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, ab.buffer.id)
}

func (ab *ArrayBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// TODO: GL_STATIC_DRAW, etc other kinds
func (ab *ArrayBuffer) Data(vertices []float32) {
	ab.Bind()
	defer ab.Unbind()
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), unsafe.Pointer(&vertices[0]), gl.STATIC_DRAW)
}
