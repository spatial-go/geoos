package geojson

import "github.com/spatial-go/geoos/space"

// BBox is for the geojson bbox attribute which is an array with all axes
// of the most southwesterly point followed by all axes of the more northeasterly point.
type BBox []float64

// NewBBox creates a bbox from a a bound.
func NewBBox(b space.Bound) BBox {
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

// Bound returns the space.Bound for the BBox.
func (bb BBox) Bound() space.Bound {
	if !bb.Valid() {
		return space.Bound{}
	}
	mid := len(bb) / 2
	return space.Bound{
		Min: space.Point{bb[0], bb[1]},
		Max: space.Point{bb[mid], bb[mid+1]},
	}
}
