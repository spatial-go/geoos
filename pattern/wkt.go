/**
	WKT 格式相关转换方法
 */
package pattern

/*
#cgo LDFLAGS: -lgeos_c
#import "../geo/geos.c"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type WKT struct {
	handler C.GEOSContextHandle_t
}


// 对象初始化，生成geos 上下文对象
func NewWKT() WKT {
	return WKT{
		handler: C.geos_initGEOS(),
	}
}
// 销毁geos 上下文环境
func (wkt *WKT) Destroy() {
	C.finishGEOS_r(wkt.handler)
	wkt.handler = nil
}

// 生成WKT格式字符串表示
func (wkt *WKT) ToWKTStr(g *C.GEOSGeometry) (string, error) {
	w := C.GEOSWKTWriter_create_r(wkt.handler)
	defer C.GEOSWKTWriter_destroy_r(wkt.handler, w)
	cstr := C.GEOSWKTWriter_write_r(wkt.handler, w, g)
	if cstr == nil {
		return "", errors.New("writer to wkt is null")
	}
	return C.GoString(cstr), nil
}

// 根据WKT格式字符串生成集合对象
func (wkt *WKT) FromWKTStr(wktstr string) *C.GEOSGeometry {
	reader := C.GEOSWKTReader_create_r(wkt.handler)
	defer C.GEOSWKTReader_destroy_r(wkt.handler, reader)
	cs := C.CString(wktstr)
	geometry := C.GEOSWKTReader_read_r(wkt.handler, reader, cs)
	C.free(unsafe.Pointer(cs))
	return geometry
}
