// Package wkb is for decoding ESRI's Well Known Binary (WKB) format
// specification at https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary
package wkb

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"

	"github.com/spatial-go/geoos/space"
)

// byteOrder represents little or big endian encoding.
// We don't use binary.ByteOrder because that is an interface
// that leaks to the heap all over the place.
type byteOrder int

const bigEndian byteOrder = 0
const littleEndian byteOrder = 1

const (
	pointType              uint32 = 1
	lineStringType         uint32 = 2
	polygonType            uint32 = 3
	multiPointType         uint32 = 4
	multiLineStringType    uint32 = 5
	multiPolygonType       uint32 = 6
	geometryCollectionType uint32 = 7
)

const (
	// limits so that bad data can't come in and preallocate tons of memory.
	// Well formed data with less elements will allocate the correct amount just fine.
	maxPointsAlloc = 10000
	maxMultiAlloc  = 100
)

// DefaultByteOrder is the order used for marshalling or encoding
// is none is specified.
var DefaultByteOrder binary.ByteOrder = binary.LittleEndian

func (b byteOrder) Uint32(buf []byte) uint32 {
	if b == littleEndian {
		return DefaultByteOrder.Uint32(buf)
	}
	return binary.BigEndian.Uint32(buf)
}

func (b byteOrder) Uint64(buf []byte) uint64 {
	if b == littleEndian {
		return DefaultByteOrder.Uint64(buf)
	}
	return binary.BigEndian.Uint64(buf)
}

func (b byteOrder) PutUint32(buf []byte, u uint32) {
	if b == littleEndian {
		DefaultByteOrder.PutUint32(buf, u)
	} else {
		binary.BigEndian.PutUint32(buf, u)
	}
}

func (b byteOrder) PutUint64(buf []byte, u uint64) {
	if b == littleEndian {
		DefaultByteOrder.PutUint64(buf, u)
	} else {
		binary.BigEndian.PutUint64(buf, u)
	}
}

// An Encoder will encode a geometry as WKB to the writer given at
// creation time.
type Encoder struct {
	buf []byte

	w     io.Writer
	order byteOrder
}

// MustMarshal will encode the geometry and panic on error.
// Currently there is no reason to error during geometry marshalling.
func MustMarshal(geom space.Geometry, bo ...byteOrder) []byte {
	d, err := Marshal(geom, bo...)
	if err != nil {
		panic(err)
	}

	return d
}

// Marshal encodes the geometry with the given byte order.
func Marshal(geom space.Geometry, bo ...byteOrder) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, geomLength(geom)))

	e := NewEncoder(buf)
	if len(bo) > 0 {
		e.order = bo[0]
	}

	err := e.Encode(geom)
	if err != nil {
		return nil, err
	}

	if buf.Len() == 0 {
		return nil, nil
	}

	return buf.Bytes(), nil
}

// NewEncoder creates a new Encoder for the given writer.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:     w,
		order: 1,
	}
}

// Encode will write the geometry encoded as WKB to the given writer.
func (e *Encoder) Encode(geom space.Geometry) error {
	if geom == nil || geom.IsEmpty() {
		return nil
	}

	switch g := geom.(type) {
	// deal with types that are not supported by wkb
	case space.Ring:
		if g == nil {
			return nil
		}
		geom = space.Polygon{g}
	case space.Bound:
		if g.Max == nil || g.Min == nil {
			return nil
		}
		geom = g.ToPolygon()
	}

	b := []byte{byte(e.order)}

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

// Decoder can decoder WKB geometry off of the stream.
type Decoder struct {
	r io.Reader
}

// Unmarshal will decode the type into a Geometry.
func Unmarshal(data []byte) (space.Geometry, error) {
	order, typ, data, err := unmarshalByteOrderType(data)
	if err != nil {
		return nil, err
	}

	switch typ {
	case pointType:
		return unmarshalPoint(order, data[5:])
	case multiPointType:
		return unmarshalMultiPoint(order, data[5:])
	case lineStringType:
		return unmarshalLineString(order, data[5:])
	case multiLineStringType:
		return unmarshalMultiLineString(order, data[5:])
	case polygonType:
		return unmarshalPolygon(order, data[5:])
	case multiPolygonType:
		return unmarshalMultiPolygon(order, data[5:])
	case geometryCollectionType:
		g, err := NewDecoder(bytes.NewReader(data)).Decode()
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil, ErrNotWKB
		}

		return g, err
	}

	return nil, ErrUnsupportedGeometry
}

// NewDecoder will create a new WKB decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

// Decode will decode the next geometry off of the stream.
func (d *Decoder) Decode() (space.Geometry, error) {
	buf := make([]byte, 8)
	order, typ, err := readByteOrderType(d.r, buf)
	if err != nil {
		return nil, err
	}

	switch typ {
	case pointType:
		return readPoint(d.r, order, buf)
	case multiPointType:
		return readMultiPoint(d.r, order, buf)
	case lineStringType:
		return readLineString(d.r, order, buf)
	case multiLineStringType:
		return readMultiLineString(d.r, order, buf)
	case polygonType:
		return readPolygon(d.r, order, buf)
	case multiPolygonType:
		return readMultiPolygon(d.r, order, buf)
	case geometryCollectionType:
		return readCollection(d.r, order, buf)
	}

	return nil, ErrUnsupportedGeometry
}

func readByteOrderType(r io.Reader, buf []byte) (byteOrder, uint32, error) {
	// the byte order is the first byte
	if _, err := r.Read(buf[:1]); err != nil {
		return 0, 0, err
	}

	var order byteOrder
	if buf[0] == 0 {
		order = bigEndian
	} else if buf[0] == 1 {
		order = littleEndian
	} else {
		return 0, 0, ErrNotWKB
	}

	// the type which is 4 bytes
	typ, err := readUint32(r, order, buf[:4])

	if err != nil {
		return 0, 0, err
	}

	return order, typ, nil
}

func readUint32(r io.Reader, order byteOrder, buf []byte) (uint32, error) {
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return unmarshalUint32(order, buf), nil
}

func unmarshalByteOrderType(buf []byte) (byteOrder, uint32, []byte, error) {
	order, typ, err := byteOrderType(buf)
	if err == nil {
		return order, typ, buf, nil
	}

	if len(buf) < 6 {
		return 0, 0, nil, err
	}

	// The prefix is incorrect, let's see if this is data in
	// MySQL's SRID+WKB format. So truncate the SRID prefix.
	buf = buf[4:]
	order, typ, err = byteOrderType(buf)
	if err != nil || typ > 7 {
		return 0, 0, nil, ErrNotWKB
	}

	return order, typ, buf, nil
}

func byteOrderType(buf []byte) (byteOrder, uint32, error) {
	if len(buf) < 6 {
		return 0, 0, ErrNotWKB
	}

	var order byteOrder
	if buf[0] == 0 {
		order = bigEndian
	} else if buf[0] == 1 {
		order = littleEndian
	} else {
		return 0, 0, ErrNotWKB
	}

	// the type which is 4 bytes
	typ := unmarshalUint32(order, buf[1:])
	return order, typ, nil
}

func unmarshalUint32(order byteOrder, buf []byte) uint32 {
	if order == littleEndian {
		return binary.LittleEndian.Uint32(buf)
	}
	return binary.BigEndian.Uint32(buf)
}

// geomLength helps to do preallocation during a marshal.
func geomLength(geom space.Geometry) int {
	switch g := geom.(type) {
	case space.Point:
		return 21
	case space.MultiPoint:
		return 9 + 21*len(g)
	case space.LineString:
		return 9 + 16*len(g)
	case space.MultiLineString:
		sum := 0
		for _, ls := range g {
			sum += 9 + 16*len(ls)
		}

		return 9 + sum
	case space.Polygon:
		sum := 0
		for _, r := range g {
			sum += 4 + 16*len(r)
		}

		return 9 + sum
	case space.MultiPolygon:
		sum := 0
		for _, c := range g {
			sum += geomLength(c)
		}

		return 9 + sum
	case space.Collection:
		sum := 0
		for _, c := range g {
			sum += geomLength(c)
		}

		return 9 + sum
	}

	return 0
}

// GeomFromWKBHexStr convert hex string to GEOSGeometry
func GeomFromWKBHexStr(wkbHex string) (space.Geometry, error) {
	wkbStr := HexToBytes(wkbHex)
	//wkbStr, err := hex.DecodeString(wkbHex)

	ewkb := &EWKBDecoder{r: bytes.NewReader(wkbStr)}
	g, _ := ewkb.Decode()

	return g, nil
}

// GeomToWKBHexStr ...
func GeomToWKBHexStr(geom space.Geometry) (wkbHex string, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, geomLength(geom)))

	BufferedWriter(buf, []space.Geometry{geom})
	bytes0 := buf.Bytes()

	return hex.EncodeToString(bytes0), err
}
