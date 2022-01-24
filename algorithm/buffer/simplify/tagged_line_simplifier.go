package simplify

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// TaggedLineStringSimplifier Simplifies a TaggedLineString, preserving topology
//  (in the sense that no new intersections are introduced). Uses the recursive Douglas-Peucker algorithm.
type TaggedLineStringSimplifier struct {
	inputIndex, outputIndex *LineSegmentIndex
	line                    *TaggedLineString
	linePts                 []matrix.Matrix
	distanceTolerance       float64
}

// Simplify Simplifies the given TaggedLineString  using the distance tolerance specified.
func (t *TaggedLineStringSimplifier) Simplify(line *TaggedLineString) {
	t.line = line
	t.linePts = line.GetParentMatrixes()
	t.simplifySection(0, len(t.linePts)-1, 0)
}

func (t *TaggedLineStringSimplifier) simplifySection(i, j, depth int) {
	depth++
	sectionIndex := []int{0, 0}
	if (i + 1) == j {
		newSeg := t.line.GetSegment(i)
		t.line.AddToResult(newSeg.LineSegment)
		// leave this segment in the input index, for efficiency
		return
	}

	isValidToSimplify := true

	if t.line.GetResultSize() < t.line.MinimumSize {
		worstCaseSize := depth + 1
		if worstCaseSize < t.line.MinimumSize {
			isValidToSimplify = false
		}
	}

	distance := []float64{0, 0}
	furthestPtIndex := t.findFurthestPoint(t.linePts, i, j, distance)
	// flattening must be less than distanceTolerance
	if distance[0] > t.distanceTolerance {
		isValidToSimplify = false
	}
	// test if flattened section would cause intersection
	candidateSeg := &matrix.LineSegment{}
	candidateSeg.P0 = t.linePts[i]
	candidateSeg.P1 = t.linePts[j]
	sectionIndex[0] = i
	sectionIndex[1] = j
	if t.hasBadIntersection(t.line, sectionIndex, candidateSeg) {
		isValidToSimplify = false
	}
	if isValidToSimplify {
		newSeg := t.flatten(i, j)
		t.line.AddToResult(newSeg)
		return
	}
	t.simplifySection(i, furthestPtIndex, depth)
	t.simplifySection(furthestPtIndex, j, depth)
}

func (t *TaggedLineStringSimplifier) findFurthestPoint(pts []matrix.Matrix, i, j int, maxDistance []float64) int {
	seg := &matrix.LineSegment{}
	seg.P0 = pts[i]
	seg.P1 = pts[j]
	maxDist := -1.0
	maxIndex := i
	for k := i + 1; k < j; k++ {
		midPt := pts[k]
		distance := measure.PlanarDistance(midPt, matrix.LineMatrix{seg.P0, seg.P1})
		if distance > maxDist {
			maxDist = distance
			maxIndex = k
		}
	}
	maxDistance[0] = maxDist
	return maxIndex
}

// flatten Flattens a section of the line between
func (t *TaggedLineStringSimplifier) flatten(start, end int) *matrix.LineSegment {
	// make a new segment for the simplified geometry
	p0 := t.linePts[start]
	p1 := t.linePts[end]
	newSeg := &matrix.LineSegment{P0: p0, P1: p1}
	// update the indexes
	t.Remove(t.line, start, end)
	t.outputIndex.AddSegment(newSeg)
	return newSeg
}

func (t *TaggedLineStringSimplifier) hasBadIntersection(parentLine *TaggedLineString,
	sectionIndex []int,
	candidateSeg *matrix.LineSegment) bool {
	if t.hasBadOutputIntersection(candidateSeg) {
		return true
	}
	if t.hasBadInputIntersection(parentLine, sectionIndex, candidateSeg) {
		return true
	}
	return false
}

func (t *TaggedLineStringSimplifier) hasBadOutputIntersection(candidateSeg *matrix.LineSegment) bool {
	querySegs := t.outputIndex.Query(candidateSeg)
	for _, v := range querySegs {
		querySeg := v
		if HasInteriorIntersection(querySeg, candidateSeg) {
			return true
		}
	}
	return false
}
func (t *TaggedLineStringSimplifier) hasBadInputIntersection(parentLine *TaggedLineString,
	sectionIndex []int,
	candidateSeg *matrix.LineSegment) bool {
	querySegs := t.inputIndex.Query(candidateSeg)
	for _, v := range querySegs {
		querySeg := v
		if HasInteriorIntersection(querySeg, candidateSeg) {
			if IsInLineSection(parentLine, sectionIndex, querySeg) {
				continue
			}
			return true
		}
	}
	return false
}

// IsInLineSection Tests whether a segment is in a section of a TaggedLineString
func IsInLineSection(
	line *TaggedLineString,
	sectionIndex []int,
	seg *matrix.LineSegment) bool {
	for i, v := range line.Segs {
		if seg.P0.Equals(v.P0) && seg.P1.Equals(v.P1) {
			if i >= sectionIndex[0] && i < sectionIndex[1] {
				return true
			}
		}
	}
	return false
}

// HasInteriorIntersection ..
func HasInteriorIntersection(seg0, seg1 *matrix.LineSegment) bool {
	return relate.IsIntersectionLineSegment(seg0, seg1)
}

// Remove Remove the segs in the section of the line
func (t *TaggedLineStringSimplifier) Remove(line *TaggedLineString,
	start, end int) {
	for i := start; i < end; i++ {
		seg := line.GetSegment(i)
		if seg == nil {
			continue
		}
		t.inputIndex.Remove(seg)
	}
}

// TaggedLinesSimplifier Simplifies a collection of TaggedLineStrings, preserving topology
//  (in the sense that no new intersections are introduced).
type TaggedLinesSimplifier struct {
	inputIndex, outputIndex *LineSegmentIndex
	distanceTolerance       float64
}

// Simplify Simplify a collection of TaggedLineStrings
func (t *TaggedLinesSimplifier) Simplify(taggedLines []*TaggedLineString) {
	for _, v := range taggedLines {
		t.inputIndex.Add(v.ParentLine)
	}
	for _, v := range taggedLines {
		tlss := &TaggedLineStringSimplifier{inputIndex: t.inputIndex,
			outputIndex:       t.outputIndex,
			distanceTolerance: t.distanceTolerance,
		}
		tlss.Simplify(v)
	}
}
