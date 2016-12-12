package main

import (
	"runtime"

	"github.com/PieterD/crap/roguelike/gli"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	1.0, 0.0, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	0.0, 1.0, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
	0.0, 0.0, 1.0, 1.0,
}

var vertexShaderText = `
#version 110
attribute vec4 position;
attribute vec4 color;
varying vec4 theColor;
void main() {
	gl_Position = position;
	theColor = color;
}
`

var fragmentShaderText = `
#version 110
varying vec4 theColor;
void main() {
	gl_FragColor = theColor;
}
`

func main() {
	err := glfw.Init()
	Panic(err)
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(800, 600, "Roguelike", nil, nil)
	defer window.Destroy()
	Panic(err)
	window.MakeContextCurrent()

	err = gl.Init()
	Panic(err)

	program, err := gli.NewProgram(vertexShaderText, fragmentShaderText)
	Panic(err)
	defer program.Delete()

	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	vbo, err := gli.NewBuffer(vertexData,
		gli.BufferAccessFrequency(gli.DYNAMIC))
	Panic(err)
	defer vbo.Delete()

	vao.Enable(4, vbo, program.Attrib("position"),
		gli.VAOStride(8))
	vao.Enable(4, vbo, program.Attrib("color"),
		gli.VAOStride(8),
		gli.VAOOffset(4))

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		program.Use()
		vao.Use()
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
