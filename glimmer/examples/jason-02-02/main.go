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
	vertex   *glimmer.Shader
	fragment *glimmer.Shader
	program  *glimmer.Program
	buffer   *glimmer.Buffer
	vPointer *glimmer.ArrayPointer
	cPointer *glimmer.ArrayPointer
}

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;
smooth out vec4 theColor;
void main() {
	gl_Position = position;
	// theColor = position;
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

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	0.0, 1.0, 0.0, 1.0,
	0.0, 0.0, 1.0, 1.0,
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)

	glfw.SwapInterval(1)

	p.vertex, err = glimmer.CreateVertexShader(vertexShaderText)
	Panicf(err, "Error compiling vertex shader: %v", err)

	p.fragment, err = glimmer.CreateFragmentShader(fragmentShaderText)
	Panicf(err, "Error compiling fragment shader: %v", err)

	p.program, err = glimmer.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	p.buffer = glimmer.CreateBuffer().FloatData(vertexData)
	p.vPointer = p.buffer.Pointer(4, false, 0, 0)
	p.cPointer = p.buffer.Pointer(4, false, 0, 48)

	return glimmer.GetError()
}

func (p *Profile) End() {
	p.program.Delete()
	p.fragment.Delete()
	p.vertex.Delete()
	p.buffer.Delete()
}

func (p *Profile) Draw(w *glfw.Window) error {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	p.program.Use()
	p.program.AttributeByName("position", p.vPointer)
	p.program.AttributeByName("color", p.cPointer)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	return nil
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}