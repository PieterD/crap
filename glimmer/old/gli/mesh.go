package gli

import (
	"fmt"
	"reflect"
)

type Mesh struct {
	typ   reflect.Type
	attrs []meshAttribute

	drawMode  iDrawMode
	indexConv meshConverter
}

type meshAttribute struct {
	name   string
	buffer int
	conv   meshConverter
	format FullFormat
}

func (mb *meshBuilder) Build() (*Mesh, error) {
	if mb.err != nil {
		return nil, mb.err
	}
	ic, err := indexConvert(mb.indexType)
	if err != nil {
		return nil, err
	}
	mesh := &Mesh{
		typ:       mb.typ,
		drawMode:  mb.drawMode,
		indexConv: ic,
	}
	for bufnum, buffer := range mb.buffers {
		for _, name := range buffer {
			attr, ok := mb.attrs[name]
			if !ok {
				return nil, fmt.Errorf("MeshBuilder: Attempted to add the attribute '%s' which was not defined", name)
			}
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
	attrs     map[string]*meshBuilderAttribute
	buffers   [][]string
	indexType iIndexType
	drawMode  iDrawMode
	err       error
}

type meshBuilderAttribute struct {
	name   string
	idx    []int
	typ    reflect.Type
	format FullFormat
	conv   meshConverter
}

func NewMeshBuilder(iface interface{}) *meshBuilder {
	typ := reflect.TypeOf(iface)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	mb := &meshBuilder{
		typ:      typ,
		attrs:    make(map[string]*meshBuilderAttribute),
		drawMode: DrawPoints,
	}

	if typ.Kind() != reflect.Struct {
		mb.err = fmt.Errorf("MeshBuilder: Expected the provided type to be a struct or a pointer to a struct")
		return mb
	}

	err := mb.seedFields(typ)
	if err != nil {
		mb.err = err
		return nil
	}
	return mb
}

func (mb *meshBuilder) seedFields(typ reflect.Type) error {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("MeshBuilder: Anonymous fields are only allowed to be structs or pointers to struct")
	}

	var buffers []string
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
		if name == "-" {
			continue
		}
		_, ok := mb.attrs[name]
		if ok {
			return fmt.Errorf("MeshBuilder: Multiple attributes named '%s'", name)
		}
		format, err := defaultFormat(field.Type)
		if err != nil {
			return err
		}
		mb.attrs[name] = &meshBuilderAttribute{
			name:   name,
			idx:    field.Index,
			typ:    field.Type,
			format: format,
		}
		buffers = append(buffers, name)
	}
	mb.buffers = append(mb.buffers, buffers)
	return nil
}

func (mb *meshBuilder) Format(name string, format FullFormat) *meshBuilder {
	if mb.err != nil {
		return mb
	}
	attr, ok := mb.attrs[name]
	if !ok {
		mb.err = fmt.Errorf("MeshBuilder: Invalid attribute name '%s' supplied to MeshBuilder.Format", name)
		return mb
	}
	attr.format = format
	return mb
}

func (mb *meshBuilder) Interleave(names ...string) *meshBuilder {
	if mb.err != nil {
		return mb
	}

	nmap := make(map[string]struct{})
	for _, name := range names {
		nmap[name] = struct{}{}
	}

	var lastlist []string
	var newlists [][]string
	for _, buflist := range mb.buffers {
		var newlist []string
		for _, bufname := range buflist {
			_, selected := nmap[bufname]
			if selected {
				lastlist = append(lastlist, bufname)
				delete(nmap, bufname)
			} else {
				newlist = append(newlist, bufname)
			}
		}
		if len(newlist) > 0 {
			newlists = append(newlists, newlist)
		}
	}
	if len(lastlist) > 0 {
		newlists = append(newlists, lastlist)
	}

	for name := range nmap {
		mb.err = fmt.Errorf("MeshBuilder: Unknown attribute name '%s' in Interleave", name)
	}

	mb.buffers = newlists

	return mb
}

func (mb *meshBuilder) Index(indextype iIndexType) *meshBuilder {
	if mb.err == nil {
		mb.indexType = indextype
	}
	return mb
}

func (mb *meshBuilder) Mode(drawmode iDrawMode) *meshBuilder {
	if mb.err == nil {
		mb.drawMode = drawmode
	}
	return mb
}
