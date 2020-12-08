package coder

import (
	"encoding/hex"
	"github.com/spatial-go/geos/geo"
)

// 根据wkb 二进制生成集合对象
func FromWKB(wkbByte []byte) (geo.GEOSGeometry, error) {
	cwkb := geo.GoByteArrayToCCharArray(wkbByte)
	reader := geo.WKBReaderFactory()
	defer geo.WKBReaderDestroy(reader)
	return geo.EncodeWKBToGeom(reader, cwkb)
}

// 生成WKB格式字符串表示
func ToWKB(g geo.GEOSGeometry) ([]byte, error) {
	writer := geo.WKBWriterFactory()
	defer geo.WKBWriterDestroy(writer)
	return geo.DecodeWKBToArray(writer, g)
}

// 生成WKB Hex格式字符串表达
func ToWKBHex(g geo.GEOSGeometry) (string, error) {
	w := geo.WKBWriterFactory()
	defer geo.WKBWriterDestroy(w)
	return geo.DecodeWKBToHexStr(w, g)
}

// 根据hex格式字符串生成集合对象
func FromWKBHex(wkbHex string) (geo.GEOSGeometry, error) {
	wkbstr, err := hex.DecodeString(wkbHex)
	if err != nil {
		return nil, err
	}
	array := geo.GoByteArrayToCCharArray(wkbstr)
	r := geo.WKBReaderFactory()
	defer geo.WKBReaderDestroy(r)
	return geo.EncodeHexToGeom(r, array)

}
