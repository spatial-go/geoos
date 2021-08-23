package quadtree

import "math"

// MinBinaryExponent ...
const MinBinaryExponent = -50

// IntervalSize Provides a test for whether an interval is
//  so small it should be considered as zero for the purposes of
//  inserting it into a binary tree.
type IntervalSize struct {
}

// IsZeroWidth Computes whether the interval [min, max] is effectively zero width.
//  I.e. the width of the interval is so much less than the
//  location of the interval that the midpoint of the interval cannot be
//  represented precisely.
func IsZeroWidth(min, max float64) bool {
	width := max - min
	if width == 0.0 {
		return true
	}
	maxAbs := math.Max(math.Abs(min), math.Abs(max))
	scaledInterval := width / maxAbs
	level, _ := math.Frexp(scaledInterval)
	level--
	return level <= MinBinaryExponent
}
