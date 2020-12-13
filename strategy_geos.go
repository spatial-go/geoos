package geos

import (
	"github.com/spatial-go/geos/geo"
)

type GEOSAlgorithm struct{}

func (G GEOSAlgorithm) Area(g Geometry) (float64, error) {
	s := MarshalString(g)
	return geo.Area(s)
}

func (G GEOSAlgorithm) Boundary(g Geometry) (Geometry, error) {
	s := MarshalString(g)
	wkt, e := geo.Boundary(s)
	if e != nil {
		return nil, e
	}
	geometry, e := UnmarshalString(wkt)
	if e != nil {
		return nil, e
	}
	return geometry, nil
}
func (G GEOSAlgorithm) Centroid(g Geometry) (Geometry, error) {
	s := MarshalString(g)
	centroid, e := geo.Centroid(s)
	if e != nil {
		return nil, e
	}
	geometry, e := UnmarshalString(centroid)
	if e != nil {
		return nil, e
	}
	return geometry, nil
}

func (G GEOSAlgorithm) IsSimple(g Geometry) (bool, error) {
	s := MarshalString(g)
	return geo.IsSimple(s)
}

func (G GEOSAlgorithm) Envelope() (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) ConvexHull() (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) UnaryUnion() (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) PointOnSurface() (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) LineMerge() (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Simplify(tolerance float64) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) SimplifyP(tolerance float64) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) UniquePoints() (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) SharedPaths(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Snap(other *Geometry, tolerance float64) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Intersection(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Difference(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) SymDifference(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Union(other *Geometry) (*Geometry, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Disjoint(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Touches(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Intersects(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Crosses(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Within(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Contains(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Overlaps(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Equals(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Covers(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) CoveredBy(other *Geometry) (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) IsEmpty() (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) IsRing() (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) HasZ() (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) IsClosed() (bool, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) SRID() (int, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) SetSRID(srid int) {
	panic("implement me")
}

func (G GEOSAlgorithm) NGeometry() (int, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Buffer(g Geometry, width float64, quadsegs int32) Geometry {
	panic("implement me")
}

func (G GEOSAlgorithm) EqualsExact(s Geometry, d Geometry, tolerance float64) bool {
	panic("implement me")
}

func (G GEOSAlgorithm) Length() (float64, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Distance(s Geometry, d Geometry) (float64, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) HausdorffDistance(s Geometry, d Geometry) (float64, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) HausdorffDistanceDensify(s Geometry, d Geometry, densifyFrac float64) (float64, error) {
	panic("implement me")
}

func (G GEOSAlgorithm) Relate(s Geometry, d Geometry, ) {
	panic("implement me")
}
