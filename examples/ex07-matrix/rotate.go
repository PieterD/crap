package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

var AnimateRotate = []TimeAnim{RotateNull, RotateX, RotateY, RotateZ, RotateAxis}

func RotateNull(elapsed float64) mgl32.Mat4 {
	m := mgl32.Ident4()
	m.SetCol(3, mgl32.Vec4{0, 0, -25, 1})
	return m
}

func RotateX(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3DX(Angle(elapsed, 2))
	m.SetCol(3, mgl32.Vec4{-5, -5, -25, 1})
	return m
}

func RotateY(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3DY(Angle(elapsed, 2))
	m.SetCol(3, mgl32.Vec4{-5, +5, -25, 1})
	return m
}

func RotateZ(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3DZ(Angle(elapsed, 2))
	m.SetCol(3, mgl32.Vec4{5, 5, -25, 1})
	return m
}

func RotateAxis(elapsed float64) mgl32.Mat4 {
	m := mgl32.HomogRotate3D(Angle(elapsed, 2), mgl32.Vec3{1, 1, 1}.Normalize())
	m.SetCol(3, mgl32.Vec4{5, -5, -25, 1})
	return m
}

func Angle(elapsed, loop float64) float32 {
	scale := 3.14159 * 2 / loop
	cur := math.Mod(elapsed, loop)
	return float32(scale * cur)
}
