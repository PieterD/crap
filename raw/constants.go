package raw

type Enum byte

const (
	ShaderTypeVertex Enum = iota + 1
	ShaderTypeGeometry
	ShaderTypeFragment
)

const (
	AccessTypeStaticDraw Enum = iota + 1
	AccessTypeStaticRead
	AccessTypeStaticCopy
	AccessTypeStreamDraw
	AccessTypeStreamRead
	AccessTypeStreamCopy
	AccessTypeDynamicDraw
	AccessTypeDynamicRead
	AccessTypeDynamicCopy
)

const (
	DataFormatByte Enum = iota + 1
	DataFormatUnsignedByte
	DataFormatShort
	DataFormatUnsignedShort
	DataFormatInt
	DataFormatUnsignedInt
	DataFormatHalfFloat
	DataFormatFloat
	DataFormatDouble
	DataFormatFixed
	DataFormatIntRev_2_10_10_10
	DataFormatUnsignedIntRev_2_10_10_10
	DataFormatUnsignedIntRev_10F_11F_11F
)

const (
	IndexFormatByte Enum = iota + 1
	IndexFormatShort
	IndexFormatInt
)

// Some data types are allowed only for uniforms, others for both uniforms
// and attributes. Data types less than DataTypeAttributeLimit are okay for both,
// but data types greater than it are suitable only for uniforms.
var DataTypeAttributeLimit = DataTypeBool

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
