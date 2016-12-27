package gli

import "github.com/go-gl/gl/v3.3-core/gl"

type Sampler interface {
	Id() uint32
	Delete()
	SetWrap(s, t, r SamplerWrapMode)
	SetFilter(min SamplerMinFilter, mag SamplerMagFilter)
	SetBorderColor(r, g, b, a float32)
	SetCompare(mode SamplerCompareMode, function DepthFunc)
	SetLod(min, max int32)
}

type iSampler struct {
	id uint32
}

func (context iContext) CreateSampler() Sampler {
	var sampler uint32
	gl.GenSamplers(1, &sampler)
	return iSampler{id: sampler}
}

func (context iContext) BindSampler(sampler Sampler, textureunit uint32) {
	gl.BindSampler(textureunit, sampler.Id())
}

func (context iContext) UnbindSampler(textureunit uint32) {
	gl.BindSampler(textureunit, 0)
}

func (sampler iSampler) Id() uint32 {
	return sampler.id
}

func (sampler iSampler) Delete() {
	gl.DeleteSamplers(1, &sampler.id)
}

func (sampler iSampler) SetWrap(s, t, r SamplerWrapMode) {
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_WRAP_S, int32(s))
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_WRAP_T, int32(t))
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_WRAP_R, int32(r))
}

func (sampler iSampler) SetFilter(min SamplerMinFilter, mag SamplerMagFilter) {
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_MIN_FILTER, int32(min))
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_MAG_FILTER, int32(mag))
}

func (sampler iSampler) SetBorderColor(r, g, b, a float32) {
	var rgba [4]float32
	rgba[0] = r
	rgba[1] = g
	rgba[2] = b
	rgba[3] = a
	gl.SamplerParameterfv(sampler.id, gl.TEXTURE_BORDER_COLOR, &rgba[0])
}

func (sampler iSampler) SetCompare(mode SamplerCompareMode, function DepthFunc) {
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_COMPARE_MODE, int32(mode))
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_COMPARE_FUNC, int32(function))
}

func (sampler iSampler) SetLod(min, max int32) {
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_MIN_LOD, min)
	gl.SamplerParameteri(sampler.id, gl.TEXTURE_MAX_LOD, max)
}
