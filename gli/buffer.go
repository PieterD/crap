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
