package prog

import (
	"fmt"
	"github.com/PieterD/glimmer/gli"
)

func (coll *programCollection) calcLocations() error {
	attrMap := make(map[string]programAttribute)
	for _, group := range coll.groups {
		location := uint(0)
		for _, program := range group.programs {
			for _, buffer := range program.buffers {
				for _, attr := range buffer.attrs {
					attr2, ok := attrMap[attr.name]
					if !ok {
						attrMap[attr.name] = *attr
						group.locations = append(group.locations, gli.AttributeLocation{
							Name:     attr.name,
							Location: uint32(location),
						})
						location += attr.typ.Location()
					} else {
						if *attr != attr2 {
							return fmt.Errorf("Attribute %s defined in more than one way: %#v and %#v", attr.name, attr, attr2)
						}
					}
				}
			}
		}
	}
	return nil
}
