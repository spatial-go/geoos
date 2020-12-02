#include <geos_c.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>

void geos_notice_handler(const char *fmt, ...);
void geos_error_handler(const char *fmt, ...);
char *geos_get_last_error(void);
GEOSContextHandle_t geos_initGEOS();
