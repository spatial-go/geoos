package relate

import (
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Computer Computes the topological relationship between two Geometries.
type Computer struct {
	Arg            []matrix.Steric // the arg(s) of the operation
	IntersectBound bool
}

func (r *Computer) computeIM(arg []matrix.Steric, intersectBound bool) *matrix.IntersectionMatrix {
	r.Arg = arg
	r.IntersectBound = intersectBound
	im := matrix.IntersectionMatrixDefault()
	// since Geometries are finite and embedded in a 2-D space, the EE element must always be 2
	im.Set(calc.ImExterior, calc.ImExterior, 2)
	// if the Geometries don't overlap there is nothing to do
	if !intersectBound {
		r.computeDisjointIM(im)
		return im
	}
	if arg[0].Equals(arg[1]) {
		switch arg[0].(type) {
		case matrix.Matrix:
			im.SetAtLeastString("0FFFFFFF2")
		case matrix.LineMatrix:
			im.SetAtLeastString("1FFF0FFF2")
		case matrix.PolygonMatrix:
			im.SetAtLeastString("2FFF1FFF2")
		}
		return im
	}
	switch r := arg[0].(type) {
	case matrix.Matrix:
		rp := &PointRelate{r, arg[1]}
		return rp.IntersectionMatrix(im)
	case matrix.LineMatrix:
		rp := &LineRelate{r, arg[1]}
		return rp.IntersectionMatrix(im)
	case matrix.PolygonMatrix:
		rp := &PolygonRelate{r, arg[1]}
		return rp.IntersectionMatrix(im)
	}

	r.computeProperIntersectionIM(im)
	return im
}

// computeDisjointIM If the Geometries are disjoint, we need to enter their dimension and
// boundary dimension in the Ext rows in the IM
func (r *Computer) computeDisjointIM(im *matrix.IntersectionMatrix) {
	ga := r.Arg[0]
	if !ga.IsEmpty() {
		im.Set(calc.ImInterior, calc.ImExterior, ga.Dimensions())
		im.Set(calc.ImBoundary, calc.ImExterior, ga.BoundaryDimensions())
	}
	gb := r.Arg[1]
	if !gb.IsEmpty() {
		im.Set(calc.ImExterior, calc.ImInterior, gb.Dimensions())
		im.Set(calc.ImExterior, calc.ImBoundary, gb.BoundaryDimensions())
	}
}

func (r *Computer) computeProperIntersectionIM(im *matrix.IntersectionMatrix) {
	// If a proper intersection is found, we can set a lower bound on the IM.
	dimA := r.Arg[0].Dimensions()
	dimB := r.Arg[1].Dimensions()

	// For Geometry's of dim 0 there can never be proper intersections.

	// If edge segments of Areas properly intersect, the areas must properly overlap.
	if dimA == 2 && dimB == 2 {
		im.SetAtLeastString("212101212")
	} else if dimA == 2 && dimB == 1 {
		//im.SetAtLeast("FFF0FFFF2")
		im.SetAtLeastString("1FFFFF1FF")
	} else if dimA == 1 && dimB == 2 {
		//im.SetAtLeast("F0FFFFFF2")
		im.SetAtLeastString("1F1FFFFFF")

	} else if dimA == 1 && dimB == 1 {
		im.SetAtLeastString("0FFFFFFFF")
	}
}

//TODO
//  updateIM update the IM with the sum of the IMs for each component
// func (r *Computer) updateIM(im *matrix.IntersectionMatrix) {
// 	//TODO
// }
