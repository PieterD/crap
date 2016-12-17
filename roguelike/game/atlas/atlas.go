package atlas

import (
	"image"

	"github.com/PieterD/crap/roguelike/game/atlas/feature"
	"github.com/PieterD/crap/roguelike/grid"
)

type Glyph struct {
	Code int
	Fore grid.Color
	Back grid.Color
}

func Translate(screen image.Rectangle, center image.Point, atlas image.Rectangle) image.Point {
	tl := center.Sub(screen.Max.Div(2))
	return tl
}

type Atlas struct {
	cells  map[image.Point]Cell
	bounds image.Rectangle
}

func New() *Atlas {
	atlas := &Atlas{
		cells: make(map[image.Point]Cell),
		bounds: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: 100, Y: 100},
		},
	}
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			atlas.cells[image.Point{X: x, Y: y}] = Cell{feature: feature.Floor}
		}
	}
	for i := 0; i < 100; i++ {
		atlas.cells[image.Point{X: i, Y: 0}] = Cell{feature: feature.Wall}
		atlas.cells[image.Point{X: i, Y: 99}] = Cell{feature: feature.Wall}
		atlas.cells[image.Point{X: 0, Y: i}] = Cell{feature: feature.Wall}
		atlas.cells[image.Point{X: 99, Y: i}] = Cell{feature: feature.Wall}
	}
	for x := 0; x < 100; x += 10 {
		for y := 0; y < 100; y += 10 {
			atlas.cells[image.Point{X: x, Y: y}] = Cell{feature: feature.Wall}
		}
	}
	return atlas
}

func (atlas *Atlas) Bounds() image.Rectangle {
	return atlas.bounds
}

func (atlas *Atlas) Glyph(p image.Point) Glyph {
	return atlas.cells[p].Glyph()
}
