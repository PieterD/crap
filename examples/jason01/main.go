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
	shader, err := glimmer.CreateShader(glimmer.VertexShader, []byte(`
	#version 330
	layout(location = 0) in vec4 position;
	ajepgoaejg
	void main() {
		gl_Position = position;
	}
	`))
	if err != nil {
		fmt.Printf("Error compiling shader: %v\n", err)
	}
	shader.Delete()
}

func main() {
	glimmer.Run(Profile{})
}
