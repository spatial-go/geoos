package buffer

import (
	"container/list"
	"sort"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// ConvexHullComputer Computes the convex hull of a Geometry.
// The convex hull is the smallest convex Geometry that contains all the
//  points in the input Geometry.
type ConvexHullComputer struct {
	inputPts []matrix.Matrix
}

// ConvexHull Returns a geometry that represents the convex hull of the input geometry.
// The returned geometry contains the minimal number of points needed to
// represent the convex hull.  In particular, no more than two consecutive
// points will be collinear.
func ConvexHull(geom matrix.Steric) matrix.Steric {
	return ConvexHullWithGeom(geom).ConvexHull()
}

// ConvexHullWithGeom Create a new convex hull construction for the input geometry.
func ConvexHullWithGeom(geom matrix.Steric) *ConvexHullComputer {
	ch := &ConvexHullComputer{}
	ch.inputPts = extractMatrixes(geom)
	return ch
}

func extractMatrixes(geom matrix.Steric) []matrix.Matrix {
	filter := matrix.UniqueArrayFilter{IsNotChange: true}
	_ = geom.Filter(&filter)
	return filter.Matrixes()
}

// ConvexHull Returns a geometry that represents the convex hull of the input geometry.
// The returned geometry contains the minimal number of points needed to
// represent the convex hull.  In particular, no more than two consecutive
// points will be collinear.
func (c *ConvexHullComputer) ConvexHull() matrix.Steric {
	reducedPts := c.inputPts
	switch length := len(c.inputPts); {
	case length == 0:
		return matrix.Collection{}
	case length == 1:
		return c.inputPts[0]
	case length == 2:
		return matrix.LineMatrix{c.inputPts[0], c.inputPts[1]}
	case length > 50:
		reducedPts = c.reduce()
	}
	// sort points for Graham scan.
	sortedPts := c.preSort(reducedPts)

	// Use Graham scan to find convex hull.
	cHS := c.grahamScan(sortedPts)

	// Convert stack to an array.
	cH := c.toMatrixArray(cHS)

	// Convert array to appropriate output geometry.
	return c.lineOrPolygon(cH)
}

// toMatrixArray An alternative to Stack.toArray, which is not present.
func (c *ConvexHullComputer) toMatrixArray(stack *list.List) []matrix.Matrix {
	ms := make([]matrix.Matrix, stack.Len())
	for i := 0; i < len(ms); i++ {
		el := stack.Front()
		ms[i] = el.Value.(matrix.Matrix)
		stack.Remove(el)
	}
	return ms
}

// reduce Uses a heuristic to reduce the number of points scanned to compute the hull.
func (c *ConvexHullComputer) reduce() []matrix.Matrix {
	polyPts := c.computeOctRing()

	// unable to compute interior polygon for some reason
	if polyPts == nil {
		return c.inputPts
	}
	ls := matrix.LineMatrix{}
	for _, v := range polyPts {
		ls = append(ls, v)
	}
	reducedSet := list.New()
	for _, v := range c.inputPts {
		if !relate.InLineMatrix(v, ls) {
			reducedSet.PushBack(v)
		}
	}
	reducedPts := c.toMatrixArray(reducedSet)

	// ensure that computed array has at least 3 points (not necessarily unique)
	if len(reducedPts) < 3 {
		return c.padArray3(reducedPts)
	}
	return reducedPts
}

func (c *ConvexHullComputer) padArray3(pts []matrix.Matrix) []matrix.Matrix {
	pad := make([]matrix.Matrix, 3)
	for i := 0; i < len(pad); i++ {
		if i < len(pad) {
			pad[i] = pts[i]
		} else {
			pad[i] = pts[0]
		}
	}
	return pad
}

func (c *ConvexHullComputer) preSort(pts []matrix.Matrix) []matrix.Matrix {

	// find the lowest point in the set. If two or more points have
	// the same minimum y matrix choose the one with the minimum x.
	// This focal point is put in array location pts[0].
	for i := 1; i < len(pts); i++ {
		if (pts[i][1] < pts[0][1]) || ((pts[i][1] == pts[0][1]) && (pts[i][0] < pts[0][0])) {
			t := pts[0]
			pts[0] = pts[i]
			pts[i] = t
		}
	}

	// sort the points radially around the focal point.
	rc := &RadialComparator{pts[0], pts}
	sort.Sort(rc)

	//radialSort(pts);
	return rc.pts
}

// grahamScan Uses the Graham Scan algorithm to compute the convex hull vertices.
func (c *ConvexHullComputer) grahamScan(ms []matrix.Matrix) *list.List {

	ps := list.New()
	ps.PushBack(ms[0])
	ps.PushBack(ms[1])
	ps.PushBack(ms[2])
	for i := 3; i < len(ms); i++ {
		el := ps.Back()
		p := el.Value.(matrix.Matrix)
		ps.Remove(el)
		// check for empty stack to guard against robustness problems
		for ps.Len() != 0 && OrientationIndex(ps.Back().Value.(matrix.Matrix), p, ms[i]) > 0 {
			el := ps.Back()
			p = el.Value.(matrix.Matrix)
			ps.Remove(el)
		}
		ps.PushBack(p)
		ps.PushBack(ms[i])
	}
	ps.PushBack(ms[0])
	return ps
}

// isBetween returns  whether the three matrixes are collinear and c2 lies between c1 and c3 inclusive
func (c *ConvexHullComputer) isBetween(c1, c2, c3 matrix.Matrix) bool {
	if OrientationIndex(c1, c2, c3) != 0 {
		return false
	}
	if c1[0] != c3[0] {
		if c1[0] <= c2[0] && c2[0] <= c3[0] {
			return true
		}
		if c3[0] <= c2[0] && c2[0] <= c1[0] {
			return true
		}
	}
	if c1[1] != c3[1] {
		if c1[1] <= c2[1] && c2[1] <= c3[1] {
			return true
		}
		if c3[1] <= c2[1] && c2[1] <= c1[1] {
			return true
		}
	}
	return false
}

func (c *ConvexHullComputer) computeOctRing() []matrix.Matrix {
	octPts := c.computeOctPts()
	// points must all lie in a line
	if len(octPts) < 3 {
		return nil
	}
	if len(octPts) > 0 {
		octPts = append(octPts, octPts[0])
	}
	return octPts
}

func (c *ConvexHullComputer) computeOctPts() []matrix.Matrix {
	pts := make([]matrix.Matrix, 8)
	for j := 0; j < 8; j++ {
		pts[j] = c.inputPts[0]
	}
	for i := 1; i < len(c.inputPts); i++ {
		if c.inputPts[i][0] < pts[0][0] {
			pts[0] = c.inputPts[i]
		}
		if c.inputPts[i][0]-c.inputPts[i][1] < pts[1][0]-pts[1][1] {
			pts[1] = c.inputPts[i]
		}
		if c.inputPts[i][1] > pts[2][1] {
			pts[2] = c.inputPts[i]
		}
		if c.inputPts[i][0]+c.inputPts[i][1] > pts[3][0]+pts[3][1] {
			pts[3] = c.inputPts[i]
		}
		if c.inputPts[i][0] > pts[4][0] {
			pts[4] = c.inputPts[i]
		}
		if c.inputPts[i][0]-c.inputPts[i][1] > pts[5][0]-pts[5][1] {
			pts[5] = c.inputPts[i]
		}
		if c.inputPts[i][1] < pts[6][1] {
			pts[6] = c.inputPts[i]
		}
		if c.inputPts[i][0]+c.inputPts[i][1] < pts[7][0]+pts[7][1] {
			pts[7] = c.inputPts[i]
		}
	}
	return pts

}

func (c *ConvexHullComputer) lineOrPolygon(ms []matrix.Matrix) matrix.Steric {

	cms := c.cleanRing(ms)
	if len(cms) == 3 {
		return matrix.LineMatrix{cms[0], cms[1]}
	}
	ls := matrix.LineMatrix{}
	for _, v := range cms {
		ls = append(ls, v)
	}
	if !matrix.Matrix(ls[0]).Equals(matrix.Matrix(ls[len(ls)-1])) {
		ls = append(ls, ls[0])
	}

	return matrix.PolygonMatrix{ls}
}

func (c *ConvexHullComputer) cleanRing(ms []matrix.Matrix) []matrix.Matrix {
	cleanedRing := []matrix.Matrix{}
	previousDistinctMatrix := matrix.Matrix{}
	for i := 0; i <= len(ms)-2; i++ {
		currentMatrix := ms[i]
		nextMatrix := ms[i+1]
		if currentMatrix.Equals(nextMatrix) {
			continue
		}
		if len(previousDistinctMatrix) > 1 &&
			c.isBetween(previousDistinctMatrix, currentMatrix, nextMatrix) {
			continue
		}
		cleanedRing = append(cleanedRing, currentMatrix)
		previousDistinctMatrix = currentMatrix
	}
	cleanedRing = append(cleanedRing, ms[len(ms)-1])
	return cleanedRing
}

// RadialComparator Compares  Matrixes for their angle and distance relative to an origin.
type RadialComparator struct {
	origin matrix.Matrix
	pts    []matrix.Matrix
}

// Len ...
func (r *RadialComparator) Len() int {
	return len(r.pts)
}

// Less ...
func (r *RadialComparator) Less(i, j int) bool {
	return r.polarCompare(r.origin, r.pts[i], r.pts[j])
}

// Swap ...
func (r RadialComparator) Swap(i, j int) {
	r.pts[i], r.pts[j] = r.pts[j], r.pts[i]

}

// polarCompare Given two points p and q compare them with respect to their radial
func (r *RadialComparator) polarCompare(o, p, q matrix.Matrix) bool {
	dxp := p[0] - o[0]
	dyp := p[1] - o[1]
	dxq := q[0] - o[0]
	dyq := q[1] - o[1]

	orient := OrientationIndex(o, p, q)

	if orient == calc.CounterClockWise {
		return false
	}
	if orient == calc.ClockWise {
		return true
	}
	// points are collinear - check distance
	op := dxp*dxp + dyp*dyp
	oq := dxq*dxq + dyq*dyq
	if op < oq {
		return true
	}
	if op > oq {
		return false
	}
	return true
}

// OrientationIndex Returns the index of the direction of the point q relative to
// a vector specified by p1-p2.
func OrientationIndex(p1, p2, q matrix.Matrix) int {

	// fast filter for orientation index
	// avoids use of slow extended-precision arithmetic in many cases
	index := orientationIndexFilter(p1[0], p1[1], p2[0], p2[1], q[0], q[1])
	if index <= 1 {
		return index
	}
	// normalize matrixes
	dx1 := calc.ValueOf(p2[0]).SelfAddOne(-p1[0])
	dy1 := calc.ValueOf(p2[1]).SelfAddOne(-p1[1])
	dx2 := calc.ValueOf(q[0]).SelfAddOne(-p2[0])
	dy2 := calc.ValueOf(q[1]).SelfAddOne(-p2[1])

	dyx := dy1.SelfMultiplyPair(dx2)
	// sign of determinant - unrolled for performance
	return dx1.SelfMultiplyPair(dy2).SelfSubtract(dyx.Hi, dyx.Lo).Signum()
}

// orientationIndexFilter A filter for computing the orientation index of three matrixes.
func orientationIndexFilter(pax, pay,
	pbx, pby, pcx, pcy float64) int {
	detSum := 0.0

	detLeft := (pax - pcx) * (pby - pcy)
	detRight := (pay - pcy) * (pbx - pcx)
	det := detLeft - detRight

	if detLeft > 0.0 {
		if detRight <= 0.0 {
			return signum(det)
		}
		detSum = detLeft + detRight
	} else if detLeft < 0.0 {
		if detRight >= 0.0 {
			return signum(det)
		}
		detSum = -detLeft - detRight
	} else {
		return signum(det)
	}

	errbound := SafeEpsilon * detSum
	if (det >= errbound) || (-det >= errbound) {
		return signum(det)
	}

	return 2
}

func signum(x float64) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
