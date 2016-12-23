package atlas

import "github.com/PieterD/crap/roguelike/game/atlas/aspect"

func GenCave(atlas *Atlas) {
	w := atlas.bounds.Max.X
	h := atlas.bounds.Max.Y
	for x := 0; x < w; x++ {
		atlas.setFeature(x, 0, aspect.Wall)
		atlas.setFeature(x, 1, aspect.Wall)
		atlas.setFeature(x, 2, aspect.Wall)
		atlas.setFeature(x, h-1, aspect.Wall)
		atlas.setFeature(x, h-2, aspect.Wall)
		atlas.setFeature(x, h-3, aspect.Wall)
	}
	for y := 0; y < h; y++ {
		atlas.setFeature(0, y, aspect.Wall)
		atlas.setFeature(1, y, aspect.Wall)
		atlas.setFeature(2, y, aspect.Wall)
		atlas.setFeature(w-1, y, aspect.Wall)
		atlas.setFeature(w-2, y, aspect.Wall)
		atlas.setFeature(w-3, y, aspect.Wall)
	}
}
