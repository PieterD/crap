package main

import (
	"fmt"

	"github.com/PieterD/crap/glimmer"
	. "github.com/PieterD/crap/glimmer/pan"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	glimmer.DefaultProfile
	vertexShader   *glimmer.Shader
	fragmentShader *glimmer.Shader
	program        *glimmer.Program
	pointer        *glimmer.ArrayPointer
}

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
void main() {
	gl_Position = position;
}
`

var fragmentShaderText = `
#version 330
out vec4 outputColor;
void main() {
	outputColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
}
`

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)

	glfw.SwapInterval(1)
	p.vertexShader, err = glimmer.CreateShader(glimmer.VertexShader, vertexShaderText)
	Panicf(err, "Error compiling vertex shader: %v", err)

	p.fragmentShader, err = glimmer.CreateShader(glimmer.FragmentShader, fragmentShaderText)
	Panicf(err, "Error compiling fragment shader: %v", err)

	p.program, err = glimmer.CreateProgram(p.vertexShader, p.fragmentShader)
	Panicf(err, "Error linking program: %v", err)

	p.pointer = glimmer.CreateBuffer().FloatData(vertexData).Pointer(4, false, 0, 0)

	return glimmer.GetError()
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (p *Profile) End() {
	p.program.Delete()
	p.fragmentShader.Delete()
	p.vertexShader.Delete()
	p.pointer.Buffer().Delete()
}

func (p *Profile) Draw(w *glfw.Window) error {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	p.program.Use()
	p.program.AttributeByName("position", p.pointer)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	return nil
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
