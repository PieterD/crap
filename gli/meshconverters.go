package gli

import (
	"fmt"
	"math"
	"reflect"
)

type meshConverter func(v reflect.Value, b []byte) []byte

func fieldConvert(typ reflect.Type, idx []int, format iDataFormat) (meshConverter, error) {
	conv, err := dataConvert(typ, format)
	if err != nil {
		return nil, err
	}
	return func(v reflect.Value, b []byte) []byte {
		v = v.FieldByIndex(idx)
		return conv(v, b)
	}, nil
}

func dataConvert(typ reflect.Type, format iDataFormat) (meshConverter, error) {
	switch typ.Kind() {
	case reflect.Float32, reflect.Float64:
		return convertFloat(typ, format, true)
	case reflect.Int8:
		return convertInt(typ, 1, format, true)
	case reflect.Int16:
		return convertInt(typ, 2, format, true)
	case reflect.Int32:
		return convertInt(typ, 4, format, true)
	case reflect.Uint8:
		return convertUint(typ, 1, format, true)
	case reflect.Uint16:
		return convertUint(typ, 2, format, true)
	case reflect.Uint32:
		return convertUint(typ, 4, format, true)
	case reflect.Array:
		return convertArray(typ, format)
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid struct field type for attribute: %v", typ)
	}
}

func convertFloat(typ reflect.Type, format iDataFormat, clear bool) (meshConverter, error) {
	switch format {
	case FmHalfFloat:
		return func(v reflect.Value, b []byte) []byte {
			bits := float2half(float32(v.Float()))
			return toBits2(bits, b, clear)
		}, nil
	case FmFixed:
		// https://www.khronos.org/assets/uploads/developers/library/coping_with_fixed_point-Bry.pdf
		return func(v reflect.Value, b []byte) []byte {
			bits := uint32(v.Float() * 65536)
			return toBits4(bits, b, clear)
		}, nil
	case FmFloat:
		return func(v reflect.Value, b []byte) []byte {
			bits := math.Float32bits(float32(v.Float()))
			return toBits4(bits, b, clear)
		}, nil
	case FmDouble:
		return func(v reflect.Value, b []byte) []byte {
			bits := math.Float64bits(v.Float())
			return toBits8(bits, b, clear)
		}, nil
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: conversion not supported", format, typ)
	}
}

func convertInt(typ reflect.Type, size int, format iDataFormat, clear bool) (meshConverter, error) {
	switch format {
	case FmByte:
		if size > 1 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits1(uint8(v.Int()), b, clear)
		}, nil
	case FmShort:
		if size > 2 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits2(uint16(v.Int()), b, clear)
		}, nil
	case FmInt:
		return func(v reflect.Value, b []byte) []byte {
			return toBits4(uint32(v.Int()), b, clear)
		}, nil
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: conversion not supported", format, typ)
	}
}

func convertUint(typ reflect.Type, size int, format iDataFormat, clear bool) (meshConverter, error) {
	switch format {
	case FmUByte:
		if size > 1 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits1(uint8(v.Uint()), b, clear)
		}, nil
	case FmUShort:
		if size > 2 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits2(uint16(v.Uint()), b, clear)
		}, nil
	case FmUInt:
		return func(v reflect.Value, b []byte) []byte {
			return toBits4(uint32(v.Uint()), b, clear)
		}, nil
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: conversion not supported", format, typ)
	}
}

func convertArray(typ reflect.Type, format iDataFormat) (meshConverter, error) {
	length := typ.Len()
	if length == 0 {
		return nil, fmt.Errorf("MeshBuilder: Empty array is not a valid field type for attribute: %v", typ)
	}
	if length > 4 {
		return nil, fmt.Errorf("MeshBuilder: Array greater than size 4 is not a valid field type for attribute: %v", typ)
	}
	var one meshConverter
	var err error
	//TODO: Int and UInt Rev stuff
	switch typ.Elem().Kind() {
	case reflect.Float32, reflect.Float64:
		one, err = convertFloat(typ, format, false)
	case reflect.Int8:
		one, err = convertInt(typ, 1, format, false)
	case reflect.Int16:
		one, err = convertInt(typ, 2, format, false)
	case reflect.Int32:
		one, err = convertInt(typ, 4, format, false)
	case reflect.Uint8:
		one, err = convertUint(typ, 1, format, false)
	case reflect.Uint16:
		one, err = convertUint(typ, 2, format, false)
	case reflect.Uint32:
		one, err = convertUint(typ, 4, format, false)
	default:
		err = fmt.Errorf("MeshBuilder: Invalid struct field type for attribute: %v", typ)
	}
	if err != nil {
		return nil, err
	}
	return func(v reflect.Value, b []byte) []byte {
		b = b[:0]
		for i := 0; i < length; i++ {
			one(v.Index(i), b)
		}
		return b
	}, nil
}

func toBits1(bits uint8, b []byte, clear bool) []byte {
	if clear {
		b = b[:0]
	}
	b = append(b, byte(bits>>0))
	return b
}

func toBits2(bits uint16, b []byte, clear bool) []byte {
	if clear {
		b = b[:0]
	}
	b = append(b, byte(bits>>8))
	b = append(b, byte(bits>>0))
	return b
}

func toBits4(bits uint32, b []byte, clear bool) []byte {
	if clear {
		b = b[:0]
	}
	b = append(b, byte(bits>>24))
	b = append(b, byte(bits>>16))
	b = append(b, byte(bits>>8))
	b = append(b, byte(bits>>0))
	return b
}

func toBits8(bits uint64, b []byte, clear bool) []byte {
	if clear {
		b = b[:0]
	}
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
