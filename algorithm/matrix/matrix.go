package matrix

import "errors"

// Matrix is a one-dimensional matrix.
type Matrix []float64

// LineMatrix is a two-dimensional matrix.
type LineMatrix [][]float64

// PolygonMatrix is a three-dimensional matrix.
type PolygonMatrix [][][]float64

// MultiPolygonMatrix is a four-dimensional matrix.
type MultiPolygonMatrix [][][][]float64

// Compare returns 0 if m1==m2,1 if positive ,-1 else
// Compares Coordinate for order.
func Compare(m1, m2 Matrix) (int, error) {
	// If one is nil, the other must also be nil.
	if (m1 == nil) != (m2 == nil) {
		return -2, errors.New("Matrix is nil")
	}

	if m1[0] < m2[0] {
		return -1, nil
	}
	if m1[0] > m2[0] {
		return 1, nil
	}
	if m1[1] < m2[1] {
		return -1, nil
	}
	if m1[1] > m2[1] {
		return 1, nil
	}
	return 0, nil
}

// Equal returns  true if the two Matrix are equal
func Equal(m1, m2 Matrix) bool {
	// If one is nil, the other must also be nil.
	if (m1 == nil) != (m2 == nil) {
		return false
	}

	if len(m1) != len(m2) {
		return false
	}

	for i := range m1 {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}

// EqualLine returns  true if the two LineMatrix are equal
func EqualLine(m1, m2 LineMatrix) bool {
	// If one is nil, the other must also be nil.
	if (m1 == nil) != (m2 == nil) {
		return false
	}

	if len(m1) != len(m2) {
		return false
	}

	for i := range m1 {
		if !Equal(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// EqualPolygon returns  true if the two PolygonMatrix are equal
func EqualPolygon(m1, m2 PolygonMatrix) bool {
	// If one is nil, the other must also be nil.
	if (m1 == nil) != (m2 == nil) {
		return false
	}

	if len(m1) != len(m2) {
		return false
	}

	for i := range m1 {
		if !EqualLine(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// EqualMultiPolygon returns  true if the two MultiPolygonMatrix are equal
func EqualMultiPolygon(m1, m2 MultiPolygonMatrix) bool {
	// If one is nil, the other must also be nil.
	if (m1 == nil) != (m2 == nil) {
		return false
	}

	if len(m1) != len(m2) {
		return false
	}

	for i := range m1 {
		if !EqualPolygon(m1[i], m2[i]) {
			return false
		}
	}
	return true
}
