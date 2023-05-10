#include "driver-interface.h"
#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>
#include <memory.h>

#include <locale.h> 

driver_context* new_driver(unsigned char* parameter) {
    srand((unsigned)time(NULL));
    setlocale(LC_ALL, "en_US.UTF-8");
    //printf("%s", parameter);
    driver_context* dc = (driver_context*)malloc(sizeof(driver_context));
    if (dc != NULL) {
        memset(dc, 0, sizeof(driver_context));
        dc->request_total_count = 3;
        dc->request_data_length = 0;
        dc->response_data_length = 0;
        dc->parameter_buf[0] = CONTROL_COMMAND_TYPE_TRUN_OFF - 1;
    }
    return dc;
}

void request(driver_context* context) {
    context->request_data_buffer[0] = 0x01;
    context->request_data_buffer[1] = 0x02;
    context->request_data_buffer[2] = 0x03;
    context->request_data_buffer[3] = 0x04;
    context->request_data_buffer[4] = 0x05;
    context->request_data_length = 5;
}

void control(driver_context* context, driver_control_value* control_value) {
    if (control_value->x == 41 &&
        control_value->y == 42 &&
        control_value->z == 43) {
        if (control_value->command_type == CONTROL_COMMAND_TYPE_TRUN_ON) {
            context->parameter_buf[0] = CONTROL_COMMAND_TYPE_TRUN_ON-1;
            context->control_data_buffer[5] = 0x66;
        }
        if (control_value->command_type == CONTROL_COMMAND_TYPE_TRUN_OFF) {
            context->parameter_buf[0] = CONTROL_COMMAND_TYPE_TRUN_OFF-1;
            context->control_data_buffer[5] = 0x68;
        }

        context->control_data_buffer[0] = 0x31;
        context->control_data_buffer[1] = 0x32;
        context->control_data_buffer[2] = 0x33;
        context->control_data_buffer[3] = 0x34;
        context->control_data_buffer[4] = 0x35;
        context->control_data_length = 6;
    }
}

void response(driver_context* context, response_data_analysis_result* result) {
    if (result) {
        memset(result, 0, sizeof(response_data_analysis_result));

        result->analysis_result = RESPONSE_DATA_ANALYSIS_SUCCESS;
        result->channel_count = 3;
        result->discard_count = 0;

        create_ai(result, 0, 1, 2, 3, rand() % 100);
        create_di(result, 1, 11, 12, 13, context->parameter_buf[0]);
        create_si(result, 2, 21, 22, 23, "hello world。你好，世界！");
    }
}

void timeout(driver_context* context){
    
}

void destroy_driver(driver_context* context) {
    free(context);
}