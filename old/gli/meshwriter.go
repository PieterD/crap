package gli

import (
	"io"
	"math"
)

type MeshWriter struct {
	w   io.Writer
	buf [8]byte
	err error
}

func NewMeshWriter(w io.Writer) *MeshWriter {
	return &MeshWriter{
		w: w,
	}
}

func (mw *MeshWriter) GetError() error {
	return mw.err
}

func (mw *MeshWriter) SetError(err error) {
	mw.err = err
}

func (mw *MeshWriter) put(num int) {
	if mw.err != nil {
		return
	}
	_, mw.err = mw.w.Write(mw.buf[:num])
}

func (mw *MeshWriter) Pad(num int) {
	for i := 0; i < num; i++ {
		mw.buf[i] = 0
	}
	mw.put(num)
}

func (mw *MeshWriter) Write(b []byte) (n int, err error) {
	if mw.err != nil {
		return 0, err
	}
	n, mw.err = mw.w.Write(b)
	return n, mw.err
}

func (mw *MeshWriter) PutUint8(bits uint8) *MeshWriter {
	mw.buf[0] = bits
	mw.put(1)
	return mw
}

func (mw *MeshWriter) PutUint16(bits uint16) *MeshWriter {
	mw.buf[0] = byte(bits >> 8)
	mw.buf[1] = byte(bits)
	mw.put(2)
	return mw
}

func (mw *MeshWriter) PutUint32(bits uint32) *MeshWriter {
	mw.buf[0] = byte(bits >> 24)
	mw.buf[1] = byte(bits >> 16)
	mw.buf[2] = byte(bits >> 8)
	mw.buf[3] = byte(bits)
	mw.put(4)
	return mw
}

func (mw *MeshWriter) PutUint64(bits uint64) *MeshWriter {
	mw.buf[0] = byte(bits >> 56)
	mw.buf[1] = byte(bits >> 48)
	mw.buf[2] = byte(bits >> 40)
	mw.buf[3] = byte(bits >> 32)
	mw.buf[4] = byte(bits >> 24)
	mw.buf[5] = byte(bits >> 16)
	mw.buf[6] = byte(bits >> 8)
	mw.buf[7] = byte(bits)
	mw.put(8)
	return mw
}

func (mw *MeshWriter) PutFloat32(f float32) *MeshWriter {
	mw.PutUint32(math.Float32bits(f))
	return mw
}

func (mw *MeshWriter) PutFloat64(f float64) *MeshWriter {
	mw.PutUint64(math.Float64bits(f))
	return mw
}

func (mw *MeshWriter) PutInt8(bits int8) *MeshWriter {
	return mw.PutUint8(uint8(bits))
}

func (mw *MeshWriter) PutInt16(bits int16) *MeshWriter {
	return mw.PutUint16(uint16(bits))
}

func (mw *MeshWriter) PutInt32(bits int32) *MeshWriter {
	return mw.PutUint32(uint32(bits))
}

func (mw *MeshWriter) PutInt64(bits int64) *MeshWriter {
	return mw.PutUint64(uint64(bits))
}
