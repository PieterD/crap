package window

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type EventFunc func(w *Window) error
type CommandFunc func() error

type Window struct {
	EventChan   <-chan Event
	CommandChan chan<- CommandFunc

	win *glfw.Window

	ec chan Event
	cc chan CommandFunc
	dc chan struct{}
}

type WindowConfig struct {
	Title     string
	Width     int
	Height    int
	Resizable bool
}

func DefaultConfig() WindowConfig {
	return WindowConfig{
		Title:     "Glimmer App",
		Width:     640,
		Height:    480,
		Resizable: true,
	}
}

func NewWindow(ef EventFunc, config WindowConfig) error {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		return err
	}
	defer glfw.Terminate()
	doresize := glfw.False
	if config.Resizable {
		doresize = glfw.True
	}
	glfw.WindowHint(glfw.Resizable, doresize)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)

	win, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
	if err != nil {
		return err
	}
	defer win.Destroy()
	win.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		win.Destroy()
		return err
	}

	ec := make(chan Event, 500)
	defer close(ec)
	cc := make(chan CommandFunc, 500)
	dc := make(chan struct{})
	w := &Window{
		EventChan:   ec,
		CommandChan: cc,
		win:         win,
		ec:          ec,
		cc:          cc,
		dc:          dc,
	}
	go func() {
		defer close(cc)
		defer close(dc)
		err := ef(w)
		if err != nil {
			cc <- func() error { return err }
		}
	}()

	win.SetCharCallback(func(iwin *glfw.Window, char rune) {
		select {
		case <-dc:
		case ec <- EventChar{
			Char: char,
		}:
		}
	})

	win.SetCloseCallback(func(iwin *glfw.Window) {
		win.SetShouldClose(false)
		select {
		case ec <- EventClose{}:
		case <-dc:
		}
	})

	win.SetCursorEnterCallback(func(iwin *glfw.Window, entered bool) {
		if entered {
			select {
			case <-dc:
			case ec <- EventCursorEnter{}:
			}
		} else {
			select {
			case <-dc:
			case ec <- EventCursorExit{}:
			}
		}
	})

	win.SetCursorPosCallback(func(iwin *glfw.Window, x float64, y float64) {
		select {
		case <-dc:
		case ec <- EventCursorPos{
			X: x,
			Y: y,
		}:
		}
	})

	win.SetDropCallback(func(iwin *glfw.Window, names []string) {
		select {
		case <-dc:
		case ec <- EventDrop{
			Names: names,
		}:
		}
	})

	win.SetFocusCallback(func(iwin *glfw.Window, focused bool) {
		if focused {
			select {
			case <-dc:
			case ec <- EventFocusGained{}:
			}
		} else {
			select {
			case <-dc:
			case ec <- EventFocusLost{}:
			}
		}
	})

	// Missing: framebuffer callback. No idea what that is.

	win.SetIconifyCallback(func(iwin *glfw.Window, iconified bool) {
		if iconified {
			select {
			case <-dc:
			case ec <- EventIconified{}:
			}
		} else {
			select {
			case <-dc:
			case ec <- EventDeIconified{}:
			}
		}
	})

	win.SetKeyCallback(func(iwin *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			select {
			case <-dc:
			case ec <- EventKeyPress{
				Key:      Key(key),
				ScanCode: scancode,
				Mod:      ModifierKey(mod),
			}:
			}
		case glfw.Release:
			select {
			case <-dc:
			case ec <- EventKeyRelease{
				Key:      Key(key),
				ScanCode: scancode,
				Mod:      ModifierKey(mod),
			}:
			}
		case glfw.Repeat:
			select {
			case <-dc:
			case ec <- EventKeyRepeat{
				Key:      Key(key),
				ScanCode: scancode,
				Mod:      ModifierKey(mod),
			}:
			}
		}
	})

	win.SetMouseButtonCallback(func(iwin *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			select {
			case <-dc:
			case ec <- EventMouseButtonPress{
				Mod:    ModifierKey(mod),
				Button: MouseButton(button),
			}:
			}
		case glfw.Release:
			select {
			case <-dc:
			case ec <- EventMouseButtonRelease{
				Mod:    ModifierKey(mod),
				Button: MouseButton(button),
			}:
			}
		}
	})

	// Missing: Pos

	win.SetRefreshCallback(func(iwin *glfw.Window) {
		select {
		case <-dc:
		case ec <- EventRefresh{}:
		}
	})

	win.SetScrollCallback(func(iwin *glfw.Window, x float64, y float64) {
		select {
		case <-dc:
		case ec <- EventScroll{
			X: x,
			Y: y,
		}:
		}
	})

	win.SetSizeCallback(func(iwin *glfw.Window, width int, height int) {
		select {
		case <-dc:
		case ec <- EventSize{
			Width:  width,
			Height: height,
		}:
		}
	})

	return w.loop()
}

func (w *Window) loop() error {
	for {
		glfw.PollEvents()
		select {
		case command, ok := <-w.cc:
			if !ok {
				return nil
			}
			err := command()
			if err != nil {
				return err
			}
		case <-time.After(time.Millisecond):
		}
	}

	return nil
}

func (w *Window) Swap() {
	w.cc <- func() error {
		w.win.SwapBuffers()
		return nil
	}
}

func (w *Window) Event(e Event) {
	go func() {
		w.ec <- e
	}()
}

func (w *Window) Future(cf CommandFunc) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		w.cc <- func() error {
			defer close(ch)
			return cf()
		}
	}()
	return ch
}
