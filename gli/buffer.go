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

func (buffer iBuffer) DataSlice(iface interface{}) {
	size := checkSlice(iface)
	val := reflect.ValueOf(iface)
	num := val.Len()
	ptr := unsafe.Pointer(val.Pointer())
	BindBuffer(buffer.targethint, buffer)
	gl.BufferData(uint32(buffer.targethint), size*num, ptr, uint32(buffer.accesshint))
	UnbindBuffer(buffer.targethint)
}

func checkSlice(iface interface{}) (size int) {
	typ := reflect.TypeOf(iface)
	if typ.Kind() != reflect.Slice && typ.Kind() != reflect.Array {
		panic(fmt.Errorf("DataSlice expected a slice or array type, got %v", typ.String()))
	}
	typ = typ.Elem()
	size = int(typ.Size())
	for {
		switch typ.Kind() {
		case reflect.Array:
			continue
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			return
		default:
			panic(fmt.Errorf("DataSlice expected slice or array of (arrays of) fixed int, uint or float, got slice of %v", typ.String()))
		}
		typ = typ.Elem()
	}
}
