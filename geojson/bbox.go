package geojson

import (
	"github.com/spatial-go/geoos"
)

// BBox is for the geojson bbox attribute which is an array with all axes
// of the most southwesterly point followed by all axes of the more northeasterly point.
type BBox []float64

// NewBBox creates a bbox from a a bound.
func NewBBox(b geoos.Bound) BBox {
	return []float64{
		b.Min.X(), b.Min.Y(),
		b.Max.X(), b.Max.Y(),
	}
}

// Valid checks if the bbox is present and has at least 4 elements.
func (bb BBox) Valid() bool {
	if bb == nil {
		return false
	}

	return len(bb) >= 4 && len(bb)%2 == 0
}

// Bound returns the geoos.Bound for the BBox.
func (bb BBox) Bound() geoos.Bound {
	if !bb.Valid() {
		return geoos.Bound{}
	}
	mid := len(bb) / 2
	return geoos.Bound{
		Min: geoos.Point{bb[0], bb[1]},
		Max: geoos.Point{bb[mid], bb[mid+1]},
	}
}
