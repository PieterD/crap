package gli

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestFieldArrayConvert(t *testing.T) {
	v := struct {
		Float3 [3]float32
		Float0 [0]float32
		Float5 [5]float32
	}{
		Float3: [3]float32{3, 6, 9},
	}

	var exp []byte
	exp = toBits4(math.Float32bits(v.Float3[0]), exp)
	exp = toBits4(math.Float32bits(v.Float3[1]), exp)
	exp = toBits4(math.Float32bits(v.Float3[2]), exp)
	testFieldConvert(t, v, 0, FmFloat.Full(3), exp)

	_, err := fieldConvert(reflect.TypeOf(v), []int{0}, FmFloat.Full(2))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error with [3]float32 field and float[2] format")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{1}, FmFloat.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error with [0]float32 field")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{2}, FmFloat.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error with [5]float32 field")
	}
}

func TestFieldSimpleConvert(t *testing.T) {
	v := struct {
		Float32 float32
		Int8    int8
		Uint8   uint8
		Int16   int16
		Uint16  uint16
		Int32   int32
		Uint32  uint32
		Float64 float64
		Bytes   []byte
	}{
		Float32: 1.0,
		Int8:    -100,
		Uint8:   32,
		Int16:   -6000,
		Uint16:  6000,
		Int32:   -116000,
		Uint32:  116000,
		Float64: 1.0,
	}

	var exp []byte

	exp = toBits4(uint32(math.Float32bits(v.Float32)), nil)
	testFieldConvert(t, v, 0, FmFloat.Full(1), exp)

	exp = toBits1(uint8(v.Int8), nil)
	testFieldConvert(t, v, 1, FmByte.Full(1), exp)
	exp = toBits1(uint8(v.Uint8), nil)
	testFieldConvert(t, v, 2, FmUByte.Full(1), exp)

	exp = toBits2(uint16(v.Int16), nil)
	testFieldConvert(t, v, 3, FmShort.Full(1), exp)
	exp = toBits2(uint16(v.Uint16), nil)
	testFieldConvert(t, v, 4, FmUShort.Full(1), exp)

	exp = toBits4(uint32(v.Int32), nil)
	testFieldConvert(t, v, 5, FmInt.Full(1), exp)
	exp = toBits4(uint32(v.Uint32), nil)
	testFieldConvert(t, v, 6, FmUInt.Full(1), exp)

	exp = toBits8(math.Float64bits(float64(v.Float64)), nil)
	testFieldConvert(t, v, 7, FmDouble.Full(1), exp)

	// Errors
	_, err := fieldConvert(reflect.TypeOf(v), []int{8}, FmFloat.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error with []byte field")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{0}, FmByte.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting float32 to Byte[1]")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{1}, FmFloat.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting int8 to Float[1]")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{2}, FmFloat.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting uint8 to Float[1]")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{0}, FmFloat.Full(2))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting float32 to Float[2]")
	}

	// Array size errors
	_, err = fieldConvert(reflect.TypeOf(v), []int{5}, FmShort.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting int32 to Short[1]")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{6}, FmUShort.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting uint32 to UShort[1]")
	}

	_, err = fieldConvert(reflect.TypeOf(v), []int{5}, FmByte.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting int32 to Byte[1]")
	}
	_, err = fieldConvert(reflect.TypeOf(v), []int{6}, FmUByte.Full(1))
	if err == nil {
		t.Fatalf("bad fieldconvert: expected error converting uint32 to UByte[1]")
	}
}

func testFieldConvert(t *testing.T, v interface{}, idx int, format FullFormat, exp []byte) {
	f, err := fieldConvert(reflect.TypeOf(v), []int{idx}, format)
	if err != nil {
		t.Fatalf("bad fieldConvert: %v", err)
	}
	bw := bytes.NewBuffer(nil)
	mw := NewMeshWriter(bw)
	f(reflect.ValueOf(v), mw)
	got := bw.Bytes()
	if !bytes.Equal(got, exp) {
		t.Fatalf("bad fieldConvert: expected %v, got %v", exp, got)
	}
}

func TestDefaultFormat(t *testing.T) {
	tests := []struct {
		iface interface{}
		fmt   FullFormat
		err   bool
	}{
		{float32(1.0), FmFloat.Full(1), false},
		{float64(1.0), FmFloat.Full(1), false},
		{int8(1), FmByte.Full(1), false},
		{uint8(1), FmUByte.Full(1), false},
		{int16(1), FmShort.Full(1), false},
		{uint16(1), FmUShort.Full(1), false},
		{int32(1), FmInt.Full(1), false},
		{uint32(1), FmUInt.Full(1), false},
		{[]byte(nil), FullFormat{}, true},
		{[0]byte{}, FullFormat{}, true},
		{[1]byte{}, FmUByte.Full(1), false},
		{[2]float32{}, FmFloat.Full(2), false},
		{[3]int32{}, FmInt.Full(3), false},
		{[4]int8{}, FmByte.Full(4), false},
		{[5]byte{}, FullFormat{}, true},
	}

	for _, test := range tests {
		typ := reflect.TypeOf(test.iface)
		format, err := defaultFormat(typ)
		if test.err && err == nil {
			t.Fatalf("bad defaultFormat: for [%v, %v] expected error, got nothing", typ, test.fmt)
		}
		if !test.err && err != nil {
			t.Fatalf("bad defaultFormat: for [%v, %v] expected no error, got %v", typ, test.fmt, err)
		}
		if format != test.fmt {
			t.Fatalf("bad defaultFormat: for [%v, %v] got %v", typ, test.fmt, format)
		}
	}
}

func toBits1(bits uint8, b []byte) []byte {
	b = append(b, byte(bits>>0))
	return b
}

func toBits2(bits uint16, b []byte) []byte {
	b = append(b, byte(bits>>8))
	b = append(b, byte(bits>>0))
	return b
}

func toBits4(bits uint32, b []byte) []byte {
	b = append(b, byte(bits>>24))
	b = append(b, byte(bits>>16))
	b = append(b, byte(bits>>8))
	b = append(b, byte(bits>>0))
	return b
}

func toBits8(bits uint64, b []byte) []byte {
	b = append(b, byte(bits>>56))
	b = append(b, byte(bits>>48))
	b = append(b, byte(bits>>40))
	b = append(b, byte(bits>>32))
	b = append(b, byte(bits>>24))
	b = append(b, byte(bits>>16))
	b = append(b, byte(bits>>8))
	b = append(b, byte(bits>>0))
	return b
}
