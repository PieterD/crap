package main

import (
	"fmt"

	"github.com/PieterD/glimmer/gli"
	. "github.com/PieterD/glimmer/pan"
	"github.com/PieterD/glimmer/window"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	window.DefaultProfile
	vertex   gli.Shader
	fragment gli.Shader
	program  gli.Program
	buffer   gli.Buffer
	vao      gli.VertexArrayObject
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)
	glfw.SwapInterval(1)
	gli.ClearColor(0, 0, 0, 0)

	// Load texture
	td, err := gli.TextureFromFile("orange.png")
	Panicf(err, "Failed to load image orange.png: %v", err)

	// Set up texture
	texture := gli.CreateTexture(gli.Texture2d)
	gli.ActiveTexture(0)
	gli.BindTexture(texture)
	texture.Data(td)

	// Set up sampler
	var sampler uint32
	gl.GenSamplers(1, &sampler)
	gl.BindSampler(0, sampler) // Set sampler to texture unit 0
	gl.SamplerParameteri(sampler, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.SamplerParameteri(sampler, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.SamplerParameteri(sampler, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.SamplerParameteri(sampler, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

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
	p.vao.Enable(attributes.ByName("vertTexCoord"), p.buffer, textureExtent)

	// Set up uniforms
	uniforms := p.program.Uniforms()
	uniforms.ByName("sampler").Sampler(0) // Set texture unit 0

	return gli.GetError()
}

func (p *Profile) End() {
	gli.SafeDelete(p.program, p.vao, p.fragment, p.vertex, p.buffer)
}

func (p *Profile) Draw(w *glfw.Window) error {
	gli.Clear(gli.ColorBufferBit)
	gli.Draw(p.program, p.vao, triangleObject)
	return gli.GetError()
}

func main() {
	err := window.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
