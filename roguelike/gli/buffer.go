package gli

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type Buffer struct {
	id   uint32
	data iDataDesc
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

func (buffer *Buffer) Delete() {
	gl.DeleteBuffers(1, &buffer.id)
}

func NewBuffer(idata interface{}) (*Buffer, error) {
	var id uint32
	gl.GenBuffers(1, &id)
	gl.BindBuffer(gl.ARRAY_BUFFER, id)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	data, err := resolveData(idata)
	if err != nil {
		return nil, err
	}
	gl.BufferData(gl.ARRAY_BUFFER, data.siz*data.length, data.ptr, gl.STATIC_DRAW)

	return &Buffer{
		id:   id,
		data: data.iDataDesc,
	}, nil
}

func resolveData(idata interface{}) (iData, error) {
	var d iData
	switch data := idata.(type) {
	case []float32:
		d.typ = gl.FLOAT
		d.siz = 4
		d.length = len(data)
		d.ptr = gl.Ptr(data)
	default:
		return iData{}, fmt.Errorf("Unusable data type for buffer")
	}
	return d, nil
}
