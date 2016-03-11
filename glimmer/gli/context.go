package gli

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func Init() {
	gl.Init()
}

type iContext struct {
}

var Current Context = CreateContext()

func CreateContext() Context {
	return iContext{}
}

type Context interface {
	CreateShader(shaderType ShaderType, source ...string) (Shader, error)
	CreateProgram(shaders ...Shader) (Program, error)
	CreateVertexArrayObject() VertexArrayObject
	CreateBuffer(accesshint BufferAccessTypeHint, targethint BufferTarget) Buffer
	BindProgram(program Program)
	UnbindProgram()
	BindVertexArrayObject(vao VertexArrayObject)
	UnbindVertexArrayObject()
	BindBuffer(target BufferTarget, buffer Buffer)
	UnbindBuffer(target BufferTarget)
	DrawArrays(mode DrawMode, program Program, array ArrayInstance)
	ClearColor(r, g, b, a float32)
	Clear(bits ...ClearBit)
	Enable(cap Capability)
	Disable(cap Capability)
	EnableIndex(cap Capability, index uint32)
	DisableIndex(cap Capability, index uint32)
	EnableCulling(frontface bool, backface bool, clockwise bool)
	DisableCulling()
}

func CreateShader(shaderType ShaderType, source ...string) (Shader, error) {
	return Current.CreateShader(shaderType, source...)
}
func CreateProgram(shaders ...Shader) (Program, error) {
	return Current.CreateProgram(shaders...)
}
func CreateVertexArrayObject() VertexArrayObject {
	return Current.CreateVertexArrayObject()
}
func CreateBuffer(accesshint BufferAccessTypeHint, targethint BufferTarget) Buffer {
	return Current.CreateBuffer(accesshint, targethint)
}
func BindProgram(program Program) {
	Current.BindProgram(program)
}
func UnbindProgram() {
	Current.UnbindProgram()
}
func BindVertexArrayObject(vao VertexArrayObject) {
	Current.BindVertexArrayObject(vao)
}
func UnbindVertexArrayObject() {
	Current.UnbindVertexArrayObject()
}
func BindBuffer(target BufferTarget, buffer Buffer) {
	Current.BindBuffer(target, buffer)
}
func UnbindBuffer(target BufferTarget) {
	Current.UnbindBuffer(target)
}
func DrawArrays(mode DrawMode, program Program, array ArrayInstance) {
	Current.DrawArrays(mode, program, array)
}
func (context iContext) DrawArrays(mode DrawMode, program Program, array ArrayInstance) {
	BindVertexArrayObject(array.Vao)
	BindProgram(program)
	gl.DrawArrays(uint32(mode), int32(array.First), int32(array.Count))
	UnbindProgram()
	UnbindVertexArrayObject()
}
func ClearColor(r, g, b, a float32) {
	Current.ClearColor(r, g, b, a)
}
func (context iContext) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}
func Clear(bits ...ClearBit) {
	Current.Clear(bits...)
}
func (context iContext) Clear(bits ...ClearBit) {
	var b uint32
	for _, bit := range bits {
		b |= uint32(bit)
	}
	gl.Clear(b)
}
func Enable(cap Capability) {
	Current.Enable(cap)
}
func (context iContext) Enable(cap Capability) {
	gl.Enable(uint32(cap))
}
func Disable(cap Capability) {
	Current.Disable(cap)
}
func (context iContext) Disable(cap Capability) {
	gl.Disable(uint32(cap))
}
func EnableIndex(cap Capability, index uint32) {
	Current.EnableIndex(cap, index)
}
func (context iContext) EnableIndex(cap Capability, index uint32) {
	gl.Enablei(uint32(cap), index)
}
func DisableIndex(cap Capability, index uint32) {
	Current.DisableIndex(cap, index)
}
func (context iContext) DisableIndex(cap Capability, index uint32) {
	gl.Disablei(uint32(cap), index)
}
func EnableCulling(front bool, back bool, clockwise bool) {
	Current.EnableCulling(front, back, clockwise)
}
func (context iContext) EnableCulling(frontface bool, backface bool, clockwise bool) {
	context.Enable(CullFace)
	if frontface && backface {
		gl.CullFace(gl.FRONT_AND_BACK)
	} else if frontface {
		gl.CullFace(gl.FRONT)
	} else if backface {
		gl.CullFace(gl.BACK)
	}
	if clockwise {
		gl.FrontFace(gl.CW)
	} else {
		gl.FrontFace(gl.CCW)
	}
}
func DisableCulling() {
	Current.DisableCulling()
}
func (context iContext) DisableCulling() {
	context.Disable(CullFace)
}
