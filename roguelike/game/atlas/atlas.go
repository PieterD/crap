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
	if screen.Max.X > atlas.Max.X {
		tl.X = -(screen.Max.X - atlas.Max.X) / 2
	} else {
		if tl.X < 0 {
			tl.X = 0
		}
		if tl.X >= atlas.Max.X-screen.Max.X {
			tl.X = atlas.Max.X - screen.Max.X
		}
	}
	if screen.Max.Y > atlas.Max.Y {
		tl.Y = -(screen.Max.Y - atlas.Max.Y) / 2
	} else {
		if tl.Y < 0 {
			tl.Y = 0
		}
		if tl.Y >= atlas.Max.Y-screen.Max.Y {
			tl.Y = atlas.Max.Y - screen.Max.Y
		}
	}
	return tl
}

type Atlas struct {
	cells  map[image.Point]Cell
	bounds image.Rectangle
}

func New() *Atlas {
	w := 100
	h := 100
	atlas := &Atlas{
		cells: make(map[image.Point]Cell),
		bounds: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: w, Y: h},
		},
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			atlas.cells[image.Point{X: x, Y: y}] = Cell{feature: feature.Floor}
		}
	}
	for x := 0; x < w; x++ {
		atlas.cells[image.Point{X: x, Y: 0}] = Cell{feature: feature.Wall}
		atlas.cells[image.Point{X: x, Y: h - 1}] = Cell{feature: feature.Wall}
	}
	for y := 0; y < h; y++ {
		atlas.cells[image.Point{X: 0, Y: y}] = Cell{feature: feature.Wall}
		atlas.cells[image.Point{X: w - 1, Y: y}] = Cell{feature: feature.Wall}
	}
	for x := 0; x < w; x += 10 {
		for y := 0; y < h; y += 10 {
			atlas.cells[image.Point{X: x, Y: y}] = Cell{feature: feature.Wall}
		}
	}
	return atlas
}

func (atlas *Atlas) Bounds() image.Rectangle {
	return atlas.bounds
}

func (atlas *Atlas) Glyph(p image.Point) Glyph {
	cell := atlas.cells[p]
	switch cell.feature {
	case feature.Wall:
		return Glyph{
			Code: doubleWall[atlas.wallrune(p)],
			Fore: grid.White,
			Back: grid.Black,
		}
	case feature.Floor:
		return Glyph{
			Code: 176,
			Fore: grid.DarkGray,
			Back: grid.Black,
		}
	default:
		return Glyph{
			Code: 32,
			Fore: grid.Black,
			Back: grid.Black,
		}
	}
}

func (atlas *Atlas) Passable(p image.Point) bool {
	return atlas.cells[p].feature.Passable
}

var doubleWall = []int{233, 186, 205, 200, 201, 187, 188, 204, 203, 185, 202, 206}
var wallRune = []int{0, 1, 2, 3, 1, 1, 4, 7, 2, 6, 2, 10, 5, 9, 8, 11}

func (atlas *Atlas) wallrune(p image.Point) int {
	x := image.Point{X: 1}
	y := image.Point{Y: 1}
	bits := 0
	if atlas.cells[p.Sub(y)].feature.Wallable {
		bits |= 1
	}
	if atlas.cells[p.Add(x)].feature.Wallable {
		bits |= 2
	}
	if atlas.cells[p.Add(y)].feature.Wallable {
		bits |= 4
	}
	if atlas.cells[p.Sub(x)].feature.Wallable {
		bits |= 8
	}
	return wallRune[bits]
}