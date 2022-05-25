package matrix

import (
	"bytes"

	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/calc"
)

// IntersectionMatrix  a Dimensionally Extended Nine-Intersection Model (DE-9IM) matrix.
// DE-9IM matrix values (such as "212FF1FF2")
// specify the topological relationship between two Geometries.
// DE-9IM matrices are 3x3 matrices with integer entries.
// The matrix indices {0,1,2} represent the topological locations
// that occur in a geometry (Interior, Boundary, Exterior).
// For a description of the DE-9IM and the spatial predicates derived from it,
// see the following references:
// <a href="http://www.opengis.org/techno/specs.htm"> OGC 99-049 OpenGIS Simple Features Specification for SQL</a>
// <a href="http://portal.opengeospatial.org/files/?artifact_id=25355">
// OGC 06-103r4 OpenGIS Implementation Standard for Geographic information - Simple feature access - Part 1: Common architecture</a>
type IntersectionMatrix struct {
	matrix [][]int
}

// IntersectionMatrixDefault Creates an IntersectionMatrix with  FALSE  dimension values.
func IntersectionMatrixDefault() *IntersectionMatrix {
	im := &IntersectionMatrix{}
	im.matrix = make([][]int, 3)
	for i := range im.matrix {
		im.matrix[i] = make([]int, 3)
	}
	im.SetAll(calc.ImFalse)
	return im
}

// SetAll Changes the elements of this IntersectionMatrix to dimensionValue
func (im *IntersectionMatrix) SetAll(dimensionValue int) {
	for ai := 0; ai < 3; ai++ {
		for bi := 0; bi < 3; bi++ {
			im.matrix[ai][bi] = dimensionValue
		}
	}
}

// IsDisjoint Tests if this matrix matches [FF*FF****].
func (im *IntersectionMatrix) IsDisjoint() bool {
	return im.matrix[calc.ImInterior][calc.ImInterior] == calc.ImFalse &&
		im.matrix[calc.ImInterior][calc.ImBoundary] == calc.ImFalse &&
		im.matrix[calc.ImBoundary][calc.ImInterior] == calc.ImFalse &&
		im.matrix[calc.ImBoundary][calc.ImBoundary] == calc.ImFalse
}

// IsIntersects Tests if isDisjoint returns false.
func (im *IntersectionMatrix) IsIntersects() bool {
	return !im.IsDisjoint()
}

// IsTouches Tests if this matrix matches
// [FT*******], [F**T*****] or [F***T****].
func (im *IntersectionMatrix) IsTouches(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if dimensionOfGeometryA > dimensionOfGeometryB {
		//no need to get transpose because pattern matrix is symmetrical
		return im.IsTouches(dimensionOfGeometryB, dimensionOfGeometryA)
	}
	if (dimensionOfGeometryA == calc.ImA && dimensionOfGeometryB == calc.ImA) ||
		(dimensionOfGeometryA == calc.ImL && dimensionOfGeometryB == calc.ImL) ||
		(dimensionOfGeometryA == calc.ImL && dimensionOfGeometryB == calc.ImA) ||
		(dimensionOfGeometryA == calc.ImP && dimensionOfGeometryB == calc.ImA) ||
		(dimensionOfGeometryA == calc.ImP && dimensionOfGeometryB == calc.ImL) {
		return im.matrix[calc.ImInterior][calc.ImInterior] == calc.ImFalse &&
			(isTrue(im.matrix[calc.ImInterior][calc.ImBoundary]) ||
				isTrue(im.matrix[calc.ImBoundary][calc.ImInterior]) ||
				isTrue(im.matrix[calc.ImBoundary][calc.ImBoundary]))
	}
	return false
}

// IsCrosses Tests whether this geometry crosses the specified geometry.
func (im *IntersectionMatrix) IsCrosses(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if (dimensionOfGeometryA == calc.ImP && dimensionOfGeometryB == calc.ImL) ||
		(dimensionOfGeometryA == calc.ImP && dimensionOfGeometryB == calc.ImA) ||
		(dimensionOfGeometryA == calc.ImL && dimensionOfGeometryB == calc.ImA) {
		return isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) &&
			isTrue(im.matrix[calc.ImInterior][calc.ImExterior])
	}
	if (dimensionOfGeometryA == calc.ImL && dimensionOfGeometryB == calc.ImP) ||
		(dimensionOfGeometryA == calc.ImA && dimensionOfGeometryB == calc.ImP) ||
		(dimensionOfGeometryA == calc.ImA && dimensionOfGeometryB == calc.ImL) {
		return isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) &&
			isTrue(im.matrix[calc.ImExterior][calc.ImInterior])
	}
	if dimensionOfGeometryA == calc.ImL && dimensionOfGeometryB == calc.ImL {
		return im.matrix[calc.ImInterior][calc.ImInterior] == 0
	}
	return false
}

// IsWithin  Tests whether this matrix matches [T*F**F***].
func (im *IntersectionMatrix) IsWithin() bool {
	return isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) &&
		im.matrix[calc.ImInterior][calc.ImExterior] == calc.ImFalse &&
		im.matrix[calc.ImBoundary][calc.ImExterior] == calc.ImFalse
}

// IsContains  Tests whether this matrix matches [T*****FF*[.
func (im *IntersectionMatrix) IsContains() bool {
	return isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) &&
		im.matrix[calc.ImExterior][calc.ImInterior] == calc.ImFalse &&
		im.matrix[calc.ImExterior][calc.ImBoundary] == calc.ImFalse
}

// IsCovers Tests if this matrix matches
//    [T*****FF*]
// or [*T****FF*]
// or [***T**FF*]
// or [****T*FF*]
func (im *IntersectionMatrix) IsCovers() bool {
	hasPointInCommon :=
		isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) ||
			isTrue(im.matrix[calc.ImInterior][calc.ImBoundary]) ||
			isTrue(im.matrix[calc.ImBoundary][calc.ImInterior]) ||
			isTrue(im.matrix[calc.ImBoundary][calc.ImBoundary])

	return hasPointInCommon &&
		im.matrix[calc.ImExterior][calc.ImInterior] == calc.ImFalse &&
		im.matrix[calc.ImExterior][calc.ImBoundary] == calc.ImFalse
}

// IsCoveredBy Tests if this matrix matches
//  [T*F**F***]
// or [*TF**F***]
// or [**FT*F***]
// or [**F*TF***]
func (im *IntersectionMatrix) IsCoveredBy() bool {
	hasPointInCommon := isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) ||
		isTrue(im.matrix[calc.ImInterior][calc.ImBoundary]) ||
		isTrue(im.matrix[calc.ImBoundary][calc.ImInterior]) ||
		isTrue(im.matrix[calc.ImBoundary][calc.ImBoundary])

	return hasPointInCommon &&
		im.matrix[calc.ImInterior][calc.ImExterior] == calc.ImFalse &&
		im.matrix[calc.ImBoundary][calc.ImExterior] == calc.ImFalse
}

// IsEquals Tests whether the argument dimensions are equal and
// this matrix matches the pattern <tt>[T*F**FFF*]</tt>.
func (im *IntersectionMatrix) IsEquals(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if dimensionOfGeometryA != dimensionOfGeometryB {
		return false
	}
	return isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) &&
		im.matrix[calc.ImInterior][calc.ImExterior] == calc.ImFalse &&
		im.matrix[calc.ImBoundary][calc.ImExterior] == calc.ImFalse &&
		im.matrix[calc.ImExterior][calc.ImInterior] == calc.ImFalse &&
		im.matrix[calc.ImExterior][calc.ImBoundary] == calc.ImFalse
}

// IsOverlaps Tests if this matrix matches
//   [T*T***T**](for two points or two surfaces)
//   [1*T***T**] (for two curves)
func (im *IntersectionMatrix) IsOverlaps(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if (dimensionOfGeometryA == calc.ImP && dimensionOfGeometryB == calc.ImP) ||
		(dimensionOfGeometryA == calc.ImA && dimensionOfGeometryB == calc.ImA) {
		return isTrue(im.matrix[calc.ImInterior][calc.ImInterior]) &&
			isTrue(im.matrix[calc.ImInterior][calc.ImExterior]) &&
			isTrue(im.matrix[calc.ImExterior][calc.ImInterior])
	}
	if dimensionOfGeometryA == calc.ImL && dimensionOfGeometryB == calc.ImL {
		return im.matrix[calc.ImInterior][calc.ImInterior] == 1 &&
			isTrue(im.matrix[calc.ImInterior][calc.ImExterior]) &&
			isTrue(im.matrix[calc.ImExterior][calc.ImInterior])
	}
	return false
}

// Matches Tests whether this matrix matches the given matrix pattern.
func (im *IntersectionMatrix) Matches(pattern string) (bool, error) {
	if len(pattern) != 9 {
		return false, algorithm.ErrorShouldBeLength9(pattern)
	}
	for ai := 0; ai < 3; ai++ {
		for bi := 0; bi < 3; bi++ {
			if !matches(im.matrix[ai][bi], pattern[3*ai+bi]) {
				return false, nil
			}
		}
	}
	return true, nil
}

// ToString Returns a nine-character String representation of this IntersectionMatrix
func (im *IntersectionMatrix) ToString() string {
	strByte := make([]byte, 9)
	for ai := 0; ai < 3; ai++ {
		for bi := 0; bi < 3; bi++ {
			strByte[3*ai+bi], _ = toDimensionSymbol(im.matrix[ai][bi])
		}
	}
	return string(strByte)
}

// Set Changes the value of one of this IntersectionMatrixes  elements.
func (im *IntersectionMatrix) Set(row, column, dimensionValue int) {
	im.matrix[row][column] = dimensionValue
}

// SetString Changes the elements of this IntersectionMatrix to the dimension symbols in dimensionSymbols.
func (im *IntersectionMatrix) SetString(dimensionSymbols string) {
	for i := 0; i < len(dimensionSymbols); i++ {
		row, col := i/3, i%3
		dv, _ := toDimensionValue(dimensionSymbols[i])
		im.matrix[row][col] = dv
	}
}

// SetAtLeast  Changes the specified element to minimumDimensionValue if the element is less.
func (im *IntersectionMatrix) SetAtLeast(row, column, minimumDimensionValue int) {
	if im.matrix[row][column] < minimumDimensionValue {
		im.matrix[row][column] = minimumDimensionValue
	}
}

// SetAtLeastIfValid If row >= 0 and column >= 0, changes the specified element to minimumDimensionValue
// if the element is less. Does nothing if row <0 or column < 0.
func (im *IntersectionMatrix) SetAtLeastIfValid(row, column, minimumDimensionValue int) {
	if row >= 0 && column >= 0 {
		im.SetAtLeast(row, column, minimumDimensionValue)
	}
}

// SetAtLeastString  For each element in this IntersectionMatrix, changes the
//  element to the corresponding minimum dimension symbol if the element is less.
func (im *IntersectionMatrix) SetAtLeastString(minimumDimensionSymbols string) {
	for i := 0; i < len(minimumDimensionSymbols); i++ {
		row, col := i/3, i%3
		dv, _ := toDimensionValue(minimumDimensionSymbols[i])
		im.SetAtLeast(row, col, dv)
	}
}

// Transpose  this IntersectionMatrix.
func (im *IntersectionMatrix) Transpose() *IntersectionMatrix {
	temp := im.matrix[1][0]
	im.matrix[1][0] = im.matrix[0][1]
	im.matrix[0][1] = temp
	temp = im.matrix[2][0]
	im.matrix[2][0] = im.matrix[0][2]
	im.matrix[0][2] = temp
	temp = im.matrix[2][1]
	im.matrix[2][1] = im.matrix[1][2]
	im.matrix[1][2] = temp
	return im
}

//  toDimensionSymbol Converts the dimension value to a dimension symbol, for example, TRUE = 'T'
func toDimensionSymbol(dimensionValue int) (byte, error) {
	switch dimensionValue {
	case calc.ImFalse:
		return calc.ImSymFalse, nil
	case calc.ImTrue:
		return calc.ImSymTrue, nil
	case calc.ImNotCare:
		return calc.ImSymNotCare, nil
	case calc.ImP:
		return calc.ImSymP, nil
	case calc.ImL:
		return calc.ImSymL, nil
	case calc.ImA:
		return calc.ImSymA, nil
	default:
		return byte('_'), algorithm.ErrorUnknownDimension(dimensionValue)
	}
}

// toDimensionValue Converts the dimension symbol to a dimension value, for example, '*' = NotCare
func toDimensionValue(dimensionSymbol byte) (int, error) {
	switch []byte(bytes.ToUpper([]byte{dimensionSymbol}))[0] {
	case calc.ImSymFalse:
		return calc.ImFalse, nil
	case calc.ImSymTrue:
		return calc.ImTrue, nil
	case calc.ImSymNotCare:
		return calc.ImNotCare, nil
	case calc.ImSymP:
		return calc.ImP, nil
	case calc.ImSymL:
		return calc.ImL, nil
	case calc.ImSymA:
		return calc.ImA, nil
	default:
		return -1, algorithm.ErrorUnknownDimension(dimensionSymbol)
	}
}

// matches Tests if the dimension value satisfies the dimension symbol.
func matches(actualDimensionValue int, requiredDimensionSymbol byte) bool {
	if requiredDimensionSymbol == calc.ImSymNotCare {
		return true
	}
	if requiredDimensionSymbol == calc.ImSymTrue && (actualDimensionValue >= 0 || actualDimensionValue == calc.ImTrue) {
		return true
	}
	if requiredDimensionSymbol == calc.ImSymFalse && actualDimensionValue == calc.ImFalse {
		return true
	}
	if requiredDimensionSymbol == calc.ImSymP && actualDimensionValue == calc.ImP {
		return true
	}
	if requiredDimensionSymbol == calc.ImSymL && actualDimensionValue == calc.ImL {
		return true
	}
	if requiredDimensionSymbol == calc.ImSymA && actualDimensionValue == calc.ImA {
		return true
	}
	return false
}

func isTrue(actualDimensionValue int) bool {
	if actualDimensionValue >= 0 || actualDimensionValue == calc.ImTrue {
		return true
	}
	return false
}
