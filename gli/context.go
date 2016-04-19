package gli

import (
	"fmt"

	"github.com/PieterD/glimmer/raw"
)

type Deletable interface {
	Delete()
}

type Context struct {
	r raw.Raw

	attributeIndexCounter int
	attributeIndexMap     map[string]attributeIndexType
}

type attributeIndexType struct {
	index int
	typ   iDataType
}

func New(r raw.Raw) *Context {
	return &Context{
		r:                 r,
		attributeIndexMap: make(map[string]attributeIndexType),
	}
}

func (ctx *Context) Init() error {
	err := ctx.r.Init()
	if err != nil {
		return fmt.Errorf("Failed to initialize opengl: %v", err)
	}
	return nil
}

func (ctx *Context) Viewport(x, y, width, height int) {
	ctx.r.Viewport(x, y, width, height)
}

func (ctx *Context) ClearColor(r, g, b, a float32) {
	ctx.r.ClearColor(r, g, b, a)
}

func (ctx *Context) SafeDelete(deletables ...Deletable) {
	for _, deletable := range deletables {
		//TODO: Does this work, or do I need to unpack the interface to get at the inner nil?
		if deletable != nil {
			deletable.Delete()
		}
	}
}

func (ctx *Context) VertexAttribute(datatype iDataType, names ...string) {
	if len(names) == 0 {
		return
	}
	if !datatype.ValidAttribute() {
		panic(fmt.Errorf("Invalid vertex attribute type: %s", datatype.String()))
	}
	index := ctx.attributeIndexCounter
	for _, name := range names {
		_, ok := ctx.attributeIndexMap[name]
		if ok {
			panic(fmt.Errorf("Set the same attribute more than once: %s %s", datatype.String(), name))
		}
		ctx.attributeIndexMap[name] = attributeIndexType{
			index: index,
			typ:   datatype,
		}
	}
	ctx.attributeIndexCounter++
}
