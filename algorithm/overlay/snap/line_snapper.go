package snap

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// LineSnapper ...
type LineSnapper struct {
	srcPts        matrix.LineMatrix
	snapPts       matrix.Steric
	snapTolerance float64
}

// Snaps the vertices and segments of the source LineString
// to the given set of snap vertices.
func (l *LineSnapper) snapTo() matrix.LineMatrix {
	l.snapVertices()
	l.snapSegments()
	return l.srcPts
}

// Snap source vertices to vertices in the target.
func (l *LineSnapper) snapVertices() {
	// try snapping vertices
	// if src is a ring then don't snap final vertex
	end := len(l.srcPts)
	isClosed := l.srcPts.IsClosed()
	if isClosed {
		end--
	}

	for i := 0; i < end; i++ {
		srcPt := l.srcPts[i]
		snapVert := l.findSnapForVertex(srcPt)
		if snapVert != nil {
			// update src with snap pt
			l.srcPts[i] = snapVert
			// keep final closing point in synch (rings only)
			if i == 0 && isClosed {
				l.srcPts[len(l.srcPts)-1] = snapVert
			}
		}
	}
}

func (l *LineSnapper) findSnapForVertex(pt matrix.Matrix) matrix.Matrix {
	for _, v := range matrix.TransMatrixs(l.snapPts) {
		spt := matrix.Matrix(v)
		// if point is already equal to a src pt, don't snap
		if pt.Equals(spt) {
			return nil
		}
		if dist := measure.PlanarDistance(pt, spt); dist < l.snapTolerance {
			return spt
		}
	}
	return nil
}

// Snap segments of the source to nearby snap vertices.
// Source segments are "cracked" at a snap vertex.
// A single input segment may be snapped several times
// to different snap vertices.
// <p>
// For each distinct snap vertex, at most one source segment
// is snapped to.  This prevents "cracking" multiple segments
// at the same point, which would likely cause
// topology collapse when being used on polygonal linework.
func (l *LineSnapper) snapSegments() {
	sPts := matrix.TransMatrixs(l.snapPts)
	// guard against empty input
	if len(sPts) == 0 {
		return
	}

	distinctPtCount := len(sPts)

	// check for duplicate snap pts when they are sourced from a linear ring.
	// TODO: Need to do this better - need to check *all* snap points for dups (using a Set?)
	if sPts[0].Equals(sPts[distinctPtCount-1]) {
		distinctPtCount--
	}

	for i := 0; i < distinctPtCount; i++ {
		snapPt := sPts[i]
		index := l.findSegmentIndexToSnap(snapPt)

		// If a segment to snap to was found, "crack" it at the snap pt.
		// The new pt is inserted immediately into the src segment list,
		// so that subsequent snapping will take place on the modified segments.
		// Duplicate points are not added.
		if index >= 0 {
			l.srcPts = append(l.srcPts, l.srcPts[:index]...)
			l.srcPts = append(l.srcPts, snapPt)
			l.srcPts = append(l.srcPts, l.srcPts[index:]...)
		}
	}
}

// Finds a src segment which snaps to (is close to) the given snap point.
// <p>
// Only a single segment is selected for snapping.
// This prevents multiple segments snapping to the same snap vertex,
// which would almost certainly cause invalid geometry
// to be created.
// (The heuristic approach to snapping used here
// is really only appropriate when
// snap pts snap to a unique spot on the src geometry.)
// <p>
// Also, if the snap vertex occurs as a vertex in the src coordinate list,
// no snapping is performed.
func (l *LineSnapper) findSegmentIndexToSnap(snapPt matrix.Matrix) int {
	minDist := math.MaxFloat64
	snapIndex := -1
	for i := 0; i < len(l.srcPts)-1; i++ {
		p0 := matrix.Matrix(l.srcPts[i])
		p1 := matrix.Matrix(l.srcPts[i+1])

		// Check if the snap pt is equal to one of the segment endpoints.
		// If the snap pt is already in the src list, don't snap at all.
		if p0.Equals(snapPt) || p1.Equals(snapPt) {
			return -1
		}

		dist := measure.DistanceSegmentToPoint(snapPt, p0, p1, measure.PlanarDistance)
		if dist < l.snapTolerance && dist < minDist {
			minDist = dist
			snapIndex = i
		}
	}
	return snapIndex
}
