// Package matrix Define spatial matrix base.
package matrix

import (
	"math"

	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/calc"
)

// A Collection is a collection of sterices that is also a Steric.
type Collection []Steric

func CollectionFromMultiLineMatrix(ml []LineMatrix) Collection {
	coll := make(Collection, len(ml))
	for i, v := range ml {
		coll[i] = v
	}
	return coll
}

// Dimensions returns the max of the dimensions of the collection.
func (c Collection) Dimensions() int {
	max := -1
	for _, g := range c {
		if d := g.Dimensions(); d > max {
			max = d
		}
	}
	return max
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (c Collection) BoundaryDimensions() int {
	dimension := calc.ImFalse
	for _, g := range c {
		if g.BoundaryDimensions() > dimension {
			dimension = g.BoundaryDimensions()
		}
	}
	return dimension
}

// Boundary returns the closure of the combinatorial boundary of this Collection.
func (c Collection) Boundary() (Steric, error) {
	return nil, algorithm.ErrNotSupportCollection
}

// Nums ...
func (c Collection) Nums() int {
	return len(c)
}

// IsEmpty returns true if the Matrix is empty.
func (c Collection) IsEmpty() bool {
	return len(c) == 0
}

// Bound returns a bound around the Collection.
func (c Collection) Bound() Bound {
	if len(c) == 0 {
		return []Matrix{}
	}
	b := c[0].Bound()
	for i := 1; i < len(c); i++ {
		bound := c[1].Bound()
		b[0][0] = math.Min(b[0][0], bound[0][0])
		b[0][1] = math.Min(b[0][1], bound[0][1])
		b[1][0] = math.Min(b[1][0], bound[1][0])
		b[1][1] = math.Min(b[1][1], bound[1][1])
	}

	return b
}

// Equals returns  true if the two Collection are equal
func (c Collection) Equals(ms Steric) bool {
	if mm, ok := ms.(Collection); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (c == nil) {
			return false
		}

		if len(mm) != len(c) {
			return false
		}

		for i := range mm {
			if !c[i].Equals(mm[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// Proximity returns true if the Steric represents the Proximity Geometry or vector.
func (c Collection) Proximity(ms Steric) bool {
	if mm, ok := ms.(Collection); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (c == nil) {
			return false
		}

		if len(mm) != len(c) {
			return false
		}

		for i := range mm {
			if !c[i].Equals(mm[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// EqualsExact returns  true if the two Collection are equalexact
func (c Collection) EqualsExact(ms Steric, tolerance float64) bool {
	if mm, ok := ms.(Collection); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (c == nil) {
			return false
		}

		if len(mm) != len(c) {
			return false
		}

		for i := range mm {
			if !c[i].EqualsExact(mm[i], tolerance) {
				return false
			}
		}
		return true
	}
	return false
}

// Filter Performs an operation with the provided .
func (c Collection) Filter(f Filter) Steric {
	if f.IsChanged() {
		mc := c[:0]
		for _, v := range c {
			g := v.Filter(f)
			mc = append(mc, g)
		}
		return mc
	}
	for _, v := range c {
		_ = v.Filter(f)
	}
	return c
}
