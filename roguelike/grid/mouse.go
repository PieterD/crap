package grid

import (
	"fmt"
	"image"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type mouseTranslator struct {
	grid   *Grid
	last   image.Point
	down   bool
	drag   bool
	button MouseButton
	start  image.Point
}

func newMouseTranslator(grid *Grid) *mouseTranslator {
	return &mouseTranslator{
		grid: grid,
		last: image.Point{X: -1, Y: -1},
	}
}

func (trans *mouseTranslator) Pos(posx, posy float64) (MouseEvent, bool) {
	x := (int(posx) - trans.grid.padx)
	y := (int(posy) - trans.grid.pady)

	if x < 0 || y < 0 {
		return MouseEvent{}, false
	}

	x /= trans.grid.runewidth
	y /= trans.grid.runeheight

	if trans.last.X == x && trans.last.Y == y {
		return MouseEvent{}, false
	}
	if x >= trans.grid.cols || y >= trans.grid.rows {
		return MouseEvent{}, false
	}

	trans.last.X = x
	trans.last.Y = y

	if trans.down && !trans.drag {
		trans.drag = true
		return MouseEvent{
			pos:    trans.last,
			drag:   StartDrag,
			start:  trans.start,
			button: trans.button,
		}, true
	}

	if trans.down && trans.drag {
		return MouseEvent{
			pos:    trans.last,
			drag:   ContinueDrag,
			start:  trans.start,
			button: trans.button,
		}, true
	}

	return MouseEvent{
		pos: trans.last,
	}, true
}

type MouseButton int

const (
	MouseButtonLeft   MouseButton = MouseButton(glfw.MouseButtonLeft)
	MouseButtonRight  MouseButton = MouseButton(glfw.MouseButtonRight)
	MouseButtonMiddle MouseButton = MouseButton(glfw.MouseButtonMiddle)
)

func (trans *mouseTranslator) Button(gbutton glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) (MouseEvent, bool) {
	fmt.Printf("button %#v %#v\n", gbutton, action)
	if gbutton != glfw.MouseButtonLeft && gbutton != glfw.MouseButtonRight && gbutton != glfw.MouseButtonMiddle {
		return MouseEvent{}, false
	}
	button := MouseButton(gbutton)
	if action == glfw.Press {
		if trans.down {
			return MouseEvent{}, false
		}
		trans.down = true
		trans.drag = false
		trans.button = button
		trans.start = trans.last
		return MouseEvent{}, false
	}
	if action == glfw.Release {
		if !trans.down || trans.button != button {
			return MouseEvent{}, false
		}
		var e MouseEvent
		if trans.drag {
			e = MouseEvent{
				pos:    trans.last,
				drag:   EndDrag,
				start:  trans.start,
				button: trans.button,
			}
		} else {
			e = MouseEvent{
				pos:    trans.last,
				press:  true,
				button: trans.button,
			}
		}
		trans.down = false
		trans.drag = false
		return e, true
	}
	return MouseEvent{}, false
}

type MouseEvent struct {
	pos    image.Point
	drag   DragState
	start  image.Point
	press  bool
	button MouseButton
}

type DragEvent struct {
	From   image.Point
	To     image.Point
	State  DragState
	Button MouseButton
}

type DragState int

const (
	NoDrag DragState = iota
	StartDrag
	ContinueDrag
	EndDrag
)

func (e MouseEvent) Pos() image.Point {
	return e.pos
}

func (e MouseEvent) Click() (MouseButton, bool) {
	if e.press {
		return e.button, true
	}
	return 0, false
}

func (e MouseEvent) Drag() (DragEvent, bool) {
	if e.drag == NoDrag {
		return DragEvent{}, false
	}
	return DragEvent{
		From:   e.start,
		To:     e.pos,
		State:  e.drag,
		Button: e.button,
	}, true
}
