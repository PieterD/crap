package vision

import "image"

type Map interface {
	SetVisible(p image.Point)
	IsTransparent(p image.Point) bool
}
