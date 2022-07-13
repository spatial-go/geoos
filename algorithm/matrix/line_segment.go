package matrix

import (
	"math"

	"github.com/spatial-go/geoos/algorithm"
)

// LineSegment is line.
type LineSegment struct {
	P0, P1 Matrix
}

// PointAlong Computes  point  that lies a given
// fraction along the line defined by this segment.
// A fraction of 0.0 returns the start point of the segment;
// a fraction of 1.0 returns the end point of the segment.
// If the fraction is &lt; 0.0 or &gt; 1.0 the point returned
// will lie before the start or beyond the end of the segment.
func (l *LineSegment) PointAlong(segmentLengthFraction float64) Matrix {
	coord := Matrix{0, 0}
	coord[0] = l.P0[0]
	coord[1] = l.P0[1]
	coord[0] = l.P0[0] + segmentLengthFraction*(l.P1[0]-l.P0[0])
	coord[1] = l.P0[1] + segmentLengthFraction*(l.P1[1]-l.P0[1])
	return coord
}

// PointAlongOffset Computes point that lies a given
// fraction along the line defined by this segment and offset from
// the segment by a given distance.
// A fraction of 0.0 offsets from the start point of the segment;
// a fraction of 1.0 offsets from the end point of the segment.
// The computed point is offset to the left of the line if the offset distance is
// positive, to the right if negative.
func (l *LineSegment) PointAlongOffset(segmentLengthFraction, offsetDistance float64) (Matrix, error) {
	// the point on the segment line
	segX := l.P0[0] + segmentLengthFraction*(l.P1[0]-l.P0[0])
	segY := l.P0[1] + segmentLengthFraction*(l.P1[1]-l.P0[1])

	dx := l.P1[0] - l.P0[0]
	dy := l.P1[1] - l.P0[1]
	lenXY := math.Sqrt(dx*dx + dy*dy)
	ux := 0.0
	uy := 0.0
	if offsetDistance != 0.0 {
		if lenXY <= 0.0 {
			return nil, algorithm.ErrComputeOffsetZero
		}

		// u is the vector that is the length of the offset, in the direction of the segment
		ux = offsetDistance * dx / lenXY
		uy = offsetDistance * dy / lenXY
	}

	// the offset point is the seg point plus the offset vector rotated 90 degrees CCW
	offsetX := segX - uy
	offsetY := segY + ux

	coord := Matrix{0, 0}
	coord[0] = offsetX
	coord[1] = offsetY
	return coord, nil
}

// LineArray returns the LineArray
func LineArray(l LineMatrix) (lines []*LineSegment) {
	for i := 0; i < len(l)-1; i++ {
		lines = append(lines, &LineSegment{Matrix(l[i]), Matrix(l[i+1])})
	}
	return
}
