package gli

type Extent struct {
	size       int
	offset     int
	stride     int
	components int
	normalize  bool
	num        int
}

type ExtentBuilder struct {
	size  int
	start int
	pos   int
	ext   []*Extent
}

func ExtentBuild(componentsize int) ExtentBuilder {
	return ExtentBuilder{
		size: componentsize,
	}
}

func (b ExtentBuilder) Skip(components int) ExtentBuilder {
	b.pos += components * b.size
	return b
}

func (b ExtentBuilder) Size(componentsize int) ExtentBuilder {
	b.size = componentsize
	return b
}

func (b ExtentBuilder) Ext(components int, ext *Extent) ExtentBuilder {
	b.ext = append(b.ext, ext)
	ext.size = b.size
	ext.offset = b.pos
	ext.components = components
	b.pos += components
	return b
}

func (b ExtentBuilder) Seq(num int) ExtentBuilder {
	if num > 0 {
		stride := b.pos - b.start
		for _, ext := range b.ext {
			ext.num = num
			ext.stride = stride
		}
		b.pos += stride * num
	}
	b.start = b.pos
	b.ext = b.ext[:0]
	return b
}
