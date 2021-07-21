package buffer

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// CurveBuilder Computes the raw offset curve for a
// single Geometry component (ring, line or point).
type CurveBuilder struct {
	*Curve
	Curves   []Curve
	distance float64
}

// LineCurve  This method handles single points as well as LineStrings.
// LineStrings are assumed <b>not</b> to be closed (the function will not
// fail for closed lines, but will generate superfluous line caps).
func (c *CurveBuilder) LineCurve(pts matrix.LineMatrix, distance float64,
	leftLoc, rightLoc int) matrix.LineMatrix {
	c.distance = distance

	if len(pts) <= 1 {
		c.computePointCurve(pts[0])
	} else {
		if c.parameters.IsSingleSided {
			isRightSide := distance < 0.0
			c.computeSingleSidedBufferCurve(pts, isRightSide)
		} else {
			c.computeLineBufferCurve(pts)
		}
	}

	lineCoord := c.Curve.Line

	c.AddCurve(lineCoord, leftLoc, rightLoc)
	return lineCoord
}

// RingCurve This method handles the degenerate cases of single points and lines,
// as well as valid rings.
func (c *CurveBuilder) RingCurve(pts matrix.LineMatrix, distance float64, side int, leftLoc, rightLoc int) matrix.LineMatrix {

	c.distance = distance

	if len(pts) <= 2 {
		return c.LineCurve(pts, distance, leftLoc, rightLoc)
	}
	if distance == 0.0 {
		copy := make(matrix.LineMatrix, len(pts))
		for i := 0; i < len(copy); i++ {
			copy[i] = matrix.Matrix(pts[i])
		}
		return copy
	}

	c.computeRingBufferCurve(pts, side)
	lineCoord := c.Curve.Line

	c.AddCurve(lineCoord, leftLoc, rightLoc)
	return lineCoord
}

// AddCurve Creates a SegmentString for a coordinate list which is a raw offset curve,
// and adds it to the list of buffer curves.
// The SegmentString is tagged with a Label giving the topology of the curve.
// The curve may be oriented in either direction.
// If the curve is oriented CW, the locations will be:
// <br>Left: Location.EXTERIOR
// <br>Right: Location.INTERIOR
func (c *CurveBuilder) AddCurve(pts matrix.LineMatrix, leftLoc, rightLoc int) {
	// don't add null or trivial curves
	if pts == nil || len(pts) < 2 {
		return
	}
	// add the edge for a coordinate list which is a raw offset curve
	curve := &Curve{Line: pts, leftLoc: leftLoc, rightLoc: rightLoc, parameters: c.parameters}
	c.Curves = append(c.Curves, *curve)
}

// isRingCurveInverted Tests whether the offset curve for a ring is fully inverted.
// An inverted ("inside-out") curve occurs in some specific situations
// involving a buffer distance which should result in a fully-eroded (empty) buffer.
// It can happen that the sides of a small, convex polygon
// produce offset segments which all cross one another to form
// a curve with inverted orientation.
// This happens at buffer distances slightly greater than the distance at
// which the buffer should disappear.
// The inverted curve will produce an incorrect non-empty buffer (for a shell)
// or an incorrect hole (for a hole).
// It must be discarded from the set of offset curves used in the buffer.
// Heuristics are used to reduce the number of cases which area checked,
// for efficiency and correctness.
func (c *CurveBuilder) isRingCurveInverted(pts matrix.LineMatrix, distance float64) bool {
	if distance == 0.0 {
		return false
	}
	/**
	 * Only proper rings can invert.
	 */
	if len(pts) <= 3 {
		return false
	}
	/**
	 * Heuristic based on low chance that a ring with many vertices will invert.
	 * This low limit ensures this test is fairly efficient.
	 */
	if len(pts) >= calc.MaxRingSize {
		return false
	}

	/**
	 * An inverted curve has no more points than the input ring.
	 * This also eliminates concave inputs (which will produce fillet arcs)
	 */
	if len(c.Curves) > len(pts) {
		return false
	}

	/**
	 * Check if the curve vertices are all closer to the input ring
	 * than the buffer distance.
	 * If so, the curve is NOT a valid buffer curve.
	 */
	distTol := calc.NearnessFactor * math.Abs(distance)

	maxDist := 0.0
	for _, v := range c.Curve.Line {
		dist := measure.DistanceLineToPoint(pts, v, measure.PlanarDistance)
		if dist > maxDist {
			maxDist = dist
		}
	}

	isCurveTooClose := maxDist < distTol
	return isCurveTooClose
}

func (c *CurveBuilder) computePointCurve(pt matrix.Matrix) {
	switch c.parameters.EndCapStyle {
	case calc.CAPROUND:
		c.Curve.CreateCircle(pt, c.distance)
	case calc.CAPSQUARE:
		c.Curve.CreateSquare(pt, c.distance)
		// otherwise curve is empty (e.g. for a butt cap);
	}
}

func (c *CurveBuilder) computeLineBufferCurve(pts matrix.LineMatrix) {
	distTol := c.distance * c.parameters.SimplifyFactor

	//--------- compute points for left side of line
	// Simplify the appropriate side of the line before generating
	simp := &LineSimplifier{inputLine: pts}
	simp1 := simp.Simplify(distTol)
	// MD - used for testing only (to eliminate simplification)
	//    Coordinate[] simp1 = inputPts;

	n1 := len(simp1) - 1
	c.Curve.initSideSegments(simp1[0], simp1[1], calc.LEFT)
	for i := 2; i <= n1; i++ {
		c.Curve.Line = append(c.Curve.Line, simp1[i])
	}
	c.Curve.Line = append(c.Curve.Line, c.Curve.offset1[1])
	// add line cap for end of line
	c.AddLineEndCap(simp1[n1-1], simp1[n1], c.distance)

	//---------- compute points for right side of line
	// Simplify the appropriate side of the line before generating
	simp2 := simp.Simplify(-distTol)
	// MD - used for testing only (to eliminate simplification)
	//    Coordinate[] simp2 = inputPts;
	n2 := len(simp2) - 1

	// since we are traversing line in opposite order, offset position is still LEFT
	c.Curve.initSideSegments(simp2[n2], simp2[n2-1], calc.LEFT)
	for i := n2 - 2; i >= 0; i-- {
		c.Curve.Line = append(c.Curve.Line, simp2[i])
	}
	c.Curve.Line = append(c.Curve.Line, c.Curve.offset1[1])
	// add line cap for start of line
	c.Curve.AddLineEndCap(simp2[1], simp2[0], c.distance)

	c.Curve.CloseRing()
}

func (c *CurveBuilder) computeSingleSidedBufferCurve(pts matrix.LineMatrix, isRightSide bool) {
	distTol := c.distance * c.parameters.SimplifyFactor

	if isRightSide {
		// add original line
		c.Curve.Line = append(c.Curve.Line, pts...)

		//---------- compute points for right side of line
		// Simplify the appropriate side of the line before generating
		simp := &LineSimplifier{inputLine: pts}
		simp2 := simp.Simplify(-distTol)
		// MD - used for testing only (to eliminate simplification)
		//    Coordinate[] simp2 = inputPts;
		n2 := len(simp2) - 1

		// since we are traversing line in opposite order, offset position is still LEFT
		c.Curve.initSideSegments(simp2[n2], simp2[n2-1], calc.LEFT)
		c.Curve.Line = append(c.Curve.Line, c.Curve.offset1[1])
		for i := n2 - 2; i >= 0; i-- {
			c.Curve.Line = append(c.Curve.Line, c.Curve.Line[i])
		}
	} else {
		// add original line
		c.Curve.Line = append(c.Curve.Line, pts...)

		//--------- compute points for left side of line
		// Simplify the appropriate side of the line before generating
		simp := &LineSimplifier{inputLine: pts}
		simp1 := simp.Simplify(distTol)
		// MD - used for testing only (to eliminate simplification)
		//      Coordinate[] simp1 = inputPts;

		n1 := len(simp1) - 1
		c.Curve.initSideSegments(simp1[0], simp1[1], calc.LEFT)
		c.Curve.Line = append(c.Curve.Line, c.Curve.offset1[1])
		for i := 2; i <= n1; i++ {
			c.Curve.Line = append(c.Curve.Line, c.Curve.Line[i])
		}
	}
	c.Curve.Line = append(c.Curve.Line, c.Curve.offset1[1])
	c.Curve.CloseRing()
}

func (c *CurveBuilder) computeRingBufferCurve(pts matrix.LineMatrix, side int) {
	// simplify input line to improve performance
	distTol := c.distance * c.parameters.SimplifyFactor
	// ensure that correct side is simplified
	if side == calc.RIGHT {
		distTol = -distTol
	}
	simp := &LineSimplifier{inputLine: pts}
	simp1 := simp.Simplify(distTol)

	n := len(simp1) - 1
	c.Curve.initSideSegments(simp1[n-1], simp1[0], side)
	for i := 1; i <= n; i++ {
		c.Curve.Line = append(c.Curve.Line, simp1[i])
	}
	c.Curve.CloseRing()
}
