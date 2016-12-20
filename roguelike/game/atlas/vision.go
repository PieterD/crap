package atlas

import (
	"image"

	"github.com/PieterD/crap/roguelike/vision"
)

func (atlas *Atlas) Vision(source image.Point) {
	atlas.visibility++
	vision.ShadowCast(atlas, source)
}
