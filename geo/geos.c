#include <geos_c.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>

void geos_notice_handler(const char *fmt, ...);
void geos_error_handler(const char *fmt, ...);
char *geos_get_last_error(void);
GEOSContextHandle_t geos_initGEOS();

void geos_notice_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "NOTICE: ");
    vfprintf(stderr, fmt, ap);
    va_end(ap);
}

#define ERRLEN 256

char geos_last_err[ERRLEN];

void geos_error_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    vsnprintf(geos_last_err, (size_t) ERRLEN, fmt, ap);
    va_end(ap);
}

char *geos_get_last_error(void) {
    return geos_last_err;
}

GEOSContextHandle_t geos_initGEOS() {
    return initGEOS_r(geos_notice_handler, geos_error_handler);
}
