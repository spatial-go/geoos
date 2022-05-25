package wkb

import (
	"errors"
	"io"
	"math"

	"github.com/spatial-go/geoos/space"
)

func unmarshalLineString(order byteOrder, data []byte) (space.LineString, error) {
	ps, err := unmarshalPoints(order, data)
	if err != nil {
		return nil, err
	}
	var line space.LineString
	for _, p := range ps {
		line = append(line, p)
	}
	return line, nil
}

func readLineString(r io.Reader, order byteOrder, buf []byte) (space.LineString, error) {
	num, err := readUint32(r, order, buf[:4])
	if err != nil {
		return nil, err
	}

	alloc := num
	if alloc > maxPointsAlloc {
		// invalid data can come in here and allocate tons of memory.
		alloc = maxPointsAlloc
	}
	result := make(space.LineString, 0, alloc)

	for i := 0; i < int(num); i++ {
		p, err := readPoint(r, order, buf)
		if err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}

func (e *Encoder) writeLineString(ls space.LineString) error {
	e.order.PutUint32(e.buf, lineStringType)
	e.order.PutUint32(e.buf[4:], uint32(len(ls)))
	_, err := e.w.Write(e.buf[:8])
	if err != nil {
		return err
	}

	for _, p := range ls {
		e.order.PutUint64(e.buf, math.Float64bits(p[0]))
		e.order.PutUint64(e.buf[8:], math.Float64bits(p[1]))
		_, err = e.w.Write(e.buf)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalMultiLineString(order byteOrder, data []byte) (space.MultiLineString, error) {
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
	result := make(space.MultiLineString, 0, alloc)

	for i := 0; i < int(num); i++ {
		ls, err := scanLineString(data)
		if err != nil {
			return nil, err
		}

		data = data[16*len(ls)+9:]
		result = append(result, ls)
	}

	return result, nil
}

func readMultiLineString(r io.Reader, order byteOrder, buf []byte) (space.MultiLineString, error) {
	num, err := readUint32(r, order, buf[:4])
	if err != nil {
		return nil, err
	}

	alloc := num
	if alloc > maxMultiAlloc {
		// invalid data can come in here and allocate tons of memory.
		alloc = maxMultiAlloc
	}
	result := make(space.MultiLineString, 0, alloc)

	for i := 0; i < int(num); i++ {
		lOrder, typ, err := readByteOrderType(r, buf)
		if err != nil {
			return nil, err
		}

		if typ != lineStringType {
			return nil, errors.New("expect multilines to contains lines, did not find a line")
		}

		ls, err := readLineString(r, lOrder, buf)
		if err != nil {
			return nil, err
		}

		result = append(result, ls)
	}

	return result, nil
}

func (e *Encoder) writeMultiLineString(mls space.MultiLineString) error {
	e.order.PutUint32(e.buf, multiLineStringType)
	e.order.PutUint32(e.buf[4:], uint32(len(mls)))
	_, err := e.w.Write(e.buf[:8])
	if err != nil {
		return err
	}

	for _, ls := range mls {
		err := e.Encode(ls)
		if err != nil {
			return err
		}
	}

	return nil
}
