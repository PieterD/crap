package main

import (
	"fmt"
	"image/png"
	"os"
)

func NewCharset(filename string, w int, h int) *Charset {
	handle, err := os.Open(filename)
	Panic(err)
	defer handle.Close()
	img, err := png.Decode(handle)
	Panic(err)
	if img.Bounds().Min.X != 0 || img.Bounds().Min.Y != 0 {
		panic(fmt.Errorf("Charset top-left is not 0,0"))
	}
	if img.Bounds().Max.X%w != 0 || img.Bounds().Max.Y%h != 0 {
		panic(fmt.Errorf("Charset size %d,%d not divisible by expected character size %d,%d",
			img.Bounds().Max.X, img.Bounds().Max.Y, w, h))
	}
	return &Charset{
		RuneWidth:  w,
		RuneHeight: h,
		Image:      img,
	}
}
