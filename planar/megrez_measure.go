package planar

import (
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space"
)

// Area returns the area of a polygonal geometry.
func (g *megrezAlgorithm) Area(geom space.Geometry) (float64, error) {
	switch geom.GeoJSONType() {
	case space.TypePolygon:
		return geom.(space.Polygon).Area()
	case space.TypeMultiPolygon:
		return geom.(space.MultiPolygon).Area()
	default:
		return 0.0, nil
	}
}

// Distance returns the minimum 2D Cartesian (planar) distance between two geometries, in projected units (spatial ref units).
func (g *megrezAlgorithm) Distance(geom1, geom2 space.Geometry) (float64, error) {
	return geom1.Distance(geom2)
}

// SphericalDistance calculates spherical distance
// To get real distance in m
func (g *megrezAlgorithm) SphericalDistance(geom1, geom2 space.Geometry) (float64, error) {
	return geom1.SpheroidDistance(geom2)
}

// HausdorffDistance returns the Hausdorff distance between two geometries, a measure of how similar
// or dissimilar 2 geometries are. Implements algorithm for computing a distance metric which can be
// thought of as the "Discrete Hausdorff Distance". This is the Hausdorff distance restricted
// to discrete points for one of the geometries
func (g *megrezAlgorithm) HausdorffDistance(geom1, geom2 space.Geometry) (float64, error) {
	return (&measure.HausdorffDistance{}).Distance(geom1.ToMatrix(), geom2.ToMatrix()), nil
}

// HausdorffDistanceDensify computes the Hausdorff distance with an additional densification fraction amount
func (g *megrezAlgorithm) HausdorffDistanceDensify(geom1, geom2 space.Geometry, densifyFrac float64) (float64, error) {
	return (&measure.HausdorffDistance{}).DistanceDensifyFrac(geom1.ToMatrix(), geom2.ToMatrix(), densifyFrac)
}

// Length returns the 2D Cartesian length of the geometry if it is a LineString, MultiLineString
func (g *megrezAlgorithm) Length(geom space.Geometry) (float64, error) {
	return geom.Length(), nil
}

// NGeometry returns the number of component geometries.
func (g *megrezAlgorithm) NGeometry(geom space.Geometry) (int, error) {
	return geom.Nums(), nil
}
