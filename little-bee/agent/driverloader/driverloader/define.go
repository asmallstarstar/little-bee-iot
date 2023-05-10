package driverloader

const MAX_REQUEST_BUFFER_SIZE uint32 = 512
const MAX_CHANNEL_COUNT uint32 = 256
const MAX_RESPONES_BUFFER_SIZE uint32 = 2048
const PARAMETER_BUFFER_SIZE uint32 = 256
const STRING_VALUE_BUFFER_SIZE uint32 = 64

const (
	VALUE_TYPE_UNKNOWN uint32 = 0
	VALUE_TYPE_AI      uint32 = 1
	VALUE_TYPE_DI      uint32 = 2
	VALUE_TYPE_SI      uint32 = 3
)

const (
	RESPONSE_DATA_ANALYSIS_SUCCESS uint32 = 0
	RESPONSE_DATA_CONTINUE_READ    uint32 = 1
	RESPONSE_DATA_DISCARD_ALL      uint32 = 2
	RESPONSE_DATA_DISCARD_PART     uint32 = 3
	RESPONSE_DATA_VALIDATION_ERROR uint32 = 4
)

type DriverContext struct {
	RequestTotalCount   uint32
	RequestCurrentIndex uint32
	ParameterBuf        [PARAMETER_BUFFER_SIZE]byte
	RequestDataBuffer   [MAX_REQUEST_BUFFER_SIZE]byte //request buffer
	RequestDataLength   uint32
	ControlDataBuffer   [MAX_REQUEST_BUFFER_SIZE]byte //control buffer
	ControlDataLength   uint32
	ResponseDataBuffer  [MAX_RESPONES_BUFFER_SIZE]byte //response buffer
	ResponseDataLength  uint32
}

type DriverControlValue struct {
	X            uint32
	Y            uint32
	Z            uint32
	CommandType  uint32
	DigitalValue uint32
	AnalogValue  float64
	StringValue  [STRING_VALUE_BUFFER_SIZE]byte
}

type DriverChannelValue struct {
	X            uint32
	Y            uint32
	Z            uint32
	ValueType    uint32
	AnalogValue  float64
	DigitalValue uint32
	StringLen    uint32
	StringValue  [STRING_VALUE_BUFFER_SIZE]byte
}

type ResponseDataAnalysisResult struct {
	AnalysisResult uint32
	DiscardCount   uint32
	ChannelCount   uint32
	ChannelValue   [MAX_CHANNEL_COUNT]DriverChannelValue
}
