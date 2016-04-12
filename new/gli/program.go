package gli

import "fmt"

type DefProgram struct {
	inputs   []defVar
	outputs  []defVar
	uniforms []defVar
}

func CorrectProgram(program *DefProgram, err error) *DefProgram {
	if err != nil {
		panic(err)
	}
	return program
}

func NewProgram(shaderlist ...*DefShader) (*DefProgram, error) {
	var vertexShader, geometryShader, fragmentShader *DefShader
	for _, shader := range shaderlist {
		switch shader.stype {
		case vertexShaderType:
			if vertexShader != nil {
				return nil, fmt.Errorf("Multiple vertex shaders")
			}
			vertexShader = shader
		case geometryShaderType:
			if geometryShader != nil {
				return nil, fmt.Errorf("Multiple geometry shaders")
			}
			geometryShader = shader
		case fragmentShaderType:
			if fragmentShader != nil {
				return nil, fmt.Errorf("Multiple fragment shaders")
			}
			fragmentShader = shader
		}
	}

	if vertexShader == nil {
		return nil, fmt.Errorf("Missing vertex shader")
	}
	if fragmentShader == nil {
		return nil, fmt.Errorf("Missing fragment shader")
	}
	if vertexShader.inputs == nil {
		return nil, fmt.Errorf("Vertex shader missing inputs")
	}

	var shaders []*DefShader
	shaders = append(shaders, vertexShader)
	if geometryShader != nil {
		shaders = append(shaders, geometryShader)
	}
	shaders = append(shaders, fragmentShader)

	uniforms, err := connect(shaders, make(map[defVar]struct{}))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect shaders: %v", err)
	}

	return &DefProgram{
		inputs:   shaders[0].inputs,
		uniforms: uniforms,
		outputs:  shaders[len(shaders)-1].outputs,
	}, nil
}

func connect(shaders []*DefShader, uniformMap map[defVar]struct{}) ([]defVar, error) {
	if len(shaders) > 0 {
		for _, uniform := range shaders[0].uniforms {
			uniformMap[uniform] = struct{}{}
		}
	}
	if len(shaders) < 2 {
		uniforms := make([]defVar, 0, len(uniformMap))
		for uniform := range uniformMap {
			uniforms = append(uniforms, uniform)
		}
		return uniforms, nil
	}
	from := shaders[0]
	to := shaders[1]
	m := make(map[defVar]struct{})
	for _, output := range from.outputs {
		m[output] = struct{}{}
	}
	for _, input := range to.inputs {
		_, ok := m[input]
		if !ok {
			return nil, fmt.Errorf("Input identifier/type not matched to output in previous stage: %s %s", input.datatype.String(), input.name)
		}
		delete(m, input)
	}
	for output := range m {
		return nil, fmt.Errorf("Output identifier/type not matched to input in next stage: %s %s", output.datatype.String(), output.name)
	}

	return connect(shaders[1:], uniformMap)
}
