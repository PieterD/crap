package main

import (
	"fmt"

	"github.com/PieterD/crap/glimmer"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	glimmer.DefaultProfile
	vertexShader   *glimmer.Shader
	fragmentShader *glimmer.Shader
	program        *glimmer.Program
	vertexArray    *glimmer.VertexArray
	buffer         *glimmer.Buffer
	arrayPointer   *glimmer.ArrayPointer
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
		outputColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
	}
`

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
}

func (p *Profile) PostCreation(w *glfw.Window) (err error) {
	p.DefaultProfile.PostCreation(w)
	p.vertexShader, err = glimmer.CreateShader(glimmer.VertexShader, vertexShaderText)
	if err != nil {
		return fmt.Errorf("Error compiling vertex shader: %v", err)
	}

	p.fragmentShader, err = glimmer.CreateShader(glimmer.FragmentShader, fragmentShaderText)
	if err != nil {
		return fmt.Errorf("Error compiling fragment shader: %v\n", err)
	}

	p.program, err = glimmer.CreateProgram(p.vertexShader, p.fragmentShader)
	if err != nil {
		return fmt.Errorf("Error linking program: %v\n", err)
	}

	p.vertexArray = glimmer.CreateVertexArray()
	p.buffer = glimmer.CreateBuffer()
	p.buffer.UseAsArrayBuffer()
	p.buffer.FloatData(vertexData)
	p.arrayPointer = p.buffer.Pointer(4, false, 0, 0)

	return glimmer.GetError()
}

func (p *Profile) End() {
	p.program.Delete()
	p.fragmentShader.Delete()
	p.vertexShader.Delete()
	p.vertexArray.Delete()
	p.buffer.Delete()
}

func (p *Profile) Draw(w *glfw.Window) error {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	p.program.Use()
	p.vertexArray.Enable(0, p.arrayPointer)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	return nil
}

func main() {
	err := glimmer.Run(&Profile{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
