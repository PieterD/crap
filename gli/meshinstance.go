package gli

import (
	"bytes"
	"fmt"
	"reflect"
)

type MeshInstance struct {
	mesh     *Mesh
	buffers  []meshInstanceBuffer
	indexbuf meshInstanceBuffer
	vertices int
	indices  int
}

type meshInstanceBuffer struct {
	meshWriter *MeshWriter
	buffer     *bytes.Buffer
	gpubuf     *Buffer
}

func (instance *MeshInstance) Transmit(ctx *Context) {
	for i := range instance.buffers {
		instance.buffers[i].gpubuf = ctx.NewBuffer(StaticDraw, BindArrayBuffer)
		instance.buffers[i].gpubuf.Data(instance.buffers[i].buffer.Bytes())
	}
	if instance.indices > 0 {
		instance.indexbuf.gpubuf = ctx.NewBuffer(StaticDraw, BindElementArrayBuffer)
		instance.indexbuf.gpubuf.Data(instance.indexbuf.buffer.Bytes())
	}
}

func (mesh *Mesh) Instance() *MeshInstance {
	bufnum := 0
	for _, attr := range mesh.attrs {
		if attr.buffer > bufnum {
			bufnum = attr.buffer
		}
	}

	instance := &MeshInstance{
		mesh:    mesh,
		buffers: make([]meshInstanceBuffer, bufnum+1),
	}

	for i := range instance.buffers {
		instance.buffers[i].buffer = bytes.NewBuffer(nil)
		instance.buffers[i].meshWriter = NewMeshWriter(instance.buffers[i].buffer)
	}

	instance.indexbuf.buffer = bytes.NewBuffer(nil)
	instance.indexbuf.meshWriter = NewMeshWriter(instance.indexbuf.buffer)

	return instance
}

type Object struct {
	instance   *MeshInstance
	vertexbase int
	vertexnum  int
	indexbase  int
	indexnum   int
}

func (instance *MeshInstance) Object(vertex interface{}, index interface{}) (*Object, error) {
	vertexv := reflect.ValueOf(vertex)
	if vertexv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("MeshBuilder: Object vertex data type is not a slice but a %v", vertexv.Type())
	}
	if vertexv.Type().Elem() != instance.mesh.typ {
		return nil, fmt.Errorf("MeshBuilder, Object vertex data type is not a slice of %v but a slice of %v", instance.mesh.typ, vertexv.Type().Elem())
	}
	var indexv reflect.Value
	if index == nil {
		indexv = reflect.ValueOf([]uint32(nil))
	} else {
		indexv = reflect.ValueOf(index)
	}
	if indexv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("MeshBuilder: Object index data type is not a slice but a %v", indexv.Kind())
	}
	switch indexv.Type().Elem().Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
	default:
		return nil, fmt.Errorf("MeshBuilder: Object index data type is not an unsigned integer type but a %v", indexv.Type().Elem().Kind())
	}

	object := &Object{
		instance:   instance,
		vertexbase: instance.vertices,
		vertexnum:  vertexv.Len(),
		indexbase:  instance.indices,
		indexnum:   indexv.Len(),
	}
	for i := 0; i < indexv.Len(); i++ {
		instance.mesh.indexConv(indexv.Index(i), instance.indexbuf.meshWriter)
		if instance.indexbuf.meshWriter.GetError() != nil {
			return nil, instance.indexbuf.meshWriter.GetError()
		}
	}
	for i := 0; i < vertexv.Len(); i++ {
		v := vertexv.Index(i)
		for _, attr := range instance.mesh.attrs {
			buf := instance.buffers[attr.buffer]
			length := buf.buffer.Len()
			siz := attr.format.DataFormat.Size()
			rem := length % siz
			if rem != 0 {
				buf.meshWriter.Pad(siz - rem)
			}
			attr.conv(v, buf.meshWriter)
			if buf.meshWriter.GetError() != nil {
				return nil, buf.meshWriter.GetError()
			}
		}
	}
	instance.indices += object.indexnum
	instance.vertices += object.vertexnum
	return object, nil
}
