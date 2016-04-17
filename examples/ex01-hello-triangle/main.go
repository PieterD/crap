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
	buffer  *gli.ArrayBuffer
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
uniform float colorshift;

void main() {
	outputColor = theColor * colorshift;
}
`

func (p *Profile) PostCreation(w *glfw.Window) error {
	glfw.SwapInterval(1)
	p.ctx = gli.New(raw33.Raw{})
	err := p.ctx.Init()
	if err != nil {
		return fmt.Errorf("Failed to initialize context: %v", err)
	}
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

	p.buffer = p.ctx.NewArrayBuffer()

	attributes := p.program.Attributes()
	position := attributes.ByName("position")
	dt, as := position.Type()
	fmt.Printf("position: %s, %d\n", dt, as)

	uniforms := p.program.Uniforms()
	colorshift := uniforms.ByName("colorshift")
	dt, as = colorshift.Type()
	fmt.Printf("colorshift: %s, %d\n", dt, as)

	return nil
}

func (p *Profile) End() {
	p.ctx.SafeDelete(p.buffer, p.vshader, p.fshader, p.program)
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
