package atlas

import (
	"image"

	"github.com/PieterD/crap/roguelike/game/atlas/aspect"
	"github.com/PieterD/crap/roguelike/grid"
	"github.com/PieterD/crap/roguelike/vision"
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
	cells   []Cell
	bounds  image.Rectangle
	visible uint64
}

func (atlas *Atlas) cell(p image.Point) *Cell {
	if !p.In(atlas.bounds) {
		return &Cell{}
	}
	return &atlas.cells[p.X+p.Y*atlas.bounds.Max.X]
}

func New() *Atlas {
	w := 100
	h := 100
	atlas := &Atlas{
		cells: make([]Cell, w*h),
		bounds: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: w, Y: h},
		},
		visible: 1,
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			atlas.setFeature(x, y, aspect.Floor)
		}
	}
	for x := 0; x < w; x++ {
		atlas.setFeature(x, 0, aspect.Wall)
		atlas.setFeature(x, h-1, aspect.Wall)
	}
	for y := 0; y < h; y++ {
		atlas.setFeature(0, y, aspect.Wall)
		atlas.setFeature(w-1, y, aspect.Wall)
	}
	GenTestlevel(atlas)
	//GenCave(atlas)
	return atlas
}

func (atlas *Atlas) Bounds() image.Rectangle {
	return atlas.bounds
}

func (atlas *Atlas) GetFeature(pos image.Point) aspect.Feature {
	return atlas.cell(pos).feature
}

func (atlas *Atlas) SetFeature(pos image.Point, feature aspect.Feature) {
	atlas.cell(pos).feature = feature
}

func (atlas *Atlas) setFeature(x, y int, ft aspect.Feature) {
	atlas.SetFeature(image.Point{X: x, Y: y}, ft)
}

func (atlas *Atlas) Glyph(p image.Point) Glyph {
	cell := atlas.cell(p)
	glyph := Glyph{
		Code: 32,
		Fore: grid.Black,
		Back: grid.Black,
	}
	if atlas.cell(p).seen {
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
		}
		if !atlas.IsVisible(p) {
			if glyph.Fore != grid.Black {
				glyph.Fore = grid.VeryDarkGray
			}
			glyph.Back = grid.Black
		}
	}
	return glyph
}

func (atlas *Atlas) IsPassable(p image.Point) bool {
	return atlas.cell(p).feature.Passable
}

func (atlas *Atlas) IsTransparent(p image.Point) bool {
	return atlas.cell(p).feature.Transparent
}

func (atlas *Atlas) SetVisible(p image.Point) {
	atlas.cell(p).visible = atlas.visible
	atlas.cell(p).seen = true
}

func (atlas *Atlas) IsVisible(p image.Point) bool {
	return atlas.cell(p).visible == atlas.visible
}

func (atlas *Atlas) Vision(source image.Point) {
	atlas.visible++
	vision.ShadowCastPar(atlas, vision.EndlessRadius(), source)
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
	if atlas.cell(p.Sub(y)).IsWallable() {
		bits |= 1
	}
	if atlas.cell(p.Add(x)).IsWallable() {
		bits |= 2
	}
	if atlas.cell(p.Add(y)).IsWallable() {
		bits |= 4
	}
	if atlas.cell(p.Sub(x)).IsWallable() {
		bits |= 8
	}
	return runes[wallRune[bits]]
}

//var floorRune = []int{44, 46, 96, 249, 250}
//var floorRune = []int{44, 46, 96, 249, 39}
//var floorRune = []int{44, 46, 96, 249, 39, 250, 250}
var floorRune = []int{250, 44, 250, 46, 250, 96, 250, 249, 250, 39, 250}

//var floorRune = []int{250}

func (atlas *Atlas) floorrune(p image.Point, runes []int) int {
	x := uint64(p.X)
	y := uint64(p.Y)
	ui := ((x<<32)|y)*(x^y) + y - x
	return runes[ui%uint64(len(runes))]
}
