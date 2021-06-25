package geoos

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Polygon is a closed area. The first LineString is the outer ring.
// The others are the holes. Each LineString is expected to be closed
// ie. the first point matches the last.
type Polygon []Ring

// GeoJSONType returns the GeoJSON type for the polygon.
func (p Polygon) GeoJSONType() string {
	return TypePolygon
}

// Dimensions returns 2 because a Polygon is a 2d object.
func (p Polygon) Dimensions() int {
	return 2
}

// Nums num of polygons
func (p Polygon) Nums() int {
	return 1
}

// // Area returns the area of this polygonal geometry
// func (p Polygon) Area() (float64, error) {
// 	s := NormalStrategy()
// 	return s.Area(p)
// }

// // Boundary returns the closure of the combinatorial boundary of this Geometry
// func (p Polygon) Boundary() (Geometry, error) {
// 	s := NormalStrategy()
// 	return s.Boundary(p)
// }

// Bound returns a bound around the polygon.
func (p Polygon) Bound() Bound {
	if len(p) == 0 {
		return emptyBound
	}
	return p[0].Bound()
}

// Equal compares two polygons. Returns true if lengths are the same
// and all points are Equal.
func (p Polygon) Equal(polygon Polygon) bool {
	if len(p) != len(polygon) {
		return false
	}
	for i := range p {
		if !p[i].Equal(polygon[i]) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry.
func (p Polygon) Area() (float64, error) {
	return algorithm.AreaOfPolygon(p.ToMatrix()), nil
}

func (p Polygon) ToMatrix() matrix.MatrixPolygon {
	var matrix matrix.MatrixPolygon
	for _, line := range p {
		var matrix1 [][]float64
		for _, point := range line {
			matrix1 = append(matrix1, []float64{point.X(), point.Y()})
		}
		matrix = append(matrix, matrix1)
	}
	return matrix
}
