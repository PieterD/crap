package prog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PieterD/glimmer/gli"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type jsonGroups struct {
	Version int                                       `json:"version"`
	Groups  map[string]map[string]map[string][]string `json:"groups"`
}

type programCollection struct {
	version int
	text    map[string]string
	groups  map[string]*programGroup
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
	usage string
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
	if jGroups.Version == 0 {
		return nil, fmt.Errorf("Missing or invalid 'version' in json file '%s'", jsonpath)
	}
	if len(jGroups.Groups) == 0 {
		return nil, fmt.Errorf("Missing 'groups' in json file '%s'", jsonpath)
	}
	coll.version = jGroups.Version
	for jGroup, jPrograms := range jGroups.Groups {
		if strings.Contains(jGroup, ".") {
			return nil, fmt.Errorf("Group name '%s' not allowed: contains '.'", jGroup)
		}
		if jGroup == "" {
			return nil, fmt.Errorf("Empty group name found")
		}
		group, ok := coll.groups[jGroup]
		if !ok {
			group = &programGroup{
				name:     jGroup,
				programs: make(map[string]*program),
			}
			coll.groups[jGroup] = group
		}
		for jProgram, jFields := range jPrograms {
			if strings.Contains(jProgram, ".") {
				return nil, fmt.Errorf("Group %s: Program name '%s' not allowed: contains '.'", jGroup, jProgram)
			}
			if jProgram == "" {
				return nil, fmt.Errorf("Group %s: Empty program name found", jGroup)
			}
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

func (coll *programCollection) Program(name string) (gli.ProgramDef, error) {
	split := strings.Split(name, ".")
	if len(split) != 2 || split[0] == "" || split[1] == "" {
		return gli.ProgramDef{}, fmt.Errorf("Invalid program identifier '%s': expected 'group.program'", name)
	}
	group, ok := coll.groups[split[0]]
	if !ok {
		return gli.ProgramDef{}, fmt.Errorf("Unknown program group: '%s'", split[0])
	}
	program, ok := group.programs[split[1]]
	if !ok {
		return gli.ProgramDef{}, fmt.Errorf("Unknown program: '%s.%s'", split[0], split[1])
	}
	p := gli.ProgramDef{
		Name:  program.name,
		Group: group.name,
	}
	p.Locations = append(p.Locations, group.locations...)
	for _, buffer := range program.buffers {
		nBuffer := gli.BufferDef{Name: buffer.name}
		switch buffer.usage {
		case "STATIC":
			nBuffer.Usage = gli.BufferUsageStaticDraw
		case "DYNAMIC":
			nBuffer.Usage = gli.BufferUsageDynamicDraw
		case "STREAM":
			nBuffer.Usage = gli.BufferUsageStreamDraw
		default:
			return gli.ProgramDef{}, fmt.Errorf("Program %s.%s: Invalid buffer usage '%s'", group.name, program.name, buffer.usage)
		}
		for _, attr := range buffer.attrs {
			nAttr := gli.ProgramResource{
				Name:     attr.name,
				Index:    attr.location,
				Resource: gli.ResourceTypeAttribute,
				Type:     attr.typ,
			}
			nBuffer.Attributes = append(nBuffer.Attributes, nAttr)
		}
		p.Buffers = append(p.Buffers, nBuffer)
	}

	if program.vertexShader != nil {
		buf := bytes.NewBuffer(nil)
		fmt.Fprintf(buf, "#version %d\n#define DYNAMIC(X)\n#define STATIC(X)\n#define STREAM(X)\n\n", coll.version)
		for _, filename := range program.vertexShader.files {
			fmt.Fprintln(buf, "")
			buf.WriteString(coll.text[filename])
		}
		p.Shaders = append(p.Shaders, gli.ShaderDef{
			Type:   gli.ShaderTypeVertex,
			Source: []string{buf.String()},
		})
	}
	if program.geometryShader != nil {
		buf := bytes.NewBuffer(nil)
		fmt.Fprintf(buf, "#version %d\n", coll.version)
		for _, filename := range program.geometryShader.files {
			fmt.Fprintln(buf, "")
			buf.WriteString(coll.text[filename])
		}
		p.Shaders = append(p.Shaders, gli.ShaderDef{
			Type:   gli.ShaderTypeGeometry,
			Source: []string{buf.String()},
		})
	}
	if program.fragmentShader != nil {
		buf := bytes.NewBuffer(nil)
		fmt.Fprintf(buf, "#version %d\n", coll.version)
		for _, filename := range program.fragmentShader.files {
			fmt.Fprintln(buf, "")
			buf.WriteString(coll.text[filename])
		}
		p.Shaders = append(p.Shaders, gli.ShaderDef{
			Type:   gli.ShaderTypeFragment,
			Source: []string{buf.String()},
		})
	}

	return p, nil
}
