package geos

import (
	"bytes"
	"fmt"
	"strings"
)

// UnmarshalString encode to geom
func UnmarshalString(s string) (Geometry, error) {
	p := Parser{NewLexer(strings.NewReader(s))}
	return p.Parse()
}

// MarshalString decode to string
func MarshalString(g Geometry) string {
	buf := bytes.NewBuffer(nil)
	wkt(buf, g)
	return buf.String()
}

func wkt(buf *bytes.Buffer, geom Geometry) {
	switch g := geom.(type) {
	case Point:
		_, _ = fmt.Fprintf(buf, "POINT(%g %g)", g.Lat(), g.Lon())
	case MultiPoint:
		if len(g) == 0 {
			buf.Write([]byte(`MULTIPOINT EMPTY`))
			return
		}

		buf.Write([]byte(`MULTIPOINT(`))
		for i, p := range g {
			if i != 0 {
				buf.WriteByte(',')
			}
			_, _ = fmt.Fprintf(buf, "(%g %g)", p.Lat(), p.Lon())
		}
		buf.WriteByte(')')
	case LineString:
		if len(g) == 0 {
			buf.Write([]byte(`LINESTRING EMPTY`))
			return
		}

		buf.Write([]byte(`LINESTRING`))
		writeLineString(buf, g)
	case MultiLineString:
		if len(g) == 0 {
			buf.Write([]byte(`MULTILINESTRING EMPTY`))
			return
		}

		buf.Write([]byte(`MULTILINESTRING(`))
		for i, ls := range g {
			if i != 0 {
				buf.WriteByte(',')
			}
			writeLineString(buf, ls)
		}
		buf.WriteByte(')')
	case Ring:
		wkt(buf, Polygon{g})
	case Polygon:
		if len(g) == 0 {
			buf.Write([]byte(`POLYGON EMPTY`))
			return
		}

		buf.Write([]byte(`POLYGON(`))
		for i, r := range g {
			if i != 0 {
				buf.WriteByte(',')
			}
			writeLineString(buf, LineString(r))
		}
		buf.WriteByte(')')
	case MultiPolygon:
		if len(g) == 0 {
			buf.Write([]byte(`MULTIPOLYGON EMPTY`))
			return
		}

		buf.Write([]byte(`MULTIPOLYGON(`))
		for i, p := range g {
			if i != 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte('(')
			for j, r := range p {
				if j != 0 {
					buf.WriteByte(',')
				}
				writeLineString(buf, LineString(r))
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	//case Collection:
	//	if len(g) == 0 {
	//		buf.Write([]byte(`GEOMETRYCOLLECTION EMPTY`))
	//		return
	//	}
	//	buf.Write([]byte(`GEOMETRYCOLLECTION(`))
	//	for i, c := range g {
	//		if i != 0 {
	//			buf.WriteByte(',')
	//		}
	//		wkt(buf, c)
	//	}
	//	buf.WriteByte(')')
	//case Bound:
	//	wkt(buf, g.ToPolygon())
	default:
		panic("unsupported type")
	}
}

func writeLineString(buf *bytes.Buffer, ls LineString) {
	buf.WriteByte('(')
	for i, p := range ls {
		if i != 0 {
			buf.WriteByte(',')
		}

		_, _ = fmt.Fprintf(buf, "%g %g", p.Lat(), p.Lon())
	}
	buf.WriteByte(')')
}
