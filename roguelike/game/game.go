package game

import (
	"fmt"
	"image"

	"github.com/PieterD/crap/roguelike/grid"
	"github.com/PieterD/crap/roguelike/grid/gridutil"
)

type Game struct {
	die  bool
	x, y int
	dir  int
}

func (g *Game) Draw(d grid.DrawableGrid) {
	gridutil.SingleBox(d, d.GridSize(), grid.White, grid.Black)
	gridutil.Text(d, image.Point{X: 2, Y: 2}, "Hello moo", grid.Red, grid.Black)
	var arrow int
	switch g.dir {
	case 0:
		arrow = 24
	case 1:
		arrow = 26
	case 2:
		arrow = 25
	case 3:
		arrow = 27
	}
	d.Set(g.x, g.y, arrow, grid.Green, grid.Black)
}

func (g *Game) Char(r rune) {
	fmt.Printf("char %d(%c)\n", r, r)
}

func (g *Game) Key(e grid.KeyEvent) {
	fmt.Printf("key %s %#v\n", e.Key(), e)
	switch e.Key() {
	case grid.KeyEscape:
		g.die = true
	case grid.KeyUp:
		g.forward()
	case grid.KeyLeft:
		g.dir = (g.dir + 3) % 4
	case grid.KeyRight:
		g.dir = (g.dir + 1) % 4
	case grid.KeyDown:
		g.dir = (g.dir + 2) % 4
	}
}

func (g *Game) Mouse(e grid.MouseEvent) {
	fmt.Printf("mouse %#v\n", e)
}

func (g *Game) forward() {
	switch g.dir {
	case 0:
		g.y--
	case 1:
		g.x++
	case 2:
		g.y++
	case 3:
		g.x--
	}
}

func (g *Game) Fin(last bool) bool {
	//fmt.Printf("finish %t\n", last)
	return g.die
}

func New() *Game {
	return &Game{
		x:   1,
		y:   1,
		dir: 2,
	}
}
