package gli

import (
	"fmt"
	"reflect"
)

type MeshInstance struct {
}

func (mesh *Mesh) Instance() *MeshInstance {
	return &MeshInstance{}
}

type Mesh struct {
	typ   reflect.Type
	attrs []meshAttribute

	drawMode  iDrawMode
	indexType iIndexType
}

type meshAttribute struct {
	name   string
	buffer int
	conv   meshConverter
	format FullFormat
}

func (mb *meshBuilder) Build() (*Mesh, error) {
	mesh := &Mesh{
		typ:       mb.typ,
		drawMode:  mb.drawMode,
		indexType: mb.indexType,
	}
	for bufnum, buffer := range mb.buffers {
		for _, name := range buffer {
			attr, ok := mb.attrs[name]
			if !ok {
				return nil, fmt.Errorf("MeshBuilder: Attempted to add the attribute '%s' which was not defined", name)
			}
			if attr.added {
				return nil, fmt.Errorf("MeshBuilder: Attempted to add the attribute '%s' more than once", name)
			}
			if !attr.ok {
				return nil, fmt.Errorf("MeshBuilder: Attempted to define attribute '%s' which has no corresponding struct field", name)
			}
			if !attr.customFormat {
				format, err := defaultFormat(attr.typ)
				if err != nil {
					return nil, err
				}
				attr.format = format
			}
			attr.added = true
			mb.attrs[name] = attr
			conv, err := fieldConvert(mb.typ, attr.idx, attr.format)
			if err != nil {
				return nil, err
			}
			mesh.attrs = append(mesh.attrs, meshAttribute{
				name:   attr.name,
				buffer: bufnum,
				conv:   conv,
				format: attr.format,
			})
		}
	}
	return mesh, nil
}

type meshBuilder struct {
	typ       reflect.Type
	fields    map[string]meshBuilderField
	attrs     map[string]*meshBuilderAttribute
	buffers   [][]string
	indexType iIndexType
	drawMode  iDrawMode
}

type meshBuilderField struct {
	name string
	typ  reflect.Type
	idx  []int
}

type meshBuilderAttribute struct {
	name         string
	idx          []int
	typ          reflect.Type
	format       FullFormat
	ok           bool
	added        bool
	customFormat bool
}

func NewMeshBuilder(iface interface{}) (*meshBuilder, error) {
	typ := reflect.TypeOf(iface)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("MeshBuilder: Expected the provided type to be a struct or a pointer to a struct")
	}

	mb := &meshBuilder{
		typ:      typ,
		fields:   make(map[string]meshBuilderField),
		attrs:    make(map[string]*meshBuilderAttribute),
		drawMode: DrawPoints,
	}
	err := mb.seedFields(typ)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (mb *meshBuilder) seedFields(typ reflect.Type) error {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("MeshBuilder: Anonymous fields are only allowed to be structs or pointers to struct")
	}

	numfield := typ.NumField()
	for i := 0; i < numfield; i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue
		}
		if field.Anonymous {
			mb.seedFields(field.Type)
			continue
		}
		name := field.Tag.Get("meshattr")
		if name == "" {
			name = field.Name
		}
		_, ok := mb.fields[name]
		if ok {
			return fmt.Errorf("MeshBuilder: Multiple attributes named '%s'", name)
		}
		mb.fields[name] = meshBuilderField{
			name: name,
			typ:  field.Type,
			idx:  field.Index,
		}
	}
	return nil
}

func (mb *meshBuilder) Attribute(name string) *meshBuilderAttribute {
	field, ok := mb.fields[name]
	attr := &meshBuilderAttribute{
		name: name,
		idx:  field.idx,
		typ:  field.typ,
		ok:   ok,
	}
	mb.attrs[name] = attr
	return attr
}

func (attr *meshBuilderAttribute) Format(format FullFormat) {
	attr.format = format
	attr.customFormat = true
}

func (mb *meshBuilder) Interleave(names ...string) *meshBuilder {
	mb.buffers = append(mb.buffers, names)
	return mb
}

func (mb *meshBuilder) Index(indextype iIndexType) *meshBuilder {
	mb.indexType = indextype
	return mb
}

func (mb *meshBuilder) Mode(drawmode iDrawMode) *meshBuilder {
	mb.drawMode = drawmode
	return mb
}
