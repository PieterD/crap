package glimmer

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()
}

type Profile interface {
	InitialSize() (width int, height int)
	InitialTitle() string
	PreCreation()
	PostCreation(w *glfw.Window)
	EventFocus(w *glfw.Window, focused bool)
	EventResize(w *glfw.Window, width int, height int)
	EventMousePos(w *glfw.Window, x float64, y float64)
	EventMouseKey(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey)
	EventKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
	EventRune(w *glfw.Window, char rune)
	Draw(w *glfw.Window)
	Cycle(w *glfw.Window)
}

type DefaultProfile struct {
}

func (p DefaultProfile) InitialSize() (int, int) {
	return 640, 480
}
func (p DefaultProfile) InitialTitle() string {
	return "Default Title"
}
func (p DefaultProfile) PreCreation() {
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
}
func (p DefaultProfile) PostCreation(w *glfw.Window) {
	glfw.SwapInterval(1)
}
func (p DefaultProfile) EventFocus(w *glfw.Window, focused bool)            {}
func (p DefaultProfile) EventResize(w *glfw.Window, width int, height int)  {}
func (p DefaultProfile) EventMousePos(w *glfw.Window, x float64, y float64) {}
func (p DefaultProfile) EventMouseKey(w *glfw.Window, key glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
}
func (p DefaultProfile) EventKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
}
func (p DefaultProfile) EventRune(w *glfw.Window, char rune) {}
func (p DefaultProfile) Draw(w *glfw.Window)                 {}
func (p DefaultProfile) Cycle(w *glfw.Window) {
	p.Draw(w)
	w.SwapBuffers()
	glfw.PollEvents()
}

func Run(p Profile) error {
	err := glfw.Init()
	if err != nil {
		return err
	}
	defer glfw.Terminate()

	p.PreCreation()

	width, height := p.InitialSize()
	title := p.InitialTitle()
	w, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return err
	}
	defer w.Destroy()

	w.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		return err
	}

	p.PostCreation(w)

	w.SetFocusCallback(p.EventFocus)
	w.SetSizeCallback(p.EventResize)
	w.SetCursorPosCallback(p.EventMousePos)
	w.SetMouseButtonCallback(p.EventMouseKey)
	w.SetKeyCallback(p.EventKey)
	w.SetCharCallback(p.EventRune)

	for !w.ShouldClose() {
		p.Cycle(w)
	}
	return nil
}
