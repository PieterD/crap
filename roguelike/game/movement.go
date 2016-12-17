package game

import "fmt"

type turnDirection int

const (
	turnLeft   turnDirection = 3
	turnRight  turnDirection = 1
	turnAround turnDirection = 2
)

func (g *Game) CmdMove() {
	cpos := g.pos
	npos := cpos
	switch g.dir {
	case 0:
		npos.Y--
	case 1:
		npos.X++
	case 2:
		npos.Y++
	case 3:
		npos.X--
	}
	if g.atlas.Passable(npos) {
		g.pos = npos
	} else {
		fmt.Printf("thump!\n")
	}
}

func (g *Game) CmdTurn(td turnDirection) {
	g.dir = (g.dir + int(td)) % 4
}
