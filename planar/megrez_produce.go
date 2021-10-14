package planar

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/overlay/snap"
	"github.com/spatial-go/geoos/coordtransform"
	"github.com/spatial-go/geoos/space"
)

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (g *megrezAlgorithm) Boundary(geom space.Geometry) (space.Geometry, error) {
	return geom.Boundary()
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (g *megrezAlgorithm) Buffer(geom space.Geometry, width float64, quadsegs int) (geometry space.Geometry) {
	buff := buffer.Buffer(geom.ToMatrix(), width, quadsegs)
	switch b := buff.(type) {
	case matrix.LineMatrix:
		return space.LineString(b)
	case matrix.PolygonMatrix:
		return space.Polygon(b)
	}
	return nil
}

// BufferInMeter Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (g *MegrezAlgorithm) BufferInMeter(geom space.Geometry, width float64, quadsegs int) (geometry space.Geometry) {
	if geom == nil {
		return
	}
	centroid := geom.Centroid()
	width = measure.MercatorDistance(width, centroid.Lat())
	transformer := coordtransform.NewTransformer(coordtransform.LLTOMERCATOR)
	geomMatrix, _ := transformer.TransformGeometry(geom.ToMatrix())
	geom = space.TransGeometry(geomMatrix)
	geometry = g.Buffer(geom, width, quadsegs)
	if geometry != nil {
		transformer.CoordType = coordtransform.MERCATORTOLL
		geomMatrix, _ := transformer.TransformGeometry(geometry.ToMatrix())
		geometry = space.TransGeometry(geomMatrix)
	}
	return
}

// Centroid  computes the geometric center of a geometry, or equivalently, the center of mass of the geometry as a POINT.
// For [MULTI]POINTs, this is computed as the arithmetic mean of the input coordinates.
// For [MULTI]LINESTRINGs, this is computed as the weighted length of each line segment.
// For [MULTI]POLYGONs, "weight" is thought in terms of area.
// If an empty geometry is supplied, an empty GEOMETRYCOLLECTION is returned.
// If NULL is supplied, NULL is returned.
// If CIRCULARSTRING or COMPOUNDCURVE are supplied, they are converted to linestring with CurveToLine first,
// then same than for LINESTRING
func (g *megrezAlgorithm) Centroid(geom space.Geometry) (space.Geometry, error) {
	if geom == nil || geom.IsEmpty() {
		return nil, nil
	}
	return space.Centroid(geom), nil
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (g *megrezAlgorithm) ConvexHull(geom space.Geometry) (space.Geometry, error) {
	result := buffer.ConvexHullWithGeom(geom.ToMatrix()).ConvexHull()
	return space.TransGeometry(result), nil
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (g *megrezAlgorithm) Envelope(geom space.Geometry) (space.Geometry, error) {
	switch geom.GeoJSONType() {
	case space.TypePoint:
		return geom, nil
	default:
		return geom.Bound().ToPolygon(), nil
	}
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (g *megrezAlgorithm) PointOnSurface(geom space.Geometry) (space.Geometry, error) {
	m := buffer.InteriorPoint(geom.ToMatrix())
	return space.Point(m), nil
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (g *megrezAlgorithm) Simplify(geom space.Geometry, tolerance float64) (space.Geometry, error) {
	result := simplify.Simplify(geom.ToMatrix(), tolerance)
	return space.TransGeometry(result), nil
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (g *megrezAlgorithm) SimplifyP(geom space.Geometry, tolerance float64) (space.Geometry, error) {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(geom.ToMatrix(), tolerance)
	return space.TransGeometry(result), nil
}

// Snap the vertices and segments of a geometry to another space.Geometry's vertices.
// A snap distance tolerance is used to control where snapping is performed.
// The result geometry is the input geometry with the vertices snapped.
// If no snapping occurs then the input geometry is returned unchanged.
func (g *megrezAlgorithm) Snap(input, reference space.Geometry, tolerance float64) (space.Geometry, error) {
	result := snap.Snap(input.ToMatrix(), reference.ToMatrix(), tolerance)
	return space.TransGeometry(result[0]), nil
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (g *megrezAlgorithm) UniquePoints(geom space.Geometry) (space.Geometry, error) {
	return geom.UniquePoints(), nil
}
