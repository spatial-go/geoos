package wkb

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

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
	order          byteOrder
	inputDimension int
	Srid           uint32
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
	return d.order.Uint32(buf), nil
}

// Decode returns geometry,it will decode the next geometry off of the stream.
func (d *EWKBDecoder) Decode() (space.Geometry, error) {

	// determine byte order
	byteOrderWKB, _ := d.readByte()

	// always set byte order, since it may change from geometry to geometry
	d.order = byteOrder(byteOrderWKB)

	//if not strict and not XDR or NDR, then we just use the dis default set at the
	//start of the geometry (if a multi-geometry).  This  allows WBKReader to work
	//with Spatialite native BLOB WKB, as well as other WKB variants that might just
	//specify endian-ness at the start of the multigeometry.
	typeInt, _ := d.readInt32()

	// To get geometry type mask out EWKB flag bits,and use only low 3 digits of type word.
	// This supports both EWKB and ISO/OGC.
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
		d.Srid, _ = d.readInt32()
		//fmt.Println(srid)
	}

	var buf = make([]byte, 8)
	order := byteOrder(d.order)

	var geom space.Geometry
	var err error
	switch uint32(geometryType) {
	case pointType:
		geom, err = readPoint(d.r, order, buf)
	case lineStringType:
		geom, err = readLineString(d.r, order, buf)
	case polygonType:
		geom, err = readPolygon(d.r, order, buf)
	case multiPointType:
		geom, err = readMultiPoint(d.r, order, buf)
	case multiLineStringType:
		geom, err = readMultiLineString(d.r, order, buf)
	case multiPolygonType:
		geom, err = readMultiPolygon(d.r, order, buf)
	case geometryCollectionType:
		geom, err = readCollection(d.r, order, buf)
	default:
		return nil, ErrUnknownWKBType
	}
	if err != nil {
		return nil, err
	}
	valid, err := space.CreateElementValidWithCoordSys(geom, int(d.Srid))
	return valid, err
}

// HexToBytes Converts a hexadecimal string to a byte array. The hexadecimal digit symbols are case-insensitive.
func HexToBytes(hex string) []byte {
	byteLen := len(hex) / 2
	bytes := make([]byte, byteLen)

	for i := 0; i < byteLen; i++ {
		i2 := 2 * i
		//TODO
		// if i2+1 > len(hex) {
		// 	//throw new IllegalArgumentException("Hex string has odd length");
		// }

		nib1 := hexToInt(hex[i2])
		nib0 := hexToInt(hex[i2+1])
		b := byte((nib1 << 4)) + byte(nib0)
		bytes[i] = b
	}
	return bytes
}

func hexToInt(hex byte) uint64 {
	nib, _ := strconv.ParseUint(string(hex), 16, 32)
	return nib
}
