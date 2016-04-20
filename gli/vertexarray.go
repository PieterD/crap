package gli

type VertexArray struct {
	ctx *Context
	id  uint32
}

func (ctx *Context) NewVertexArray() *VertexArray {
	vaoid := ctx.r.VertexArrayCreate()
	return &VertexArray{
		ctx: ctx,
		id:  vaoid,
	}
}

func (vao *VertexArray) Delete() {
	vao.ctx.r.VertexArrayDelete(vao.id)
}

func (vao *VertexArray) bind() {
	if vao.ctx.currentVertexArray == vao {
		return
	}
	vao.ctx.currentVertexArray = vao
	vao.ctx.r.VertexArrayBind(vao.id)
}
