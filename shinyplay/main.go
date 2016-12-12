package main

import (
	"fmt"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Width:  480,
			Height: 640,
		})
		Panic(err)

		g := Game{}

		g.Init(s, w)

		publish := false
		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					g.Exit()
					return
				}
			case key.Event:
				fmt.Printf("key\n")
				g.Key(e)
			case mouse.Event:
				fmt.Printf("%f, %f\n", e.X, e.Y)
			case paint.Event:
				fmt.Printf("paint %t\n", e.External)
				publish = true
			case size.Event:
				fmt.Printf("resize %d, %d\n", e.WidthPx, e.HeightPx)
				g.Resize(e.WidthPx, e.HeightPx)
			}

			if publish {
				publish = false
				g.Draw()
			}
		}
	})
}
