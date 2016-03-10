package gli

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Buffer interface {
	Id() uint32
	Delete()
	DataSlice(iface interface{}) SliceData
	Hints() (BufferAccessTypeHint, BufferTarget)
}

type iBuffer struct {
	id         uint32
	targethint BufferTarget
	accesshint BufferAccessTypeHint
}

type BufferAccessTypeHint uint32

const (
	StaticDraw  BufferAccessTypeHint = gl.STATIC_DRAW
	StaticRead  BufferAccessTypeHint = gl.STATIC_READ
	StaticCopy  BufferAccessTypeHint = gl.STATIC_COPY
	StreamDraw  BufferAccessTypeHint = gl.STREAM_DRAW
	StreamRead  BufferAccessTypeHint = gl.STREAM_READ
	StreamCopy  BufferAccessTypeHint = gl.STREAM_COPY
	DynamicDraw BufferAccessTypeHint = gl.DYNAMIC_DRAW
	DynamicRead BufferAccessTypeHint = gl.DYNAMIC_READ
	DynamicCopy BufferAccessTypeHint = gl.DYNAMIC_COPY
)

type BufferTarget uint32

const (
	ArrayBuffer             BufferTarget = gl.ARRAY_BUFFER
	AtomicCounterBuffer     BufferTarget = gl.ATOMIC_COUNTER_BUFFER
	CopyReadBuffer          BufferTarget = gl.COPY_READ_BUFFER
	CopyWriteBuffer         BufferTarget = gl.COPY_WRITE_BUFFER
	DrawIndirectBuffer      BufferTarget = gl.DRAW_INDIRECT_BUFFER
	DispatchIndirectBuffer  BufferTarget = gl.DISPATCH_INDIRECT_BUFFER
	ElementArrayBuffer      BufferTarget = gl.ELEMENT_ARRAY_BUFFER
	PixelPackBuffer         BufferTarget = gl.PIXEL_PACK_BUFFER
	PixelUnpackBuffer       BufferTarget = gl.PIXEL_UNPACK_BUFFER
	QueryBuffer             BufferTarget = gl.QUERY_BUFFER
	ShaderStorageBuffer     BufferTarget = gl.SHADER_STORAGE_BUFFER
	TextureBuffer           BufferTarget = gl.TEXTURE_BUFFER
	TransformFeedbackBuffer BufferTarget = gl.TRANSFORM_FEEDBACK_BUFFER
	UniformBuffer           BufferTarget = gl.UNIFORM_BUFFER
)

type VertexDimension uint32

const (
	Vertex1d VertexDimension = 1
	Vertex2d VertexDimension = 2
	Vertex3d VertexDimension = 3
	Vertex4d VertexDimension = 4
)

func (context iContext) BindBuffer(target BufferTarget, buffer Buffer) {
	gl.BindBuffer(uint32(target), buffer.Id())
}

func (context iContext) UnbindBuffer(target BufferTarget) {
	gl.BindBuffer(uint32(target), 0)
}

func (context iContext) CreateBuffer(accesshint BufferAccessTypeHint, targethint BufferTarget) Buffer {
	var id uint32
	gl.GenBuffers(1, &id)
	return iBuffer{id: id, targethint: targethint, accesshint: accesshint}
}

func (buffer iBuffer) Id() uint32 {
	return buffer.id
}

func (buffer iBuffer) Delete() {
	gl.DeleteBuffers(1, &buffer.id)
}

func (buffer iBuffer) Hints() (BufferAccessTypeHint, BufferTarget) {
	return buffer.accesshint, buffer.targethint
}

func (buffer iBuffer) DataSlice(iface interface{}) SliceData {
	ptr, size, length, typ := checkSlice(iface)
	BindBuffer(buffer.targethint, buffer)
	gl.BufferData(uint32(buffer.targethint), size*length, ptr, uint32(buffer.accesshint))
	UnbindBuffer(buffer.targethint)
	return SliceData{
		Buffer: buffer,
		Type:   convertBasicType(typ),
		Size:   size,
		Length: length,
	}
}

func checkSlice(iface interface{}) (ptr unsafe.Pointer, size int, length int, typ reflect.Type) {
	val := reflect.ValueOf(iface)
	typ = val.Type()
	if typ.Kind() != reflect.Slice && typ.Kind() != reflect.Array {
		panic(fmt.Errorf("DataSlice expected a slice or array type, got %v", typ))
	}
	typ = typ.Elem()
	ptr = unsafe.Pointer(val.Pointer())
	length = val.Len()
	for {
		switch typ.Kind() {
		case reflect.Array:
			length *= typ.Len()
			continue
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Float32, reflect.Float64:
			size = int(typ.Size())
			return
		default:
			panic(fmt.Errorf("DataSlice expected slice or array of (arrays of) fixed int, uint or float, got slice of %v", typ))
		}
		typ = typ.Elem()
	}
}

func convertBasicType(typ reflect.Type) DataType {
	switch typ.Kind() {
	case reflect.Int8:
		return GlByte
	case reflect.Uint8:
		return GlUByte
	case reflect.Int16:
		return GlShort
	case reflect.Uint16:
		return GlUShort
	case reflect.Int32:
		return GlInt
	case reflect.Uint32:
		return GlUInt
	case reflect.Float32:
		return GlFloat
	case reflect.Float64:
		return GlDouble
	default:
		panic(fmt.Errorf("Unable to convert '%v' to opengl type", typ))
	}
}

type SliceData struct {
	Buffer Buffer
	Type   DataType
	Size   int
	Length int
}

type DataPointer struct {
	Buffer     Buffer
	Type       DataType
	Size       int
	Length     int
	Components VertexDimension
	Normalize  bool
	Stride     int
	Start      int
}

func (data SliceData) Pointer(components VertexDimension, normalize bool, stride int, start int) DataPointer {
	return DataPointer{
		Buffer:     data.Buffer,
		Type:       data.Type,
		Size:       data.Size,
		Length:     data.Length,
		Components: components,
		Normalize:  normalize,
		Stride:     stride * data.Size,
		Start:      start * data.Size,
	}
}

func (data SliceData) Sub(iface interface{}, offset int) {
	ptr, size, length, typ := checkSlice(iface)
	conv := convertBasicType(typ)
	if conv != data.Type {
		panic(fmt.Errorf("Invalid SliceData.Sub: Given type %d does not match original %d", conv, data.Type))
	}
	_, targethint := data.Buffer.Hints()
	BindBuffer(targethint, data.Buffer)
	gl.BufferSubData(uint32(targethint), offset, size*length, ptr)
	UnbindBuffer(targethint)
}
