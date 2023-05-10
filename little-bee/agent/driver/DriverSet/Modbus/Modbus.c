#include "driver-interface.h"
#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>
#include <memory.h>
#include <string.h>
#include <locale.h> 

driver_context* new_driver(unsigned char* parameter) {
    setlocale(LC_ALL, "en_US.UTF-8");
    printf("%s", parameter);
    driver_context* dc = (driver_context*)malloc(sizeof(driver_context));
    if (dc != NULL) {

    }
    return dc;
}

void request(driver_context* context) {

}

void control(driver_context* context, driver_control_value* control_value) {

}

void response(driver_context* context, response_data_analysis_result* result) {
    
}

void timeout(driver_context* context) {

}

void destroy_driver(driver_context* context) {
    free(context);
}