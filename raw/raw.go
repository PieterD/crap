package raw

type Enum byte

const (
	ShaderTypeVertex Enum = iota + 1
	ShaderTypeGeometry
	ShaderTypeFragment
)

const (
	DataTypeFloat Enum = iota + 1
	DataTypeFloatVec2
	DataTypeFloatVec3
	DataTypeFloatVec4
	DataTypeFloatMat2
	DataTypeFloatMat3
	DataTypeFloatMat4
	DataTypeFloatMat2x3
	DataTypeFloatMat2x4
	DataTypeFloatMat3x2
	DataTypeFloatMat3x4
	DataTypeFloatMat4x2
	DataTypeFloatMat4x3
	DataTypeInt
	DataTypeIntVec2
	DataTypeIntVec3
	DataTypeIntVec4
	DataTypeUnsignedInt
	DataTypeUnsignedIntVec2
	DataTypeUnsignedIntVec3
	DataTypeUnsignedIntVec4
	DataTypeDouble
	DataTypeDoubleVec2
	DataTypeDoubleVec3
	DataTypeDoubleVec4
	DataTypeDoubleMat2
	DataTypeDoubleMat3
	DataTypeDoubleMat4
	DataTypeDoubleMat2x3
	DataTypeDoubleMat2x4
	DataTypeDoubleMat3x2
	DataTypeDoubleMat3x4
	DataTypeDoubleMat4x2
	DataTypeDoubleMat4x3

	DataTypeBool
	DataTypeBool2
	DataTypeBool3
	DataTypeBool4
	DataTypeSampler1d
	DataTypeSampler2d
	DataTypeSampler3d
	DataTypeSamplerCube
	DataTypeSamplerCubeShadow
	DataTypeSampler1dShadow
	DataTypeSampler2dShadow
	DataTypeSampler1dArray
	DataTypeSampler2dArray
	DataTypeSampler1dArrayShadow
	DataTypeSampler2dArrayShadow
	DataTypeSampler2dMultisample
	DataTypeSampler2dMultisampleArray
	DataTypeSamplerBuffer
	DataTypeSampler2dRect
	DataTypeSampler2dRectShadow
	DataTypeIntSampler1d
	DataTypeIntSampler2d
	DataTypeIntSampler3d
	DataTypeIntSamplerCube
	DataTypeIntSampler1dArray
	DataTypeIntSampler2dArray
	DataTypeIntSampler2dMultisample
	DataTypeIntSampler2dMultisampleArray
	DataTypeIntSamplerBuffer
	DataTypeIntSampler2dRect
	DataTypeUIntSampler1d
	DataTypeUIntSampler2d
	DataTypeUIntSampler3d
	DataTypeUIntSamplerCube
	DataTypeUIntSampler1dArray
	DataTypeUIntSampler2dArray
	DataTypeUIntSampler2dMultisample
	DataTypeUIntSampler2dMultisampleArray
	DataTypeUIntSamplerBuffer
	DataTypeUIntSampler2dRect
	DataTypeImage1d
	DataTypeImage2d
	DataTypeImage3d
	DataTypeImageCube
	DataTypeImage1dArray
	DataTypeImage2dArray
	DataTypeImage2dMultisample
	DataTypeImage2dMultisampleArray
	DataTypeImageBuffer
	DataTypeImage2dRect
	DataTypeInt1d
	DataTypeInt2d
	DataTypeInt3d
	DataTypeIntCube
	DataTypeInt1dArray
	DataTypeInt2dArray
	DataTypeInt2dMultisample
	DataTypeInt2dMultisampleArray
	DataTypeIntBuffer
	DataTypeInt2dRect
	DataTypeUInt1d
	DataTypeUInt2d
	DataTypeUInt3d
	DataTypeUIntCube
	DataTypeUInt1dArray
	DataTypeUInt2dArray
	DataTypeUInt2dMultisample
	DataTypeUInt2dMultisampleArray
	DataTypeUIntBuffer
	DataTypeUInt2dRect
	DataTypeUIntAtomicCounter
)

type Raw interface {
	Init() error
	Viewport(x, y, width, height int)
	ClearColor(r, g, b, a float32)

	ShaderCreate(iShadertype Enum) (shaderid uint32, err error)
	ShaderDelete(shaderid uint32)
	ShaderSource(shaderid uint32, source ...string)
	ShaderCompile(shaderid uint32)
	ShaderCompileStatus(shaderid uint32) bool
	ShaderInfoLogLength(shaderid uint32) int
	ShaderInfoLog(shaderid uint32, buf []byte) []byte

	ProgramCreate() (programid uint32, err error)
	ProgramDelete(programid uint32)
	ProgramAttachShader(programid uint32, shaderid uint32)
	ProgramLink(programid uint32)
	ProgramLinkStatus(programid uint32) bool
	ProgramInfoLogLength(programid uint32) int
	ProgramInfoLog(programid uint32, buf []byte) []byte

	ProgramAttributeNum(programid uint32) int
	ProgramAttributeMaxLength(programid uint32) int
	ProgramAttribute(programid uint32, index int, buf []byte) (namebytes []byte, datatype Enum, size int)
	ProgramAttributeLocation(programid uint32, namebytes []byte) (location int, ok bool)

	ProgramUniformNum(programid uint32) int
	ProgramUniformMaxLength(programid uint32) int
	ProgramUniform(programid uint32, index int, buf []byte) (namebytes []byte, datatype Enum, size int)
	ProgramUniformLocation(programid uint32, namebytes []byte) (location int, ok bool)
}
