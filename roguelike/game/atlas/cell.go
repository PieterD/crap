package atlas

import (
	"github.com/PieterD/crap/roguelike/game/atlas/feature"
	"github.com/PieterD/crap/roguelike/grid"
)

type Cell struct {
	// Floor, wall, etc
	feature feature.Feature

	// Chest, water, altar, lava, etc
	//furnishings []Furnishing

	// Sword, armor, food, gold
	//objects     []Object
}

func (cell Cell) Glyph() Glyph {
	switch cell.feature {
	case feature.Wall:
		return Glyph{
			Code: 35,
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
			Fore: grid.White,
			Back: grid.Black,
		}
	}
}
