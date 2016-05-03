package gli

import "unsafe"

type Buffer struct {
	ctx    *Context
	id     uint32
	access iAccessType
	target iBindTarget
}

func (ctx *Context) NewBuffer(access iAccessType, target iBindTarget) *Buffer {
	bufferid := ctx.r.BufferCreate()
	return &Buffer{
		ctx:    ctx,
		id:     bufferid,
		access: access,
		target: target,
	}
}

func (buffer *Buffer) Delete() {
	buffer.ctx.r.BufferDelete(buffer.id)
}

func (buffer *Buffer) Data(data []byte) {
	ptr := unsafe.Pointer(&data[0])
	buffer.bind(buffer.target)
	buffer.ctx.r.BufferData(buffer.target.t, len(data), ptr, buffer.access.t)
}

func (buffer *Buffer) SubData(offset int, data []byte) {
	ptr := unsafe.Pointer(&data[0])
	buffer.bind(buffer.target)
	buffer.ctx.r.BufferSubData(buffer.target.t, offset, len(data), ptr)
}

func (buffer *Buffer) bind(target iBindTarget) {
	curbind, ok := buffer.ctx.currentBuffer[target]
	if ok && curbind == buffer {
		return
	}
	buffer.ctx.currentBuffer[target] = buffer
	buffer.ctx.r.BufferBind(buffer.id, target.t)
}

func (buffer *Buffer) unbind(target iBindTarget) {
	curbind, ok := buffer.ctx.currentBuffer[target]
	if ok && curbind == buffer {
		buffer.ctx.r.BufferBind(0, target.t)
	}
}
