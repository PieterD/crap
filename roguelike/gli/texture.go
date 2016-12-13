package gli

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

type Texture struct {
	id   uint32
	size image.Point
}

func LoadImage(file string) (*image.RGBA, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return FixImage(img), nil
}

func FixImage(img image.Image) *image.RGBA {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return rgba
}

func (texture *Texture) Id() uint32 {
	return texture.id
}

func (texture *Texture) Use(unit int) {
	gl.ActiveTexture(gl.TEXTURE0 + uint32(unit))
	gl.BindTexture(gl.TEXTURE_2D, texture.id)
}

func (texture *Texture) Size() image.Point {
	return texture.size
}

func (texture *Texture) Delete() {
	gl.DeleteTextures(1, &texture.id)
}

func NewTexture(img *image.RGBA, opts ...TextureOption) (*Texture, error) {
	if img.Stride != img.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride in texture image")
	}
	opt := textureOption{
		filterMin: LINEAR,
		filterMag: LINEAR,
		wrap_s:    REPEAT,
		wrap_t:    REPEAT,
	}
	for _, o := range opts {
		o(&opt)
	}

	var id uint32
	gl.GenTextures(1, &id)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, id)
	defer gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(opt.filterMin))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(opt.filterMag))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int32(opt.wrap_s))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int32(opt.wrap_t))
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(img.Rect.Size().X),
		int32(img.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix))
	return &Texture{
		id:   id,
		size: img.Rect.Size(),
	}, nil
}

type textureOption struct {
	filterMin TextureFilterEnum
	filterMag TextureFilterEnum
	wrap_s    TextureWrapEnum
	wrap_t    TextureWrapEnum
}

type TextureOption func(opt *textureOption)

func TextureFilter(min, mag TextureFilterEnum) TextureOption {
	return func(opt *textureOption) {
		opt.filterMin = min
		opt.filterMag = mag
	}
}

type TextureFilterEnum uint32

const (
	NEAREST TextureFilterEnum = gl.NEAREST
	LINEAR  TextureFilterEnum = gl.LINEAR
)

func TextureWrap(wrap_s, wrap_t TextureWrapEnum) TextureOption {
	return func(opt *textureOption) {
		opt.wrap_s = wrap_s
		opt.wrap_t = wrap_t
	}
}

type TextureWrapEnum uint32

const (
	CLAMP_TO_EDGE   TextureWrapEnum = gl.CLAMP_TO_EDGE
	REPEAT          TextureWrapEnum = gl.REPEAT
	MIRRORED_REPEAT TextureWrapEnum = gl.MIRRORED_REPEAT
)
