package prog

import (
	"encoding/json"
	"fmt"
	"github.com/PieterD/glimmer/gli"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type jsonGroups map[string]map[string]map[string][]string

type programCollection struct {
	text   map[string]string
	groups map[string]*programGroup
}

type programGroup struct {
	name      string
	programs  map[string]*program
	locations []gli.AttributeLocation
}

type programAttribute struct {
	name     string
	buffer   string
	typ      gli.DataType
	location uint
}

type programAttributeBuffer struct {
	name  string
	attrs []*programAttribute
}

type program struct {
	name           string
	buffers        map[string]*programAttributeBuffer
	vertexShader   *programShader
	geometryShader *programShader
	fragmentShader *programShader
}

type programShader struct {
	files []string
}

func ReadPrograms(jsonpath string) (gli.ProgramCollection, error) {
	base := filepath.Dir(jsonpath)
	fileHandle, err := os.Open(jsonpath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open program json file '%s': %v", jsonpath, err)
	}
	defer fileHandle.Close()
	var jGroups jsonGroups
	json.NewDecoder(fileHandle).Decode(&jGroups)
	filenames := make(map[string]struct{})
	coll := &programCollection{
		text:   make(map[string]string),
		groups: make(map[string]*programGroup),
	}
	for jGroup, jPrograms := range jGroups {
		group, ok := coll.groups[jGroup]
		if !ok {
			group = &programGroup{
				name:     jGroup,
				programs: make(map[string]*program),
			}
			coll.groups[jGroup] = group
		}
		for jProgram, jFields := range jPrograms {
			if _, ok := group.programs[jProgram]; ok {
				return nil, fmt.Errorf("Program %s.%s defined multiple times", jGroup, jProgram)
			}
			program := &program{
				name: jProgram,
			}
			group.programs[jProgram] = program
			for jField, jFiles := range jFields {
				shader := &programShader{}
				switch strings.ToLower(jField) {
				case "vertex":
					program.vertexShader = shader
				case "geometry":
					program.geometryShader = shader
				case "fragment":
					program.fragmentShader = shader
				default:
					return nil, fmt.Errorf("Program %s.%s required unknown shader type '%s'", jGroup, jProgram, jField)
				}
				if len(jFiles) == 0 {
					return nil, fmt.Errorf("Shader %s.%s.%s contains no files", jGroup, jProgram, jField)
				}
				for _, jfile := range jFiles {
					if jfile == "" {
						return nil, fmt.Errorf("Shader %s.%s.%s contains empty path", jGroup, jProgram, jField)
					}
					if jfile[0] == '/' {
						return nil, fmt.Errorf("Shader %s.%s.%s contains absolute path '%s'", jGroup, jProgram, jField, jfile)
					}

					file := filepath.Join(base, jfile)
					filenames[file] = struct{}{}
					shader.files = append(shader.files, file)
				}
			}
			if program.vertexShader == nil {
				return nil, fmt.Errorf("Program %s.%s has no vertex shader", jGroup, jProgram)
			}
			if program.fragmentShader == nil {
				return nil, fmt.Errorf("Program %s.%s has no fragment shader", jGroup, jProgram)
			}
		}
	}
	for filename, _ := range filenames {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("Failed to read shader text file '%s': %v", filename, content)
		}
		coll.text[filename] = string(content)
	}

	err = coll.parseAttrs()
	if err != nil {
		return nil, err
	}

	err = coll.calcLocations()
	if err != nil {
		return nil, err
	}

	return coll, nil
}

func (coll *programCollection) Program(name string) gli.ProgramDef {
	return nil
}
