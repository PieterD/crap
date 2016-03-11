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
	time     gli.ProgramUniform
}

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
uniform float duration;
uniform float time;

void main() {
	float timeScale = 3.14159f * 2.0f / duration;
	float currTime = mod(time, duration);
	vec4 totalOffset = vec4(
		cos(currTime * timeScale) * 0.5f,
		sin(currTime * timeScale) * 0.5f,
		0.0f,
		0.0f);
	gl_Position = position + totalOffset;
}
`

var fragmentShaderText = `
#version 330
out vec4 outputColor;
uniform float fragDuration;
uniform float time;

const vec4 firstColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
const vec4 secondColor = vec4(0.0f, 1.0f, 0.0f, 1.0f);
void main() {
	float currTime = mod(time, fragDuration);
	float currLerp = currTime / fragDuration;
	outputColor = mix(firstColor, secondColor, currLerp);
}
`

var vertexData = []float32{
	0.25, 0.25, 0.0, 1.0,
	0.25, -0.25, 0.0, 1.0,
	-0.25, -0.25, 0.0, 1.0,
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

	position := p.program.Attributes().ByName("position")
	uniforms := p.program.Uniforms()
	uniforms.ByName("duration").Float(5.0)
	uniforms.ByName("fragDuration").Float(10.0)
	p.time = uniforms.ByName("time")

	data := p.buffer.DataSlice(vertexData)
	p.vao.Enable(position, data.Pointer(gli.Vertex4d, false, 0, 0))

	gli.ClearColor(0, 0, 0, 0)

	return glimmer.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.buffer)
}

func (p *Profile) Draw(w *glfw.Window) error {
	p.time.Float(float32(glfw.GetTime()))
	gli.Clear(gli.ColorBufferBit)
	gli.DrawArrays(gli.Triangles, p.program, p.vao.Instance(0, 3))
	return glimmer.GetError()
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
