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

	currentArrayBuffer *ArrayBuffer
	currentVertexArray *VertexArray

	attributeIndexCounter int
	attributeIndexMap     map[string]attributeIndexType
}

type attributeIndexType struct {
	index int
	typ   FullType
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

func (ctx *Context) VertexAttribute(fulltype FullType, names ...string) {
	if len(names) == 0 {
		return
	}
	if !fulltype.DataType.ValidAttribute() {
		panic(fmt.Errorf("Invalid vertex attribute type: %s", fulltype.String()))
	}
	index := ctx.attributeIndexCounter
	for _, name := range names {
		_, ok := ctx.attributeIndexMap[name]
		if ok {
			panic(fmt.Errorf("Set the same attribute more than once: %s %s", fulltype.String(), name))
		}
		ctx.attributeIndexMap[name] = attributeIndexType{
			index: index,
			typ:   fulltype,
		}
	}
	ctx.attributeIndexCounter += calculateLocationSize(fulltype)
}

func calculateLocationSize(fulltype FullType) int {
	var size int
	switch fulltype.DataType {
	case Float, Float2, Float3, Float4:
		size = 1
	case FloatMat2, FloatMat2x3, FloatMat2x4:
		size = 2
	case FloatMat3, FloatMat3x2, FloatMat3x4:
		size = 3
	case FloatMat4, FloatMat4x2, FloatMat4x3:
		size = 4
	case Int, Int2, Int3, Int4, UInt, UInt2, UInt3, UInt4:
		size = 1
	case Double, Double2, Double3, Double4:
		size = 1
	case DoubleMat2, DoubleMat2x3, DoubleMat2x4:
		size = 2
	case DoubleMat3, DoubleMat3x2, DoubleMat3x4:
		size = 3
	case DoubleMat4, DoubleMat4x2, DoubleMat4x3:
		size = 4
	default:
		panic(fmt.Errorf("Invalid attribute type when calculating size: %s", fulltype.String()))
	}
	if fulltype.ArraySize > 0 {
		size *= int(fulltype.ArraySize)
	}
	return size
}
