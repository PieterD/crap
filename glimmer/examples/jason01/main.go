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

func (p *Profile) PostCreation(w *glfw.Window) {
	p.DefaultProfile.PostCreation(w)
	var err error
	p.vertexShader, err = glimmer.CreateShader(glimmer.VertexShader, `
	#version 330
	layout(location = 0) in vec4 position;
	void main() {
		gl_Position = position;
	}
	`)
	if err != nil {
		fmt.Printf("Error compiling vertex shader: %v\n", err)
		return
	}

	p.fragmentShader, err = glimmer.CreateShader(glimmer.FragmentShader, `
	#version 330
	out vec4 outputColor;
	void main() {
		outputColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
	}
	`)
	if err != nil {
		fmt.Printf("Error compiling fragment shader: %v\n", err)
		return
	}

	p.program, err = glimmer.CreateProgram(p.vertexShader, p.fragmentShader)
	if err != nil {
		fmt.Printf("Error linking program: %v\n", err)
		return
	}

	p.vertexArray = glimmer.CreateVertexArray()
	p.buffer = glimmer.CreateBuffer()
	p.buffer.UseAsArrayBuffer()
	p.buffer.FloatData([]float32{
		0.75, 0.75, 0.0, 1.0,
		0.75, -0.75, 0.0, 1.0,
		-0.75, -0.75, 0.0, 1.0,
	})
	p.arrayPointer = p.buffer.Pointer(4, false, 0, 0)

	fmt.Printf("initialization: %v\n", glimmer.GetError())
}

func (p *Profile) End() {
	p.program.Delete()
	p.fragmentShader.Delete()
	p.vertexShader.Delete()
	p.vertexArray.Delete()
	p.buffer.Delete()
}

func (p *Profile) Draw(w *glfw.Window) {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	p.program.Use()
	p.vertexArray.Enable(0, p.arrayPointer)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func main() {
	glimmer.Run(&Profile{})
}
