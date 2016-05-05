package win

import (
	"fmt"
	"github.com/go-gl/glfw/v3.1/glfw"
	"runtime"
	"time"
)

type Window struct {
	glfwWindow *glfw.Window
}
type WindowConfig struct {
	Title     string
	Width     int
	Height    int
	Resizable bool
	TickSpeed time.Duration
	MinFps    int
}

func DefaultConfig() WindowConfig {
	return WindowConfig{
		Title:     "Glimmer App",
		Width:     640,
		Height:    480,
		Resizable: true,
		TickSpeed: time.Millisecond * 5,
		MinFps:    1,
	}
}

type EventHandler interface {
	Init() error
	Tick() error
	Anim(ticks int) error
	Draw() error
	Drop(ticks int) error
	EventChar(char rune)
	EventClose()
	EventRefresh()
	EventDragDrop(names []string)
	EventCursorEnter()
	EventCursorExit()
	EventCursorPos(x, y float64)
	EventFocusGained()
	EventFocusLost()
	EventIconified()
	EventDeiconified()
	EventKeyPress(key Key, scancode int, mod ModifierKey)
	EventKeyRepeat(key Key, scancode int, mod ModifierKey)
	EventKeyRelease(key Key, scancode int, mod ModifierKey)
	EventMouseButtonPress(button MouseButton, mod ModifierKey)
	EventMouseButtonRelease(button MouseButton, mod ModifierKey)
	EventScroll(x, y float64)
	EventSize(width, height int)
}

func New(config WindowConfig, handler EventHandler) error {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		return fmt.Errorf("Failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate()
	doresize := glfw.False
	if config.Resizable {
		doresize = glfw.True
	}
	glfw.WindowHint(glfw.Resizable, doresize)

	win, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
	if err != nil {
		return err
	}
	defer win.Destroy()
	callbacks(win, handler)

	win.MakeContextCurrent()
	glfw.SwapInterval(1)

	handler.Init()

	frameSkip := 0
	var timeAccumulated time.Duration = 0
	for !win.ShouldClose() {
		ticks := 0
		start := time.Now()
		//TODO: This way of frame skipping is bad.
		if timeAccumulated > time.Second/time.Duration(config.MinFps) {
			frameSkip++
		} else if frameSkip > 0 {
			frameSkip--
		}
		if timeAccumulated < config.TickSpeed {
			glfw.PollEvents()
		}
		skipped := 0
		for timeAccumulated >= config.TickSpeed {
			timeAccumulated -= config.TickSpeed
			if skipped < frameSkip {
				skipped++
				continue
			}
			ticks++
			glfw.PollEvents()
			err := handler.Tick()
			if err != nil {
				return fmt.Errorf("Error ticking game state: %v", err)
			}
		}
		if skipped > 0 {
			err := handler.Drop(skipped)
			if err != nil {
				return fmt.Errorf("Error skipping %d frames: %v", skipped, err)
			}
		}
		if ticks > 0 {
			err := handler.Anim(ticks)
			if err != nil {
				return fmt.Errorf("Error animating %d ticks: %v", ticks, err)
			}
		}
		err := handler.Draw()
		if err != nil {
			return fmt.Errorf("Error drawing: %v", err)
		}
		win.SwapBuffers()
		timeAccumulated += time.Now().Sub(start)
	}

	return nil
}

func callbacks(win *glfw.Window, handler EventHandler) {
	win.SetCharCallback(func(iwin *glfw.Window, char rune) {
		handler.EventChar(char)
	})

	win.SetCloseCallback(func(iwin *glfw.Window) {
		handler.EventClose()
	})

	win.SetCursorEnterCallback(func(iwin *glfw.Window, entered bool) {
		if entered {
			handler.EventCursorEnter()
		} else {
			handler.EventCursorEnter()
		}
	})

	win.SetCursorPosCallback(func(iwin *glfw.Window, x float64, y float64) {
		handler.EventCursorPos(x, y)
	})

	win.SetDropCallback(func(iwin *glfw.Window, names []string) {
		handler.EventDragDrop(names)
	})

	win.SetFocusCallback(func(iwin *glfw.Window, focused bool) {
		if focused {
			handler.EventFocusGained()
		} else {
			handler.EventFocusLost()
		}
	})

	// Missing: framebuffer callback. No idea what that is.

	win.SetIconifyCallback(func(iwin *glfw.Window, iconified bool) {
		if iconified {
			handler.EventIconified()
		} else {
			handler.EventDeiconified()
		}
	})

	win.SetKeyCallback(func(iwin *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			handler.EventKeyPress(Key(key), scancode, ModifierKey(mod))
		case glfw.Release:
			handler.EventKeyRelease(Key(key), scancode, ModifierKey(mod))
		case glfw.Repeat:
			handler.EventKeyRepeat(Key(key), scancode, ModifierKey(mod))
		}
	})

	win.SetMouseButtonCallback(func(iwin *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			handler.EventMouseButtonPress(MouseButton(button), ModifierKey(mod))
		case glfw.Release:
			handler.EventMouseButtonRelease(MouseButton(button), ModifierKey(mod))
		}
	})

	// Missing: Pos

	win.SetRefreshCallback(func(iwin *glfw.Window) {
		handler.EventRefresh()
	})

	win.SetScrollCallback(func(iwin *glfw.Window, x float64, y float64) {
		handler.EventScroll(x, y)
	})

	win.SetSizeCallback(func(iwin *glfw.Window, width int, height int) {
		handler.EventSize(width, height)
	})
}
