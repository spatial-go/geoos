package wkb

import (
	"io"
	"log"
	"math"

	"github.com/spatial-go/geoos/space"
)

// BufferedWriter returns []Geometry from reader.
func BufferedWriter(bufferedWriter io.Writer, geoms []space.Geometry) {
	ewkb := &EWKBEncoder{
		Encoder: NewEncoder(bufferedWriter),
		Srid:    space.WGS84,
	}
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
func (e *EWKBEncoder) writeGeometryType(geometryType uint32) {
	//  flag3D := (outputDimension == 3) ? 0x80000000 : 0;
	flag3D := uint32(0)
	typeInt := geometryType | flag3D
	if e.Srid != 0 {
		typeInt |= 0x20000000
	} else {
		typeInt |= 0
	}
	buf := make([]byte, 4)
	e.order.PutUint32(buf, uint32(typeInt))
	if _, err := e.w.Write(buf); err != nil {
		log.Println(err)
	}
	if e.Srid != 0 {
		e.order.PutUint32(buf, e.Srid)
		if _, err := e.w.Write(buf); err != nil {
			log.Println(err)
		}
	}
}
func (e *EWKBEncoder) writePoint(p space.Point) (err error) {
	e.writeGeometryType(pointType)

	e.order.PutUint64(e.buf, math.Float64bits(p[0]))
	e.order.PutUint64(e.buf[8:], math.Float64bits(p[1]))
	_, err = e.w.Write(e.buf)
	return err
}

// TODO rewrite others types
