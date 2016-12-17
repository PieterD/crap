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

	wallTest(atlas)

	return atlas
}

func wallTest(atlas *Atlas) {
	// specials
	atlas.setFeature(10, 10, feature.Wall)

	atlas.setFeature(10, 20, feature.Wall)
	atlas.setFeature(10, 19, feature.Wall)
	atlas.setFeature(10, 21, feature.Wall)

	atlas.setFeature(10, 30, feature.Wall)
	atlas.setFeature(9, 30, feature.Wall)
	atlas.setFeature(11, 30, feature.Wall)

	atlas.setFeature(10, 40, feature.Wall)
	atlas.setFeature(9, 40, feature.Wall)
	atlas.setFeature(11, 40, feature.Wall)
	atlas.setFeature(10, 39, feature.Wall)
	atlas.setFeature(10, 41, feature.Wall)

	// doubles
	atlas.setFeature(20, 10, feature.Wall)
	atlas.setFeature(20, 9, feature.Wall)

	atlas.setFeature(20, 20, feature.Wall)
	atlas.setFeature(21, 20, feature.Wall)

	atlas.setFeature(20, 30, feature.Wall)
	atlas.setFeature(20, 31, feature.Wall)

	atlas.setFeature(20, 40, feature.Wall)
	atlas.setFeature(19, 40, feature.Wall)

	// triples
	atlas.setFeature(30, 10, feature.Wall)
	atlas.setFeature(30, 9, feature.Wall)
	atlas.setFeature(31, 10, feature.Wall)

	atlas.setFeature(30, 20, feature.Wall)
	atlas.setFeature(31, 20, feature.Wall)
	atlas.setFeature(30, 21, feature.Wall)

	atlas.setFeature(30, 30, feature.Wall)
	atlas.setFeature(29, 30, feature.Wall)
	atlas.setFeature(30, 31, feature.Wall)

	atlas.setFeature(30, 40, feature.Wall)
	atlas.setFeature(30, 39, feature.Wall)
	atlas.setFeature(29, 40, feature.Wall)

	// quads
	atlas.setFeature(40, 10, feature.Wall)
	atlas.setFeature(40, 9, feature.Wall)
	atlas.setFeature(40, 11, feature.Wall)
	atlas.setFeature(41, 10, feature.Wall)

	atlas.setFeature(40, 20, feature.Wall)
	atlas.setFeature(41, 20, feature.Wall)
	atlas.setFeature(40, 21, feature.Wall)
	atlas.setFeature(39, 20, feature.Wall)

	atlas.setFeature(40, 30, feature.Wall)
	atlas.setFeature(40, 29, feature.Wall)
	atlas.setFeature(39, 30, feature.Wall)
	atlas.setFeature(40, 31, feature.Wall)

	atlas.setFeature(40, 40, feature.Wall)
	atlas.setFeature(40, 39, feature.Wall)
	atlas.setFeature(41, 40, feature.Wall)
	atlas.setFeature(39, 40, feature.Wall)
}

func (atlas *Atlas) Bounds() image.Rectangle {
	return atlas.bounds
}

func (atlas *Atlas) setFeature(x, y int, ft feature.Feature) {
	cell := atlas.cells[image.Point{X: x, Y: y}]
	cell.feature = ft
	atlas.cells[image.Point{X: x, Y: y}] = cell
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

var singleWall = []int{79, 179, 196, 192, 218, 191, 217, 195, 194, 180, 193, 197}
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
