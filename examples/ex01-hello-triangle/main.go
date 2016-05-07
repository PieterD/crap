package main

import (
	"fmt"

	_ "github.com/PieterD/glimmer/driver/gl330"
	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/prog"
	"github.com/PieterD/glimmer/win"
)

type Profile struct {
	win.DefaultEventHandler
	programs gli.ProgramCollection
	ctx      *gli.Context
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
	program, err := p.programs.Program("MyGroup.MyProgram")
	Panic(err)

	var shaders []gli.ShaderId
	for _, shader := range program.Shaders {
		id, err := p.ctx.Driver().ShaderCreate(shader.Type, shader.Source...)
		Panic(err)
		shaders = append(shaders, id)
	}
	prog, err := p.ctx.Driver().ProgramCreate(program.Locations, shaders...)
	Panic(err)

	attrs, err := p.ctx.Driver().ProgramAttributes(prog)
	Panic(err)
	for _, attr := range attrs {
		fmt.Printf("%#v\n", attr)
	}
	Panic(err)
	return nil
}

func (p *Profile) FrameDraw() error {
	return nil
}

func main() {
	programs, err := prog.ReadPrograms("shaders/programs.json")
	Panic(err)
	err = win.New(win.DefaultConfig(), &Profile{programs: programs})
	Panic(err)
}
