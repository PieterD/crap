package gli

import "unsafe"

type ArrayBuffer struct {
	ctx    *Context
	id     uint32
	size   int
	length int
	typ    iDataType
}

func (ctx *Context) NewArrayBuffer() *ArrayBuffer {
	bufferid := ctx.r.ArrayBufferCreate()
	return &ArrayBuffer{
		ctx: ctx,
		id:  bufferid,
	}
}

func (buffer *ArrayBuffer) Delete() {
	buffer.ctx.r.ArrayBufferDelete(buffer.id)
}

func (buffer *ArrayBuffer) DataFloat(accesstype iAccessType, data []float32) {
	buffer.size = 4
	buffer.length = len(data)
	buffer.typ = Float
	ptr := unsafe.Pointer(&data[0])
	buffer.bind()
	buffer.ctx.r.ArrayBufferData(buffer.id, len(data)*4, ptr, accesstype.t)
}

func (buffer *ArrayBuffer) SubFloat(offset int, data []float32) {
	ptr := unsafe.Pointer(&data[0])
	buffer.bind()
	buffer.ctx.r.ArrayBufferSubData(buffer.id, offset*4, len(data)*4, ptr)
}

func (buffer *ArrayBuffer) bind() {
	if buffer.ctx.currentArrayBuffer == buffer {
		return
	}
	buffer.ctx.currentArrayBuffer = buffer
	buffer.ctx.r.ArrayBufferBind(buffer.id)
}
