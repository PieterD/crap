package grid

import "github.com/go-gl/glfw/v3.2/glfw"

type mouseTranslator struct {
	lastx float64
	lasty float64
}

func (trans *mouseTranslator) Pos(posx, posy float64) (MouseEvent, bool) {
	return MouseEvent{
		x: int(posx),
		y: int(posy),
	}, true
}

func (trans *mouseTranslator) Button(button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) (MouseEvent, bool) {
	return MouseEvent{}, true
}

type MouseEvent struct {
	x, y int
}

func (e MouseEvent) X() int {
	return e.x
}

func (e MouseEvent) Y() int {
	return e.y
}
