package wallify

import (
	"image"
)

type Wallifier interface {
	IsWallable(p image.Point) bool
	IsSeen(p image.Point) bool
}

//var singleWall = []int{79, 179, 196, 192, 218, 191, 217, 195, 194, 180, 193, 197}
//var singleWall = []int{9, 179, 196, 192, 218, 191, 217, 195, 194, 180, 193, 197}
var SingleWall = []int{233, 179, 196, 192, 218, 191, 217, 195, 194, 180, 193, 197}
var DoubleWall = []int{233, 186, 205, 200, 201, 187, 188, 204, 203, 185, 202, 206}
var wallRune = []int{0, 1, 2, 3, 1, 1, 4, 7, 2, 6, 2, 10, 5, 9, 8, 11}
var invisiWall = []int{5, 10, 3, 6, 12, 9, 1, 2, 4, 8}

func Wallify(w Wallifier, p image.Point, runes []int) int {
	x := image.Point{X: 1}
	y := image.Point{Y: 1}
	bits := 0
	if w.IsWallable(p.Sub(y)) {
		bits |= 1
	}
	if w.IsWallable(p.Add(x)) {
		bits |= 2
	}
	if w.IsWallable(p.Add(y)) {
		bits |= 4
	}
	if w.IsWallable(p.Sub(x)) {
		bits |= 8
	}
	switch bits {
	case 7, 14, 13, 11, 15:
		bits2 := 0
		if w.IsWallable(p.Sub(y)) && w.IsSeen(p.Sub(y)) {
			bits2 |= 1
		}
		if w.IsWallable(p.Add(x)) && w.IsSeen(p.Add(x)) {
			bits2 |= 2
		}
		if w.IsWallable(p.Add(y)) && w.IsSeen(p.Add(y)) {
			bits2 |= 4
		}
		if w.IsWallable(p.Sub(x)) && w.IsSeen(p.Sub(x)) {
			bits2 |= 8
		}
		if bits2 == 0 {
			for _, iw := range invisiWall {
				if bits&iw == iw {
					bits = iw
					break
				}
			}
		} else {
			bits = bits2
		}
	}
	return runes[wallRune[bits]]
}
