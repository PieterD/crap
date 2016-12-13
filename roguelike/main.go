package main

import (
	"runtime"

	_ "image/png"

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
	1.0, 0.0,
	0.75, -0.75, 0.0, 1.0,
	0.0, 1.0, 0.0, 1.0,
	1.0, 1.0,
	-0.75, -0.75, 0.0, 1.0,
	0.0, 0.0, 1.0, 1.0,
	0.0, 1.0,
}

var vertexShaderText = `
#version 110
attribute vec4 position;
attribute vec4 color;
attribute vec2 texCoord;
varying vec4 theColor;
varying vec2 theTexCoord;
void main() {
	gl_Position = position;
	theColor = color;
	theTexCoord = texCoord;
}
`

var fragmentShaderText = `
#version 110
varying vec4 theColor;
varying vec2 theTexCoord;
uniform sampler2D tex;
void main() {
	gl_FragColor = theColor * texture2D(tex, theTexCoord);
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
		gli.VAOStride(10))
	vao.Enable(4, vbo, program.Attrib("color"),
		gli.VAOStride(10),
		gli.VAOOffset(4))
	vao.Enable(2, vbo, program.Attrib("texCoord"),
		gli.VAOStride(10),
		gli.VAOOffset(8))

	img, err := gli.LoadImage("resources/rogue_yun_16x16.png")
	Panic(err)
	texture := gli.NewTexture(img,
		gli.TextureFilter(gli.LINEAR, gli.LINEAR),
		gli.TextureWrap(gli.CLAMP_TO_EDGE, gli.CLAMP_TO_EDGE))
	program.Use()
	gl.Uniform1i(program.Uniform("tex").Location(), 0)

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		program.Use()
		vao.Use()
		texture.Use(0)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
