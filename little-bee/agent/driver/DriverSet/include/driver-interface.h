#ifndef DRIVER_INTERFACE_H
#define DRIVER_INTERFACE_H

#include <string.h>

#if defined (_WIN32)
   #if defined(DRIVER_LIB_EXPORT)
     #define  DRIVERLIB_EXPORT __declspec(dllexport)
   #else
     #define  DRIVERLIB_EXPORT __declspec(dllimport)
   #endif
#else
   #define DRIVERLIB_EXPORT
#endif

#define SIGNAL_DI_VALUE_0  0
#define SIGNAL_DI_VALUE_1  1

#define MAX_REQUEST_BUFFER_SIZE    512
#define MAX_CHANNEL_COUNT          256
#define MAX_RESPONES_BUFFER_SIZE   2048
#define PARAMETER_BUFFER_SIZE      256
#define STRING_VALUE_BUFFER_SIZE   64

#define VALUE_TYPE_UNKNOWN 0
#define VALUE_TYPE_AI      1;
#define VALUE_TYPE_DI      2;
#define VALUE_TYPE_SI      3;

#define RESPONSE_DATA_ANALYSIS_SUCCESS 0
#define RESPONSE_DATA_CONTINUE_READ    1
#define RESPONSE_DATA_DISCARD_ALL      2
#define RESPONSE_DATA_DISCARD_PART     3
#define RESPONSE_DATA_VALIDATION_ERROR 4

#define CONTROL_COMMAND_TYPE_UNKNOWN   0 
#define CONTROL_COMMAND_TYPE_TRUN_ON   1
#define CONTROL_COMMAND_TYPE_TRUN_OFF  2 
#define CONTROL_COMMAND_TYPE_STRIKE    3 
#define CONTROL_COMMAND_TYPE_ANALOG    4 
#define CONTROL_COMMAND_TYPE_STRING    5 

typedef struct driver_context_t
{
	unsigned int    request_total_count;
	unsigned int    request_current_index;
	unsigned char   parameter_buf[PARAMETER_BUFFER_SIZE];
	unsigned char   request_data_buffer[MAX_REQUEST_BUFFER_SIZE]; //request buffer
	unsigned int    request_data_length;
	unsigned char   control_data_buffer[MAX_REQUEST_BUFFER_SIZE]; //control buffer
	unsigned int    control_data_length;
	unsigned char   response_data_buffer[MAX_RESPONES_BUFFER_SIZE]; //response buffer
	unsigned int    response_data_length;
} driver_context;

typedef struct driver_control_value_t {
	unsigned int   x;
	unsigned int   y;
	unsigned int   z;
	unsigned char  command_type;
	int            digital_value;
	double         analog_value;
	unsigned char  string_value[STRING_VALUE_BUFFER_SIZE];
} driver_control_value;

typedef struct driver_channel_value_t {
	unsigned int    x;
	unsigned int    y;
	unsigned int    z;
	unsigned int   value_type;
	double         analog_value;
	unsigned int   digital_value;
	unsigned int   string_len;
	unsigned char   string_value[STRING_VALUE_BUFFER_SIZE];
} driver_channel_value;

typedef struct response_data_analysis_result_t {
	unsigned int   analysis_result;
	unsigned int   discard_count;
	unsigned int   channel_count;
	driver_channel_value channel_value[MAX_CHANNEL_COUNT];
}response_data_analysis_result;

void create_ai(response_data_analysis_result* result, int index, int x, int y, int z, double value) {
	result->channel_value[index].x = x;
	result->channel_value[index].y = y;
	result->channel_value[index].z = z;
	result->channel_value[index].value_type = VALUE_TYPE_AI;
	result->channel_value[index].analog_value = value;
}

void create_di(response_data_analysis_result* result, int index, int x, int y, int z, int value) {
	result->channel_value[index].x = x;
	result->channel_value[index].y = y;
	result->channel_value[index].z = z;
	result->channel_value[index].value_type = VALUE_TYPE_DI;
	result->channel_value[index].digital_value = value;
}

void create_si(response_data_analysis_result* result, int index, int x, int y, int z, char* value) {
	result->channel_value[index].x = x;
	result->channel_value[index].y = y;
	result->channel_value[index].z = z;
	result->channel_value[index].value_type = VALUE_TYPE_SI;
	result->channel_value[index].string_len = (unsigned int)strlen(value);
	strcpy_s(result->channel_value[index].string_value, STRING_VALUE_BUFFER_SIZE, value);
}

#ifdef __cplusplus
extern "C" {
#endif

	DRIVERLIB_EXPORT driver_context* new_driver(unsigned char *parameter);
	DRIVERLIB_EXPORT void request(driver_context* context);
	DRIVERLIB_EXPORT void control(driver_context* context,driver_control_value * control_value);
	DRIVERLIB_EXPORT void response(driver_context* context, response_data_analysis_result* result);
	DRIVERLIB_EXPORT void timeout(driver_context* context);
	DRIVERLIB_EXPORT void destroy_driver(driver_context* context);

#ifdef __cplusplus
}
#endif

#endif //DRIVER_INTERFACE_H