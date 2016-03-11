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
	height    float32
	vertex    gli.Shader
	fragment  gli.Shader
	program   gli.Program
	buffer    gli.Buffer
	vao       gli.VertexArrayObject
	uniHeight gli.ProgramUniform
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
uniform float height;
void main() {
	float lerpValue = gl_FragCoord.y / height;
	outputColor = mix(vec4(1.0f, 1.0f, 1.0f, 1.0f), vec4(0.2f, 0.2f, 0.2f, 1.0f), lerpValue);
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

	p.program, err = gli.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	p.vao = gli.CreateVertexArrayObject()
	p.buffer = gli.CreateBuffer(gli.StaticDraw, gli.ArrayBuffer)

	attr := p.program.Attributes().ByName("position")
	pointer := p.buffer.DataSlice(vertexData).Pointer(gli.Vertex4d, false, 0, 0)
	p.vao.Enable(attr, pointer)

	p.uniHeight = p.program.Uniforms().ByName("height")

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
	p.uniHeight.Float(p.height)
	gli.DrawArrays(gli.Triangles, p.program, p.vao.Instance(0, 3))
	return nil
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	p.DefaultProfile.EventResize(w, width, height)
	p.height = float32(height)
}

func main() {
	err := glimmer.Run(&Profile{height: 640.0})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
