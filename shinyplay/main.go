package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

type Charset struct {
	RuneWidth  int
	RuneHeight int
	Image      image.Image
}

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

func copyImage(dst draw.Image, src image.Image) {
	draw.Draw(dst, dst.Bounds(), src, image.Point{X: 0, Y: 0}, draw.Src)
}

func main() {
	cs := NewCharset("IBM437-16x16.png", 16, 16)

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Width:  480,
			Height: 640,
		})
		Panic(err)
		defer w.Release()
		buffer, err := s.NewBuffer(cs.Image.Bounds().Max)
		Panic(err)
		defer buffer.Release()
		texture, err := s.NewTexture(image.Point{X: 480, Y: 640})
		Panic(err)
		defer texture.Release()

		copyImage(buffer.RGBA(), cs.Image)
		texture.Upload(image.Point{}, buffer, image.Rectangle{Min: image.Point{X: 16, Y: 0}, Max: image.Point{X: 32, Y: 16}})

		publish := false
		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}
			case mouse.Event:
				fmt.Printf("%f, %f\n", e.X, e.Y)
			case paint.Event:
				fmt.Printf("paint\n")
				publish = true
			case size.Event:
				fmt.Printf("size\n")
			}

			if publish {
				w.Copy(image.Point{}, texture, texture.Bounds(), draw.Src, nil)
				w.Publish()
				publish = false
			}
		}
	})
}
