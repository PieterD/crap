package gli

import (
	"fmt"

	"github.com/PieterD/glimmer/raw"
)

type Context struct {
	r raw.Raw
}

func New(r raw.Raw) *Context {
	return &Context{
		r: r,
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
