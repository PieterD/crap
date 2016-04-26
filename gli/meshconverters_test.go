package gli

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

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
	exp = toBits4(math.Float32bits(v.Float32), nil)
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
	got := f(reflect.ValueOf(v), nil)
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

func TestToBits8(t *testing.T) {
	var b []byte
	b = toBits8(0x3256100112345678, b)
	b = toBits8(0x3F32123487654321, b)
	if b[0] != 0x32 || b[1] != 0x56 || b[2] != 0x10 || b[3] != 0x01 {
		t.Fatalf("bad toBits8")
	}
	if b[4] != 0x12 || b[5] != 0x34 || b[6] != 0x56 || b[7] != 0x78 {
		t.Fatalf("bad toBits8")
	}
	if b[8] != 0x3F || b[9] != 0x32 || b[10] != 0x12 || b[11] != 0x34 {
		t.Fatalf("bad toBits8")
	}
	if b[12] != 0x87 || b[13] != 0x65 || b[14] != 0x43 || b[15] != 0x21 {
		t.Fatalf("bad toBits8")
	}
}

func TestToBits4(t *testing.T) {
	var b []byte
	b = toBits4(0x32561001, b)
	b = toBits4(0x3F321234, b)
	if b[0] != 0x32 || b[1] != 0x56 || b[2] != 0x10 || b[3] != 0x01 {
		t.Fatalf("bad toBits4")
	}
	if b[4] != 0x3F || b[5] != 0x32 || b[6] != 0x12 || b[7] != 0x34 {
		t.Fatalf("bad toBits4")
	}
}

func TestToBits2(t *testing.T) {
	var b []byte
	b = toBits2(0x1001, b)
	b = toBits2(0x3F32, b)
	if b[0] != 0x10 || b[1] != 0x01 || b[2] != 0x3F || b[3] != 0x32 {
		t.Fatalf("bad toBits2")
	}
}

func TestToBits1(t *testing.T) {
	var b []byte
	b = toBits1(1, b)
	b = toBits1(2, b)
	if b[0] != 1 || b[1] != 2 {
		t.Fatalf("bad toBits1")
	}
}
