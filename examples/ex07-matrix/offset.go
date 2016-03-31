package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

var AnimateOffset = []TimeAnim{OffsetStationary, OffsetOval, OffsetCircle}

func Offset(vec mgl32.Vec4) mgl32.Mat4 {
	m := mgl32.Ident4()
	m.SetCol(3, vec)
	return m
}

func OffsetStationary(elapsed float64) mgl32.Mat4 {
	return Offset(mgl32.Vec4{0, 0, -20, 1})
}

func OffsetOval(elapsed float64) mgl32.Mat4 {
	const fLoopDuration = 3
	const fScale = 3.14159 * 2 / fLoopDuration
	fCurrTimeThroughLoop := math.Mod(elapsed, fLoopDuration)
	return Offset(mgl32.Vec4{
		float32(math.Cos(fCurrTimeThroughLoop*fScale) * 4),
		float32(math.Sin(fCurrTimeThroughLoop*fScale) * 6),
		-20,
		1,
	})
}

func OffsetCircle(elapsed float64) mgl32.Mat4 {
	const fLoopDuration = 12
	const fScale = 3.14159 * 2 / fLoopDuration
	fCurrTimeThroughLoop := math.Mod(elapsed, fLoopDuration)
	return Offset(mgl32.Vec4{
		float32(math.Cos(fCurrTimeThroughLoop*fScale) * 5),
		-3.5,
		float32(math.Sin(fCurrTimeThroughLoop*fScale)*5 - 20),
		1,
	})
}
