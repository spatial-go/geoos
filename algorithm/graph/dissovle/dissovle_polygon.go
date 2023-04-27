// Package dissovle Slice a geometric polygon.
package dissovle

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/graph/clipping"
	"github.com/spatial-go/geoos/algorithm/graph/de9im"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// PolygonDissovle returns a Geometry containing the dissovle.
//
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func PolygonDissovle(poly matrix.PolygonMatrix, diss matrix.Steric) (result matrix.Steric, err error) {
	switch d := diss.(type) {
	case matrix.LineMatrix:
		clip := graph.ClipHandle(d, poly)
		switch im := de9im.IMByClip(clip); {
		case !im.IsIntersects():
			return matrix.Collection{poly}, algorithm.ErrUnknownType(diss)
		}
		gu, _ := clip.Union()
		interLines, _ := clipping.Intersection(d, poly)
		result := matrix.Collection{}
		switch line := interLines.(type) {
		case matrix.LineMatrix:
			if de9im.IM(line, matrix.LineMatrix(poly[0])).IsCoveredBy() {
				return matrix.Collection{poly}, algorithm.ErrUnknownType(diss)
			}
			if lines, err := dissovleLink(gu, line); err == nil {

				for _, v := range lines {
					poly := matrix.PolygonMatrix{v}
					result = append(result, poly)
				}
				return result, nil
			}
		case matrix.Collection:
			for _, interLine := range line {
				if inter, ok := interLine.(matrix.LineMatrix); ok {
					if de9im.IM(inter, matrix.LineMatrix(poly[0])).IsCoveredBy() {
						return matrix.Collection{poly}, algorithm.ErrUnknownType(diss)
					}
					if lines, err := dissovleLink(gu, inter); err == nil {

						for _, v := range lines {
							poly := matrix.PolygonMatrix{v}
							result = append(result, poly)
						}
						return result, nil
					}
				}
			}

		default:
		}
	case matrix.PolygonMatrix:
		// TODO:
	}
	return matrix.Collection{poly}, algorithm.ErrUnknownType(diss)
}
