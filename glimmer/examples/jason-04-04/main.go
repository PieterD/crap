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

	perspectiveMatrix gli.ProgramUniform
}

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;
uniform vec2 offset;
uniform mat4 perspectiveMatrix;

smooth out vec4 theColor;

void main() {
	vec4 cameraPos = position + vec4(offset.x, offset.y, 0.0, 0.0);
	gl_Position = perspectiveMatrix * cameraPos;
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
	p.buffer = gli.CreateBuffer(gli.StaticDraw, gli.ArrayBuffer)

	data := p.buffer.DataSlice(vertexData)

	pm := gli.PerspectiveMatrix(0.5, 3.0, 1.0, 640, 480)

	attributes := p.program.Attributes()
	p.vao.Enable(attributes.ByName("position"), data.Pointer(gli.Vertex4d, false, 0, 0))
	p.vao.Enable(attributes.ByName("color"), data.Pointer(gli.Vertex4d, false, 0, len(vertexData)/2))
	uniforms := p.program.Uniforms()
	uniforms.ByName("offset").Float(0.5, 0.5)
	p.perspectiveMatrix = uniforms.ByName("perspectiveMatrix")
	p.perspectiveMatrix.Float(pm[:]...)

	gli.ClearColor(0, 0, 0, 0)
	gli.EnableCulling(false, true, true)

	return glimmer.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.buffer)
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	pm := gli.PerspectiveMatrix(0.5, 3.0, 1.0, width, height)
	p.perspectiveMatrix.Float(pm[:]...)
	gl.Viewport(0, 0, int32(width), int32(height))
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
