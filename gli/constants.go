package gli

import "github.com/go-gl/gl/v3.3-core/gl"

const (
	INVALID_INDEX uint32 = gl.INVALID_INDEX
)

type ProgramParameter uint32

const (
	PROGRAM_DELETE_STATUS                ProgramParameter = gl.DELETE_STATUS
	LINK_STATUS                          ProgramParameter = gl.LINK_STATUS
	VALIDATE_STATUS                      ProgramParameter = gl.VALIDATE_STATUS
	PROGRAM_INFO_LOG_LENGTH              ProgramParameter = gl.INFO_LOG_LENGTH
	ATTACHED_SHADERS                     ProgramParameter = gl.ATTACHED_SHADERS
	ACTIVE_ATTRIBUTES                    ProgramParameter = gl.ACTIVE_ATTRIBUTES
	ACTIVE_ATTRIBUTE_MAX_LENGTH          ProgramParameter = gl.ACTIVE_ATTRIBUTE_MAX_LENGTH
	ACTIVE_UNIFORMS                      ProgramParameter = gl.ACTIVE_UNIFORMS
	ACTIVE_UNIFORM_MAX_LENGTH            ProgramParameter = gl.ACTIVE_UNIFORM_MAX_LENGTH
	ACTIVE_UNIFORM_BLOCKS                ProgramParameter = gl.ACTIVE_UNIFORM_BLOCKS
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH ProgramParameter = gl.ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH
)

type UniformParameter uint32

const (
	UNIFORM_TYPE                        UniformParameter = gl.UNIFORM_TYPE
	UNIFORM_SIZE                        UniformParameter = gl.UNIFORM_SIZE
	UNIFORM_NAME_LENGTH                 UniformParameter = gl.UNIFORM_NAME_LENGTH
	UNIFORM_BLOCK_INDEX                 UniformParameter = gl.UNIFORM_BLOCK_INDEX
	UNIFORM_OFFSET                      UniformParameter = gl.UNIFORM_OFFSET
	UNIFORM_ARRAY_STRIDE                UniformParameter = gl.UNIFORM_ARRAY_STRIDE
	UNIFORM_MATRIX_STRIDE               UniformParameter = gl.UNIFORM_MATRIX_STRIDE
	UNIFORM_IS_ROW_MAJOR                UniformParameter = gl.UNIFORM_IS_ROW_MAJOR
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX UniformParameter = gl.UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX
)

type DataType uint32

const (
	GlBool                 DataType = gl.BOOL
	GlByte                 DataType = gl.BYTE
	GlUByte                DataType = gl.UNSIGNED_BYTE
	GlShort                DataType = gl.SHORT
	GlUShort               DataType = gl.UNSIGNED_SHORT
	GlInt                  DataType = gl.INT
	GlUInt                 DataType = gl.UNSIGNED_INT
	GlFloat                DataType = gl.FLOAT
	GlHalfFloat            DataType = gl.HALF_FLOAT
	GlFixed                DataType = gl.FIXED
	GlInt_2_10_10_10_REV   DataType = gl.INT_2_10_10_10_REV
	GlUInt_2_10_10_10_REV  DataType = gl.UNSIGNED_INT_2_10_10_10_REV
	GlUInt_10F_11F_11F_REV DataType = gl.UNSIGNED_INT_10F_11F_11F_REV
	GlDouble               DataType = gl.DOUBLE

	GlBoolV2   DataType = gl.BOOL_VEC2
	GlBoolV3   DataType = gl.BOOL_VEC3
	GlBoolV4   DataType = gl.BOOL_VEC4
	GlIntV2    DataType = gl.INT_VEC2
	GlIntV3    DataType = gl.INT_VEC3
	GlIntV4    DataType = gl.INT_VEC4
	GlUIntV2   DataType = gl.UNSIGNED_INT_VEC2
	GlUIntV3   DataType = gl.UNSIGNED_INT_VEC3
	GlUIntV4   DataType = gl.UNSIGNED_INT_VEC4
	GlFloatV2  DataType = gl.FLOAT_VEC2
	GlFloatV3  DataType = gl.FLOAT_VEC3
	GlFloatV4  DataType = gl.FLOAT_VEC4
	GlDoubleV2 DataType = gl.DOUBLE_VEC2
	GlDoubleV3 DataType = gl.DOUBLE_VEC3
	GlDoubleV4 DataType = gl.DOUBLE_VEC4

	GlFloatMat2    DataType = gl.FLOAT_MAT2
	GlFloatMat2x3  DataType = gl.FLOAT_MAT2x3
	GlFloatMat2x4  DataType = gl.FLOAT_MAT2x4
	GlFloatMat3x2  DataType = gl.FLOAT_MAT3x2
	GlFloatMat3    DataType = gl.FLOAT_MAT3
	GlFloatMat3x4  DataType = gl.FLOAT_MAT3x4
	GlFloatMat4x2  DataType = gl.FLOAT_MAT4x2
	GlFloatMat4x3  DataType = gl.FLOAT_MAT4x3
	GlFloatMat4    DataType = gl.FLOAT_MAT4
	GlDoubleMat2   DataType = gl.DOUBLE_MAT2
	GlDoubleMat2x3 DataType = gl.DOUBLE_MAT2x3
	GlDoubleMat2x4 DataType = gl.DOUBLE_MAT2x4
	GlDoubleMat3x2 DataType = gl.DOUBLE_MAT3x2
	GlDoubleMat3   DataType = gl.DOUBLE_MAT3
	GlDoubleMat3x4 DataType = gl.DOUBLE_MAT3x4
	GlDoubleMat4x2 DataType = gl.DOUBLE_MAT4x2
	GlDoubleMat4x3 DataType = gl.DOUBLE_MAT4x3
	GlDoubleMat4   DataType = gl.DOUBLE_MAT4

	GlSampler1dShadow      DataType = gl.SAMPLER_1D_SHADOW
	GlSampler2dShadow      DataType = gl.SAMPLER_2D_SHADOW
	GlSampler1dArrayShadow DataType = gl.SAMPLER_1D_ARRAY_SHADOW
	GlSampler2dArrayShadow DataType = gl.SAMPLER_2D_ARRAY_SHADOW
	GlSampler2dRectShadow  DataType = gl.SAMPLER_2D_RECT_SHADOW
	GlSamplerCubeShadow    DataType = gl.SAMPLER_CUBE_SHADOW

	GlSampler1d                 DataType = gl.SAMPLER_1D
	GlSampler2d                 DataType = gl.SAMPLER_2D
	GlSampler3d                 DataType = gl.SAMPLER_3D
	GlSamplerCube               DataType = gl.SAMPLER_CUBE
	GlSampler1dArray            DataType = gl.SAMPLER_1D_ARRAY
	GlSampler2dArray            DataType = gl.SAMPLER_2D_ARRAY
	GlSampler2dMultisample      DataType = gl.SAMPLER_2D_MULTISAMPLE
	GlSampler2dMultisampleArray DataType = gl.SAMPLER_2D_MULTISAMPLE_ARRAY
	GlSampler2dRect             DataType = gl.SAMPLER_2D_RECT
	// GlSamplerBuffer             DataType = gl.SAMPLER_Buffer

	GlIntSampler1d                 DataType = gl.INT_SAMPLER_1D
	GlIntSampler2d                 DataType = gl.INT_SAMPLER_2D
	GlIntSampler3d                 DataType = gl.INT_SAMPLER_3D
	GlIntSamplerCube               DataType = gl.INT_SAMPLER_CUBE
	GlIntSampler1dArray            DataType = gl.INT_SAMPLER_1D_ARRAY
	GlIntSampler2dArray            DataType = gl.INT_SAMPLER_2D_ARRAY
	GlIntSampler2dMultisample      DataType = gl.INT_SAMPLER_2D_MULTISAMPLE
	GlIntSampler2dMultisampleArray DataType = gl.INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	GlIntSampler2dRect             DataType = gl.INT_SAMPLER_2D_RECT
	// GlIntSamplerBuffer             DataType = gl.INT_SAMPLER_Buffer

	GlUintSampler1d                 DataType = gl.UNSIGNED_INT_SAMPLER_1D
	GlUintSampler2d                 DataType = gl.UNSIGNED_INT_SAMPLER_2D
	GlUintSampler3d                 DataType = gl.UNSIGNED_INT_SAMPLER_3D
	GlUintSamplerCube               DataType = gl.UNSIGNED_INT_SAMPLER_CUBE
	GlUintSampler1dArray            DataType = gl.UNSIGNED_INT_SAMPLER_1D_ARRAY
	GlUintSampler2dArray            DataType = gl.UNSIGNED_INT_SAMPLER_2D_ARRAY
	GlUintSampler2dMultisample      DataType = gl.UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE
	GlUintSampler2dMultisampleArray DataType = gl.UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY
	GlUintSampler2dRect             DataType = gl.UNSIGNED_INT_SAMPLER_2D_RECT
	// GlUintSamplerBuffer             DataType = gl.UNSIGNED_INT_SAMPLER_Buffer
)

type BufferAccessTypeHint uint32

const (
	StaticDraw  BufferAccessTypeHint = gl.STATIC_DRAW
	StaticRead  BufferAccessTypeHint = gl.STATIC_READ
	StaticCopy  BufferAccessTypeHint = gl.STATIC_COPY
	StreamDraw  BufferAccessTypeHint = gl.STREAM_DRAW
	StreamRead  BufferAccessTypeHint = gl.STREAM_READ
	StreamCopy  BufferAccessTypeHint = gl.STREAM_COPY
	DynamicDraw BufferAccessTypeHint = gl.DYNAMIC_DRAW
	DynamicRead BufferAccessTypeHint = gl.DYNAMIC_READ
	DynamicCopy BufferAccessTypeHint = gl.DYNAMIC_COPY
)

type BufferTarget uint32

const (
	ArrayBuffer             BufferTarget = gl.ARRAY_BUFFER
	AtomicCounterBuffer     BufferTarget = gl.ATOMIC_COUNTER_BUFFER
	CopyReadBuffer          BufferTarget = gl.COPY_READ_BUFFER
	CopyWriteBuffer         BufferTarget = gl.COPY_WRITE_BUFFER
	DrawIndirectBuffer      BufferTarget = gl.DRAW_INDIRECT_BUFFER
	DispatchIndirectBuffer  BufferTarget = gl.DISPATCH_INDIRECT_BUFFER
	ElementArrayBuffer      BufferTarget = gl.ELEMENT_ARRAY_BUFFER
	PixelPackBuffer         BufferTarget = gl.PIXEL_PACK_BUFFER
	PixelUnpackBuffer       BufferTarget = gl.PIXEL_UNPACK_BUFFER
	QueryBuffer             BufferTarget = gl.QUERY_BUFFER
	ShaderStorageBuffer     BufferTarget = gl.SHADER_STORAGE_BUFFER
	TextureBuffer           BufferTarget = gl.TEXTURE_BUFFER
	TransformFeedbackBuffer BufferTarget = gl.TRANSFORM_FEEDBACK_BUFFER
	UniformBuffer           BufferTarget = gl.UNIFORM_BUFFER
)

type ShaderType uint32

const (
	VertexShader         ShaderType = gl.VERTEX_SHADER
	GeometryShader       ShaderType = gl.GEOMETRY_SHADER
	FragmentShader       ShaderType = gl.FRAGMENT_SHADER
	ComputeShader        ShaderType = gl.COMPUTE_SHADER
	TessControlShader    ShaderType = gl.TESS_CONTROL_SHADER
	TessEvaluationShader ShaderType = gl.TESS_EVALUATION_SHADER
)

type ShaderParameter uint32

const (
	SHADER_TYPE          ShaderParameter = gl.SHADER_TYPE
	SHADER_DELETE_STATUS ShaderParameter = gl.DELETE_STATUS
	COMPILE_STATUS       ShaderParameter = gl.COMPILE_STATUS
	INFO_LOG_LENGTH      ShaderParameter = gl.INFO_LOG_LENGTH
	SHADER_SOURCE_LENGTH ShaderParameter = gl.SHADER_SOURCE_LENGTH
)

type Bool uint32

const (
	TRUE  Bool = gl.TRUE
	FALSE Bool = gl.FALSE
)

type DrawMode uint32

const (
	Points                 DrawMode = gl.POINTS
	LineStrip              DrawMode = gl.LINE_STRIP
	LineLoop               DrawMode = gl.LINE_LOOP
	Lines                  DrawMode = gl.LINES
	LineStripAdjacency     DrawMode = gl.LINE_STRIP_ADJACENCY
	LinesAdjacency         DrawMode = gl.LINES_ADJACENCY
	TriangleStrip          DrawMode = gl.TRIANGLE_STRIP
	TriangleFan            DrawMode = gl.TRIANGLE_FAN
	Triangles              DrawMode = gl.TRIANGLES
	TriangleStripAdjacency DrawMode = gl.TRIANGLE_STRIP_ADJACENCY
	TrianglesAdjacency     DrawMode = gl.TRIANGLES_ADJACENCY
	Patches                DrawMode = gl.PATCHES
)

type ClearBit uint32

const (
	ColorBufferBit   ClearBit = gl.COLOR_BUFFER_BIT
	DepthBufferBit   ClearBit = gl.DEPTH_BUFFER_BIT
	StencilBufferBit ClearBit = gl.STENCIL_BUFFER_BIT
)

type Capability uint32

const (
	Blend                      Capability = gl.BLEND
	ClipDistance0              Capability = gl.CLIP_DISTANCE0
	ClipDistance1              Capability = gl.CLIP_DISTANCE1
	ClipDistance2              Capability = gl.CLIP_DISTANCE2
	ClipDistance3              Capability = gl.CLIP_DISTANCE3
	ClipDistance4              Capability = gl.CLIP_DISTANCE4
	ClipDistance5              Capability = gl.CLIP_DISTANCE5
	ClipDistance6              Capability = gl.CLIP_DISTANCE6
	ClipDistance7              Capability = gl.CLIP_DISTANCE7
	ColorLogicOp               Capability = gl.COLOR_LOGIC_OP
	CullFace                   Capability = gl.CULL_FACE
	DebugOutput                Capability = gl.DEBUG_OUTPUT
	DebugOutputSynchronous     Capability = gl.DEBUG_OUTPUT_SYNCHRONOUS
	DepthClamp                 Capability = gl.DEPTH_CLAMP
	DepthTest                  Capability = gl.DEPTH_TEST
	Dither                     Capability = gl.DITHER
	FramebufferSRGB            Capability = gl.FRAMEBUFFER_SRGB
	LineSmooth                 Capability = gl.LINE_SMOOTH
	MultiSample                Capability = gl.MULTISAMPLE
	PolygonOffsetFill          Capability = gl.POLYGON_OFFSET_FILL
	PolygonOffsetLine          Capability = gl.POLYGON_OFFSET_LINE
	PolygonOffsetPoint         Capability = gl.POLYGON_OFFSET_POINT
	PolygonSmooth              Capability = gl.POLYGON_SMOOTH
	PrimitiveRestart           Capability = gl.PRIMITIVE_RESTART
	PrimitiveRestartFixedIndex Capability = gl.PRIMITIVE_RESTART_FIXED_INDEX
	RasterizerDiscard          Capability = gl.RASTERIZER_DISCARD
	SampleAlphaToCoverage      Capability = gl.SAMPLE_ALPHA_TO_COVERAGE
	SampleAlphaToOne           Capability = gl.SAMPLE_ALPHA_TO_ONE
	SampleCoverage             Capability = gl.SAMPLE_COVERAGE
	// SampleShading              Capability = gl.SAMPLE_SHADING
	SampleMask             Capability = gl.SAMPLE_MASK
	ScissorTest            Capability = gl.SCISSOR_TEST
	StencilTest            Capability = gl.STENCIL_TEST
	TextureCubeMapSeamless Capability = gl.TEXTURE_CUBE_MAP_SEAMLESS
	ProgramPointSize       Capability = gl.PROGRAM_POINT_SIZE
)

type IndexedCapability uint32

const (
	IndexedBlend       IndexedCapability = gl.BLEND
	IndexedScissorTest IndexedCapability = gl.SCISSOR_TEST
)

type DepthFunc uint32

const (
	DepthNever        DepthFunc = gl.NEVER
	DepthLess         DepthFunc = gl.LESS
	DepthLessEqual    DepthFunc = gl.LEQUAL
	DepthEqual        DepthFunc = gl.EQUAL
	DepthGreaterEqual DepthFunc = gl.GEQUAL
	DepthGreater      DepthFunc = gl.GREATER
	DepthAlways       DepthFunc = gl.ALWAYS
)
