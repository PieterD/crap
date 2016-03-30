package gli

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"os"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Texture interface {
	Id() uint32
	Target() TextureTarget
	Data(data TextureData)
}

type iTexture struct {
	id     uint32
	target TextureTarget
}

func (context iContext) CreateTexture(target TextureTarget) Texture {
	var texture uint32
	gl.GenTextures(1, &texture)
	return iTexture{id: texture, target: target}
}

func (context iContext) ActiveTexture(i uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + i)
}

func (context iContext) BindTexture(texture Texture) {
	gl.BindTexture(uint32(texture.Target()), texture.Id())
}

func (texture iTexture) Id() uint32 {
	return texture.id
}

func (texture iTexture) Target() TextureTarget {
	return texture.target
}

func (texture iTexture) Data(data TextureData) {
	gl.TexImage2D(uint32(texture.target), 0, int32(data.InternalFormat), int32(data.Width), int32(data.Height), 0, data.ExternalFormat, uint32(data.Type), data.Ptr)
}

type TextureData struct {
	Width          uint32
	Height         uint32
	InternalFormat uint32
	ExternalFormat uint32
	Type           DataType
	Ptr            unsafe.Pointer
}

func TextureFromFile(path string) (TextureData, error) {
	f, err := os.Open(path)
	if err != nil {
		return TextureData{}, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return TextureData{}, err
	}
	return TextureFromImage(img)
}

func TextureFromImage(img image.Image) (TextureData, error) {
	rgba, ok := img.(*image.RGBA)
	if !ok || rgba.Stride != rgba.Rect.Size().X*4 {
		rgba = image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	}
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return TextureData{}, fmt.Errorf("Unsupported stride from new RGBA image")
	}
	return TextureData{
		Width:          uint32(rgba.Rect.Size().X),
		Height:         uint32(rgba.Rect.Size().Y),
		InternalFormat: gl.RGBA,
		ExternalFormat: gl.RGBA,
		Type:           GlUByte,
		Ptr:            gl.Ptr(rgba.Pix),
	}, nil
}
