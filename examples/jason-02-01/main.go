package main

import (
	"fmt"

	"github.com/PieterD/crap/glimmer"
	. "github.com/PieterD/crap/glimmer/pan"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	glimmer.DefaultProfile
	vertex   *glimmer.Shader
	fragment *glimmer.Shader
	program  *glimmer.Program
	pointer  *glimmer.ArrayPointer
}

var vertexShaderText = `
#version 330
layout(location = 0) in vec4 position;
void main() {
	gl_Position = position;
}
`

var fragmentShaderText = `
#version 330
out vec4 outputColor;
void main() {
	float lerpValue = gl_FragCoord.y / 500.0f;
	outputColor = mix(vec4(1.0f, 1.0f, 1.0f, 1.0f), vec4(0.2f, 0.2f, 0.2f, 1.0f), lerpValue);
}
`

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	defer Recover(&err)

	glfw.SwapInterval(1)

	p.vertex, err = glimmer.CreateVertexShader(vertexShaderText)
	Panicf(err, "Error compiling vertex shader: %v", err)

	p.fragment, err = glimmer.CreateFragmentShader(fragmentShaderText)
	Panicf(err, "Error compiling fragment shader: %v", err)

	p.program, err = glimmer.CreateProgram(p.vertex, p.fragment)
	Panicf(err, "Error linking program: %v", err)

	p.pointer = glimmer.CreateBuffer().FloatData(vertexData).Pointer(4, false, 0, 0)

	return glimmer.GetError()
}

func (p *Profile) End() {
	p.program.Delete()
	p.fragment.Delete()
	p.vertex.Delete()
	p.pointer.Buffer().Delete()
}

func (p *Profile) Draw(w *glfw.Window) error {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	p.program.Use()
	p.program.AttributeByName("position", p.pointer)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	return nil
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
