// Package wkb is for decoding ESRI's Well Known Binary (WKB) format
// specification at https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary
package wkb

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/spatial-go/geoos/algorithm/calc/bytevalues"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/space"
)

// ErrUnknownWKBType ...
var ErrUnknownWKBType = fmt.Errorf("Unknown WKB type ")

// ErrAttempt ...
var ErrAttempt = fmt.Errorf("Attempt to read past end of input")

// BufferedReader returns []Geometry from reader.
func BufferedReader(bufferedReader io.Reader) []space.Geometry {
	geoms := []space.Geometry{}
	for {
		buf := make([]byte, 64)
		if line, err := bufferedReader.Read(buf); err == io.EOF {
			break
		} else if err == nil {
			if line == 0 {
				continue
			}

			ewkb := &EWKBDecoder{r: bytes.NewReader(buf)}
			g, _ := ewkb.Decode()
			geoms = append(geoms, g)
		}
	}
	return geoms
}

// EWKBDecoder Decoder can decoder EWKB geometry off of the stream.
type EWKBDecoder struct {
	r              io.Reader
	byteOrder      int
	ordValues      []float64
	inputDimension int

	// true if structurally invalid input should be reported rather than repaired.
	// At some point this could be made client-controllable.
	isStrict bool
}

func (d *EWKBDecoder) readByte() (byte, error) {
	buf := make([]byte, 1)
	if num, err := d.r.Read(buf); err == nil {
		if num < len(buf) {
			return '0', ErrAttempt
		}
	}
	return buf[0], nil
}
func (d *EWKBDecoder) readInt32() (uint32, error) {
	buf := make([]byte, 4)
	if num, err := d.r.Read(buf); err == nil {
		if num < len(buf) {
			return 0, ErrAttempt
		}
	}
	return bytevalues.GetInt32(buf, d.byteOrder), nil
}
func (d *EWKBDecoder) readInt64() (uint64, error) {
	buf := make([]byte, 8)
	if num, err := d.r.Read(buf); err == nil {
		if num < len(buf) {
			return 0, ErrAttempt
		}
	}
	return bytevalues.GetInt64(buf, d.byteOrder), nil
}

func (d *EWKBDecoder) readFloat32() (float32, error) {
	buf := make([]byte, 4)
	if num, err := d.r.Read(buf); err == nil {
		if num < len(buf) {
			return 0, ErrAttempt
		}
	}
	return bytevalues.GetFloat32(buf, d.byteOrder), nil
}

func (d *EWKBDecoder) readFloat64() (float64, error) {
	buf := make([]byte, 8)
	if num, err := d.r.Read(buf); err == nil {
		if num < len(buf) {
			return 0, ErrAttempt
		}
	}
	return bytevalues.GetFloat64(buf, d.byteOrder), nil
}

// Decode returns geometry,it will decode the next geometry off of the stream.
func (d *EWKBDecoder) Decode() (space.Geometry, error) {

	// determine byte order
	byteOrderWKB, _ := d.readByte()

	// always set byte order, since it may change from geometry to geometry
	d.byteOrder = int(byteOrderWKB)

	//if not strict and not XDR or NDR, then we just use the dis default set at the
	//start of the geometry (if a multi-geometry).  This  allows WBKReader to work
	//with Spatialite native BLOB WKB, as well as other WKB variants that might just
	//specify endian-ness at the start of the multigeometry.

	typeInt, _ := d.readInt32()

	/**
	 * To get geometry type mask out EWKB flag bits,
	 * and use only low 3 digits of type word.
	 * This supports both EWKB and ISO/OGC.
	 */
	geometryType := (typeInt & 0xffff) % 1000

	// handle 3D and 4D WKB geometries
	// geometries with Z coordinates have the 0x80 flag (postgis EWKB)
	// or are in the 1000 range (Z) or in the 3000 range (ZM) of geometry type (ISO/OGC 06-103r4)
	hasZ := ((int64(typeInt)&0x80000000) != 0 || (typeInt&0xffff)/1000 == 1 || (typeInt&0xffff)/1000 == 3)
	// geometries with M coordinates have the 0x40 flag (postgis EWKB)
	// or are in the 1000 range (M) or in the 3000 range (ZM) of geometry type (ISO/OGC 06-103r4)
	hasM := ((typeInt&0x40000000) != 0 || (typeInt&0xffff)/1000 == 2 || (typeInt&0xffff)/1000 == 3)
	//System.out.println(typeInt + " - " + geometryType + " - hasZ:" + hasZ);
	d.inputDimension = 2
	if hasZ {
		d.inputDimension++
	}
	if hasM {
		d.inputDimension++
	}
	// determine if SRID are present (EWKB only)
	hasSRID := (typeInt & 0x20000000) != 0
	if hasSRID {
		_, _ = d.readInt32()
		//fmt.Println(srid)
	}

	// only allocate ordValues buffer if necessary
	if d.ordValues == nil || len(d.ordValues) < d.inputDimension {
		d.ordValues = make([]float64, d.inputDimension)
	}
	var geom space.Geometry
	var buf = make([]byte, 8)
	switch uint32(geometryType) {
	case pointType:
		geom, _ = readPoint(d.r, byteOrder(d.byteOrder), buf)
		break
	case lineStringType:
		geom = d.readLineString()
		break
	case polygonType:
		geom = d.readPolygon()
		break
	case multiPointType:
		geom = d.readMultiPoint()
		break
	case multiLineStringType:
		geom = d.readMultiLineString()
		break
	case multiPolygonType:
		geom = d.readMultiPolygon()
		break
	case geometryCollectionType:
		geom = d.readGeometryCollection()
		break
	default:
		return nil, ErrUnknownWKBType
	}
	return geom, nil
}

func (d *EWKBDecoder) readPoint() space.Point {
	pts := d.readMatrixes(1)
	// If X and Y are NaN create a empty point
	if math.IsNaN(pts[0][0]) || math.IsNaN(pts[0][1]) {
		return space.Point(matrix.Matrix{})
	}
	return space.Point(pts[0])
}

func (d *EWKBDecoder) readMatrixes(size int) []matrix.Matrix {
	pts := []matrix.Matrix{}
	for i := 0; i < size; i++ {
		d.readMatrix()
		pt := matrix.Matrix{}
		pt = append(pt, d.ordValues...)
		pts = append(pts, pt)
	}
	return pts
}

/**
 * Reads a coordinate value with the specified dimensionality.
 * Makes the X and Y ordinates precise according to the precision model
 * in use.
 * @throws ParseException
 */
func (d *EWKBDecoder) readMatrix() {
	for i := 0; i < d.inputDimension; i++ {

		rf, _ := d.readFloat64()

		d.ordValues[i] = float64(rf)

	}
}

func (d *EWKBDecoder) readLineString() space.LineString {
	size := d.readNumField("numCoords")
	pts := d.readMatrixes(size)
	ls := space.LineString{}
	for _, v := range pts {
		ls = append(ls, v)
	}
	return ls
}

func (d *EWKBDecoder) readLinearRing() space.Ring {
	size := d.readNumField("numCoords")
	pts := d.readMatrixes(size)
	ls := space.LineString{}
	for _, v := range pts {
		ls = append(ls, v)
	}
	ls = append(ls, ls[0])
	return space.Ring(ls)
}

func (d *EWKBDecoder) readPolygon() space.Polygon {
	numRings := d.readNumField("")
	poly := space.Polygon{}
	for i := 0; i < numRings; i++ {
		poly = append(poly, d.readLinearRing())
	}
	return poly
}
func (d *EWKBDecoder) readMultiPoint() space.MultiPoint {
	numGeom := d.readNumField("")
	geoms := space.MultiPoint{}
	for i := 0; i < numGeom; i++ {
		g := d.readPoint()
		geoms = append(geoms, g)
	}
	return geoms
}

func (d *EWKBDecoder) readMultiLineString() space.MultiLineString {
	numGeom := d.readNumField("")
	geoms := space.MultiLineString{}
	for i := 0; i < numGeom; i++ {
		g := d.readLineString()
		geoms = append(geoms, g)
	}
	return geoms
}

func (d *EWKBDecoder) readMultiPolygon() space.MultiPolygon {
	numGeom := d.readNumField("")
	geoms := space.MultiPolygon{}
	for i := 0; i < numGeom; i++ {
		g := d.readPolygon()
		geoms = append(geoms, g)
	}
	return geoms
}

func (d *EWKBDecoder) readGeometryCollection() space.Collection {
	numGeom := d.readNumField("")
	geoms := space.Collection{}
	for i := 0; i < numGeom; i++ {
		g, _ := d.Decode()
		geoms = append(geoms, g)
	}
	return geoms
}

func (d *EWKBDecoder) readNumField(fieldName string) int {
	num, _ := d.readInt32()
	return int(num)
}

// HexToBytes Converts a hexadecimal string to a byte array. The hexadecimal digit symbols are case-insensitive.
func HexToBytes(hex string) []byte {
	byteLen := len(hex) / 2
	bytes := make([]byte, byteLen)

	for i := 0; i < byteLen; i++ {
		i2 := 2 * i
		if i2+1 > len(hex) {
			//throw new IllegalArgumentException("Hex string has odd length");
		}

		nib1 := hexToInt(hex[i2])
		nib0 := hexToInt(hex[i2+1])
		b := byte((nib1 << 4)) + byte(nib0)
		bytes[i] = b
	}
	return bytes
}

func hexToInt(hex byte) uint64 {
	nib, _ := strconv.ParseUint(string(hex), 16, 32)
	if nib < 0 {
		return 0
	}
	return nib
}
