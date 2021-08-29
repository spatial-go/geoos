// Package wkb is for decoding ESRI's Well Known Binary (WKB) format
// specification at https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary
package wkb

import (
	"io"

	"github.com/spatial-go/geoos/space"
)

// BufferedWriter returns []Geometry from reader.
func BufferedWriter(bufferedWriter io.Writer, geoms []space.Geometry) {
	ewkb := &EWKBEncoder{Encoder: NewEncoder(bufferedWriter)}
	for _, v := range geoms {
		_ = ewkb.Encode(v)
	}
}

// EWKBEncoder Decoder can decoder EWKB geometry off of the stream.
type EWKBEncoder struct {
	*Encoder
	Srid uint32
}

// Encode will write the geometry encoded as WKB to the given writer.
func (e *EWKBEncoder) Encode(geom space.Geometry) error {
	if geom == nil || geom.IsEmpty() {
		return nil
	}

	switch g := geom.(type) {
	// deal with types that are not supported by wkb
	case space.Ring:
		geom = space.Polygon{g}
	case space.Bound:
		if g.Max == nil || g.Min == nil {
			return nil
		}
		geom = g.ToPolygon()
	}

	var b []byte
	if e.order == littleEndian {
		b = []byte{1}
	} else {
		b = []byte{0}
	}

	_, err := e.w.Write(b)
	if err != nil {
		return err
	}

	if e.Srid != 0 {
		buf := make([]byte, 4)
		e.order.PutUint32(buf, e.Srid)

		_, err := e.w.Write(buf)
		if err != nil {
			return err
		}
	}

	if e.buf == nil {
		e.buf = make([]byte, 16)
	}

	switch g := geom.(type) {
	case space.Point:
		return e.writePoint(g)
	case space.MultiPoint:
		return e.writeMultiPoint(g)
	case space.LineString:
		return e.writeLineString(g)
	case space.MultiLineString:
		return e.writeMultiLineString(g)
	case space.Polygon:
		return e.writePolygon(g)
	case space.MultiPolygon:
		return e.writeMultiPolygon(g)
	case space.Collection:
		return e.writeCollection(g)
	}

	return ErrUnknownWKBType
}
