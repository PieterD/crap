package gli

import "github.com/go-gl/gl/v3.3-core/gl"

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
	CreateTexture(target TextureTarget) Texture
	CreateSampler() Sampler
	BindProgram(program Program)
	UnbindProgram()
	BindVertexArrayObject(vao VertexArrayObject)
	UnbindVertexArrayObject()
	BindBuffer(target BufferTarget, buffer Buffer)
	UnbindBuffer(target BufferTarget)
	BindTexture(texture Texture)
	BindSampler(sampler Sampler, textureunit uint32)
	UnbindSampler(textureunit uint32)
	ActiveTexture(i uint32)
	Draw(program Program, vao VertexArrayObject, object Object)
	ClearColor(r, g, b, a float32)
	ClearDepth(d float32)
	Clear(bits ...ClearBit)
	Enable(cap Capability)
	Disable(cap Capability)
	EnableIndex(cap Capability, index uint32)
	DisableIndex(cap Capability, index uint32)
	IsEnabled(cap Capability) bool
	IsEnabledIndex(cap Capability, index uint32) bool
	EnableCulling(frontface bool, backface bool, clockwise bool)
	DisableCulling()
	EnableDepth(depthfunc DepthFunc, mask bool, nearRange float32, farRange float32)
	DisableDepth()
	Viewport(x, y int32, width, height int32)
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
func CreateTexture(target TextureTarget) Texture {
	return Current.CreateTexture(target)
}
func CreateSampler() Sampler {
	return Current.CreateSampler()
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
func BindTexture(texture Texture) {
	Current.BindTexture(texture)
}
func BindSampler(sampler Sampler, textureunit uint32) {
	Current.BindSampler(sampler, textureunit)
}
func UnbindSampler(textureunit uint32) {
	Current.UnbindSampler(textureunit)
}
func ActiveTexture(i uint32) {
	Current.ActiveTexture(i)
}
func ClearColor(r, g, b, a float32) {
	Current.ClearColor(r, g, b, a)
}
func (context iContext) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}
func ClearDepth(d float32) {
	Current.ClearDepth(d)
}
func (context iContext) ClearDepth(d float32) {
	gl.ClearDepthf(d)
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
func IsEnabled(cap Capability) bool {
	return Current.IsEnabled(cap)
}
func (context iContext) IsEnabled(cap Capability) bool {
	return gl.IsEnabled(uint32(cap))
}
func IsEnabledIndex(cap Capability, index uint32) bool {
	return Current.IsEnabledIndex(cap, index)
}
func (context iContext) IsEnabledIndex(cap Capability, index uint32) bool {
	return gl.IsEnabledi(uint32(cap), index)
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
func EnableDepth(depthfunc DepthFunc, mask bool, nearRange float32, farRange float32) {
	Current.EnableDepth(depthfunc, mask, nearRange, farRange)
}
func (context iContext) EnableDepth(depthfunc DepthFunc, mask bool, nearRange float32, farRange float32) {
	context.Enable(DepthTest)
	gl.DepthMask(mask)
	gl.DepthFunc(uint32(depthfunc))
	gl.DepthRangef(nearRange, farRange)
}
func DisableDepth() {
	Current.DisableDepth()
}
func (context iContext) DisableDepth() {
	context.Disable(DepthTest)
}
func Viewport(x, y int32, width, height int32) {
	Current.Viewport(x, y, width, height)
}
func (context iContext) Viewport(x, y int32, width, height int32) {
	gl.Viewport(0, 0, width, height)
}
