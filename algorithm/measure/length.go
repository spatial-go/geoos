package measure

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// OfLine Computes the length of a linestring specified by a sequence of points.
func OfLine(pts matrix.LineMatrix) float64 {
	// optimized for processing line
	if n := len(pts); n > 1 {
		length := 0.0

		for i := 0; i < len(pts)-1; i++ {
			x0, y0 := pts[i][0], pts[i][1]
			x1, y1 := pts[i+1][0], pts[i+1][1]
			dx, dy := x1-x0, y1-y0
			length += math.Sqrt(dx*dx + dy*dy)
			if i == len(pts)-2 {
				break
			}
		}
		return length
	}
	return 0.0
}
