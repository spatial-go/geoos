// Package buffer define geomtry matrix conversion.
package buffer

import (
	"math"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/calc/angle"
	"github.com/spatial-go/geoos/algorithm/graph/clipping"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// VariableLineBuffer describes a line variable buffer
type VariableLineBuffer struct {
	Line             matrix.LineMatrix
	QuadrantSegments int
}

// InterpolatedBuffer Creates a buffer polygon along a line with the buffer distance interpolated
// between a start distance and an end distance.
func (v *VariableLineBuffer) InterpolatedBuffer(startDistance, endDistance float64) matrix.Steric {
	distance := interpolate(v.Line,
		startDistance, endDistance)
	return v.DistancesBuffer(distance)
}

// DistancesBuffer Creates a buffer polygon along a line with the buffer distances.
func (v *VariableLineBuffer) DistancesBuffer(distances []float64) (buffer matrix.Steric) {

	partsGeom := matrix.Collection{}
	for i := 1; i < len(v.Line); i++ {
		dist0 := distances[i-1]
		dist1 := distances[i]

		if dist0 > 0 || dist1 > 0 {
			poly := v.segmentBuffer(v.Line[i-1], v.Line[i], dist0, dist1)
			if !poly.IsEmpty() {
				partsGeom = append(partsGeom, poly)
			}
		}
	}
	buffer, _ = clipping.UnaryUnion(partsGeom)

	if !geoos.GeoosTestTag {
		matrix.WriteMatrix("collection", partsGeom)
	}

	return
}

// Computes a list of values for the points along a line by interpolating between values for the start and end point.
// The interpolation is based on the distance of each point along the line relative to the total line length.
func interpolate(line matrix.LineMatrix, startDistance, endDistance float64) []float64 {
	startValue := math.Abs(startDistance)
	endValue := math.Abs(endDistance)
	values := make([]float64, len(line))
	values[0] = startValue
	values[len(values)-1] = endValue

	totalLen := measure.OfLine(line)
	currLen := 0.0
	for i := 1; i < len(values)-1; i++ {
		from := matrix.Matrix(line[i])
		to := matrix.Matrix(line[i-1])
		segLen := measure.PlanarDistance(from, to)
		currLen += segLen
		lenFrac := currLen / totalLen
		delta := lenFrac * (endValue - startValue)
		values[i] = startValue + delta
	}
	return values
}

// Computes a variable buffer polygon for a single segment,with the given endpoints and buffer distances.
// The individual segment buffers are unioned to form the final buffer.
func (v *VariableLineBuffer) segmentBuffer(p0, p1 matrix.Matrix,
	dist0, dist1 float64) matrix.Steric {
	/**
	 * Compute for increasing distance only, so flip if needed
	 */
	if dist0 > dist1 {
		return v.segmentBuffer(p1, p0, dist1, dist0)
	}

	// forward tangent line
	tangent := outerTangent(p0, p1, dist0, dist1)
	// tangent := matrix.LineSegment{}

	// if tangent is null then compute a buffer for largest circle

	if tangent.P0.IsEmpty() {
		center := p0
		dist := dist0
		if dist1 > dist0 {
			center = p1
			dist = dist1
		}
		return Buffer(center, dist, v.QuadrantSegments)
	}
	t0 := tangent.P0
	t1 := tangent.P1

	// reverse tangent line on other side of segment
	seg := matrix.LineSegment{P0: p0, P1: p1}
	tr0 := seg.Reflected(t0)
	tr1 := seg.Reflected(t1)

	coords := matrix.LineMatrix{t0, t1}

	// end cap
	coords = v.addCap(p1, dist1, t1, tr1, coords)

	coords = append(coords, tr1, tr0)
	// start cap
	coords = v.addCap(p0, dist0, tr0, t0, coords)

	// close
	coords = append(coords, t0)

	polygon := matrix.PolygonMatrix{coords}
	return polygon
}

// Adds a semi-circular cap CCW around the point p.
func (v *VariableLineBuffer) addCap(p matrix.Matrix, r float64, t1, t2 matrix.Matrix, coords matrix.LineMatrix) matrix.LineMatrix {

	angStart := angle.Angle(p, t1)
	angEnd := angle.Angle(p, t2)
	if angStart < angEnd {
		angStart += 2 * math.Pi
	}
	indexStart := v.capAngleIndex(angStart)
	indexEnd := v.capAngleIndex(angEnd)

	for i := indexStart; i > indexEnd; i-- {
		// use negative increment to create points CW
		ang := v.capAngle(i)
		coords = append(coords, projectPolar(p, r, ang))
	}
	return coords
}

// Computes the canonical cap point index for a given angle.
// The angle is rounded down to the next lower index.
func (v *VariableLineBuffer) capAngleIndex(ang float64) int {
	capSegAng := math.Pi / 2.0 / float64(v.QuadrantSegments)
	index := int(ang / capSegAng)
	return index
}

// Computes the angle for the given cap point index.
func (v *VariableLineBuffer) capAngle(index int) float64 {
	capSegAng := math.Pi / 2.0 / float64(v.QuadrantSegments)
	return capSegAng * float64(index)
}

func projectPolar(p matrix.Matrix, r, ang float64) matrix.Matrix {
	x := p[0] + r*snapTrig(math.Cos(ang))
	y := p[1] + r*snapTrig(math.Sin(ang))
	return matrix.Matrix{x, y}
}

// Snap trig values to integer values for better consistency.
func snapTrig(x float64) float64 {
	if x > (1 - calc.SnapTrigTol) {
		return 1
	}
	if x < (-1 + calc.SnapTrigTol) {
		return -1
	}
	if math.Abs(x) < calc.SnapTrigTol {
		return 0
	}
	return x
}

// Computes the two circumference points defining the outer tangent line between two circles.
// For the algorithm see <a href='https://en.wikipedia.org/wiki/Tangent_lines_to_circles#Outer_tangent'>Wikipedia</a>.
func outerTangent(c1, c2 matrix.Matrix, r1, r2 float64) matrix.LineSegment {
	/**
	 * If distances are inverted then flip to compute and flip result back.
	 */
	if r1 > r2 {
		seg := outerTangent(c2, c1, r2, r1)
		return seg
	}
	x1 := c1[0]
	y1 := c1[1]
	x2 := c2[0]
	y2 := c2[1]
	// TODO: handle r1 == r2?
	a3 := -math.Atan2(y2-y1, x2-x1)

	dr := r2 - r1
	d := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))

	a2 := math.Asin(dr / d)
	// check if no tangent exists
	if math.IsNaN(a2) {
		return matrix.LineSegment{}
	}
	a1 := a3 - a2

	aa := math.Pi/2 - a1
	x3 := x1 + r1*math.Cos(aa)
	y3 := y1 + r1*math.Sin(aa)
	x4 := x2 + r2*math.Cos(aa)
	y4 := y2 + r2*math.Sin(aa)

	return matrix.LineSegment{P0: matrix.Matrix{x3, y3}, P1: matrix.Matrix{x4, y4}}
}
