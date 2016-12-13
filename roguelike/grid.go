package main

import "fmt"

type Grid struct {
	screenwidth  int
	screenheight int
	runewidth    int
	runeheight   int
	texwidth     int
	texheight    int

	cols int
	rows int
	padx int
	pady int

	texcols  int
	texrows  int
	texcodes int
}

func NewGrid(runewidth, runeheight, texwidth, texheight int) (*Grid, error) {
	if texwidth%runewidth != 0 {
		return nil, fmt.Errorf("Texture width must be a multiple of the rune width")
	}
	if texheight%runeheight != 0 {
		return nil, fmt.Errorf("Texture height must be a multiple of the rune height")
	}
	grid := &Grid{
		runewidth:  runewidth,
		runeheight: runeheight,
		texwidth:   texwidth,
		texheight:  texheight,
	}

	grid.texcols = grid.texwidth / grid.runewidth
	grid.texrows = grid.texheight / grid.runeheight
	grid.texcodes = grid.texcols * grid.texrows
	return grid, nil
}

func (grid *Grid) Resize(width, height int) {
	grid.screenwidth = width
	grid.screenheight = height
	grid.cols = grid.screenwidth / grid.runewidth
	grid.rows = grid.screenheight / grid.runeheight
	grid.padx = (grid.screenwidth - grid.cols*grid.runewidth) / 2
	grid.pady = (grid.screenheight - grid.rows*grid.runeheight) / 2
}

func (grid *Grid) Coordinates() ([]float32, []uint32) {
	vert := make([]float32, grid.cols*grid.rows*2*4)
	index := make([]uint32, grid.cols*grid.rows*6)
	vpos := uint32(0)
	ipos := 0
	for y := 0; y < grid.rows; y++ {
		for x := 0; x < grid.cols; x++ {
			vstart := (y*grid.rows + x) * 2 * 4
			vcur := vert[vstart : vstart+2*4]
			vcur[0], vcur[1] = grid.fcoord(x, y+1)
			vcur[2], vcur[3] = grid.fcoord(x+1, y+1)
			vcur[4], vcur[5] = grid.fcoord(x+1, y)
			vcur[6], vcur[7] = grid.fcoord(x, y)
			icur := index[ipos : ipos+6]
			icur[0] = vpos
			icur[1] = vpos + 1
			icur[2] = vpos + 2
			icur[3] = vpos
			icur[4] = vpos + 2
			icur[5] = vpos + 3
			vpos += 4
			ipos += 6
		}
	}
	return vert, index
}

func (grid *Grid) fcoord(x, y int) (fx float32, fy float32) {
	fwidth := float32(grid.screenwidth)
	fheight := float32(grid.screenheight)
	itlx := grid.padx + x*grid.runewidth
	itly := grid.pady + y*grid.runeheight
	fx = float32(itlx)/fwidth*2.0 - 1.0
	fy = float32(itly)/fheight*2.0 - 1.0
	return fx, fy
}
