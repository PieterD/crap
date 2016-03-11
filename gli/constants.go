package gli

import "github.com/go-gl/gl/v3.3-core/gl"

type ProgramParameter uint32

const (
	PROGRAM_DELETE_STATUS       ProgramParameter = gl.DELETE_STATUS
	LINK_STATUS                 ProgramParameter = gl.LINK_STATUS
	VALIDATE_STATUS             ProgramParameter = gl.VALIDATE_STATUS
	PROGRAM_INFO_LOG_LENGTH     ProgramParameter = gl.INFO_LOG_LENGTH
	ATTACHED_SHADERS            ProgramParameter = gl.ATTACHED_SHADERS
	ACTIVE_ATTRIBUTES           ProgramParameter = gl.ACTIVE_ATTRIBUTES
	ACTIVE_ATTRIBUTE_MAX_LENGTH ProgramParameter = gl.ACTIVE_ATTRIBUTE_MAX_LENGTH
	ACTIVE_UNIFORMS             ProgramParameter = gl.ACTIVE_UNIFORMS
	ACTIVE_UNIFORM_MAX_LENGTH   ProgramParameter = gl.ACTIVE_UNIFORM_MAX_LENGTH
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

type VertexDimension uint32

const (
	Vertex1d VertexDimension = 1
	Vertex2d VertexDimension = 2
	Vertex3d VertexDimension = 3
	Vertex4d VertexDimension = 4
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
