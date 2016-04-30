package gli

type Buffer struct {
	ctx      *Context
	id       uint32
	access   iAccessType
	target   iBindTarget
	length   int
	capacity int
}

func (ctx *Context) NewBuffer(access iAccessType, targethint iBindTarget) *Buffer {
}
