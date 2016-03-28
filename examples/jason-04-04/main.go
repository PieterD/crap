package main

import (
	"fmt"

	"github.com/PieterD/glimmer/gli"
	. "github.com/PieterD/glimmer/pan"
	"github.com/PieterD/glimmer/window"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	window.DefaultProfile
	vertex   gli.Shader
	fragment gli.Shader
	program  gli.Program
	buffer   gli.Buffer
	vao      gli.VertexArrayObject

	perspectiveMatrix gli.Uniform
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)
	glfw.SwapInterval(1)
	gli.ClearColor(0, 0, 0, 0)
	gli.EnableCulling(false, true, true)
	perspective := gli.PerspectiveMatrix(1.0, 3.0, 1.0, 640, 480)

	// Set up shaders
	p.vertex, err = gli.CreateShader(gli.VertexShader, vertexShaderText)
	Panicf(err, "Error compiling vertex shader: %v", err)
	p.fragment, err = gli.CreateShader(gli.FragmentShader, fragmentShaderText)
	Panicf(err, "Error compiling fragment shader: %v", err)
	p.program, err = gli.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	// Set up vertex arrays
	p.vao = gli.CreateVertexArrayObject()
	p.buffer = gli.CreateBuffer(gli.StaticDraw, gli.ArrayBuffer).DataSlice(vertexData)
	attributes := p.program.Attributes()
	p.vao.Enable(attributes.ByName("position"), p.buffer, vertexExtent)
	p.vao.Enable(attributes.ByName("color"), p.buffer, colorExtent)

	// Set up uniforms
	uniforms := p.program.Uniforms()
	uniforms.ByName("offset").Float(0.5, 0.5)
	p.perspectiveMatrix = uniforms.ByName("perspectiveMatrix")
	p.perspectiveMatrix.Float(perspective[:]...)

	return gli.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.buffer)
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	perspective := gli.PerspectiveMatrix(1.0, 3.0, 1.0, width, height)
	p.perspectiveMatrix.Float(perspective[:]...)
	gli.Viewport(0, 0, int32(width), int32(height))
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit)
	gli.Draw(p.program, p.vao, prismObject)
	return gli.GetError()
}

func main() {
	err := window.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
