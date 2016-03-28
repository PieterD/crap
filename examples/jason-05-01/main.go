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
	vbuffer  gli.Buffer
	ibuffer  gli.Buffer
	vao1     gli.VertexArrayObject
	vao2     gli.VertexArrayObject

	offset            gli.Uniform
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
	p.vbuffer = gli.CreateBuffer(gli.StaticDraw, gli.ArrayBuffer).DataSlice(vertexData)
	p.ibuffer = gli.CreateBuffer(gli.StaticDraw, gli.ElementArrayBuffer).DataSlice(indexData)
	attributes := p.program.Attributes()
	pos := attributes.ByName("position")
	col := attributes.ByName("color")
	p.vao1 = gli.CreateVertexArrayObject()
	p.vao1.Enable(pos, p.vbuffer, vertexExtent1)
	p.vao1.Enable(col, p.vbuffer, colorExtent1)
	p.vao1.Elements(p.ibuffer)
	p.vao2 = gli.CreateVertexArrayObject()
	p.vao2.Enable(pos, p.vbuffer, vertexExtent2)
	p.vao2.Enable(col, p.vbuffer, colorExtent2)
	p.vao2.Elements(p.ibuffer)

	// Set up uniforms
	uniforms := p.program.Uniforms()
	p.offset = uniforms.ByName("offset")
	p.perspectiveMatrix = uniforms.ByName("perspectiveMatrix")
	p.perspectiveMatrix.Float(perspective[:]...)

	return gli.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao1, p.vao2, p.fragment, p.vertex, p.vbuffer, p.ibuffer)
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	perspective := gli.PerspectiveMatrix(1.0, 3.0, 1.0, width, height)
	p.perspectiveMatrix.Float(perspective[:]...)
	gli.Viewport(0, 0, int32(width), int32(height))
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit)
	p.offset.Float(0.0, 0.0, 0.0)
	gli.Draw(p.program, p.vao1, wedgeObject)
	p.offset.Float(0.0, 0.0, -1.0)
	gli.Draw(p.program, p.vao2, wedgeObject)
	return gli.GetError()
}

func main() {
	err := window.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
