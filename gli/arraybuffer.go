package gli

type ArrayBuffer struct {
	ctx        *Context
	id         uint32
	accesstype iAccessType
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
