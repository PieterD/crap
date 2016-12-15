package gli

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type Buffer struct {
	id        uint32
	bindpoint uint32
	usage     uint32
	data      iDataDesc
}

type iDataDesc struct {
	typ    uint32
	siz    int
	length int
}

type iData struct {
	iDataDesc
	ptr unsafe.Pointer
}

func (buffer *Buffer) Id() uint32 {
	return buffer.id
}

func (buffer *Buffer) Len() int {
	return buffer.data.length
}

func (buffer *Buffer) Use() {
	gl.BindBuffer(buffer.bindpoint, buffer.id)
}

func (buffer *Buffer) Delete() {
	gl.DeleteBuffers(1, &buffer.id)
}

func NewBuffer(idata interface{}, opts ...BufferOption) (*Buffer, error) {
	opt := bufferOption{
		freq:      STATIC,
		nature:    DRAW,
		bindpoint: gl.ARRAY_BUFFER,
	}
	for _, o := range opts {
		o(&opt)
	}
	var usage uint32
	switch int(opt.freq) | int(opt.nature) {
	case int(STATIC) | int(DRAW):
		usage = gl.STATIC_DRAW
	case int(STATIC) | int(READ):
		usage = gl.STATIC_READ
	case int(STATIC) | int(COPY):
		usage = gl.STATIC_COPY
	case int(STREAM) | int(DRAW):
		usage = gl.STREAM_DRAW
	case int(STREAM) | int(READ):
		usage = gl.STREAM_READ
	case int(STREAM) | int(COPY):
		usage = gl.STREAM_COPY
	case int(DYNAMIC) | int(DRAW):
		usage = gl.DYNAMIC_DRAW
	case int(DYNAMIC) | int(READ):
		usage = gl.DYNAMIC_READ
	case int(DYNAMIC) | int(COPY):
		usage = gl.DYNAMIC_COPY
	default:
		panic(fmt.Errorf("Could not resolve buffer usage from options"))
	}
	data, err := resolveData(idata)
	if err != nil {
		return nil, err
	}

	var id uint32
	gl.GenBuffers(1, &id)

	buffer := &Buffer{
		id:        id,
		data:      data.iDataDesc,
		usage:     usage,
		bindpoint: opt.bindpoint,
	}

	buffer.Upload(idata)
	return buffer, nil
}

func (buffer *Buffer) Upload(idata interface{}) {
	data, err := resolveData(idata)
	if err != nil {
		panic(err)
	}
	if data.typ != buffer.data.typ {
		panic(fmt.Errorf("buffer data type mismatch: %04X and %04X", buffer.data.typ, data.typ))
	}
	gl.BindBuffer(buffer.bindpoint, buffer.id)
	defer gl.BindBuffer(buffer.bindpoint, 0)
	gl.BufferData(buffer.bindpoint, data.siz*data.length, data.ptr, buffer.usage)
}

func (buffer *Buffer) Update(offset int, idata interface{}) {
	data, err := resolveData(idata)
	if err != nil {
		panic(err)
	}
	if data.typ != buffer.data.typ {
		panic(fmt.Errorf("buffer data type mismatch: %04X and %04X", buffer.data.typ, data.typ))
	}
	gl.BindBuffer(buffer.bindpoint, buffer.id)
	defer gl.BindBuffer(buffer.bindpoint, 0)
	gl.BufferSubData(buffer.bindpoint, offset*data.siz, data.length*data.siz, data.ptr)
}

func resolveData(idata interface{}) (iData, error) {
	var d iData
	switch data := idata.(type) {
	case []float32:
		d.typ = gl.FLOAT
		d.siz = 4
		d.length = len(data)
		d.ptr = gl.Ptr(data)
	case []uint8:
		d.typ = gl.UNSIGNED_BYTE
		d.siz = 1
		d.length = len(data)
		d.ptr = gl.Ptr(data)
	case []uint16:
		d.typ = gl.UNSIGNED_SHORT
		d.siz = 2
		d.length = len(data)
		d.ptr = gl.Ptr(data)
	case []uint32:
		d.typ = gl.UNSIGNED_INT
		d.siz = 4
		d.length = len(data)
		d.ptr = gl.Ptr(data)
	default:
		return iData{}, fmt.Errorf("Unusable data type for buffer")
	}
	return d, nil
}

type bufferOption struct {
	freq      BufferAccessFrequencyEnum
	nature    BufferAccessNatureEnum
	bindpoint uint32
}

type BufferOption func(opt *bufferOption)

func BufferAccessFrequency(e BufferAccessFrequencyEnum) BufferOption {
	return func(opt *bufferOption) {
		opt.freq = e
	}
}

type BufferAccessFrequencyEnum int

const (
	STATIC BufferAccessFrequencyEnum = iota
	STREAM
	DYNAMIC
)

func BufferAccessNature(e BufferAccessNatureEnum) BufferOption {
	return func(opt *bufferOption) {
		opt.nature = e
	}
}

type BufferAccessNatureEnum int

const (
	DRAW BufferAccessNatureEnum = iota * 8
	READ
	COPY
)

func BufferElementArray() BufferOption {
	return func(opt *bufferOption) {
		opt.bindpoint = gl.ELEMENT_ARRAY_BUFFER
	}
}
