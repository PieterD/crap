package main

import (
	"fmt"

	"github.com/PieterD/crap/glimmer"
	"github.com/PieterD/crap/glimmer/gli"
	. "github.com/PieterD/crap/glimmer/pan"
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
uniform vec2 offset;

smooth out vec4 theColor;

void main() {
	gl_Position = position + vec4(offset.x, offset.y, 0.0, 0.0);
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
	p.buffer = gli.CreateBuffer(gli.StreamDraw, gli.ArrayBuffer)

	data := p.buffer.DataSlice(vertexData)

	attributes := p.program.Attributes()
	p.vao.Enable(attributes.ByName("position"), data.Pointer(gli.Vertex4d, false, 0, 0))
	p.vao.Enable(attributes.ByName("color"), data.Pointer(gli.Vertex4d, false, 0, len(vertexData)/2))
	uniforms := p.program.Uniforms()
	uniforms.ByName("offset").Float(0.5, 0.25)

	gli.ClearColor(0, 0, 0, 0)
	gli.EnableCulling(false, true, true)

	return glimmer.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.buffer)
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit)
	gli.DrawArrays(gli.Triangles, p.program, p.vao.Instance(0, 36))
	return glimmer.GetError()
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
