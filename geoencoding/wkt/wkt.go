// Package wkt is for decoding  Well Known Text (WKT) format specification
// at https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary
package wkt

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/spatial-go/geoos/space"
)

// UnmarshalString encode to geom
func UnmarshalString(s string) (space.Geometry, error) {
	p := Parser{NewLexer(strings.NewReader(s))}
	geom, _ := p.Parse()

	t, err := p.scanToken()
	if err != nil {
		return geom, err
	}
	if t.ttype != EOF {
		return geom, fmt.Errorf("parse point unexpected token %s on pos %d", t.lexeme, t.pos)
	}

	return geom, err
}

// MarshalString decode to string
func MarshalString(geom space.Geometry) string {
	buf := bytes.NewBuffer(nil)
	wkt(buf, geom)
	return buf.String()
}

func wkt(buf *bytes.Buffer, geometry space.Geometry) {
	if geometry == nil {
		buf.Write([]byte(``))
		return
	}
	switch geometry.(type) {
	case *space.GeometryValid:
		_, _ = fmt.Fprintf(buf, "SRID=%v;", geometry.CoordinateSystem())
	}

	geom := geometry.Geom()
	switch geom.GeoJSONType() {
	case space.TypePoint:
		if geom.IsEmpty() {
			buf.Write([]byte(`POINT EMPTY`))
			return
		}
		_, _ = fmt.Fprintf(buf, "POINT(%g %g)", geom.(space.Point).Lon(), geom.(space.Point).Lat())
	case space.TypeMultiPoint:
		if geom.IsEmpty() {
			buf.Write([]byte(`MULTIPOINT EMPTY`))
			return
		}
		buf.Write([]byte(`MULTIPOINT(`))
		for i, p := range geom.(space.MultiPoint) {
			if i != 0 {
				buf.WriteByte(',')
			}
			_, _ = fmt.Fprintf(buf, "(%g %g)", p.Lon(), p.Lat())
		}
		buf.WriteByte(')')
	case space.TypeLineString:
		if geom.IsEmpty() {
			buf.Write([]byte(`LINESTRING EMPTY`))
			return
		}

		buf.Write([]byte(`LINESTRING`))
		writeLineString(buf, geom.(space.LineString))
	case space.TypeMultiLineString:
		if geom.IsEmpty() {
			buf.Write([]byte(`MULTILINESTRING EMPTY`))
			return
		}

		buf.Write([]byte(`MULTILINESTRING(`))
		for i, ls := range geom.(space.MultiLineString) {
			if i != 0 {
				buf.WriteByte(',')
			}
			writeLineString(buf, ls)
		}
		buf.WriteByte(')')
	case space.TypePolygon:
		if geom.IsEmpty() {
			buf.Write([]byte(`POLYGON EMPTY`))
			return
		}

		buf.Write([]byte(`POLYGON(`))
		for i, r := range geom.(space.Polygon) {
			if i != 0 {
				buf.WriteByte(',')
			}
			writeLineString(buf, space.LineString(r))
		}
		buf.WriteByte(')')
	case space.TypeMultiPolygon:
		if geom.IsEmpty() {
			buf.Write([]byte(`MULTIPOLYGON EMPTY`))
			return
		}

		buf.Write([]byte(`MULTIPOLYGON(`))
		for i, p := range geom.(space.MultiPolygon) {
			if i != 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte('(')
			for j, r := range p {
				if j != 0 {
					buf.WriteByte(',')
				}
				writeLineString(buf, space.LineString(r))
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case space.TypeCollection:
		if geom.IsEmpty() {
			buf.Write([]byte(`GEOMETRYCOLLECTION EMPTY`))
			return
		}
		buf.Write([]byte(`GEOMETRYCOLLECTION(`))
		for i, c := range geom.(space.Collection) {
			if i != 0 {
				buf.WriteByte(',')
			}
			wkt(buf, c)
		}
		buf.WriteByte(')')
	default:
		buf.Write([]byte(``))
	}
}

func writeLineString(buf *bytes.Buffer, ls space.LineString) {
	buf.WriteByte('(')
	for i, p := range ls.ToPointArray() {
		if i != 0 {
			buf.WriteByte(',')
		}

		_, _ = fmt.Fprintf(buf, "%g %g", p.Lon(), p.Lat())
	}
	buf.WriteByte(')')
}
