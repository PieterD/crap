package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	//glfw.SwapInterval(1)
	//getSwapSpeed(window)

	window.SetSizeCallback(func(w *glfw.Window, width int, height int) {
		fmt.Printf("resize: %d, %d\n", width, height)
		gl.Viewport(0, 0, int32(width), int32(height))
	})
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		fmt.Printf("mouse: %f, %f\n", xpos, ypos)
	})
	window.SetFocusCallback(func(w *glfw.Window, focused bool) {
		fmt.Printf("focus: %t\n", focused)
	})
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		fmt.Printf("key\n")
	})
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		fmt.Printf("mouse\n")
	})
	window.SetCharCallback(func(w *glfw.Window, char rune) {
		fmt.Printf("char: '%c'\n", char)
	})

	if err := gl.Init(); err != nil {
		panic(err)
	}

	for !window.ShouldClose() {
		draw()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, 1)
	gl.Color4f(1, 1, 1, 1)
	gl.Begin(gl.QUADS)
	gl.Vertex3f(-0.9, -0.9, 0)
	gl.Vertex3f(0.9, -0.9, 0)
	gl.Vertex3f(0.9, 0.9, 0)
	gl.Vertex3f(-0.9, 0.9, 0)
	gl.End()
}

func getSwapSpeed(window *glfw.Window) (float64, bool) {
	window.SwapBuffers()
	window.SwapBuffers()
	window.SwapBuffers()
	i := 0
	sum := 0.0
	for !window.ShouldClose() && i < 51 {
		glfw.SetTime(0)
		window.SwapBuffers()
		glfw.PollEvents()
		t := glfw.GetTime()
		fmt.Printf("time: %f = %f\n", t, 1/t)
		sum += t
		i++
	}
	sum /= float64(i)
	if sum < 0.002 {
		return 0, false
	}
	fmt.Printf("Average: %fms = %f fps\n", sum, 1/sum)
	sum *= 0.95
	fmt.Printf("Reporting: %fms\n", sum)
	return sum, true
}
