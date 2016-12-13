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
	// Initialize glfw and create window
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

	// Initialize opengl
	err = gl.Init()
	Panic(err)

	// Create shaders and program
	program, err := gli.NewProgram(vertexShaderText, fragmentShaderText)
	Panic(err)
	defer program.Delete()

	// Create Vertex ArrayObject
	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	// Create Buffer from vertex data
	vbo, err := gli.NewBuffer(vertexData,
		gli.BufferAccessFrequency(gli.DYNAMIC))
	Panic(err)
	defer vbo.Delete()

	// Set up VAO
	vao.Enable(4, vbo, program.Attrib("position"),
		gli.VAOStride(10))
	vao.Enable(4, vbo, program.Attrib("color"),
		gli.VAOStride(10),
		gli.VAOOffset(4))
	vao.Enable(2, vbo, program.Attrib("texCoord"),
		gli.VAOStride(10),
		gli.VAOOffset(8))

	// Load and initialize texture
	img, err := gli.LoadImage("resources/rogue_yun_16x16.png")
	Panic(err)
	texture, err := gli.NewTexture(img,
		gli.TextureFilter(gli.LINEAR, gli.LINEAR),
		gli.TextureWrap(gli.CLAMP_TO_EDGE, gli.CLAMP_TO_EDGE))
	Panic(err)
	defer texture.Delete()

	// Set texture unit
	program.Uniform("tex").SetInt(1)

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	program.Use()
	vao.Use()
	texture.Use(1)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
