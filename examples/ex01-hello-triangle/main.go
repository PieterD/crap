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

var vertexShaderText = `
#version 330

in vec4 position;
in vec4 color;
smooth out vec4 theColor;

void main() {
	gl_Position = position;
	theColor = color;
}
`

var fragmentShaderText = `
#version 330

smooth in vec4 theColor;
out vec4 outputColor;

void main() {
	outputColor = theColor;
}
`

func (p *Profile) PostCreation(w *glfw.Window) error {
	glfw.SwapInterval(1)
	p.ctx = gli.New(raw33.Raw{})
	err := p.ctx.Init()
	if err != nil {
		return fmt.Errorf("Failed to initialize context: %v", err)
	}

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
