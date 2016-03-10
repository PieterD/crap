package main

import (
	"fmt"

	"github.com/PieterD/crap/glimmer"
	"github.com/PieterD/crap/glimmer/gli"
	. "github.com/PieterD/crap/glimmer/pan"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	glimmer.DefaultProfile
	vertex   gli.Shader
	fragment gli.Shader
	program  *glimmer.Program
	buffer   gli.Buffer
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

	p.vertex, err = gli.CreateShader(gli.VertexShader, vertexShaderText)
	Panicf(err, "Error compiling vertex shader: %v", err)

	p.fragment, err = gli.CreateShader(gli.FragmentShader, fragmentShaderText)
	Panicf(err, "Error compiling fragment shader: %v", err)

	p.program, err = glimmer.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	p.buffer = gli.CreateBuffer(gli.StaticDraw, gli.ArrayBuffer)
	pointer := p.buffer.DataSlice(vertexData).Pointer(gli.Vertex4d, false, 0, 0)
	p.program.AttributeByName("position", pointer)

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
	p.program.Bind()
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	return nil
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
