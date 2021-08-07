package buffer

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// const ...
const (
	INIT          = 0
	DELETE        = 1
	KEEP          = 1
	NUMPTSTOCHECK = 10
	DPSAFEEPSILON = 1e-15
)

// LineSimplifier  Simplifies a buffer input line to
// remove concavities with shallow depth.
type LineSimplifier struct {
	inputLine         matrix.LineMatrix
	distanceTol       float64
	isDeleted         []byte
	angleOrientation  int
	leftLoc, rightLoc int
	parameters        *CurveParameters
}

// Simplify the input coordinate list.
// If the distance tolerance is positive,
// concavities on the LEFT side of the line are simplified.
// If the supplied distance tolerance is negative,
// concavities on the RIGHT side of the line are simplified.
func (l *LineSimplifier) Simplify(distanceTol float64) matrix.LineMatrix {
	l.distanceTol = math.Abs(distanceTol)
	if distanceTol < 0 {
		l.angleOrientation = calc.CLOCKWISE
	}

	// rely on fact that boolean array is filled with false value
	l.isDeleted = make([]byte, len(l.inputLine))

	for isChanged := false; isChanged; isChanged = l.deleteShallowConcavities() {

	}
	return l.collapseLine()
}

/**
* Uses a sliding window containing 3 vertices to detect shallow angles
* in which the middle vertex can be deleted, since it does not
* affect the shape of the resulting buffer in a significant way.
* @return
 */
func (l *LineSimplifier) deleteShallowConcavities() bool {
	/**
	 * Do not simplify end line segments of the line string.
	 * This ensures that end caps are generated consistently.
	 */
	index := 1

	midIndex := l.findNextNonDeletedIndex(index)
	lastIndex := l.findNextNonDeletedIndex(midIndex)

	isChanged := false
	for lastIndex < len(l.inputLine) {
		// test triple for shallow concavity
		isMiddleVertexDeleted := false
		if l.isDeletable(index, midIndex, lastIndex,
			l.distanceTol) {
			l.isDeleted[midIndex] = DELETE
			isMiddleVertexDeleted = true
			isChanged = true
		}
		// move simplification window forward
		if isMiddleVertexDeleted {
			index = lastIndex
		} else {
			index = midIndex
		}
		midIndex = l.findNextNonDeletedIndex(index)
		lastIndex = l.findNextNonDeletedIndex(midIndex)
	}
	return isChanged
}

/**
* Finds the next non-deleted index, or the end of the point array if none
* @param index
* @return the next non-deleted index, if any
* or inputLine.length if there are no more non-deleted indices
 */
func (l *LineSimplifier) findNextNonDeletedIndex(index int) int {
	next := index + 1
	for next < len(l.inputLine) && l.isDeleted[next] == DELETE {
		next++
	}
	return next
}

func (l *LineSimplifier) collapseLine() matrix.LineMatrix {
	coordList := matrix.LineMatrix{}
	for i := 0; i < len(l.inputLine); i++ {
		if l.isDeleted[i] != DELETE {
			coordList = append(coordList, l.inputLine[i])
		}
	}
	return coordList
}

func (l *LineSimplifier) isDeletable(i0, i1, i2 int, distanceTol float64) bool {
	p0 := l.inputLine[i0]
	p1 := l.inputLine[i1]
	p2 := l.inputLine[i2]

	if !l.isConcave(p0, p1, p2) {
		return false
	}
	if !l.isShallow(p0, p1, p2, distanceTol) {
		return false
	}

	return l.isShallowSampled(p0, p1, i0, i2, distanceTol)
}

func (l *LineSimplifier) isShallowConcavity(p0, p1, p2 matrix.Matrix, distanceTol float64) bool {
	orientation := l.orientationIndex(p0[0], p0[1], p1[0], p1[1], p2[0], p2[1])
	isAngleToSimplify := (orientation == l.angleOrientation)
	if !isAngleToSimplify {
		return false
	}
	dist := measure.DistanceSegmentToPoint(p2, p1, p0, measure.PlanarDistance)
	return dist < distanceTol
}

/**
* Checks for shallowness over a sample of points in the given section.
* This helps prevents the simplification from incrementally
* "skipping" over points which are in fact non-shallow.
*
* @param p0 start coordinate of section
* @param p2 end coordinate of section
* @param i0 start index of section
* @param i2 end index of section
* @param distanceTol distance tolerance
* @return
 */
func (l *LineSimplifier) isShallowSampled(p0, p2 matrix.Matrix, i0, i2 int, distanceTol float64) bool {
	// check every n'th point to see if it is within tolerance
	inc := (i2 - i0) / NUMPTSTOCHECK
	if inc <= 0 {
		inc = 1
	}
	for i := i0; i < i2; i += inc {
		if !l.isShallow(p0, p2, l.inputLine[i], distanceTol) {
			return false
		}
	}
	return true
}

func (l *LineSimplifier) isShallow(p0, p1, p2 matrix.Matrix, distanceTol float64) bool {
	dist := measure.DistanceSegmentToPoint(p2, p1, p0, measure.PlanarDistance)
	return dist < distanceTol
}

func (l *LineSimplifier) isConcave(p0, p1, p2 matrix.Matrix) bool {
	orientation := l.orientationIndex(p0[0], p0[1], p1[0], p1[1], p2[0], p2[1])
	isConcave := (orientation == l.angleOrientation)
	return isConcave
}

func (l *LineSimplifier) orientationIndex(p1x, p1y,
	p2x, p2y,
	qx, qy float64) int {
	// fast filter for orientation index
	// avoids use of slow extended-precision arithmetic in many cases
	index := l.orientationIndexFilter(p1x, p1y, p2x, p2y, qx, qy)
	if index <= 1 {
		return index
	}
	// normalize coordinates
	dx1 := (&calc.PairFloat{Hi: p2x, Lo: 0.0}).SelfAdd(-p1x, 0.0)
	dy1 := (&calc.PairFloat{Hi: p2y, Lo: 0.0}).SelfAdd(-p1y, 0.0)
	dx2 := (&calc.PairFloat{Hi: qx, Lo: 0.0}).SelfAdd(-p2x, 0.0)
	dy2 := (&calc.PairFloat{Hi: qy, Lo: 0.0}).SelfAdd(-p2y, 0.0)
	dy11 := dy1.SelfMultiply(dx2.Hi, dx2.Lo)

	// sign of determinant - unrolled for performance
	return dx1.SelfMultiply(dy2.Hi, dy2.Lo).SelfSubtract(dy11.Hi, dy11.Lo).Signum()
}

// orientationIndexFilter A filter for computing the orientation index of three coordinates.
func (l *LineSimplifier) orientationIndexFilter(pax, pay,
	pbx, pby, pcx, pcy float64) int {
	detsum := 0.0

	detleft := (pax - pcx) * (pby - pcy)
	detright := (pay - pcy) * (pbx - pcx)
	det := detleft - detright

	if detleft > 0.0 {
		if detright <= 0.0 {
			return calc.Signum(det)
		}
		detsum = detleft + detright
	} else if detleft < 0.0 {
		if detright >= 0.0 {
			return calc.Signum(det)
		}
		detsum = -detleft - detright
	} else {
		return calc.Signum(det)
	}

	errbound := DPSAFEEPSILON * detsum
	if (det >= errbound) || (-det >= errbound) {
		return calc.Signum(det)
	}

	return 2
}
