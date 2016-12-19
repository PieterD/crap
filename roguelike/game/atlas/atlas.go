package atlas

import (
	"image"

	"github.com/PieterD/crap/roguelike/game/atlas/aspect"
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
	cells      map[image.Point]Cell
	bounds     image.Rectangle
	visibility uint64
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
		visibility: 1,
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			atlas.cells[image.Point{X: x, Y: y}] = Cell{feature: aspect.Floor}
		}
	}
	for x := 0; x < w; x++ {
		atlas.cells[image.Point{X: x, Y: 0}] = Cell{feature: aspect.Wall}
		atlas.cells[image.Point{X: x, Y: h - 1}] = Cell{feature: aspect.Wall}
	}
	for y := 0; y < h; y++ {
		atlas.cells[image.Point{X: 0, Y: y}] = Cell{feature: aspect.Wall}
		atlas.cells[image.Point{X: w - 1, Y: y}] = Cell{feature: aspect.Wall}
	}
	for x := 0; x < w; x += 10 {
		for y := 0; y < h; y += 10 {
			atlas.cells[image.Point{X: x, Y: y}] = Cell{feature: aspect.Wall}
		}
	}

	wallTest(atlas)

	max := 5
	for i := 1; i <= max; i++ {
		atlas.setFeature(max, i, aspect.Wall)
		atlas.setFeature(i, max, aspect.Wall)
	}

	atlas.setFeature(max, 2, aspect.ClosedDoor)

	for x := 0; x <= 10; x++ {
		atlas.setFeature(45+x, 5, aspect.Wall)
		atlas.setFeature(45+x, 15, aspect.Wall)
		atlas.setFeature(45, 5+x, aspect.Wall)
		atlas.setFeature(55, 5+x, aspect.Wall)
	}
	atlas.setFeature(50, 5, aspect.ClosedDoor)
	atlas.setFeature(50, 10, aspect.Floor)

	return atlas
}

func wallTest(atlas *Atlas) {
	// specials
	atlas.setFeature(10, 10, aspect.Wall)

	atlas.setFeature(10, 20, aspect.Wall)
	atlas.setFeature(10, 19, aspect.Wall)
	atlas.setFeature(10, 21, aspect.Wall)

	atlas.setFeature(10, 30, aspect.Wall)
	atlas.setFeature(9, 30, aspect.Wall)
	atlas.setFeature(11, 30, aspect.Wall)

	atlas.setFeature(10, 40, aspect.Wall)
	atlas.setFeature(9, 40, aspect.Wall)
	atlas.setFeature(11, 40, aspect.Wall)
	atlas.setFeature(10, 39, aspect.Wall)
	atlas.setFeature(10, 41, aspect.Wall)

	// doubles
	atlas.setFeature(20, 10, aspect.Wall)
	atlas.setFeature(20, 9, aspect.Wall)

	atlas.setFeature(20, 20, aspect.Wall)
	atlas.setFeature(21, 20, aspect.Wall)

	atlas.setFeature(20, 30, aspect.Wall)
	atlas.setFeature(20, 31, aspect.Wall)

	atlas.setFeature(20, 40, aspect.Wall)
	atlas.setFeature(19, 40, aspect.Wall)

	// triples
	atlas.setFeature(30, 10, aspect.Wall)
	atlas.setFeature(30, 9, aspect.Wall)
	atlas.setFeature(31, 10, aspect.Wall)

	atlas.setFeature(30, 20, aspect.Wall)
	atlas.setFeature(31, 20, aspect.Wall)
	atlas.setFeature(30, 21, aspect.Wall)

	atlas.setFeature(30, 30, aspect.Wall)
	atlas.setFeature(29, 30, aspect.Wall)
	atlas.setFeature(30, 31, aspect.Wall)

	atlas.setFeature(30, 40, aspect.Wall)
	atlas.setFeature(30, 39, aspect.Wall)
	atlas.setFeature(29, 40, aspect.Wall)

	// quads
	atlas.setFeature(40, 10, aspect.Wall)
	atlas.setFeature(40, 9, aspect.Wall)
	atlas.setFeature(40, 11, aspect.Wall)
	atlas.setFeature(41, 10, aspect.Wall)

	atlas.setFeature(40, 20, aspect.Wall)
	atlas.setFeature(41, 20, aspect.Wall)
	atlas.setFeature(40, 21, aspect.Wall)
	atlas.setFeature(39, 20, aspect.Wall)

	atlas.setFeature(40, 30, aspect.Wall)
	atlas.setFeature(40, 29, aspect.Wall)
	atlas.setFeature(39, 30, aspect.Wall)
	atlas.setFeature(40, 31, aspect.Wall)

	atlas.setFeature(40, 40, aspect.Wall)
	atlas.setFeature(40, 39, aspect.Wall)
	atlas.setFeature(41, 40, aspect.Wall)
	atlas.setFeature(39, 40, aspect.Wall)
}

func (atlas *Atlas) Bounds() image.Rectangle {
	return atlas.bounds
}

func (atlas *Atlas) GetFeature(pos image.Point) aspect.Feature {
	return atlas.cells[pos].feature
}

func (atlas *Atlas) SetFeature(pos image.Point, feature aspect.Feature) {
	c := atlas.cells[pos]
	c.feature = feature
	atlas.cells[pos] = c
}

func (atlas *Atlas) setFeature(x, y int, ft aspect.Feature) {
	cell := atlas.cells[image.Point{X: x, Y: y}]
	cell.feature = ft
	atlas.cells[image.Point{X: x, Y: y}] = cell
}

func (atlas *Atlas) Glyph(p image.Point) Glyph {
	cell := atlas.cells[p]
	var glyph Glyph
	switch cell.feature {
	case aspect.Wall:
		glyph = Glyph{
			Code: atlas.wallrune(p, singleWall),
			Fore: grid.Gray,
			Back: grid.Black,
		}
	case aspect.Floor:
		glyph = Glyph{
			Code: atlas.floorrune(p, floorRune),
			Fore: grid.DarkGray,
			Back: grid.Black,
		}
	case aspect.ClosedDoor:
		glyph = Glyph{
			Code: 43,
			Fore: grid.DarkRed,
			Back: grid.Black,
		}
	case aspect.OpenDoor:
		glyph = Glyph{
			Code: 47,
			Fore: grid.DarkRed,
			Back: grid.Black,
		}
	default:
		glyph = Glyph{
			Code: 32,
			Fore: grid.Black,
			Back: grid.Black,
		}
	}
	if !atlas.Visible(p) {
		if glyph.Fore != grid.Black {
			glyph.Fore = grid.VeryDarkGray
		}
		glyph.Back = grid.Black
	}
	return glyph
}

func (atlas *Atlas) Passable(p image.Point) bool {
	return atlas.cells[p].feature.Passable
}

func (atlas *Atlas) Transparent(p image.Point) bool {
	return atlas.cells[p].feature.Transparent
}

func (atlas *Atlas) SetVisible(p image.Point) {
	cell := atlas.cells[p]
	cell.visibility = atlas.visibility
	atlas.cells[p] = cell
}

func (atlas *Atlas) Visible(p image.Point) bool {
	return atlas.cells[p].visibility == atlas.visibility
}

//var singleWall = []int{79, 179, 196, 192, 218, 191, 217, 195, 194, 180, 193, 197}
//var singleWall = []int{9, 179, 196, 192, 218, 191, 217, 195, 194, 180, 193, 197}
var singleWall = []int{233, 179, 196, 192, 218, 191, 217, 195, 194, 180, 193, 197}
var doubleWall = []int{233, 186, 205, 200, 201, 187, 188, 204, 203, 185, 202, 206}
var wallRune = []int{0, 1, 2, 3, 1, 1, 4, 7, 2, 6, 2, 10, 5, 9, 8, 11}

func (atlas *Atlas) wallrune(p image.Point, runes []int) int {
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
	return runes[wallRune[bits]]
}

//var floorRune = []int{44, 46, 96, 249, 250}
//var floorRune = []int{44, 46, 96, 249, 39}
//var floorRune = []int{44, 46, 96, 249, 39, 250, 250}
//var floorRune = []int{250, 44, 250, 46, 250, 96, 250, 249, 250, 39, 250}
var floorRune = []int{250}

func (atlas *Atlas) floorrune(p image.Point, runes []int) int {
	x := uint64(p.X)
	y := uint64(p.Y)
	ui := ((x<<32)|y)*(x^y) + y - x
	return runes[ui%uint64(len(runes))]
}
