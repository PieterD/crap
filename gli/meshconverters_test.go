package gli

import (
	"reflect"
	"testing"
)

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

	/*
		format, err := defaultFormat(reflect.TypeOf(float32(1.0)))
		if err != nil {
			t.Fatalf("bad defaultFormat %v", err)
		}
		if format != FmFloat.Full(1) {
			t.Fatalf("bad defaultFormat %v", format)
		}
	*/
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
