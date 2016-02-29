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
	PreCreation() error
	PostCreation(w *glfw.Window) error
	EventFocus(w *glfw.Window, focused bool)
	EventResize(w *glfw.Window, width int, height int)
	EventMousePos(w *glfw.Window, x float64, y float64)
	EventMouseKey(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey)
	EventKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
	EventRune(w *glfw.Window, char rune)
	Draw(w *glfw.Window) error
	Swap(w *glfw.Window) error
	End()
}

type DefaultProfile struct {
}

func (p DefaultProfile) InitialSize() (int, int) {
	return 640, 480
}
func (p DefaultProfile) InitialTitle() string {
	return "Default Title"
}
func (p DefaultProfile) PreCreation() error {
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	return nil
}
func (p DefaultProfile) PostCreation(w *glfw.Window) error {
	glfw.SwapInterval(1)
	return nil
}
func (p DefaultProfile) EventFocus(w *glfw.Window, focused bool) {}
func (p DefaultProfile) EventResize(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
func (p DefaultProfile) EventMousePos(w *glfw.Window, x float64, y float64) {}
func (p DefaultProfile) EventMouseKey(w *glfw.Window, key glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
}
func (p DefaultProfile) EventKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
}
func (p DefaultProfile) EventRune(w *glfw.Window, char rune) {}
func (p DefaultProfile) Draw(w *glfw.Window) error {
	return nil
}
func (p DefaultProfile) Swap(w *glfw.Window) error {
	w.SwapBuffers()
	glfw.PollEvents()
	return nil
}
func (p DefaultProfile) End() {}

func Run(p Profile) error {
	defer p.End()
	err := glfw.Init()
	if err != nil {
		return err
	}
	defer glfw.Terminate()

	err = p.PreCreation()
	if err != nil {
		return err
	}

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

	err = p.PostCreation(w)
	if err != nil {
		return err
	}

	w.SetFocusCallback(p.EventFocus)
	w.SetSizeCallback(p.EventResize)
	w.SetCursorPosCallback(p.EventMousePos)
	w.SetMouseButtonCallback(p.EventMouseKey)
	w.SetKeyCallback(p.EventKey)
	w.SetCharCallback(p.EventRune)

	for !w.ShouldClose() {
		err = p.Draw(w)
		if err != nil {
			return err
		}
		err = p.Swap(w)
		if err != nil {
			return err
		}
	}
	return nil
}
