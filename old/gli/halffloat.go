package gli

import "math"

// Fast Half Float Conversions
// Jeroen van der Zijp
// November 2008 (Revised September 2010)
// ftp://ftp.fox-toolkit.org/pub/fasthalffloatconversion.pdf

var basetable [512]uint16
var shifttable [512]byte

func init() {
	var i uint32
	var e int32
	for i = 0; i < 256; i++ {
		e = int32(i) - 127
		if e < -24 {
			basetable[i|0x000] = 0x0000
			basetable[i|0x100] = 0x8000
			shifttable[i|0x000] = 24
			shifttable[i|0x100] = 24
		} else if e < -14 {
			basetable[i|0x000] = (0x0400 >> (uint32(-e) - 14))
			basetable[i|0x100] = (0x0400 >> (uint32(-e) - 14)) | 0x8000
			shifttable[i|0x000] = byte(-e - 1)
			shifttable[i|0x100] = byte(-e - 1)
		} else if e <= 15 {
			basetable[i|0x000] = uint16(uint32(e+15) << 10)
			basetable[i|0x100] = uint16(uint32(e+15)<<10) | 0x8000
			shifttable[i|0x000] = 13
			shifttable[i|0x100] = 13
		} else if e < 128 {
			basetable[i|0x000] = 0x7C00
			basetable[i|0x100] = 0xFC00
			shifttable[i|0x000] = 24
			shifttable[i|0x100] = 24
		} else {
			basetable[i|0x000] = 0x7C00
			basetable[i|0x100] = 0xFC00
			shifttable[i|0x000] = 13
			shifttable[i|0x100] = 13
		}
	}
}

func float2half(ff float32) (hb uint16) {
	fb := math.Float32bits(ff)
	hb = basetable[(fb>>23)&0x1FF] + uint16((fb&0x007FFFFF)>>shifttable[(fb>>23)&0x01FF])
	return hb
}
