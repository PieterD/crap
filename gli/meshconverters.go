package gli

import (
	"fmt"
	"reflect"
)

type meshConverter func(v reflect.Value, mw *MeshWriter)

func indexConvert(it iIndexType) (meshConverter, error) {
	if it.t == 0 {
		return nil, nil
	}
	switch it {
	case IndexByte:
		return func(v reflect.Value, mw *MeshWriter) {
			ui := v.Uint()
			if ui >= 256 {
				mw.SetError(fmt.Errorf("Index value overflow: %d does not fit in ubyte", ui))
			} else {
				mw.PutUint8(uint8(ui))
			}
		}, nil
	case IndexShort:
		return func(v reflect.Value, mw *MeshWriter) {
			ui := v.Uint()
			if ui >= 256*256 {
				mw.SetError(fmt.Errorf("Index value overflow: %d does not fit in ushort", ui))
			} else {
				mw.PutUint16(uint16(ui))
			}
		}, nil
	case IndexInt:
		return func(v reflect.Value, mw *MeshWriter) {
			ui := v.Uint()
			if ui >= 256*256*256*256 {
				mw.SetError(fmt.Errorf("Index value overflow: %d does not fit in uint", ui))
			} else {
				mw.PutUint32(uint32(ui))
			}
		}, nil
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid index type %v", it)
	}
}

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

func fieldConvert(typ reflect.Type, idx []int, format FullFormat) (meshConverter, error) {
	conv, err := dataConvert(typ.FieldByIndex(idx).Type, format)
	if err != nil {
		return nil, err
	}
	return func(v reflect.Value, mw *MeshWriter) {
		v = v.FieldByIndex(idx)
		conv(v, mw)
	}, nil
}

func dataConvert(typ reflect.Type, format FullFormat) (meshConverter, error) {
	if typ.Kind() == reflect.Array {
		return convertArray(typ, format)
	}
	if format.Components > 1 {
		return nil, fmt.Errorf("MeshBuilder: Non-array type %v but format has components %v", typ, format)
	}
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
	default:
		return nil, fmt.Errorf("MeshBuilder: Invalid struct field type for attribute: %v", typ)
	}
}

func convertFloat(typ reflect.Type, format FullFormat) (meshConverter, error) {
	switch format.DataFormat {
	case FmHalfFloat:
		return func(v reflect.Value, mw *MeshWriter) {
			bits := float2half(float32(v.Float()))
			mw.PutUint16(bits)
		}, nil
	case FmFixed:
		// https://www.khronos.org/assets/uploads/developers/library/coping_with_fixed_point-Bry.pdf
		return func(v reflect.Value, mw *MeshWriter) {
			bits := uint32(v.Float() * 65536)
			mw.PutUint32(bits)
		}, nil
	case FmFloat:
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutFloat32(float32(v.Float()))
		}, nil
	case FmDouble:
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutFloat64(v.Float())
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
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutUint8(uint8(v.Int()))
		}, nil
	case FmShort:
		if size > 2 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutUint16(uint16(v.Int()))
		}, nil
	case FmInt:
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutUint32(uint32(v.Int()))
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
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutUint8(uint8(v.Uint()))
		}, nil
	case FmUShort:
		if size > 2 {
			return nil, fmt.Errorf("MeshBuilder: Invalid format %v for field type %v: Format too small", format, typ)
		}
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutUint16(uint16(v.Uint()))
		}, nil
	case FmUInt:
		return func(v reflect.Value, mw *MeshWriter) {
			mw.PutUint32(uint32(v.Uint()))
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
	return func(v reflect.Value, mw *MeshWriter) {
		for i := 0; i < length; i++ {
			one(v.Index(i), mw)
		}
	}, nil
}
