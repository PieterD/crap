package main

import (
	"fmt"
	"math"

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
	data     gli.SliceData
	position gli.ProgramAttribute
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
	p.buffer = gli.CreateBuffer(gli.StreamDraw, gli.ArrayBuffer)

	p.position = p.program.Attributes().ByName("position")
	p.data = p.buffer.DataSlice(vertexData)

	p.vao.Enable(p.position, p.data.Pointer(gli.Vertex4d, false, 0, 0))

	return glimmer.GetError()
}

func (p *Profile) End() {
	p.program.Delete()
	p.fragment.Delete()
	p.vertex.Delete()
	p.buffer.Delete()
}

func (p *Profile) Draw(w *glfw.Window) error {
	p.computePositionOffsets()
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gli.DrawArrays(gli.Triangles, p.program, p.vao.Instance(0, 3))
	return glimmer.GetError()
}

func (p *Profile) computePositionOffsets() {
	loopDuration := 5.0
	scale := 3.14159 * 2.0 / loopDuration
	elapsedTime := glfw.GetTime()
	_, frac := math.Modf(elapsedTime / loopDuration)
	frac *= 5
	x := float32(math.Cos(frac*scale) * 0.5)
	y := float32(math.Sin(frac*scale) * 0.5)
	newvert := make([]float32, len(vertexData))
	copy(newvert, vertexData)
	for i := 0; i < len(newvert); i += 4 {
		newvert[i+0] += x
		newvert[i+1] += y
	}
	p.data.Sub(newvert, 0)
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
