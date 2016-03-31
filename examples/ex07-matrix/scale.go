package main

import "github.com/go-gl/mathgl/mgl32"

var AnimateScale = []TimeAnim{ScaleNull, ScaleStaticUniform, ScaleStaticNonUniform, ScaleDynamicUniform, ScaleDynamicNonUniform}

func ScaleNull(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(1, 1, 1)
	m.SetCol(3, mgl32.Vec4{0, 0, -45, 1})
	return m
}

func ScaleStaticUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(4, 4, 4)
	m.SetCol(3, mgl32.Vec4{-10, -10, -45, 1})
	return m
}

func ScaleStaticNonUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(0.5, 1, 10)
	m.SetCol(3, mgl32.Vec4{-10, 10, -45, 1})
	return m
}

func ScaleDynamicUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(
		Mix(1, 4, Lerp(elapsed, 3)),
		1,
		1,
	)
	m.SetCol(3, mgl32.Vec4{10, 10, -45, 1})
	return m
}

func ScaleDynamicNonUniform(elapsed float64) mgl32.Mat4 {
	m := mgl32.Scale3D(
		Mix(1, 0.5, Lerp(elapsed, 3)),
		1,
		Mix(1, 10, Lerp(elapsed, 5)),
	)
	m.SetCol(3, mgl32.Vec4{10, -10, -45, 1})
	return m
}

func Mix(s1, s2, lerp float64) float32 {
	return float32(s1*lerp + s2*(1-lerp))
}
