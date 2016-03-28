package main

import (
	"fmt"
	"math"

	"github.com/PieterD/glimmer/gli"
	. "github.com/PieterD/glimmer/pan"
	"github.com/PieterD/glimmer/window"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
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
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)
	glfw.SwapInterval(1)
	gli.ClearColor(0, 0, 0, 0)
	gli.ClearDepth(1)
	gli.EnableCulling(false, true, true)
	gli.EnableDepth(gli.DepthLessEqual, true, 0, 1)
	perspective := gli.PerspectiveMatrix(1.0, 45.0, gli.FrustumScale(45.0), 640, 480)

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
	pm := gli.PerspectiveMatrix(1.0, 45.0, gli.FrustumScale(45.0), width, height)
	p.cameraToClipMatrix.Float(pm[:]...)
	gli.Viewport(0, 0, int32(width), int32(height))
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit, gli.DepthBufferBit)
	t := glfw.GetTime()
	for _, f := range []TimeAnim{OffsetStationary, OffsetOval, OffsetCircle} {
		offset := f(t)
		p.modelToCameraMatrix.Float(offset[:]...)
		gli.Draw(p.program, p.vao, starObject)
	}
	return gli.GetError()
}

func main() {
	err := window.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

type TimeAnim func(elapsed float64) mgl32.Mat4

func Offset(vec mgl32.Vec4) mgl32.Mat4 {
	m := mgl32.Ident4()
	m.SetCol(3, vec)
	return m
}

func OffsetStationary(elapsed float64) mgl32.Mat4 {
	return Offset(mgl32.Vec4{0, 0, -20, 1})
}

func OffsetOval(elapsed float64) mgl32.Mat4 {
	const fLoopDuration = 3
	const fScale = 3.14159 * 2 / fLoopDuration
	fCurrTimeThroughLoop := math.Mod(elapsed, fLoopDuration)
	return Offset(mgl32.Vec4{
		float32(math.Cos(fCurrTimeThroughLoop*fScale) * 4),
		float32(math.Sin(fCurrTimeThroughLoop*fScale) * 6),
		-20,
		1,
	})
}

func OffsetCircle(elapsed float64) mgl32.Mat4 {
	const fLoopDuration = 12
	const fScale = 3.14159 * 2 / fLoopDuration
	fCurrTimeThroughLoop := math.Mod(elapsed, fLoopDuration)
	return Offset(mgl32.Vec4{
		float32(math.Cos(fCurrTimeThroughLoop*fScale) * 5),
		-3.5,
		float32(math.Sin(fCurrTimeThroughLoop*fScale)*5 - 20),
		1,
	})
}
