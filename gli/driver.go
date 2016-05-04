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

type Driver interface {
	Name() string
	ShaderCreate(typ ShaderType, sources ...string) (ShaderId, error)
	ShaderDelete(id ShaderId)
	ProgramCreate(locations []AttributeLocation, shaders ...ShaderId) (ProgramId, error)
	ProgramDelete(id ProgramId)
}
