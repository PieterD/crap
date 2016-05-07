package main

import (
	"fmt"

	_ "github.com/PieterD/glimmer/driver/gl330"
	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/win"
)

type Profile struct {
	win.DefaultEventHandler
	ctx *gli.Context
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func (p *Profile) Init() error {
	ctx, err := gli.New()
	if err != nil {
		return err
	}
	p.ctx = ctx
	vs, err := p.ctx.Driver().ShaderCreate(gli.ShaderTypeVertex, vertexShaderText)
	Panic(err)
	fs, err := p.ctx.Driver().ShaderCreate(gli.ShaderTypeFragment, fragmentShaderText)
	Panic(err)
	prog, err := p.ctx.Driver().ProgramCreate(nil, vs, fs)
	Panic(err)
	attrs, err := p.ctx.Driver().ProgramAttributes(prog)
	Panic(err)
	for _, attr := range attrs {
		fmt.Printf("%#v\n", attr)
	}
	return nil
}

func (p *Profile) FrameDraw() error {
	return nil
}

func main() {
	err := win.New(win.DefaultConfig(), &Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
