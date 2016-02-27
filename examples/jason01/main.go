package main

import (
	"fmt"

	"github.com/PieterD/crap/glimmer"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Profile struct {
	glimmer.DefaultProfile
}

func (p Profile) PostCreation(w *glfw.Window) {
	p.DefaultProfile.PostCreation(w)
	vertexShader, err := glimmer.CreateShader(glimmer.VertexShader, `
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
	defer vertexShader.Delete()

	fragmentShader, err := glimmer.CreateShader(glimmer.FragmentShader, `
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
	defer fragmentShader.Delete()

	program, err := glimmer.CreateProgram(vertexShader, fragmentShader)
	if err != nil {
		fmt.Printf("Error linking program: %v\n", err)
		return
	}
	defer program.Delete()

	fmt.Printf("shader: %v\n", glimmer.GetError())
}

func main() {
	glimmer.Run(Profile{})
}
