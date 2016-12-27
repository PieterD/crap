package gli

import "math"

type Deleter interface {
	Delete()
}

func SafeDelete(deleters ...Deleter) {
	for _, deleter := range deleters {
		if deleter != nil {
			deleter.Delete()
		}
	}
}

func PerspectiveMatrix(near, far, frustumScale float64, w, h int) (mat [16]float32) {
	mat[0] = float32(frustumScale / (float64(w) / float64(h)))
	mat[5] = float32(frustumScale)
	mat[10] = float32((near + far) / (near - far))
	mat[11] = -1.0
	mat[14] = float32((2.0 * far * near) / (near - far))
	return
}

func FrustumScale(fovdeg float64) float64 {
	const degtorad = 3.14159 * 2.0 / 360.0
	fovrad := fovdeg * degtorad
	return 1.0 / math.Tan(fovrad/2.0)
}
