package gli

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Texture interface {
	Id() uint32
	Delete()
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

func (texture iTexture) Delete() {
	gl.DeleteTextures(1, &texture.id)
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
		return TextureData{}, fmt.Errorf("Failed to open texture file '%s': %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return TextureData{}, fmt.Errorf("Failed to decode image '%s': %v", err)
	}
	return TextureFromImage(img), nil
}

func TextureFromImage(img image.Image) TextureData {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	slice := make([]byte, 0, width*height*4)
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			slice = append(slice, byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8))
		}
	}
	return TextureFromBytes(slice, width, height)
}

func TextureFromBytes(slice []byte, width int, height int) TextureData {
	return TextureData{
		Width:          uint32(width),
		Height:         uint32(height),
		InternalFormat: gl.RGBA,
		ExternalFormat: gl.RGBA,
		Type:           GlUByte,
		Ptr:            gl.Ptr(slice),
	}
}
