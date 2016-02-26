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
	shader, err := glimmer.CreateShader(glimmer.VertexShader, `
	#version 330
	layout(location = 0) in vec4 position;
	void main() {
		gl_Position = position;
	}
	`)
	if err != nil {
		fmt.Printf("Error compiling shader: %v\n", err)
		return
	}
	shader.Delete()
	fmt.Printf("shader: %v\n", glimmer.GetError())
}

func main() {
	glimmer.Run(Profile{})
}
