package game

import (
	"fmt"

	"github.com/PieterD/crap/roguelike/grid"
	"github.com/PieterD/crap/roguelike/grid/gridutil"
)

type Game struct {
	die bool
}

func (g *Game) Draw(d grid.DrawableGrid) {
	gridutil.SingleBox(d, d.GridSize(), grid.White, grid.Black)
	d.Set(1, 1, 25, grid.Green, grid.Black)
}

func (g *Game) Char(r rune) {
	fmt.Printf("char %d(%c)\n", r, r)
}

func (g *Game) Key(e grid.KeyEvent) {
	fmt.Printf("key %s %#v\n", e.Key(), e)
	if e.Key() == grid.KeyEscape {
		g.die = true
	}
}

func (g *Game) Fin(last bool) bool {
	//fmt.Printf("finish %t\n", last)
	return g.die
}

func New() *Game {
	return &Game{}
}
