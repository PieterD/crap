package gli

type Mesh struct {
	name       string
	indextype  iIndexType
	triangles  []uint
	attributes map[meshIndex]*meshAttribute
}

type meshIndex struct {
	name   string
	offset int
}

type meshAttribute struct {
	format    iDataFormat
	typ       FullType
	normalize bool
	contents  []float64
}

type meshBuilder struct {
	mesh *Mesh
}

func MeshBuilder(name string) *meshBuilder {
	return &meshBuilder{
		mesh: &Mesh{
			name:       name,
			attributes: make(map[meshIndex]*meshAttribute),
		},
	}
}

func (mb *meshBuilder) Done() *Mesh {
	return mb.mesh
}

func (mb *meshBuilder) Triangles(format iIndexType) *meshTriangleBuilder {
	mb.mesh.indextype = format
	return &meshTriangleBuilder{
		mesh: mb.mesh,
	}
}

type meshTriangleBuilder struct {
	mesh *Mesh
}

func (mb *meshTriangleBuilder) Add(a, b, c uint) {
	mb.mesh.triangles = append(mb.mesh.triangles, a, b, c)
}

func (mb *meshBuilder) Attribute(name string, format iDataFormat, typ FullType, normalize bool) *meshAttributeBuilder {
	//TODO: Expand to multiple attributes {using offset} for larger datatypes like arrays and matrices
	attr := &meshAttribute{
		format:    format,
		typ:       typ,
		normalize: normalize,
	}
	mb.mesh.attributes[meshIndex{name, 0}] = attr
	return &meshAttributeBuilder{
		mesh: mb.mesh,
		name: name,
		typ:  typ,
	}
}

type meshAttributeBuilder struct {
	mesh *Mesh
	name string
	typ  FullType
}

func (mb *meshAttributeBuilder) Add(data ...float64) {
	attr := mb.mesh.attributes[meshIndex{mb.name, 0}]
	//TODO: Check data length with attribute type
	attr.contents = append(attr.contents, data...)
}
