package grid

var colorData = []float32{
	0.0, 0.0, 0.0,
	1.0, 1.0, 1.0,
	1.0, 0.0, 0.0,
	0.0, 1.0, 0.0,
	0.0, 0.0, 1.0,
}

type Color int

const (
	Black Color = iota
	White
	Red
	Green
	Blue
)
