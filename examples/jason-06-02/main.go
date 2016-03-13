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

	modelToCameraMatrix gli.ProgramUniform
	cameraToClipMatrix  gli.ProgramUniform
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)
	glfw.SwapInterval(1)
	gli.ClearColor(0, 0, 0, 0)
	gli.ClearDepth(1)
	gli.EnableCulling(false, true, true)
	gli.EnableDepth(gli.DepthLessEqual, true, 0, 1)
	perspective := gli.PerspectiveMatrix(1.0, 61.0, gli.FrustumScale(45.0), 640, 480)

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
	pm := gli.PerspectiveMatrix(1.0, 61.0, gli.FrustumScale(45.0), width, height)
	p.cameraToClipMatrix.Float(pm[:]...)
	gli.Viewport(0, 0, int32(width), int32(height))
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit, gli.DepthBufferBit)
	t := glfw.GetTime()
	for _, f := range []TimeAnim{ScaleNull, ScaleStaticUniform, ScaleStaticNonUniform, ScaleDynamicUniform, ScaleDynamicNonUniform} {
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

func Lerp(elapsed, loop float64) float64 {
	val := math.Mod(elapsed, loop) / loop
	if val > 0.5 {
		val = 1 - val
	}
	return val * 2
}

func ScaleNull(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(1, 1, 1)
	m.SetCol(3, mgl32.Vec4{0, 0, -45, 1})
	return m
}

func ScaleStaticUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(4, 4, 4)
	m.SetCol(3, mgl32.Vec4{-10, -10, -45, 1})
	return m
}

func ScaleStaticNonUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(0.5, 1, 10)
	m.SetCol(3, mgl32.Vec4{-10, 10, -45, 1})
	return m
}

func ScaleDynamicUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(
		Mix(1, 4, Lerp(elapsed, 3)),
		1,
		1,
	)
	m.SetCol(3, mgl32.Vec4{10, 10, -45, 1})
	return m
}

func ScaleDynamicNonUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(
		Mix(1, 0.5, Lerp(elapsed, 3)),
		1,
		Mix(1, 10, Lerp(elapsed, 5)),
	)
	m.SetCol(3, mgl32.Vec4{10, -10, -45, 1})
	return m
}

func Mix(s1, s2, lerp float64) float32 {
	return float32(s1*lerp + s2*(1-lerp))
}
