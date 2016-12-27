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
	vao      gli.VertexArrayObject

	modelToCameraMatrix gli.Uniform
	cameraToClipMatrix  gli.Uniform

	crane *Crane
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)
	glfw.SwapInterval(1)
	gli.ClearColor(0, 0, 0, 0)
	gli.ClearDepth(1)
	gli.EnableCulling(false, true, true)
	gli.EnableDepth(gli.DepthLessEqual, true, 0, 1)
	perspective := gli.PerspectiveMatrix(1.0, 100.0, gli.FrustumScale(45.0), 640, 480)
	p.crane = NewCrane(p)

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
	p.vao = gli.CreateVertexArrayObject()
	p.vao.Enable(attributes.ByName("position"), p.vbuffer, vertexExtent)
	p.vao.Enable(attributes.ByName("color"), p.vbuffer, colorExtent)
	p.vao.Elements(p.ibuffer)

	// Set up uniforms
	uniforms := p.program.Uniforms()
	p.modelToCameraMatrix = uniforms.ByName("modelToCameraMatrix")
	p.cameraToClipMatrix = uniforms.ByName("cameraToClipMatrix")
	p.cameraToClipMatrix.Float(perspective[:]...)

	return gli.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.vbuffer, p.ibuffer)
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	pm := gli.PerspectiveMatrix(1.0, 100.0, gli.FrustumScale(45.0), width, height)
	p.cameraToClipMatrix.Float(pm[:]...)
	gli.Viewport(0, 0, int32(width), int32(height))
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit, gli.DepthBufferBit)
	p.crane.Draw()
	return gli.GetError()
}

func main() {
	err := window.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
