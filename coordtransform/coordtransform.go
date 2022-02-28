package coordtransform

import (
	"errors"
	"sync"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// CoordType Transformer
const (
	MERCATORTOLL = "MERCATORTOLL"
	LLTOMERCATOR = "LLTOMERCATOR"
)

// Transformer ...
type Transformer struct {
	CoordType string
}

var instance *Transformer
var once sync.Once

// GetInstance ...
func GetInstance() *Transformer {
	once.Do(func() {
		instance = &Transformer{}
	})
	return instance
}

// NewTransformer ...
func NewTransformer(coordType string) *Transformer {
	return &Transformer{CoordType: coordType}
}

// TransformLatLng ...
func (t *Transformer) TransformLatLng(lng, lat float64) (float64, float64) {
	switch t.CoordType {
	case MERCATORTOLL:
		lng, lat = MercatorToLL(lng, lat)
	case LLTOMERCATOR:
		lng, lat = LLToMercator(lng, lat)
	default:
	}
	return lng, lat
}

// TransformPoint ...
func (t *Transformer) TransformPoint(point matrix.Matrix) matrix.Matrix {
	lng, lat := t.TransformLatLng(point[0], point[1])
	return matrix.Matrix{lng, lat}
}

// TransformMultiPoint ...
func (t *Transformer) TransformMultiPoint(multiPoint []matrix.Matrix) []matrix.Matrix {
	for i := range multiPoint {
		multiPoint[i] = t.TransformPoint(multiPoint[i])
	}
	return multiPoint
}

// TransformLine ...
func (t *Transformer) TransformLine(lineString matrix.LineMatrix) matrix.LineMatrix {
	for i := range lineString {
		lineString[i] = t.TransformPoint(lineString[i])
	}
	return lineString
}

// TransformPolygon ...
func (t *Transformer) TransformPolygon(polygon matrix.PolygonMatrix) matrix.PolygonMatrix {
	for i := range polygon {
		polygon[i] = t.TransformLine(polygon[i])
	}
	return polygon
}

// TransformMultiLineString ...
func (t *Transformer) TransformMultiLineString(multiLineString []matrix.LineMatrix) []matrix.LineMatrix {
	for i := range multiLineString {
		multiLineString[i] = t.TransformLine(multiLineString[i])
	}
	return multiLineString
}

// TransformMultiPolygon ...
func (t *Transformer) TransformMultiPolygon(multiPolygon []matrix.PolygonMatrix) []matrix.PolygonMatrix {
	for i := range multiPolygon {
		multiPolygon[i] = t.TransformPolygon(multiPolygon[i])
	}
	return multiPolygon
}

// TransformGeometry ...
func (t *Transformer) TransformGeometry(geom matrix.Steric) (matrix.Steric, error) {
	switch mt := geom.(type) {
	case matrix.Matrix:
		return t.TransformPoint(mt), nil
	case matrix.LineMatrix:
		return t.TransformLine(mt), nil
	case matrix.PolygonMatrix:
		return t.TransformPolygon(mt), nil
	case matrix.Collection:
		for i := range mt {
			mt[i], _ = t.TransformGeometry(mt[i])
		}
		return mt, nil
	default:
		return nil, errors.New("error geometry type")
	}
}
