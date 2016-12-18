package atlas

import (
	"image"

	"github.com/PieterD/crap/roguelike/vision"
)

func (atlas *Atlas) Vision(source image.Point) {
	atlas.visibility++
	caster := vision.NewShadowCast(atlasVision{atlas: atlas})
	caster.Vision(source)
}

type atlasVision struct {
	atlas *Atlas
}

func (v atlasVision) MakeVisible(p image.Point) {
	v.atlas.SetVisible(p)
}

func (v atlasVision) IsTransparent(p image.Point) bool {
	return v.atlas.Transparent(p)
}
