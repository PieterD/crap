package atlas

import "github.com/PieterD/crap/roguelike/game/atlas/aspect"

type Cell struct {
	// Floor, wall, etc
	feature aspect.Feature
	// Chest, water, altar, lava, etc
	//furnishing aspect.Furniture

	// Sword, armor, food, gold
	//objects     []Object
}
