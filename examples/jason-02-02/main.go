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
	program  gli.Program
	buffer   gli.Buffer
	vao      gli.VertexArrayObject
}

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;
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

	p.vertex, err = gli.CreateShader(gli.VertexShader, vertexShaderText)
	Panicf(err, "Error compiling vertex shader: %v", err)

	p.fragment, err = gli.CreateShader(gli.FragmentShader, fragmentShaderText)
	Panicf(err, "Error compiling fragment shader: %v", err)

	p.program, err = gli.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	p.vao = gli.CreateVertexArrayObject()
	p.buffer = gli.CreateBuffer(gli.StaticDraw, gli.ArrayBuffer)

	attributes := p.program.Attributes()
	position := attributes.ByName("position")
	color := attributes.ByName("color")

	data := p.buffer.DataSlice(vertexData)
	p.vao.Enable(position, data.Pointer(gli.Vertex4d, false, 0, 0))
	p.vao.Enable(color, data.Pointer(gli.Vertex4d, false, 0, 12))

	return glimmer.GetError()
}

func (p *Profile) End() {
	p.program.Delete()
	p.vao.Delete()
	p.fragment.Delete()
	p.vertex.Delete()
	p.buffer.Delete()
}

func (p *Profile) Draw(w *glfw.Window) error {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gli.BindVertexArrayObject(p.vao)
	gli.BindProgram(p.program)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gli.UnbindProgram()
	gli.UnbindVertexArrayObject()
	return glimmer.GetError()
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
