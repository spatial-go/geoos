// Package dissovle Slice a geometric polygon.
package dissovle

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// UnaryUnion returns a Geometry containing the union.
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func DissovlePolygon(poly matrix.PolygonMatrix, diss matrix.Steric) (result matrix.Collection, err error) {
	return nil, algorithm.ErrUnknownType(diss)
}
