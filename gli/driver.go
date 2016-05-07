package gli

import "fmt"

var driver Driver

func Register(d Driver) {
	if driver != nil {
		panic(fmt.Errorf("gli.Register(%s): already registered %s previously", d.Name(), driver.Name()))
	}
	driver = d
}

type AttributeLocation struct {
	Name     string
	Location uint32
}

type ProgramResource struct {
	Resource ResourceType
	Name     string
	Type     DataType
	Index    int
}

func (pr ProgramResource) GoString() string {
	return fmt.Sprintf("gli.ProgramResource{Resource: %s, Name:\"%s\", Type: %s, Index: %d}", pr.Resource, pr.Name, pr.Type.GoString(), pr.Index)
}

type Driver interface {
	Name() string
	Init() error
	ShaderCreate(typ ShaderType, sources ...string) (ShaderId, error)
	ShaderDelete(id ShaderId)
	ProgramCreate(locations []AttributeLocation, shaders ...ShaderId) (ProgramId, error)
	ProgramDelete(id ProgramId)
	ProgramAttributes(id ProgramId) ([]ProgramResource, error)
}
