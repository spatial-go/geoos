package space

import (
	"errors"

	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/overlay"
)

// const geomtype
const (
	TypePoint      = "Point"
	TypeMultiPoint = "MultiPoint"

	TypeLineString      = "LineString"
	TypeMultiLineString = "MultiLineString"

	TypePolygon      = "Polygon"
	TypeMultiPolygon = "MultiPolygon"

	TypeCollection = "GeometryCollection"

	TypeBound = "Bound"
)

// Line  straight line  .
type Line struct {
	Start, End Point
}

// IsIntersection returns intersection of a and other.
func (l *Line) IsIntersection(o *Line) bool {
	mark, _ := overlay.Intersection(&algorithm.Vertex{Matrix: matrix.Matrix(l.Start)},
		&algorithm.Vertex{Matrix: matrix.Matrix(l.End)},
		&algorithm.Vertex{Matrix: matrix.Matrix(o.Start)},
		&algorithm.Vertex{Matrix: matrix.Matrix(o.End)})
	return mark
}

// Intersection returns intersection of a and other.
func (l *Line) Intersection(o *Line) (bool, Point) {
	mark, ip := overlay.Intersection(&algorithm.Vertex{Matrix: matrix.Matrix(l.Start)},
		&algorithm.Vertex{Matrix: matrix.Matrix(l.End)},
		&algorithm.Vertex{Matrix: matrix.Matrix(o.Start)},
		&algorithm.Vertex{Matrix: matrix.Matrix(o.End)})
	return mark, Point(ip.Matrix)
}

// Element describes a geographic Element
type Element struct {
	Geometry
}

// distanceWithFunc returns distance Between the two Geometry.
func (el *Element) distanceWithFunc(g Geometry, f measure.Distance) (float64, error) {
	if el.IsEmpty() && g.IsEmpty() {
		return 0, nil
	}
	if el.IsEmpty() != g.IsEmpty() {
		return 0, errors.New("Geometry is nil")
	}
	switch g.GeoJSONType() {
	case TypePoint:
		if el.GeoJSONType() == TypePoint {
			return el.distancePointWithFunc(g, f)
		}
		elem := &Element{g}
		return elem.distanceWithFunc(el.Geometry, f)
	case TypeLineString:
		if el.GeoJSONType() == TypePoint {
			return el.distancePointWithFunc(g, f)
		} else if el.GeoJSONType() == TypeLineString {
			return el.distanceLineWithFunc(g, f)
		}
		elem := &Element{g}
		return elem.distanceWithFunc(el.Geometry, f)
	case TypePolygon:
		if el.GeoJSONType() == TypePoint {
			return el.distancePointWithFunc(g, f)
		} else if el.GeoJSONType() == TypeLineString {
			return el.distanceLineWithFunc(g, f)
		} else if el.GeoJSONType() == TypePolygon {
			var dist float64
			for _, v := range el.Geometry.(Polygon) {
				elem := &Element{LineString(v)}
				if distP, _ := elem.distanceWithFunc(g, f); dist > distP {
					dist = distP
				}
			}
			return dist, nil
		}
		elem := &Element{g}
		return elem.distanceWithFunc(el.Geometry, f)
	case TypeMultiPoint:
		var dist float64
		for _, v := range g.(MultiPoint) {
			if distP, _ := el.distanceWithFunc(v, f); dist > distP {
				dist = distP
			}
		}
		return dist, nil
	case TypeMultiLineString:
		var dist float64
		for _, v := range g.(MultiLineString) {
			if distP, _ := el.distanceWithFunc(v, f); dist > distP {
				dist = distP
			}
		}
		return dist, nil
	case TypeMultiPolygon:
		var dist float64
		for _, v := range g.(MultiPolygon) {
			if distP, _ := el.distanceWithFunc(v, f); dist > distP {
				dist = distP
			}
		}
		return dist, nil
	case TypeCollection:
		var dist float64
		for _, v := range g.(Collection) {
			if distP, err := el.distanceWithFunc(v, f); err == nil && dist > distP {
				dist = distP
			}
		}
		return dist, nil
	case TypeBound:
		elem := &Element{el.Bound().ToRing()}
		return elem.distanceWithFunc(g, f)
	default:
		return 0, nil
	}
}

// distancePointWithFunc returns distance Between the two Geometry.
func (el *Element) distancePointWithFunc(g Geometry, f measure.Distance) (float64, error) {
	switch g.GeoJSONType() {
	case TypePoint:
		return f(matrix.Matrix(el.Geometry.(Point)), matrix.Matrix(g.(Point))), nil
	case TypeLineString:
		return measure.DistanceLineToPoint(matrix.LineMatrix(g.(LineString)), matrix.Matrix(el.Geometry.(Point)), f), nil
	case TypePolygon:
		return measure.DistancePolygonToPoint(matrix.PolygonMatrix(g.(Polygon)), matrix.Matrix(el.Geometry.(Point)), f), nil
	default:
		return 0, errors.New("Wrong usage function distancePointWithFunc")
	}
}

// distanceLineWithFunc returns distance Between the two Geometry.
func (el *Element) distanceLineWithFunc(g Geometry, f measure.Distance) (float64, error) {
	switch g.GeoJSONType() {
	case TypeLineString:
		var dist float64
		if mark := IsIntersectionLineString(el.Geometry.(LineString), g.(LineString)); mark {
			return 0, nil
		}
		for _, v := range el.Geometry.(LineString) {
			elem := &Element{Point(v)}
			if distP, _ := elem.distanceWithFunc(g, f); dist > distP {
				dist = distP
			}
		}
		return dist, nil
	case TypePolygon:
		var dist float64
		for _, v := range g.(Polygon) {
			elem := &Element{LineString(v)}
			if distP, _ := elem.distanceWithFunc(el, f); dist > distP {
				dist = distP
			}
		}
		return dist, nil
	default:
		return 0, errors.New("Wrong usage function distanceLineWithFunc")
	}
}

// IsIntersectionLineString returns intersection of edge a and b.
func IsIntersectionLineString(aLine, bLine LineString) bool {
	aEdge, bEdge := &algorithm.Edge{Vertexs: []algorithm.Vertex{}}, &algorithm.Edge{Vertexs: []algorithm.Vertex{}}

	for _, v := range aLine {
		aEdge.Vertexs = append(aEdge.Vertexs, algorithm.Vertex{Matrix: v})
	}
	for _, v := range bLine {
		bEdge.Vertexs = append(aEdge.Vertexs, algorithm.Vertex{Matrix: v})
	}
	return overlay.IsIntersectionEdge(*aEdge, *bEdge)
}

// IntersectionLineString returns intersection of edge a and b.
func IntersectionLineString(aLine, bLine LineString) (bool, []Point) {
	aEdge, bEdge := &algorithm.Edge{Vertexs: []algorithm.Vertex{}}, &algorithm.Edge{Vertexs: []algorithm.Vertex{}}

	for _, v := range aLine {
		aEdge.Vertexs = append(aEdge.Vertexs, algorithm.Vertex{Matrix: v})
	}
	for _, v := range bLine {
		bEdge.Vertexs = append(aEdge.Vertexs, algorithm.Vertex{Matrix: v})
	}
	mark, ps := overlay.IntersectionEdge(*aEdge, *bEdge)
	intersectPoints := []Point{}
	for _, v := range ps {
		intersectPoints = append(intersectPoints, Point(v.Matrix))
	}
	return mark, intersectPoints
}
