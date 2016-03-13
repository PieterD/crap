package main

import (
	"fmt"
	"math"

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

	position gli.ProgramAttribute
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)
	glfw.SwapInterval(1)
	gli.ClearColor(0, 0, 0, 0)

	// Set up shaders
	p.vertex, err = gli.CreateShader(gli.VertexShader, vertexShaderText)
	Panicf(err, "Error compiling vertex shader: %v", err)
	p.fragment, err = gli.CreateShader(gli.FragmentShader, fragmentShaderText)
	Panicf(err, "Error compiling fragment shader: %v", err)
	p.program, err = gli.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	// Set up vertex arrays
	p.vao = gli.CreateVertexArrayObject()
	p.buffer = gli.CreateBuffer(gli.StreamDraw, gli.ArrayBuffer).DataSlice(vertexData)
	attributes := p.program.Attributes()
	p.position = attributes.ByName("position")
	p.vao.Enable(p.position, p.buffer, vertexExtent)

	return glimmer.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.buffer)
}

func (p *Profile) Draw(w *glfw.Window) error {
	p.computePositionOffsets()
	gli.Clear(gli.ColorBufferBit)
	gli.Draw(p.program, p.vao, triangleObject)
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
	p.buffer.SubSlice(newvert, 0)
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
