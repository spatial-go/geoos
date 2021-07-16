package buffer

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Curve A dynamic list of the vertices in a constructed offset curve.
// Automatically removes adjacent vertices
// which are closer than a given tolerance.
type Curve struct {
	s0, s1, s2                         matrix.Matrix
	side                               int
	Line, seg0, seg1, offset0, offset1 matrix.LineMatrix
	leftLoc, rightLoc                  int
	parameters                         *CurveParameters
	distance                           float64
}

// DefaultCurve Creates a default curve.
func DefaultCurve() *Curve {
	return &Curve{parameters: DefaultCurveParameters()}
}

// SetParameters Set a  parameters of curve.
func (c *Curve) SetParameters(parameters *CurveParameters) {
	c.parameters = parameters
}

// CloseRing close ring.
func (c *Curve) CloseRing() {
	if len(c.Line) < 1 {
		return
	}
	startPt := c.Line[0]
	lastPt := c.Line[len(c.Line)-1]
	if matrix.Equal(matrix.Matrix(startPt), matrix.Matrix(lastPt)) {
		return
	}
	c.Line = append(c.Line, startPt)
}

// AddPt add point to OffsetSegmentString.
func (c *Curve) AddPt(p matrix.Matrix) {
	c.Line = append(c.Line, p)
}

// CreateCircle Creates a CW circle around a point
func (c *Curve) CreateCircle(p matrix.Matrix, distance float64) matrix.LineMatrix {
	// add start point
	c.Line = append(c.Line, matrix.Matrix{p[0] + distance, p[1]})

	directionFactor := CLOCKWISE
	totalAngle := math.Abs(0.0 - ANGLE*math.Pi)
	nSegs := totalAngle/(math.Pi/2.0/float64(c.parameters.QuadrantSegments)) + 0.5

	if nSegs < 1 {
		return c.Line // no segments because angle is less than increment - nothing to do!
	}

	// choose angle increment so that each segment has equal length
	angleInc := totalAngle / nSegs

	pt := matrix.Matrix{0, 0}
	for i := 0; float64(i) < nSegs; i++ {
		angle := float64(directionFactor) * float64(i) * angleInc
		pt[0] = p[0] + distance*math.Cos(angle)
		pt[1] = p[1] + distance*math.Sin(angle)
		c.Line = append(c.Line, pt)
	}

	c.CloseRing()
	return c.Line
}

// CreateSquare Creates a CW square around a point
func (c *Curve) CreateSquare(p matrix.Matrix, distance float64) matrix.LineMatrix {
	c.Line = append(c.Line, matrix.Matrix{p[0] + distance, p[1] + distance})
	c.Line = append(c.Line, matrix.Matrix{p[0] + distance, p[1] - distance})
	c.Line = append(c.Line, matrix.Matrix{p[0] - distance, p[1] - distance})
	c.Line = append(c.Line, matrix.Matrix{p[0] - distance, p[1] + distance})

	c.CloseRing()
	return c.Line
}

// AddLineEndCap Add an end cap around point p1, terminating a line segment coming from p0
func (c *Curve) AddLineEndCap(p0, p1 matrix.Matrix, distance float64) {
	seg := matrix.LineMatrix{p0, p1}

	offsetL := make(matrix.LineMatrix, 2)
	c.computeOffsetSegment(seg, LEFT, distance, offsetL)
	offsetR := make(matrix.LineMatrix, 2)
	c.computeOffsetSegment(seg, RIGHT, distance, offsetR)

	dx := p1[0] - p0[0]
	dy := p1[1] - p0[1]
	angle := math.Atan2(dy, dx)

	switch c.parameters.EndCapStyle {
	case CAPROUND:
		// add offset seg points with a fillet between them
		c.Line = append(c.Line, offsetL[1])
		c.addDirectedFillet(p1, distance, angle+math.Pi/2, angle-math.Pi/2, CLOCKWISE)
		c.Line = append(c.Line, offsetR[1])
	case CAPFLAT:
		// only offset segment points are added
		c.Line = append(c.Line, offsetL[1])
		c.Line = append(c.Line, offsetR[1])
	case CAPSQUARE:
		// add a square defined by extensions of the offset segment endpoints
		squareCapSideOffset := matrix.Matrix{}
		squareCapSideOffset[0] = math.Abs(distance) * math.Cos(angle)
		squareCapSideOffset[1] = math.Abs(distance) * math.Sin(angle)

		squareCapLOffset := matrix.Matrix{
			offsetL[1][0] + squareCapSideOffset[0],
			offsetL[1][1] + squareCapSideOffset[1]}
		squareCapROffset := matrix.Matrix{
			offsetR[1][0] + squareCapSideOffset[0],
			offsetR[1][1] + squareCapSideOffset[1]}
		c.Line = append(c.Line, squareCapLOffset)
		c.Line = append(c.Line, squareCapROffset)
	}
}

// addDirectedFillet Adds points for a circular fillet arc
// between two specified angles.
// The start and end point for the fillet are not added -
// the caller must add them if required.
func (c *Curve) addDirectedFillet(p matrix.Matrix, radius, startAngle, endAngle float64, direction int) {
	directionFactor := 1.0
	if direction == CLOCKWISE {
		directionFactor = -1.0
	}

	totalAngle := math.Abs(startAngle - endAngle)
	filletAngleQuantum := math.Pi / 2.0 / float64(c.parameters.QuadrantSegments)
	nSegs := (totalAngle/filletAngleQuantum + 0.5)

	if nSegs < 1 {
		return // no segments because angle is less than increment - nothing to do!
	}
	// choose angle increment so that each segment has equal length
	angleInc := totalAngle / nSegs
	pt := matrix.Matrix{0, 0}
	for i := 0; float64(i) < nSegs; i++ {
		angle := startAngle + directionFactor*float64(i)*angleInc
		pt[0] = p[0] + radius*math.Cos(angle)
		pt[1] = p[1] + radius*math.Sin(angle)
		c.Line = append(c.Line, pt)
	}
}

func (c *Curve) initSideSegments(s1, s2 matrix.Matrix, side int) {
	c.s1 = s1
	c.s2 = s2
	c.side = side
	c.seg1 = matrix.LineMatrix{s1, s2}
	c.computeOffsetSegment(c.seg1, side, c.distance, c.offset1)
}

// computeOffsetSegment Compute an offset segment for an input segment on a given side and at a given distance.
// The offset points are computed in full double precision, for accuracy.
func (c *Curve) computeOffsetSegment(seg matrix.LineMatrix, side int, distance float64, offset matrix.LineMatrix) {
	sideSign := -1.0
	if side == LEFT {
		sideSign = 1.0
	}
	dx := seg[1][0] - seg[0][0]
	dy := seg[1][1] - seg[0][1]
	length := math.Sqrt(dx*dx + dy*dy)
	// u is the vector that is the length of the offset, in the direction of the segment
	ux := sideSign * distance * dx / length
	uy := sideSign * distance * dy / length
	offset[0][0] = seg[0][0] - uy
	offset[0][1] = seg[0][1] + ux
	offset[1][0] = seg[1][0] - uy
	offset[1][1] = seg[1][1] + ux
}
