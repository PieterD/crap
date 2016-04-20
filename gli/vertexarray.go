package gli

type VertexArray struct {
	ctx *Context
	id  uint32
	// attributes map[int]*vertexArrayAttribute
	enabled map[int]struct{}
}

func (ctx *Context) NewVertexArray() *VertexArray {
	vaoid := ctx.r.VertexArrayCreate()
	return &VertexArray{
		ctx:     ctx,
		id:      vaoid,
		enabled: make(map[int]struct{}),
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

func (vao *VertexArray) enable(idx int) {
	vao.bind()
	_, ok := vao.enabled[idx]
	if ok {
		return
	}
	vao.enabled[idx] = struct{}{}
	vao.ctx.r.VertexArrayEnable(idx)
}

func (vao *VertexArray) disable(idx int) {
	vao.bind()
	_, ok := vao.enabled[idx]
	if !ok {
		return
	}
	delete(vao.enabled, idx)
	vao.ctx.r.VertexArrayDisable(idx)
}

/*
func (vao *VertexArray) Attribute(name string, extent Extent, buffer *ArrayBuffer) {
	attributeindex, ok := vao.ctx.attributeIndexMap[name]
	if !ok {
		panic(fmt.Errorf("Could not bind attribute '%s', as it has not been configured", name))
	}
	idx := attributeindex.index
	typ := attributeindex.typ
	vao.bind()
	buffer.bind()
	// TODO: Finish
}
*/
