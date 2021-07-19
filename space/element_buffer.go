package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// ElementCentroid describes a geographic Element centroid
type ElementCentroid struct {
	*buffer.Centroid
}

// Centroid Computes the centroid point of a geometry.
func Centroid(geom Geometry) Point {
	elem := &ElementCentroid{&buffer.Centroid{}}
	elem.Add(geom)
	m := elem.GetCentroid()
	return Point(m)
}

// Add  Adds a Geometry to the centroid total
func (ec *ElementCentroid) Add(geom Geometry) {
	if geom == nil || geom.IsEmpty() {
		return
	}
	switch geom.GeoJSONType() {
	case TypePoint:
		ec.addPoint(geom.(Point))
	case TypeLineString:
		ec.addLineSegments(geom.(LineString))
	case TypePolygon:
		ec.addPolygon(geom.(Polygon))
	case TypeMultiPoint:
		for _, v := range geom.(MultiPoint) {
			ec.addPoint(v)
		}
	case TypeMultiLineString:
		for _, v := range geom.(MultiLineString) {
			ec.addLineSegments(v)
		}
	case TypeMultiPolygon:
		for _, v := range geom.(MultiPolygon) {
			ec.addPolygon(v)
		}
	case TypeCollection:
		for _, v := range geom.(Collection) {
			ec.Add(v)
		}
	case TypeBound:
		ec.addLineSegments(LineString(geom.Bound().ToRing()))
	default:
	}
}

// addPoint Adds a point to the point centroid accumulator.
func (ec *ElementCentroid) addPoint(pt Point) {
	ec.PtCount++
	if ec.PtCentSum == nil {
		ec.PtCentSum = make(matrix.Matrix, 2)
	}
	ec.PtCentSum[0] += pt.X()
	ec.PtCentSum[1] += pt.Y()
}

// addLineSegments Adds the line segments  to the linear centroid accumulators.
func (ec *ElementCentroid) addLineSegments(lines LineString) {
	linelen := 0.0
	if ec.LineCentSum == nil {
		ec.LineCentSum = matrix.Matrix{0, 0}
	}
	for i := 0; i < len(lines)-1; i++ {
		segmentLen, _ := Point(lines[i]).Distance(Point(lines[i+1]))
		if segmentLen == 0.0 {
			continue
		}
		linelen += segmentLen
		midx := (lines[i][0] + lines[i+1][0]) / 2
		midy := (lines[i][1] + lines[i+1][1]) / 2
		ec.LineCentSum[0] += segmentLen * midx
		ec.LineCentSum[1] += segmentLen * midy
	}
	ec.TotalLength += linelen
	if linelen == 0.0 && len(lines) > 0 {
		ec.addPoint(lines[0])
	}
}

// addPolygon Adds the polygon  to the polygon centroid accumulators.
func (ec *ElementCentroid) addPolygon(poly Polygon) {
	for i, v := range poly {
		isPositiveArea := false
		if i == 0 {
			if len(v) > 0 {
				ec.AreaBasePt = v[0]
			}
			isPositiveArea = !measure.IsCCW(v)
		} else {
			isPositiveArea = measure.IsCCW(v)
		}
		for i := 0; i < len(v)-1; i++ {
			ec.addTriangle(Point(ec.AreaBasePt), v[i], v[i+1], isPositiveArea)
		}
		ec.addLineSegments(v)
	}
}

// addTriangle Adds the Triangle  to the Triangle centroid accumulators.
func (ec *ElementCentroid) addTriangle(p0, p1, p2 Point, isPositiveArea bool) {
	sign := 1.0
	if !isPositiveArea {
		sign = -1.0
	}
	ec.TriangleCent3 = matrix.Matrix{p0.X() + p1.X() + p2.X(), p0.Y() + p1.Y() + p2.Y()}
	area2 := (p1.X()-p0.X())*(p2.Y()-p0.Y()) -
		(p2.X()-p0.X())*(p1.Y()-p0.Y())
	if ec.Cg3 == nil {
		ec.Cg3 = matrix.Matrix{0, 0}
	}
	ec.Cg3[0] += sign * area2 * ec.TriangleCent3[0]
	ec.Cg3[1] += sign * area2 * ec.TriangleCent3[1]
	ec.Areasum2 += sign * area2
}

// ElementBuffer describes a geographic Element buffer
type ElementBuffer struct {
	*buffer.CurveBuilder
	distance float64
	param    *buffer.CurveParameters
}

// Buffer Computes the set of raw offset curves for the buffer.
// Each offset curve has an attached {@link Label} indicating
// its left and right location.
func Buffer(geom Geometry, distance float64) Geometry {
	eb := ElementBuffer{}
	eb.distance = distance
	eb.Add(geom)
	bufferSeg := eb.CurveBuilder.Curves
	if len(bufferSeg) <= 0 {
		return nil
	}
	poly := Polygon{}
	for _, v := range bufferSeg {
		poly = append(poly, v.Line)
	}
	return poly
}

// Add Add a geometry to the graph.
func (eb *ElementBuffer) Add(geom Geometry) {
	if geom.IsEmpty() {
		return
	}
	if eb.param == nil || eb.param.IsEmpty() {
		eb.param = buffer.DefaultCurveParameters()
		eb.CurveBuilder = &buffer.CurveBuilder{
			Curve: buffer.DefaultCurve(),
		}
	}
	switch geom.GeoJSONType() {
	case TypePoint:
		eb.addPoint(geom.(Point))
	case TypeLineString:
		eb.addLineString(geom.(LineString))
	case TypePolygon:
		eb.addPolygon(geom.(Polygon))
	case TypeMultiPoint:
		for _, v := range geom.(MultiPoint) {
			eb.addPoint(v)
		}
	case TypeMultiLineString:
		for _, v := range geom.(MultiLineString) {
			eb.addLineString(v)
		}
	case TypeMultiPolygon:
		for _, v := range geom.(MultiPolygon) {
			eb.addPolygon(v)
		}
	case TypeCollection:
		for _, v := range geom.(Collection) {
			eb.Add(v)
		}
	case TypeBound:
		eb.addLineString(LineString(geom.Bound().ToRing()))
	default:
	}
}

// addPoint Add a Point to the graph.
func (eb *ElementBuffer) addPoint(p Point) {
	// a zero or negative width buffer of a point is empty
	if eb.distance <= 0.0 {
		return
	}
	eb.LineCurve(matrix.LineMatrix{matrix.Matrix(p)}, eb.distance, buffer.EXTERIOR, buffer.INTERIOR)
}

// addLineString Add a LineString to the graph.
func (eb *ElementBuffer) addLineString(line LineString) {
	if eb.isLineOffsetEmpty(eb.distance) {
		return
	}
	if line.IsRing() && eb.param.IsSingleSided {
		eb.LineCurve(matrix.LineMatrix(line), eb.distance, buffer.EXTERIOR, buffer.INTERIOR)
	} else {
		eb.LineCurve(matrix.LineMatrix(line), eb.distance, buffer.EXTERIOR, buffer.INTERIOR)
	}

}

// isLineOffsetEmpty Tests whether the offset curve for line or point geometries
// at the given offset distance is empty (does not exist).
// This is the case if:
// the distance is zero,
// the distance is negative, except for the case of singled-sided buffers
func (eb *ElementBuffer) isLineOffsetEmpty(distance float64) bool {
	// a zero width buffer of a line or point is empty
	if distance == 0.0 {
		return true
	}
	// a negative width buffer of a line or point is empty,
	// except for single-sided buffers, where the sign indicates the side
	if distance < 0.0 && !eb.param.IsSingleSided {
		return true
	}
	return false
}
func (eb *ElementBuffer) addPolygon(p Polygon) {
	offsetDistance := eb.distance
	offsetSide := buffer.LEFT
	if eb.distance < 0.0 {
		offsetDistance = -eb.distance
		offsetSide = buffer.RIGHT
	}

	shell := p.Shell()

	if eb.distance <= 0.0 && len(shell) < 3 {
		return
	}
	eb.addRingSide(
		shell,
		offsetDistance,
		offsetSide,
		buffer.EXTERIOR,
		buffer.INTERIOR)

	if offsetSide == buffer.LEFT {
		offsetSide = buffer.RIGHT
	}
	if offsetSide == buffer.RIGHT {
		offsetSide = buffer.LEFT
	}

	for i := 0; i < len(p.Holes()); i++ {

		hole := p[i]

		// Holes are topologically labelled opposite to the shell, since
		// the interior of the polygon lies on their opposite side
		// (on the left, if the hole is oriented CCW)
		eb.addRingSide(
			hole,
			offsetDistance,
			offsetSide,
			buffer.INTERIOR,
			buffer.EXTERIOR)
	}
}
func (eb *ElementBuffer) addRingBothSides(ring Ring, distance float64) {
	eb.addRingSide(ring, distance,
		buffer.LEFT,
		buffer.EXTERIOR, buffer.INTERIOR)
	/* Add the opposite side of the ring
	 */
	eb.addRingSide(ring, distance,
		buffer.RIGHT,
		buffer.INTERIOR, buffer.EXTERIOR)
}

// addRingSide Adds an offset curve for one side of a ring.
// The side and left and right topological location arguments
// are provided as if the ring is oriented CW.
// (If the ring is in the opposite orientation,
// this is detected and
// the left and right locations are interchanged and the side is flipped.)
func (eb *ElementBuffer) addRingSide(ring Ring, offsetDistance float64, side, cwLeftLoc, cwRightLoc int) {
	// don't bother adding ring if it is "flat" and will disappear in the output
	if offsetDistance == 0.0 && len(ring) < buffer.MinRingSize {
		return
	}

	leftLoc := cwLeftLoc
	rightLoc := cwRightLoc
	isCCW := measure.IsCCW(matrix.LineMatrix(ring))
	if len(ring) >= buffer.MinRingSize && isCCW {
		leftLoc = cwRightLoc
		rightLoc = cwLeftLoc
		if side == buffer.LEFT {
			side = buffer.RIGHT
		}
		if side == buffer.RIGHT {
			side = buffer.LEFT
		}

	}
	eb.RingCurve(matrix.LineMatrix(ring), offsetDistance, side, leftLoc, rightLoc)

}
