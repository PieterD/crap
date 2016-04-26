package gli

import (
	"fmt"
	"math"
	"reflect"
)

func defaultFormat(typ reflect.Type) (FullFormat, error) {
	var length int = 1
	etyp := typ
	if typ.Kind() == reflect.Array {
		length = typ.Len()
		if length <= 0 {
			return FullFormat{}, fmt.Errorf("MeshBuilder: array length %d is too short: %v", length, typ)
		}
		if length > 4 {
			return FullFormat{}, fmt.Errorf("MeshBuilder: array length %d is too long: %v", length, typ)
		}
		etyp = etyp.Elem()
	}
	df, ok := defaultBasicFormat(etyp)
	if !ok {
		return FullFormat{}, fmt.Errorf("MeshBuilder: Invalid field type %v", typ)
	}
	return df.Full(length), nil
}

func defaultBasicFormat(typ reflect.Type) (iDataFormat, bool) {
	switch typ.Kind() {
	case reflect.Float32, reflect.Float64:
		return FmFloat, true
	case reflect.Int8:
		return FmByte, true
	case reflect.Int16:
		return FmShort, true
	case reflect.Int32:
		return FmInt, true
	case reflect.Uint8:
		return FmUByte, true
	case reflect.Uint16:
		return FmUShort, true
	case reflect.Uint32:
		return FmUInt, true
	default:
		return iDataFormat{}, false
	}
}

type meshConverter func(v reflect.Value, b []byte) []byte

func fieldConvert(typ reflect.Type, idx []int, format FullFormat) (meshConverter, error) {
	conv, err := dataConvert(typ.FieldByIndex(idx).Type, format)
	if err != nil {
		return nil, err
	}
	return func(v reflect.Value, b []byte) []byte {
		v = v.FieldByIndex(idx)
		return conv(v, b)
	}, nil
}

func dataConvert(typ reflect.Type, format FullFormat) (meshConverter, error) {
	switch typ.Kind() {
	case reflect.Float32, reflect.Float64:
		return convertFloat(typ, format)
	case reflect.Int8:
		return convertInt(typ, 1, format)
	case reflect.Int16:
		return convertInt(typ, 2, format)
	case reflect.Int32:
		return convertInt(typ, 4, format)
	case reflect.Uint8:
		return convertUint(typ, 1, format)
	case reflect.Uint16:
		return convertUint(typ, 2, format)
	case reflect.Uint32:
		return convertUint(typ, 4, format)
	case reflect.Array:
		return convertArray(typ, format)
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid struct field type for attribute: %v", typ)
	}
}

func convertFloat(typ reflect.Type, format FullFormat) (meshConverter, error) {
	switch format.DataFormat {
	case FmHalfFloat:
		return func(v reflect.Value, b []byte) []byte {
			bits := float2half(float32(v.Float()))
			return toBits2(bits, b)
		}, nil
	case FmFixed:
		// https://www.khronos.org/assets/uploads/developers/library/coping_with_fixed_point-Bry.pdf
		return func(v reflect.Value, b []byte) []byte {
			bits := uint32(v.Float() * 65536)
			return toBits4(bits, b)
		}, nil
	case FmFloat:
		return func(v reflect.Value, b []byte) []byte {
			bits := math.Float32bits(float32(v.Float()))
			return toBits4(bits, b)
		}, nil
	case FmDouble:
		return func(v reflect.Value, b []byte) []byte {
			bits := math.Float64bits(v.Float())
			return toBits8(bits, b)
		}, nil
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: conversion not supported", format, typ)
	}
}

func convertInt(typ reflect.Type, size int, format FullFormat) (meshConverter, error) {
	switch format.DataFormat {
	case FmByte:
		if size > 1 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits1(uint8(v.Int()), b)
		}, nil
	case FmShort:
		if size > 2 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits2(uint16(v.Int()), b)
		}, nil
	case FmInt:
		return func(v reflect.Value, b []byte) []byte {
			return toBits4(uint32(v.Int()), b)
		}, nil
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: conversion not supported", format, typ)
	}
}

func convertUint(typ reflect.Type, size int, format FullFormat) (meshConverter, error) {
	switch format.DataFormat {
	case FmUByte:
		if size > 1 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits1(uint8(v.Uint()), b)
		}, nil
	case FmUShort:
		if size > 2 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, b []byte) []byte {
			return toBits2(uint16(v.Uint()), b)
		}, nil
	case FmUInt:
		return func(v reflect.Value, b []byte) []byte {
			return toBits4(uint32(v.Uint()), b)
		}, nil
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: conversion not supported", format, typ)
	}
}

func convertArray(typ reflect.Type, format FullFormat) (meshConverter, error) {
	length := typ.Len()
	if length == 0 {
		return nil, fmt.Errorf("MeshBuilder: Empty array is not a valid field type for attribute: %v", typ)
	}
	if length > 4 {
		return nil, fmt.Errorf("MeshBuilder: Array greater than size 4 is not a valid field type for attribute: %v", typ)
	}
	if length != format.Components {
		return nil, fmt.Errorf("MeshBuilder: Array length '%d' not equal to type component number '%v'", length, format)
	}
	var one meshConverter
	var err error
	//TODO: Int and UInt Rev stuff
	switch typ.Elem().Kind() {
	case reflect.Float32, reflect.Float64:
		one, err = convertFloat(typ, format)
	case reflect.Int8:
		one, err = convertInt(typ, 1, format)
	case reflect.Int16:
		one, err = convertInt(typ, 2, format)
	case reflect.Int32:
		one, err = convertInt(typ, 4, format)
	case reflect.Uint8:
		one, err = convertUint(typ, 1, format)
	case reflect.Uint16:
		one, err = convertUint(typ, 2, format)
	case reflect.Uint32:
		one, err = convertUint(typ, 4, format)
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
