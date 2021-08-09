package buffer

import (
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// ComputerBuffer describes a geographic Element buffer
type ComputerBuffer struct {
	*CurveBuilder
	distance float64
	param    *CurveParameters
}

// Buffer Computes the set of raw offset curves for the buffer.
// Each offset curve has an attached  indicating
// its left and right location.
func Buffer(geom matrix.Steric, distance float64, quadsegs int) matrix.Steric {
	eb := ComputerBuffer{}
	eb.param = DefaultCurveParameters()
	eb.param.QuadrantSegments = quadsegs
	eb.distance = distance
	eb.CurveBuilder = &CurveBuilder{
		Curve: CurveWithParameters(eb.param, eb.distance),
	}

	eb.Add(geom)
	bufferSeg := eb.CurveBuilder.Curves
	if len(bufferSeg) <= 0 {
		return nil
	}
	poly := matrix.PolygonMatrix{}
	for _, v := range bufferSeg {
		poly = append(poly, v.Line)
	}
	return poly
}

// Add Add a geometry to the graph.
func (eb *ComputerBuffer) Add(geom matrix.Steric) {
	if geom.IsEmpty() {
		return
	}
	if eb.param == nil || eb.param.IsEmpty() {
		eb.param = DefaultCurveParameters()
		eb.CurveBuilder = &CurveBuilder{
			Curve: CurveWithParameters(eb.param, eb.distance),
		}
	}
	switch st := geom.(type) {
	case matrix.Matrix:
		eb.addPoint(st)
	case matrix.LineMatrix:
		eb.addLineString(st)
	case matrix.PolygonMatrix:
		eb.addPolygon(st)
	case matrix.Collection:
		// TODO add support collection
		for i, v := range st {
			if i == 0 {
				eb.Add(v)
			}
		}
	}
}

// addPoint Add a Point to the graph.
func (eb *ComputerBuffer) addPoint(p matrix.Matrix) {
	// a zero or negative width buffer of a point is empty
	if eb.distance <= 0.0 {
		return
	}
	eb.LineCurve(matrix.LineMatrix{matrix.Matrix(p)}, eb.distance, calc.EXTERIOR, calc.INTERIOR)
}

// addLineString Add a LineString to the graph.
func (eb *ComputerBuffer) addLineString(line matrix.LineMatrix) {
	if eb.isLineOffsetEmpty(eb.distance) {
		return
	}

	if matrix.Matrix(line[0]).Equals(matrix.Matrix((line[len(line)-1]))) && eb.param.IsSingleSided {
		eb.LineCurve(matrix.LineMatrix(line), eb.distance, calc.EXTERIOR, calc.INTERIOR)
	} else {
		eb.LineCurve(matrix.LineMatrix(line), eb.distance, calc.EXTERIOR, calc.INTERIOR)
	}

}

// isLineOffsetEmpty Tests whether the offset curve for line or point geometries
// at the given offset distance is empty (does not exist).
// This is the case if:
// the distance is zero,
// the distance is negative, except for the case of singled-sided buffers
func (eb *ComputerBuffer) isLineOffsetEmpty(distance float64) bool {
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
func (eb *ComputerBuffer) addPolygon(p matrix.PolygonMatrix) {
	offsetDistance := eb.distance
	offsetSide := calc.LEFT
	if eb.distance < 0.0 {
		offsetDistance = -eb.distance
		offsetSide = calc.RIGHT
	}

	shell := p[0]

	if eb.distance <= 0.0 && len(shell) < 3 {
		return
	}
	eb.addRingSide(
		shell,
		offsetDistance,
		offsetSide,
		calc.EXTERIOR,
		calc.INTERIOR)

	if offsetSide == calc.LEFT {
		offsetSide = calc.RIGHT
	}
	if offsetSide == calc.RIGHT {
		offsetSide = calc.LEFT
	}

	for i := 1; i < len(p); i++ {

		hole := p[i]

		// Holes are topologically labelled opposite to the shell, since
		// the interior of the polygon lies on their opposite side
		// (on the left, if the hole is oriented CCW)
		eb.addRingSide(
			hole,
			offsetDistance,
			offsetSide,
			calc.INTERIOR,
			calc.EXTERIOR)
	}
}
func (eb *ComputerBuffer) addRingBothSides(ring matrix.LineMatrix, distance float64) {
	eb.addRingSide(ring, distance,
		calc.LEFT,
		calc.EXTERIOR, calc.INTERIOR)
	// Add the opposite side of the ring
	eb.addRingSide(ring, distance,
		calc.RIGHT,
		calc.INTERIOR, calc.EXTERIOR)
}

// addRingSide Adds an offset curve for one side of a ring.
// The side and left and right topological location arguments
// are provided as if the ring is oriented CW.
// (If the ring is in the opposite orientation,
// this is detected and
// the left and right locations are interchanged and the side is flipped.)
func (eb *ComputerBuffer) addRingSide(ring matrix.LineMatrix, offsetDistance float64, side, cwLeftLoc, cwRightLoc int) {
	// don't bother adding ring if it is "flat" and will disappear in the output
	if offsetDistance == 0.0 && len(ring) < calc.MinRingSize {
		return
	}

	leftLoc := cwLeftLoc
	rightLoc := cwRightLoc
	isCCW := measure.IsCCW(matrix.LineMatrix(ring))
	if len(ring) >= calc.MinRingSize && isCCW {
		leftLoc = cwRightLoc
		rightLoc = cwLeftLoc
		if side == calc.LEFT {
			side = calc.RIGHT
		}
		if side == calc.RIGHT {
			side = calc.LEFT
		}

	}
	eb.RingCurve(matrix.LineMatrix(ring), offsetDistance, side, leftLoc, rightLoc)

}
