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

var vertexTexCoords = []float32{
	1.0 / 16.0, 0.0 / 16.0,
	2.0 / 16.0, 0.0 / 16.0,
	2.0 / 16.0, 1.0 / 16.0,
	1.0 / 16.0, 1.0 / 16.0,
}

var vertexColors = []uint8{
	255, 0, 0,
	0, 255, 0,
	0, 0, 255,
	255, 255, 255,
}

var vertexShaderText = `
#version 110
attribute vec2 position;
attribute vec3 color;
attribute vec2 texCoord;
varying vec4 theColor;
varying vec2 theTexCoord;
void main() {
	gl_Position = vec4(position, 0.0, 1.0);
	theColor = vec4(color, 1.0);
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

	// Set texture unit
	program.Uniform("tex").SetInt(1)

	// Create Vertex ArrayObject
	vao, err := gli.NewVAO()
	Panic(err)
	defer vao.Delete()

	// Create Buffer of vertex colors
	colvbo, err := gli.NewBuffer(vertexColors)
	Panic(err)
	defer colvbo.Delete()

	// Create Buffer of vertex texture coordinates
	texvbo, err := gli.NewBuffer(vertexTexCoords)
	Panic(err)
	defer texvbo.Delete()

	// Create grid, and the position and index buffers
	grid, err := NewGrid(16, 16, texture.Size().X, texture.Size().Y)
	Panic(err)
	grid.Resize(width, height)
	vData, vIndex := grid.Coordinates()
	posvbo, err := gli.NewBuffer(vData)
	Panic(err)
	idxvbo, err := gli.NewBuffer(vIndex, gli.BufferElementArray())
	Panic(err)
	defer posvbo.Delete()

	// Set up VAO
	vao.Enable(2, posvbo, program.Attrib("position"))
	vao.Enable(3, colvbo, program.Attrib("color"), gli.VAONormalize())
	vao.Enable(2, texvbo, program.Attrib("texCoord"))

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	program.Use()
	vao.Use()
	texture.Use(1)
	idxvbo.Use()
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
		fmt.Printf("draw\n")
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
