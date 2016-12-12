package main

import (
	"image"
	"image/draw"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
)

type Charset struct {
	RuneWidth  int
	RuneHeight int
	Image      image.Image
}

type Game struct {
	width  int
	height int

	cs *Charset

	scr screen.Screen
	win screen.Window
	buf screen.Buffer
	tex screen.Texture
}

func (g *Game) Init(scr screen.Screen, win screen.Window) {
	var err error
	g.cs = NewCharset("rogue_yun_16x16.png", 16, 16)
	g.scr = scr
	g.win = win
	g.buf, err = g.scr.NewBuffer(g.cs.Image.Bounds().Max)
	Panic(err)
	copyImage(g.buf.RGBA(), g.cs.Image)
}

func (g *Game) Resize(width, height int) {
	if g.tex != nil {
		g.tex.Release()
	}
	var err error
	g.tex, err = g.scr.NewTexture(image.Point{X: width, Y: height})
	Panic(err)
	g.width = width
	g.height = height
}

func (g *Game) Key(e key.Event) {
	if e.Code == key.CodeEscape {
		g.win.Release()
		g.win = nil
	}
}

func (g *Game) Draw() {
	g.tex.Upload(image.Point{}, g.buf, image.Rectangle{Min: image.Point{X: 16, Y: 0}, Max: image.Point{X: 32, Y: 16}})
	g.win.Copy(image.Point{}, g.tex, g.tex.Bounds(), draw.Src, nil)
	g.win.Publish()
}

func (g *Game) Exit() {
	if g.tex != nil {
		g.tex.Release()
	}
	if g.buf != nil {
		g.buf.Release()
	}
	if g.win != nil {
		g.win.Release()
	}
}

func copyImage(dst draw.Image, src image.Image) {
	draw.Draw(dst, dst.Bounds(), src, image.Point{X: 0, Y: 0}, draw.Src)
}
