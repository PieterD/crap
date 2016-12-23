package atlas

import (
	"github.com/PieterD/crap/roguelike/game/atlas/aspect"
	"image"
	"math/rand"
)

func GenCave(atlas *Atlas) {
	w := atlas.bounds.Max.X
	h := atlas.bounds.Max.Y
	cells := make([]byte, w*h*2)
	swp := cells[w*h:]
	cells = cells[:w*h]
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := x + y*w
			if rand.Intn(100) < 45 {
				cells[idx] = 1
			} else {
				cells[idx] = 0
			}
		}
	}
	for cycle := 0; cycle < 5; cycle++ {
		swp, cells = cells, swp
		for y := 1; y < h-1; y++ {
			for x := 1; x < w-1; x++ {
				idxTo := x+y*w
				neighbors := 0
				for yy:=-1; yy<=1; yy++ {
					for xx := -1; xx <= 1; xx++ {
						idxFrom := (x+xx)+(y+yy)*w
						if swp[idxFrom] == 1{
							neighbors++
						}
					}
				}
				if neighbors >= 5 {
					cells[idxTo] = 1
				} else {
					cells[idxTo] = 0
				}
			}
		}
	}

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if cells[x+y*w] == 1 {
				atlas.cell(image.Point{X: x, Y: y}).feature = aspect.Wall
			}
		}
	}
}