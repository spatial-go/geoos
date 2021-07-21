package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/operation"
	"github.com/spatial-go/geoos/algorithm/relate"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// Line  straight line  .
type Line struct {
	Start, End Point
}

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
	vop := &operation.ValidOP{Steric: el.ToMatrix()}
	return vop.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func Centroid(geom Geometry) Point {
	cent := &buffer.Centroid{}

	if geom == nil || geom.IsEmpty() {
		return nil
	}
	cent.Add(geom.ToMatrix())
	m := cent.GetCentroid()
	return Point(m)
}

// Relate Computes the  Intersection Matrix for the spatial relationship
// between two Sterics, using the default (OGC SFS) Boundary Node Rule
func Relate(a, b Geometry) (string, error) {
	if a.IsCollection() || b.IsCollection() {
		return "", spaceerr.ErrNotSupportCollection
	}
	rel := &relate.Relationship{Arg: []matrix.Steric{a.ToMatrix(), b.ToMatrix()},
		IntersectBound: a.Bound().IntersectsBound(b.Bound())}
	im := rel.IntersectionMatrix()
	return im.ToString(), nil
}

// Distance returns distance Between the two Geometry.
func Distance(from, to Geometry, f measure.Distance) (float64, error) {
	if from == nil || from.IsEmpty() ||
		to == nil || to.IsEmpty() {
		return 0, nil
	}
	elem := &measure.ElementDistance{From: from.ToMatrix(), To: to.ToMatrix(), F: f}
	return elem.Distance()
}
