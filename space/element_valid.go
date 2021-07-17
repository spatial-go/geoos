package space

import (
	"container/ring"
)

// ElementValid describes a geographic Element Valid
type ElementValid struct {
	Geometry
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (el *ElementValid) IsClosed() bool {
	switch el.GeoJSONType() {
	case TypeLineString:
		return el.Geometry.(LineString).IsClosed()
	case TypeMultiLineString:
		return el.Geometry.(MultiLineString).IsClosed()
	}
	return true
}

// IsSimple Computes simplicity for geometries.
func (el *ElementValid) IsSimple() bool {
	if el.IsEmpty() {
		return true
	}
	switch el.GeoJSONType() {
	case TypePoint:
		return true
	case TypeLineString:
		return el.isSimpleLine()
	case TypeMultiLineString:
		return el.isSimpleMultiLine()
	case TypeMultiPoint:
		return el.isSimpleMultiPoint()
	case TypePolygon:
		return el.isSimplePolygon()
	case TypeMultiPolygon:
		return el.isSimpleMultiPolygon()
	case TypeCollection:
		return el.isSimpleCollection()
	default:
		return true

	}
}

// isSimpleMultiPoint Computes simplicity for MultiPoint geometries.
func (el *ElementValid) isSimpleMultiPoint() bool {
	points := ring.New(len(el.Geometry.(MultiPoint)))
	nonSimplePts := true
	for _, v := range el.Geometry.(MultiPoint) {
		points.Do(func(i interface{}) {
			if v.Equal(i.(Geometry)) {
				nonSimplePts = false
			}
		})
		if nonSimplePts {
			points.Value = v
			points = points.Next()
		} else {
			return false
		}
	}
	return true
}

// isSimplePolygon Computes simplicity for polygonal geometries.
// Polygonal geometries are simple if and only if
//  all of their component rings are simple.
func (el *ElementValid) isSimplePolygon() bool {
	for _, ring := range el.Geometry.(Polygon) {
		elem := ElementValid{LineString(ring)}
		if !elem.isSimpleLine() {
			return false
		}
	}
	return true
}

// isSimpleMultiPolygon Computes simplicity for multi  polygonal geometries.
// Polygonal geometries are simple if and only if
//  all of their component rings are simple.
func (el *ElementValid) isSimpleMultiPolygon() bool {
	for _, poly := range el.Geometry.(MultiPolygon) {
		elem := ElementValid{poly}
		if !elem.isSimplePolygon() {
			return false
		}
	}
	return true
}

// isSimpleCollection Computes simplicity for collection geometries.
//  geometries are simple if and only if
//  all geometrie are simple.
func (el *ElementValid) isSimpleCollection() bool {
	for _, g := range el.Geometry.(Collection) {
		elem := ElementValid{g}
		if !elem.IsSimple() {
			return false
		}
	}
	return true
}

// isSimpleLine Computes simplicity for LineString geometries.
// geometries are simple if they do not self-intersect at interior points
// (i.e. points other than the endpoints)..
func (el *ElementValid) isSimpleLine() bool {
	lines := el.Geometry.(LineString).ToLineArray()
	numLine := len(lines)
	for i, line1 := range lines {
		for j, line2 := range lines {
			if i == j || j-i == 1 || i-j == 1 {
				continue
			}
			if line1.IsIntersection(&line2) {
				if (i == 0 && j == numLine-1) ||
					(j == 0 && i == numLine-1) {
					_, ip := line1.Intersection(&line2)
					if ip.Equal(line1.Start) ||
						ip.Equal(line1.End) ||
						ip.Equal(line2.Start) ||
						ip.Equal(line2.End) {
						continue
					}
				}
				return false
			}
		}
	}
	return true
}

// isSimpleMultiLine Computes simplicity for MultiLineString geometries.
// geometries are simple if
// their elements are simple and they intersect only at points
// which are boundary points of both elements.
func (el *ElementValid) isSimpleMultiLine() bool {
	mls := el.Geometry.(MultiLineString)
	for _, v := range mls {
		elem := ElementValid{v}
		if !elem.isSimpleLine() {
			return false
		}

	}

	for _, line1 := range mls {
		for _, line2 := range mls {
			mark, ips := IntersectionLineString(line1, line2)
			if !mark {
				continue
			}
			if len(ips) > 2 {
				return false
			}
			boundarys := MultiPoint{}
			if b, err := line1.Boundary(); err == nil {
				boundarys = append(boundarys, b.(MultiPoint)...)
			}
			if b, err := line2.Boundary(); err == nil {
				boundarys = append(boundarys, b.(MultiPoint)...)
			}

			for _, v := range ips {
				for _, p := range boundarys {
					if v.Equal(p) {
						continue
					}
				}
				return false
			}
		}
	}
	return true
}
