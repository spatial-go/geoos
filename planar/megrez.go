package planar

import (
	"github.com/spatial-go/geoos/space"
)

// megrezAlgorithm algorithm implement
type megrezAlgorithm struct{}

// Equals returns TRUE if the given Geometries are "spatially equal".
func (g *megrezAlgorithm) Equals(geom1, geom2 space.Geometry) (bool, error) {
	return geom1.Equals(geom2), nil
}

// EqualsExact returns true if both geometries are Equal, as evaluated by their
// points being within the given tolerance.
func (g *megrezAlgorithm) EqualsExact(geom1, geom2 space.Geometry, tolerance float64) (bool, error) {
	return geom1.EqualsExact(geom2, tolerance), nil
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (g *megrezAlgorithm) IsClosed(geom space.Geometry) (bool, error) {
	elem := space.GeometryValid{Geometry: geom}
	return elem.IsClosed(), nil
}

// IsEmpty returns true if this space.Geometry is an empty geometry.
// If true, then this space.Geometry represents an empty geometry collection, polygon, point etc.
func (g *megrezAlgorithm) IsEmpty(geom space.Geometry) (bool, error) {
	return geom.IsEmpty(), nil
}

// IsRing returns true if the lineal geometry has the ring property.
func (g *megrezAlgorithm) IsRing(geom space.Geometry) (bool, error) {
	elem := space.GeometryValid{Geometry: geom}
	return elem.IsClosed() && elem.IsSimple(), nil
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points, such as self intersection or self tangency.
func (g *megrezAlgorithm) IsSimple(geom space.Geometry) (bool, error) {
	return geom.IsSimple(), nil
}
