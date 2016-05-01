package main

import (
	"fmt"
	"os"

	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/raw/raw33"
	"github.com/PieterD/glimmer/window"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	window.DefaultProfile
	ctx     *gli.Context
	vshader *gli.Shader
	fshader *gli.Shader
	program *gli.Program
}

func (p *Profile) PostCreation(w *glfw.Window) error {
	glfw.SwapInterval(1)
	p.ctx = gli.New(raw33.Raw{})
	err := p.ctx.Init()
	if err != nil {
		return fmt.Errorf("Failed to initialize context: %v", err)
	}
	p.ctx.VertexAttribute(gli.Float4.Full(), "Position", "position")
	p.ctx.VertexAttribute(gli.Float3.Full(), "Color", "color")

	p.ctx.ClearColor(0, 0, 0, 0)

	p.vshader, err = p.ctx.NewShader(gli.VertexShader, vertexShaderText)
	if err != nil {
		return fmt.Errorf("Failed to create vertex shader: %v", err)
	}

	p.fshader, err = p.ctx.NewShader(gli.FragmentShader, fragmentShaderText)
	if err != nil {
		return fmt.Errorf("Failed to create fragment shader: %v", err)
	}

	p.program, err = p.ctx.NewProgram(p.vshader, p.fshader)
	if err != nil {
		return fmt.Errorf("Failed to create program: %v", err)
	}

	instance, _ := DefineMyMesh()
	instance.Transmit(p.ctx)

	return nil
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	p.ctx.Viewport(0, 0, width, height)
}

func (p *Profile) Draw(w *glfw.Window) error {
	return nil
}

func main() {
	err := window.Run(&Profile{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running: %v\n", err)
	}
}
