package buffer

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/calc/angle"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// Curve A dynamic list of the vertices in a constructed offset curve.
// Automatically removes adjacent vertices
// which are closer than a given tolerance.
type Curve struct {
	s0, s1, s2                                    matrix.Matrix
	side                                          int
	Line                                          matrix.LineMatrix
	seg0, seg1, offset0, offset1                  *matrix.LineSegment
	leftLoc, rightLoc                             int
	parameters                                    *CurveParameters
	distance                                      float64
	hasNarrowConcaveAngle                         bool
	closingSegLengthFactor, minimimVertexDistance float64
}

// CurveWithDistance Creates a default curve.
func CurveWithDistance(distance float64) *Curve {
	return CurveWithParameters(DefaultCurveParameters(), distance)
}

// CurveWithParameters Creates a default curve.
func CurveWithParameters(parameters *CurveParameters, distance float64) *Curve {
	cur := &Curve{parameters: parameters, closingSegLengthFactor: 1.0}
	cur.minimimVertexDistance = distance * calc.CurveVertexSnapDistanceFactor
	cur.distance = distance
	return cur
}

// SetParameters Set a  parameters of curve.
func (c *Curve) SetParameters(parameters *CurveParameters) {
	c.parameters = parameters
}

// AddLine add a line.
func (c *Curve) AddLine(line matrix.LineMatrix) {
	for _, pt := range line {
		if len(c.Line) > 0 && measure.PlanarDistance(matrix.Matrix(c.Line[len(c.Line)-1]), matrix.Matrix(pt)) <= c.minimimVertexDistance {
			continue
		}
		c.AddPt(pt)
	}
}

// Add add a Pts.
func (c *Curve) Add(pts ...matrix.Matrix) {
	for _, pt := range pts {
		if len(c.Line) > 0 && measure.PlanarDistance(matrix.Matrix(c.Line[len(c.Line)-1]), pt) <= c.minimimVertexDistance {
			continue
		}
		c.Line = append(c.Line, matrix.Matrix{pt[0], pt[1]})
	}
}

// CloseRing close ring.
func (c *Curve) CloseRing() {
	if len(c.Line) < 1 {
		return
	}
	startPt := c.Line[0]
	lastPt := c.Line[len(c.Line)-1]
	if matrix.Matrix(startPt).Equals(matrix.Matrix(lastPt)) {
		return
	}
	c.Add(startPt)
}

// AddPt add point to OffsetSegmentString.
func (c *Curve) AddPt(p matrix.Matrix) {
	c.Add(p)
}

// CreateCircle Creates a CW circle around a point
func (c *Curve) CreateCircle(p matrix.Matrix, distance float64) matrix.LineMatrix {
	// add start point
	c.Add(matrix.Matrix{p[0] + distance, p[1]})

	startAngle, endAngle := 0.0, angle.PiTimes2
	directionFactor := calc.ClockWise
	totalAngle := math.Abs(startAngle - endAngle)
	nSegs := int(totalAngle/(math.Pi/2.0/float64(c.parameters.QuadrantSegments)) + 0.5)

	if nSegs < 1 {
		return c.Line // no segments because angle is less than increment - nothing to do!
	}

	// choose angle increment so that each segment has equal length
	angleInc := totalAngle / float64(nSegs)

	pt := matrix.Matrix{0, 0}
	for i := 1; i < nSegs; i++ {
		angle := startAngle + float64(directionFactor)*float64(i)*angleInc
		pt[0] = p[0] + distance*math.Cos(angle)
		pt[1] = p[1] + distance*math.Sin(angle)
		c.Add(pt)
		pt = matrix.Matrix{0, 0}
	}

	c.CloseRing()
	return c.Line
}

// CreateSquare Creates a CW square around a point
func (c *Curve) CreateSquare(p matrix.Matrix, distance float64) matrix.LineMatrix {
	c.Add(matrix.Matrix{p[0] + distance, p[1] + distance})
	c.Add(matrix.Matrix{p[0] + distance, p[1] - distance})
	c.Add(matrix.Matrix{p[0] - distance, p[1] - distance})
	c.Add(matrix.Matrix{p[0] - distance, p[1] + distance})

	c.CloseRing()
	return c.Line
}

// AddLineEndCap Add an end cap around point p1, terminating a line segment coming from p0
func (c *Curve) AddLineEndCap(p0, p1 matrix.Matrix, distance float64) {
	seg := &matrix.LineSegment{P0: p0, P1: p1}

	offsetL := &matrix.LineSegment{P0: matrix.Matrix{0, 0}, P1: matrix.Matrix{0, 0}}
	offsetR := &matrix.LineSegment{P0: matrix.Matrix{0, 0}, P1: matrix.Matrix{0, 0}}
	c.computeOffsetSegment(seg, calc.SideLeft, distance, offsetL)
	c.computeOffsetSegment(seg, calc.SideRight, distance, offsetR)

	dx := p1[0] - p0[0]
	dy := p1[1] - p0[1]
	angle := math.Atan2(dy, dx)

	switch c.parameters.EndCapStyle {
	case calc.CapRound:
		// add offset seg points with a fillet between them
		c.Add(offsetL.P1)
		c.addDirectedFillet(p1, distance, angle+math.Pi/2, angle-math.Pi/2, calc.ClockWise)
		c.Add(offsetR.P1)
	case calc.CapFlat:
		// only offset segment points are added
		c.Add(offsetL.P1)
		c.Add(offsetR.P1)
	case calc.CapSquare:
		// add a square defined by extensions of the offset segment endpoints
		squareCapSideOffset := matrix.Matrix{}
		squareCapSideOffset[0] = math.Abs(distance) * math.Cos(angle)
		squareCapSideOffset[1] = math.Abs(distance) * math.Sin(angle)

		squareCapLOffset := matrix.Matrix{
			offsetL.P1[0] + squareCapSideOffset[0],
			offsetL.P1[1] + squareCapSideOffset[1]}
		squareCapROffset := matrix.Matrix{
			offsetR.P1[0] + squareCapSideOffset[0],
			offsetR.P1[1] + squareCapSideOffset[1]}
		c.Add(squareCapLOffset)
		c.Add(squareCapROffset)
	}
}

// addDirectedFillet Adds points for a circular fillet arc
// between two specified angles.
// The start and end point for the fillet are not added -
// the caller must add them if required.
func (c *Curve) addDirectedFillet(p matrix.Matrix, radius, startAngle, endAngle float64, direction int) {
	directionFactor := 1.0
	if direction == calc.ClockWise {
		directionFactor = -1.0
	}

	totalAngle := math.Abs(startAngle - endAngle)
	filletAngleQuantum := math.Pi / 2.0 / float64(c.parameters.QuadrantSegments)
	nSegs := int(totalAngle/filletAngleQuantum + 0.5)

	if nSegs < 1 {
		return // no segments because angle is less than increment - nothing to do!
	}
	// choose angle increment so that each segment has equal length
	angleInc := totalAngle / float64(nSegs)
	pt := matrix.Matrix{0, 0}
	for i := 0; i < nSegs; i++ {
		angle := startAngle + directionFactor*float64(i)*angleInc
		pt[0] = p[0] + radius*math.Cos(angle)
		pt[1] = p[1] + radius*math.Sin(angle)
		c.Add(pt)
		pt = matrix.Matrix{0, 0}
	}
}

func (c *Curve) initSideSegments(s1, s2 matrix.Matrix, side int) {
	c.s1 = s1
	c.s2 = s2
	c.side = side
	c.seg1 = &matrix.LineSegment{P0: s1, P1: s2}
	c.offset0 = &matrix.LineSegment{P0: matrix.Matrix{0, 0}, P1: matrix.Matrix{0, 0}}
	c.offset1 = &matrix.LineSegment{P0: matrix.Matrix{0, 0}, P1: matrix.Matrix{0, 0}}
	c.computeOffsetSegment(c.seg1, side, c.distance, c.offset1)
}

// computeOffsetSegment Compute an offset segment for an input segment on a given side and at a given distance.
// The offset points are computed in full double precision, for accuracy.
func (c *Curve) computeOffsetSegment(seg *matrix.LineSegment, side int, distance float64, offset *matrix.LineSegment) {
	sideSign := -1.0
	if side == calc.SideLeft {
		sideSign = 1.0
	}
	dx := seg.P1[0] - seg.P0[0]
	dy := seg.P1[1] - seg.P0[1]
	length := math.Sqrt(dx*dx + dy*dy)
	// u is the vector that is the length of the offset, in the direction of the segment
	ux := sideSign * distance * dx / length
	uy := sideSign * distance * dy / length
	offset.P0[0] = seg.P0[0] - uy
	offset.P0[1] = seg.P0[1] + ux
	offset.P1[0] = seg.P1[0] - uy
	offset.P1[1] = seg.P1[1] + ux
}

func (c *Curve) addNextSegment(p matrix.Matrix, addStartPoint bool) {
	// s0-s1-s2 are the coordinates of the previous segment and the current one
	c.s0 = c.s1
	c.s1 = c.s2
	c.s2 = p
	c.seg0 = &matrix.LineSegment{P0: c.s0, P1: c.s1}
	c.computeOffsetSegment(c.seg0, c.side, c.distance, c.offset0)
	c.seg1 = &matrix.LineSegment{P0: c.s1, P1: c.s2}
	c.computeOffsetSegment(c.seg1, c.side, c.distance, c.offset1)

	// do nothing if points are equal
	if c.s1.Equals(c.s2) {
		return
	}
	orientation := OrientationIndex(c.s0, c.s1, c.s2)
	outsideTurn :=
		(orientation == calc.ClockWise && c.side == calc.SideLeft) || (orientation == calc.CounterClockWise && c.side == calc.SideRight)

	if orientation == 0 { // lines are collinear
		c.addCollinear(addStartPoint)
	} else if outsideTurn {
		c.addOutsideTurn(orientation, addStartPoint)
	} else { // inside turn
		c.addInsideTurn(orientation, addStartPoint)
	}
}

func (c *Curve) addCollinear(addStartPoint bool) {

	mark, ips := relate.Intersection(c.s0, c.s1, c.s1, c.s2)

	// if numInt is < 2, the lines are parallel and in the same direction. In
	// this case the point can be ignored, since the offset lines will also be
	// parallel.
	if mark && ips[0].IsCollinear {

		// segments are collinear but reversing.
		// Add an "end-cap" fillet
		// all the way around to other direction This case should ONLY happen
		// for LineStrings, so the orientation is always CW. (Polygons can never
		// have two consecutive segments which are parallel but reversed,
		// because that would be a self intersection.
		if c.parameters.JoinStyle == calc.JoinBevel || c.parameters.JoinStyle == calc.JoinMitre {
			if addStartPoint {
				c.Add(c.offset0.P1)
			}
			c.Add(c.offset0.P0)
		} else {
			c.addCornerFillet(c.s1, c.offset0.P1, c.offset1.P0, calc.ClockWise, c.distance)
		}
	}
}

// Add points for a circular fillet around a reflex corner.
// Adds the start and end points
func (c *Curve) addCornerFillet(p, p0, p1 matrix.Matrix, direction int, radius float64) {
	dx0 := p0[0] - p[0]
	dy0 := p0[1] - p[1]
	startAngle := math.Atan2(dy0, dx0)
	dx1 := p1[0] - p[0]
	dy1 := p1[1] - p[1]
	endAngle := math.Atan2(dy1, dx1)

	if direction == calc.ClockWise {
		if startAngle <= endAngle {
			startAngle += 2.0 * math.Pi
		}
	} else { // direction == COUNTERCLOCKWISE
		if startAngle >= endAngle {
			startAngle -= 2.0 * math.Pi
		}
	}
	c.Add(p0)
	c.addDirectedFillet(p, radius, startAngle, endAngle, direction)
	c.Add(p1)
}

// Adds the offset points for an outside (convex) turn
func (c *Curve) addOutsideTurn(orientation int, addStartPoint bool) {

	// Heuristic: If offset endpoints are very close together,
	// just use one of them as the corner vertex.
	// This avoids problems with computing mitre corners in the case
	// where the two segments are almost parallel
	// (which is hard to compute a robust intersection for).
	offsetFactor := calc.OffsetSegmentSeparationFactor
	//offsetFactor = offsetFactor * math.Pow10(int(math.Log10(c.offset0.P1[1]))) / 10.0
	if measure.PlanarDistance(c.offset0.P1, c.offset1.P0) < c.distance*offsetFactor {
		c.Add(c.offset0.P1)
		return
	}

	if c.parameters.JoinStyle == calc.JoinMitre {
		c.addMitreJoin(c.s1, c.offset0, c.offset1, c.distance)
	} else if c.parameters.JoinStyle == calc.JoinBevel {
		c.addBevelJoin(c.offset0, c.offset1)
	} else {
		// add a circular fillet connecting the endpoints of the offset segments
		if addStartPoint {
			c.Add(c.offset0.P1)
		}
		// TESTING - comment out to produce beveled joins
		c.addCornerFillet(c.s1, c.offset0.P1, c.offset1.P0, orientation, c.distance)
		c.Add(c.offset1.P0)
	}
}

// Adds the offset points for an inside (concave) turn.

func (c *Curve) addInsideTurn(orientation int, addStartPoint bool) {

	// add intersection point of offset segments (if any)
	mark, ips := relate.Intersection(c.offset0.P0, c.offset0.P1, c.offset1.P0, c.offset1.P1)

	if mark {
		c.Add(ips[0].Matrix)
	} else {

		// If no intersection is detected,
		// it means the angle is so small and/or the offset so
		// large that the offsets segments don't intersect.
		// In this case we must
		// add a "closing segment" to make sure the buffer curve is continuous,
		// fairly smooth (e.g. no sharp reversals in direction)
		// and tracks the buffer correctly around the corner. The curve connects
		// the endpoints of the segment offsets to points
		// which lie toward the centre point of the corner.
		// The joining curve will not appear in the final buffer outline, since it
		// is completely internal to the buffer polygon.
		//
		// In complex buffer cases the closing segment may cut across many other
		// segments in the generated offset curve.  In order to improve the
		// performance of the noding, the closing segment should be kept as short as possible.
		// (But not too short, since that would defeat its purpose).
		// This is the purpose of the closingSegFactor heuristic value.

		// The intersection test above is vulnerable to robustness errors; i.e. it
		// may be that the offsets should intersect very close to their endpoints,
		// but aren't reported as such due to rounding. To handle this situation
		// appropriately, we use the following test: If the offset points are very
		// close, don't add closing segments but simply use one of the offset
		// points
		c.hasNarrowConcaveAngle = true
		//System.out.println("NARROW ANGLE - distance = " + distance);

		insideFactor := calc.InsideTurnVertexSnapDistanceFactor
		insideFactor = insideFactor * math.Pow10(int(math.Log10(c.offset0.P1[1]))) / 10.0
		if measure.PlanarDistance(c.offset0.P1, c.offset1.P0) < c.distance*insideFactor {
			c.Add(c.offset0.P1)
		} else {
			// add endpoint of this segment offset
			c.Add(c.offset0.P1)

			// Add "closing segment" of required length.
			if c.closingSegLengthFactor > 0 {
				mid0 := matrix.Matrix{(c.closingSegLengthFactor*c.offset0.P1[0] + c.s1[0]) / (c.closingSegLengthFactor + 1),
					(c.closingSegLengthFactor*c.offset0.P1[1] + c.s1[1]) / (c.closingSegLengthFactor + 1)}
				c.Add(mid0)

				mid1 := matrix.Matrix{(c.closingSegLengthFactor*c.offset1.P0[0] + c.s1[0]) / (c.closingSegLengthFactor + 1),
					(c.closingSegLengthFactor*c.offset1.P0[1] + c.s1[1]) / (c.closingSegLengthFactor + 1)}
				c.Add(mid1)
			} else {

				// This branch is not expected to be used except for testing purposes.
				// It is equivalent to the JTS 1.9 logic for closing segments
				// (which results in very poor performance for large buffer distances)
				c.Add(c.s1)
			}

			// add start point of next segment offset
			c.Add(c.offset1.P0)
		}
	}
}

// Adds a mitre join connecting the two reflex offset segments.
// The mitre will be beveled if it exceeds the mitre ratio limit.
func (c *Curve) addMitreJoin(p matrix.Matrix,
	offset0, offset1 *matrix.LineSegment,
	distance float64) {

	// This computation is unstable if the offset segments are nearly collinear.
	// However, this situation should have been eliminated earlier by the check
	// for whether the offset segment endpoints are almost coincident
	_, intPt := relate.Intersection(c.offset0.P0, c.offset0.P1, offset1.P0, offset1.P1)
	if intPt != nil {
		mitreRatio := 1.0
		if distance > 0.0 {
			mitreRatio = measure.PlanarDistance(intPt[0].Matrix, p) / math.Abs(distance)
		}
		if mitreRatio <= c.parameters.MitreLimit {
			c.Add(intPt[0].Matrix)
			return
		}
	}
	// at this point either intersection failed or mitre limit was exceeded
	c.addLimitedMitreJoin(offset0, offset1, distance, c.parameters.MitreLimit)
	//      addBevelJoin(offset0, offset1);
}

// Adds a limited mitre join connecting the two reflex offset segments.
// A limited mitre is a mitre which is beveled at the distance
// determined by the mitre ratio limit.
func (c *Curve) addLimitedMitreJoin(
	offset0, offset1 *matrix.LineSegment,
	distance, mitreLimit float64) {
	basePt := c.seg0.P1

	ang0 := angle.Angle(basePt, c.seg0.P0)

	// oriented angle between segments
	angDiff := angle.BetweenOriented(c.seg0.P0, basePt, c.seg1.P1)
	// half of the interior angle
	angDiffHalf := angDiff / 2

	// angle for bisector of the interior angle between the segments
	midAng := angle.Normalize(ang0 + angDiffHalf)
	// rotating this by PI gives the bisector of the reflex angle
	mitreMidAng := angle.Normalize(midAng + math.Pi)

	// the miterLimit determines the distance to the mitre bevel
	mitreDist := mitreLimit * distance
	// the bevel delta is the difference between the buffer distance
	// and half of the length of the bevel segment
	bevelDelta := mitreDist * math.Abs(math.Sin(angDiffHalf))
	bevelHalfLen := distance - bevelDelta

	// compute the midpoint of the bevel segment
	bevelMidX := basePt[0] + mitreDist*math.Cos(mitreMidAng)
	bevelMidY := basePt[1] + mitreDist*math.Sin(mitreMidAng)
	bevelMidPt := matrix.Matrix{bevelMidX, bevelMidY}

	// compute the mitre midline segment from the corner point to the bevel segment midpoint
	mitreMidLine := matrix.LineSegment{P0: basePt, P1: bevelMidPt}

	// finally the bevel segment endpoints are computed as offsets from
	// the mitre midline
	bevelEndLeft, _ := mitreMidLine.PointAlongOffset(1.0, bevelHalfLen)
	bevelEndRight, _ := mitreMidLine.PointAlongOffset(1.0, -bevelHalfLen)

	if c.side == calc.SideLeft {
		c.Add(bevelEndLeft)
		c.Add(bevelEndRight)
	} else {
		c.Add(bevelEndRight)
		c.Add(bevelEndLeft)
	}
}

// Adds a bevel join connecting the two offset segments
// around a reflex corner.

func (c *Curve) addBevelJoin(
	offset0, offset1 *matrix.LineSegment) {
	c.Add(offset0.P1)
	c.Add(offset1.P0)
}
