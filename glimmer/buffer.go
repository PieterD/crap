package glimmer

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Buffer struct {
	id       uint32
	access   BufferAccessType
	nature   BufferNatureType
	datatype uint32
	bound    uint32
}

type BufferAccessType uint32

const (
	StaticAccess BufferAccessType = iota
	StreamAccess
	DynamicAccess
)

type BufferNatureType uint32

const (
	DrawNature BufferNatureType = iota
	ReadNature
	CopyNature
)

type DataType uint32

const (
	Float DataType = gl.FLOAT
)

func CreateBuffer() *Buffer {
	buffer := new(Buffer)
	gl.GenBuffers(1, &buffer.id)
	buffer.access = StaticAccess
	buffer.nature = DrawNature
	return buffer
}

func (buffer *Buffer) Delete() {
	gl.DeleteBuffers(1, &buffer.id)
}

func (buffer *Buffer) UseAsArrayBuffer() {
	buffer.bound = gl.ARRAY_BUFFER
}

func (buffer *Buffer) AccessNature(access BufferAccessType, nature BufferNatureType) {
	buffer.access = access
	buffer.nature = nature
}

func (buffer *Buffer) Bind() {
	gl.BindBuffer(buffer.bound, buffer.id)
}

func (buffer *Buffer) Unbind() {
	buffer.bound = 0
	gl.BindBuffer(0, buffer.id)
}

func (buffer *Buffer) Data(ptr unsafe.Pointer, size int, datatype uint32) {
	buffer.Bind()
	buffer.datatype = datatype
	gl.BufferData(gl.ARRAY_BUFFER, size, ptr, usage(buffer.access, buffer.nature))
}

func (buffer *Buffer) FloatData(vertices []float32) {
	buffer.Data(unsafe.Pointer(&vertices[0]), 4*len(vertices), gl.FLOAT)
}

func (buffer *Buffer) Pointer(datasize int, normalize bool, stride int, start int) *ArrayPointer {
	return &ArrayPointer{buffer: buffer, datasize: datasize, normalize: normalize, stride: stride, start: start}
}

type ArrayPointer struct {
	buffer    *Buffer
	datasize  int
	normalize bool
	stride    int
	start     int
}

func usage(access BufferAccessType, nature BufferNatureType) uint32 {
	if access == StaticAccess {
		switch nature {
		case DrawNature:
			return gl.STATIC_DRAW
		case ReadNature:
			return gl.STATIC_READ
		case CopyNature:
			return gl.STATIC_COPY
		}
	} else if access == StreamAccess {
		switch nature {
		case DrawNature:
			return gl.STREAM_DRAW
		case ReadNature:
			return gl.STREAM_READ
		case CopyNature:
			return gl.STREAM_COPY
		}
	} else if access == DynamicAccess {
		switch nature {
		case DrawNature:
			return gl.DYNAMIC_DRAW
		case ReadNature:
			return gl.DYNAMIC_READ
		case CopyNature:
			return gl.DYNAMIC_COPY
		}
	}
	// TODO: Possibly, an error
	return gl.STATIC_DRAW
}
