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

	modelToWorldMatrix gli.ProgramUniform
	globalMatrices     gli.UniformBlock
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)
	glfw.SwapInterval(1)
	gli.ClearColor(0, 0, 0, 0)
	gli.ClearDepth(1)
	gli.EnableCulling(false, true, true)
	gli.EnableDepth(gli.DepthLessEqual, true, 0, 1)
	// perspective := gli.PerspectiveMatrix(1.0, 61.0, gli.FrustumScale(45.0), 640, 480)

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
	p.modelToWorldMatrix = uniforms.ByName("modelToWorldMatrix")
	id := mgl32.Ident4()
	p.modelToWorldMatrix.Float(id[:]...)

	blocks := p.program.UniformBlocks()
	p.globalMatrices = blocks.ByName("GlobalMatrices")
	return gli.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.vbuffer, p.ibuffer)
}

func (p *Profile) EventResize(w *glfw.Window, width int, height int) {
	// pm := gli.PerspectiveMatrix(1.0, 61.0, gli.FrustumScale(45.0), width, height)
	// p.cameraToClipMatrix.Float(pm[:]...)
	gli.Viewport(0, 0, int32(width), int32(height))
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit, gli.DepthBufferBit)
	// t := glfw.GetTime()
	// for _, f := range []TimeAnim{RotateNull, RotateX, RotateY, RotateZ, RotateAxis} {
	// offset := f(t)
	// p.modelToCameraMatrix.Float(offset[:]...)
	gli.Draw(p.program, p.vao, starObject)
	// }
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

func RotateNull(elapsed float64) mgl32.Mat4 {
	m := mgl32.Ident4()
	m.SetCol(3, mgl32.Vec4{0, 0, -25, 1})
	return m
}

func RotateX(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3DX(Angle(elapsed, 2))
	m.SetCol(3, mgl32.Vec4{-5, -5, -25, 1})
	return m
}

func RotateY(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3DY(Angle(elapsed, 2))
	m.SetCol(3, mgl32.Vec4{-5, +5, -25, 1})
	return m
}

func RotateZ(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3DZ(Angle(elapsed, 2))
	m.SetCol(3, mgl32.Vec4{5, 5, -25, 1})
	return m
}

func RotateAxis(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3D(Angle(elapsed, 2), mgl32.Vec3{1, 1, 1}.Normalize())
	m.SetCol(3, mgl32.Vec4{5, -5, -25, 1})
	return m
}

func Angle(elapsed, loop float64) float32 {
	scale := 3.14159 * 2 / loop
	cur := math.Mod(elapsed, loop)
	return float32(scale * cur)
}
