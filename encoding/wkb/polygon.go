package wkb

import (
	"errors"
	"io"
	"math"

	"github.com/spatial-go/geoos/space"
)

func unmarshalPolygon(order byteOrder, data []byte) (space.Polygon, error) {
	if len(data) < 4 {
		return nil, ErrNotWKB
	}
	num := unmarshalUint32(order, data)
	data = data[4:]

	alloc := num
	if alloc > maxMultiAlloc {
		// invalid data can come in here and allocate tons of memory.
		alloc = maxMultiAlloc
	}
	result := make(space.Polygon, 0, alloc)

	for i := 0; i < int(num); i++ {
		ps, err := unmarshalPoints(order, data)
		if err != nil {
			return nil, err
		}

		data = data[16*len(ps)+4:]

		var line space.LineString
		for _, p := range ps {
			line = append(line, p)
		}
		result = append(result, line)
	}

	return result, nil
}

func readPolygon(r io.Reader, order byteOrder, buf []byte) (space.Polygon, error) {
	num, err := readUint32(r, order, buf[:4])
	if err != nil {
		return nil, err
	}

	alloc := num
	if alloc > maxMultiAlloc {
		// invalid data can come in here and allocate tons of memory.
		alloc = maxMultiAlloc
	}
	result := make(space.Polygon, 0, alloc)

	for i := 0; i < int(num); i++ {
		ls, err := readLineString(r, order, buf)
		if err != nil {
			return nil, err
		}

		result = append(result, space.Ring(ls))
	}

	return result, nil
}

func (e *Encoder) writePolygon(p space.Polygon) error {
	e.order.PutUint32(e.buf, polygonType)
	e.order.PutUint32(e.buf[4:], uint32(len(p)))
	_, err := e.w.Write(e.buf[:8])
	if err != nil {
		return err
	}
	for _, r := range p {
		e.order.PutUint32(e.buf, uint32(len(r)))
		_, err := e.w.Write(e.buf[:4])
		if err != nil {
			return err
		}
		for _, p := range r {
			e.order.PutUint64(e.buf, math.Float64bits(p[0]))
			e.order.PutUint64(e.buf[8:], math.Float64bits(p[1]))
			_, err = e.w.Write(e.buf)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func unmarshalMultiPolygon(order byteOrder, data []byte) (space.MultiPolygon, error) {
	if len(data) < 4 {
		return nil, ErrNotWKB
	}
	num := unmarshalUint32(order, data)
	data = data[4:]

	alloc := num
	if alloc > maxMultiAlloc {
		// invalid data can come in here and allocate tons of memory.
		alloc = maxMultiAlloc
	}
	result := make(space.MultiPolygon, 0, alloc)

	for i := 0; i < int(num); i++ {
		p, err := scanPolygon(data)
		if err != nil {
			return nil, err
		}

		l := 9
		for _, r := range p {
			l += 4 + 16*len(r)
		}
		data = data[l:]

		result = append(result, p)
	}

	return result, nil
}

func readMultiPolygon(r io.Reader, order byteOrder, buf []byte) (space.MultiPolygon, error) {
	num, err := readUint32(r, order, buf[:4])
	if err != nil {
		return nil, err
	}

	alloc := num
	if alloc > maxMultiAlloc {
		// invalid data can come in here and allocate tons of memory.
		alloc = maxMultiAlloc
	}
	result := make(space.MultiPolygon, 0, alloc)

	for i := 0; i < int(num); i++ {
		pOrder, typ, err := readByteOrderType(r, buf)
		if err != nil {
			return nil, err
		}

		if typ != polygonType {
			return nil, errors.New("expect multipolygons to contains polygons, did not find a polygon")
		}

		p, err := readPolygon(r, pOrder, buf)
		if err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}

func (e *Encoder) writeMultiPolygon(mp space.MultiPolygon) error {
	e.order.PutUint32(e.buf, multiPolygonType)
	e.order.PutUint32(e.buf[4:], uint32(len(mp)))
	_, err := e.w.Write(e.buf[:8])
	if err != nil {
		return err
	}

	for _, p := range mp {
		err := e.Encode(p)
		if err != nil {
			return err
		}
	}

	return nil
}
