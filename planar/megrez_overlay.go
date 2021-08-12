package planar

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay"
	"github.com/spatial-go/geoos/algorithm/sharedpaths"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (g *MegrezAlgorithm) Difference(geom1, geom2 space.Geometry) (space.Geometry, error) {
	if (geom1.GeoJSONType()) != (geom2.GeoJSONType()) {
		return nil, algorithm.ErrNotMatchType
	}
	var err error
	if result, err := overlay.Difference(geom1.ToMatrix(), geom2.ToMatrix()); err == nil {
		return space.TransGeometry(result), nil
	}
	return nil, err
}

// Intersection returns a geometry that represents the point set intersection of the Geometries.
func (g *MegrezAlgorithm) Intersection(geom1, geom2 space.Geometry) (intersectGeom space.Geometry, intersectErr error) {
	switch geom1.GeoJSONType() {
	case space.TypePoint:
		over := &overlay.PointOverlay{Subject: geom1.ToMatrix(), Clipping: geom2.ToMatrix()}
		if result, err := over.Intersection(); err == nil {
			intersectGeom = space.TransGeometry(result)
		} else {
			intersectErr = err
		}
	case space.TypeLineString:
		over := &overlay.LineOverlay{PointOverlay: &overlay.PointOverlay{Subject: geom1.ToMatrix(), Clipping: geom2.ToMatrix()}}
		if result, err := over.Intersection(); err == nil {
			intersectGeom = space.TransGeometry(result)
		} else {
			intersectErr = err
		}
	case space.TypePolygon:
		over := &overlay.PolygonOverlay{PointOverlay: &overlay.PointOverlay{Subject: geom1.ToMatrix(), Clipping: geom2.ToMatrix()}}
		if result, err := over.Intersection(); err != nil {
			intersectGeom = space.TransGeometry(result)
		} else {
			intersectErr = err
		}
	default:
		intersectErr = algorithm.ErrNotMatchType
	}
	return
}

// LineMerge returns a (set of) LineString(s) formed by sewing together the constituent line work of a MULTILINESTRING.
func (g *MegrezAlgorithm) LineMerge(geom space.Geometry) (space.Geometry, error) {
	if geom.GeoJSONType() != space.TypeMultiLineString {
		return nil, spaceerr.ErrNotSupportGeometry
	}
	result := overlay.LineMerge(geom.ToMatrix().(matrix.Collection))
	var lm space.MultiLineString
	for _, v := range result {
		lm = append(lm, space.LineString(v.(matrix.LineMatrix)))
	}

	return lm, nil
}

// SharedPaths returns a collection containing paths shared by the two input geometries.
// Those going in the same direction are in the first element of the collection,
// those going in the opposite direction are in the second element.
// The paths themselves are given in the direction of the first geometry.
func (g *MegrezAlgorithm) SharedPaths(geom1, geom2 space.Geometry) (string, error) {
	forwDir, backDir, _ := sharedpaths.SharedPaths(geom1.ToMatrix(), geom2.ToMatrix())
	var forw, back space.Geometry
	if forwDir == nil {
		forw = space.MultiLineString{}
	} else {
		forw = space.TransGeometry(forwDir)
	}
	if backDir == nil {
		back = space.MultiLineString{}
	} else {
		back = space.TransGeometry(backDir)
	}
	coll := space.Collection{forw, back}

	return wkt.MarshalString(coll), nil
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (g *MegrezAlgorithm) SymDifference(geom1, geom2 space.Geometry) (space.Geometry, error) {
	if geom1.GeoJSONType() != geom2.GeoJSONType() {
		return nil, algorithm.ErrNotMatchType
	}
	var err error
	if result, err := overlay.SymDifference(geom1.ToMatrix(), geom2.ToMatrix()); err == nil {
		return space.TransGeometry(result), nil
	}
	return nil, err
}

// UnaryUnion does dissolve boundaries between components of a multipolygon (invalid) and does perform union
// between the components of a geometrycollection
func (g *MegrezAlgorithm) UnaryUnion(geom space.Geometry) (space.Geometry, error) {
	if geom.GeoJSONType() == space.TypeMultiPolygon {
		result := overlay.UnaryUnion(geom.ToMatrix())
		return space.Polygon(result.(matrix.PolygonMatrix)), nil
	}
	return nil, ErrNotPolygon
}

// Union returns a new geometry representing all points in this geometry and the other.
func (g *MegrezAlgorithm) Union(geom1, geom2 space.Geometry) (space.Geometry, error) {
	if geom1.GeoJSONType() == space.TypePolygon && geom2.GeoJSONType() == space.TypePolygon {
		result := overlay.Union(matrix.PolygonMatrix(geom1.(space.Polygon)), matrix.PolygonMatrix(geom2.(space.Polygon)))
		return space.Polygon(result.(matrix.PolygonMatrix)), nil
	} else if geom1.GeoJSONType() == space.TypePoint && geom2.GeoJSONType() == space.TypePoint {
		return space.MultiPoint{geom1.(space.Point), geom2.(space.Point)}, nil
	} else if geom1.GeoJSONType() == space.TypeLineString && geom2.GeoJSONType() == space.TypeLineString {
		result := overlay.UnionLine(geom1.ToMatrix().(matrix.LineMatrix), geom2.ToMatrix().(matrix.LineMatrix))
		return space.TransGeometry(result), nil
	}
	return space.Collection{geom1, geom2}, nil
}
