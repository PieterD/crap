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
	DataSlice(iface interface{}) Buffer
	SubSlice(iface interface{}, offset int) Buffer
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

func (buffer iBuffer) DataSlice(iface interface{}) Buffer {
	ptr, size, length, _ := checkSlice(iface)
	BindBuffer(buffer.targethint, buffer)
	gl.BufferData(uint32(buffer.targethint), size*length, ptr, uint32(buffer.accesshint))
	UnbindBuffer(buffer.targethint)
	return buffer
}

func (buffer iBuffer) SubSlice(iface interface{}, offset int) Buffer {
	ptr, size, length, _ := checkSlice(iface)
	_, targethint := buffer.Hints()
	BindBuffer(targethint, buffer)
	gl.BufferSubData(uint32(targethint), offset, size*length, ptr)
	UnbindBuffer(targethint)
	return buffer
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

type Extent struct {
	Start      int
	Stride     int
	Type       DataType
	Components int
	Normalize  bool
}

type Index struct {
	Type     DataType
	Vertices int
}

type Object struct {
	Mode     DrawMode
	Start    int
	Vertices int
}
