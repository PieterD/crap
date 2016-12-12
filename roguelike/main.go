package main

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

var vertexData = []float32{
	0.75, 0.75, 0.0, 1.0,
	0.75, -0.75, 0.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
}

var colorData = []float32{
	1.0, 0.0, 0.0, 1.0,
	0.0, 1.0, 0.0, 1.0,
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
	//glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(800, 600, "Roguelike", nil, nil)
	defer window.Destroy()
	Panic(err)
	window.MakeContextCurrent()

	err = gl.Init()
	Panic(err)

	program, err := newProgram(vertexShaderText, fragmentShaderText)
	Panic(err)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var posVbo uint32
	gl.GenBuffers(1, &posVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, posVbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData)*4, gl.Ptr(vertexData), gl.STATIC_DRAW)
	posLoc := uint32(gl.GetAttribLocation(program, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(posLoc)
	gl.VertexAttribPointer(posLoc, 4, gl.FLOAT, false, 0, gl.PtrOffset(0))

	var colVbo uint32
	gl.GenBuffers(1, &colVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, colVbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(colorData)*4, gl.Ptr(colorData), gl.STATIC_DRAW)
	colLoc := uint32(gl.GetAttribLocation(program, gl.Str("color\x00")))
	gl.EnableVertexAttribArray(colLoc)
	gl.VertexAttribPointer(colLoc, 4, gl.FLOAT, false, 0, gl.PtrOffset(0))

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
