/**
WKT 格式相关转换方法
*/
package coder

import "github.com/spatial-go/geos/geo"

// 生成WKT格式字符串表示
func ToWKTStr(g geo.GEOSGeometry) (string, error) {
	w := geo.WKTWriterFactory()
	defer geo.WKTWriterDestroy(w)
	return geo.DecodeWKTToStr(w, g)
}

// 根据WKT格式字符串生成集合对象
func GeomFromWKTStr(wktstr string) geo.GEOSGeometry {
	reader := geo.WKTReaderFactory()
	defer geo.WKTReaderDestroy(reader)
	return geo.EncodeWKTToGeom(reader, wktstr)
}
