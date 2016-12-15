package main

import (
	"fmt"
	"runtime"

	_ "image/png"

	"github.com/PieterD/crap/roguelike/gli"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

var vertexData = []uint8{
	1, 0, 0, 1,
	2, 0, 0, 1,
	2, 1, 0, 1,
	1, 1, 0, 1,
}

var colorData = []float32{
	0.0, 0.0, 0.0,
	1.0, 1.0, 1.0,
	1.0, 0.0, 0.0,
	0.0, 1.0, 0.0,
	0.0, 0.0, 1.0,
}

var vertexShaderText = `
#version 110
attribute vec2 position;
attribute float foreColor;
attribute float backColor;
attribute vec2 texCoord;
uniform vec3 colorData[5];
uniform vec2 runeSize;
varying vec3 theForeColor;
varying vec3 theBackColor;
varying vec2 theTexCoord;
void main() {
	gl_Position = vec4(position, 0.0, 1.0);
	theForeColor = colorData[int(foreColor)];
	theBackColor = colorData[int(backColor)];
	theTexCoord = vec2(texCoord.x / runeSize.x, texCoord.y / runeSize.y);
}
`

var fragmentShaderText = `
#version 110
varying vec3 theForeColor;
varying vec3 theBackColor;
varying vec2 theTexCoord;
uniform sampler2D tex;
void main() {
	vec4 texColor = texture2D(tex, theTexCoord);
	gl_FragColor = vec4(mix(theBackColor, theForeColor, texColor.a), 1.0);
}
`

func main() {
	width := 800
	height := 600
	// Initialize glfw and create window
	err := glfw.Init()
	Panic(err)
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(width, height, "Roguelike", nil, nil)
	defer window.Destroy()
	Panic(err)
	window.MakeContextCurrent()

	// Initialize opengl
	err = gl.Init()
	Panic(err)

	window.SetSizeCallback(func(win *glfw.Window, w int, h int) {
		fmt.Printf("resize\n")
		width = w
		height = h
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	// Create shaders and program
	program, err := gli.NewProgram(vertexShaderText, fragmentShaderText)
	Panic(err)
	defer program.Delete()

	// Load and initialize texture
	img, err := gli.LoadImage("resources/rogue_yun_16x16.png")
	Panic(err)
	texture, err := gli.NewTexture(img,
		gli.TextureFilter(gli.LINEAR, gli.LINEAR),
		gli.TextureWrap(gli.CLAMP_TO_EDGE, gli.CLAMP_TO_EDGE))
	Panic(err)
	defer texture.Delete()

	// Create Vertex ArrayObject
	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	// Create grid, and the position and index buffers
	grid, err := NewGrid(16, 16, texture.Size().X, texture.Size().Y)
	Panic(err)
	grid.Resize(width, height)
	vData, vIndex := grid.Coordinates()
	posvbo, err := gli.NewBuffer(vData)
	Panic(err)
	defer posvbo.Delete()
	idxvbo, err := gli.NewBuffer(vIndex, gli.BufferElementArray())
	Panic(err)
	defer idxvbo.Delete()

	grid.Set(0, 0, 1, 1, 0)
	grid.Set(49, 0, 2, 1, 0)
	grid.Set(0, 1, 2, 1, 0)
	vbo, err := gli.NewBuffer(grid.VertexData(),
		gli.BufferAccessFrequency(gli.DYNAMIC))
	Panic(err)
	defer vbo.Delete()

	// Set up VAO
	vao.Enable(2, posvbo, program.Attrib("position"))
	vao.Enable(2, vbo, program.Attrib("texCoord"),
		gli.VAOStride(4))
	vao.Enable(1, vbo, program.Attrib("foreColor"),
		gli.VAOStride(4), gli.VAOOffset(2))
	vao.Enable(1, vbo, program.Attrib("backColor"),
		gli.VAOStride(4), gli.VAOOffset(3))

	// Set uniforms
	program.Uniform("tex").SetSampler(1)
	program.Uniform("colorData[0]").SetFloat(colorData...)
	program.Uniform("runeSize").SetFloat(float32(grid.RuneSize().X), float32(grid.RuneSize().Y))

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	program.Use()
	vao.Use()
	texture.Use(1)
	idxvbo.Use()
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawElements(gl.TRIANGLES, grid.Vertices(), gl.UNSIGNED_INT, gl.PtrOffset(0))
		//fmt.Printf("draw\n")
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
