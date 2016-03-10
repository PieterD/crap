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
	program  *glimmer.Program
	buffer   gli.Buffer
	data     gli.SliceData
}

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
uniform vec2 offset;

void main() {
	vec4 totalOffset = vec4(offset.x, offset.y, 0.0, 0.0);
	gl_Position = position + totalOffset;
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

	p.program, err = glimmer.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	p.buffer = gli.CreateBuffer(gli.StreamDraw, gli.ArrayBuffer)
	p.data = p.buffer.DataSlice(vertexData)
	p.program.AttributeByName("position", p.data.Pointer(gli.Vertex4d, false, 0, 0))

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
	p.program.Bind()
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	p.program.Unbind()
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
	p.program.UniformFloat2("offset", x, y)
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
