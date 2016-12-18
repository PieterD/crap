package vision

import "image"

type Map interface {
	MakeVisible(p image.Point)
	IsTransparent(p image.Point) bool
}
