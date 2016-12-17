package atlas

import "github.com/PieterD/crap/roguelike/game/atlas/feature"

type Cell struct {
	// Floor, wall, etc
	feature feature.Feature

	// Chest, water, altar, lava, etc
	//furnishings []Furnishing

	// Sword, armor, food, gold
	//objects     []Object
}
