package matrix

import (
	"bytes"

	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/calc"
)

// IntersectionMatrix  a Dimensionally Extended Nine-Intersection Model (DE-9IM) matrix.
// DE-9IM matrix values (such as "212FF1FF2")
// specify the topological relationship between two Geometrys.
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
	im.SetAll(calc.FALSE)
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

// IsDisjoint Tests if this matrix matches <code>[FF*FF****]</code>.
func (im *IntersectionMatrix) IsDisjoint() bool {
	return im.matrix[calc.INTERIOR][calc.INTERIOR] == calc.FALSE &&
		im.matrix[calc.INTERIOR][calc.BOUNDARY] == calc.FALSE &&
		im.matrix[calc.BOUNDARY][calc.INTERIOR] == calc.FALSE &&
		im.matrix[calc.BOUNDARY][calc.BOUNDARY] == calc.FALSE
}

// IsIntersects Tests if isDisjoint returns false.
func (im *IntersectionMatrix) IsIntersects() bool {
	return !im.IsDisjoint()
}

// IsTouches Tests if this matrix matches
// <code>[FT*******]</code>, <code>[F**T*****]</code> or <code>[F***T****]</code>.
func (im *IntersectionMatrix) IsTouches(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if dimensionOfGeometryA > dimensionOfGeometryB {
		//no need to get transpose because pattern matrix is symmetrical
		return im.IsTouches(dimensionOfGeometryB, dimensionOfGeometryA)
	}
	if (dimensionOfGeometryA == calc.A && dimensionOfGeometryB == calc.A) ||
		(dimensionOfGeometryA == calc.L && dimensionOfGeometryB == calc.L) ||
		(dimensionOfGeometryA == calc.L && dimensionOfGeometryB == calc.A) ||
		(dimensionOfGeometryA == calc.P && dimensionOfGeometryB == calc.A) ||
		(dimensionOfGeometryA == calc.P && dimensionOfGeometryB == calc.L) {
		return im.matrix[calc.INTERIOR][calc.INTERIOR] == calc.FALSE &&
			(isTrue(im.matrix[calc.INTERIOR][calc.BOUNDARY]) ||
				isTrue(im.matrix[calc.BOUNDARY][calc.INTERIOR]) ||
				isTrue(im.matrix[calc.BOUNDARY][calc.BOUNDARY]))
	}
	return false
}

// IsCrosses Tests whether this geometry crosses the specified geometry.
func (im *IntersectionMatrix) IsCrosses(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if (dimensionOfGeometryA == calc.P && dimensionOfGeometryB == calc.L) ||
		(dimensionOfGeometryA == calc.P && dimensionOfGeometryB == calc.A) ||
		(dimensionOfGeometryA == calc.L && dimensionOfGeometryB == calc.A) {
		return isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) &&
			isTrue(im.matrix[calc.INTERIOR][calc.EXTERIOR])
	}
	if (dimensionOfGeometryA == calc.L && dimensionOfGeometryB == calc.P) ||
		(dimensionOfGeometryA == calc.A && dimensionOfGeometryB == calc.P) ||
		(dimensionOfGeometryA == calc.A && dimensionOfGeometryB == calc.L) {
		return isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) &&
			isTrue(im.matrix[calc.EXTERIOR][calc.INTERIOR])
	}
	if dimensionOfGeometryA == calc.L && dimensionOfGeometryB == calc.L {
		return im.matrix[calc.INTERIOR][calc.INTERIOR] == 0
	}
	return false
}

// IsWithin  Tests whether this matrix matches <code>[T*F**F***]</code>.
func (im *IntersectionMatrix) IsWithin() bool {
	return isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) &&
		im.matrix[calc.INTERIOR][calc.EXTERIOR] == calc.FALSE &&
		im.matrix[calc.BOUNDARY][calc.EXTERIOR] == calc.FALSE
}

// IsContains  Tests whether this matrix matches [T*****FF*[.
func (im *IntersectionMatrix) IsContains() bool {
	return isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) &&
		im.matrix[calc.EXTERIOR][calc.INTERIOR] == calc.FALSE &&
		im.matrix[calc.EXTERIOR][calc.BOUNDARY] == calc.FALSE
}

// IsCovers Tests if this matrix matches
//    <code>[T*****FF*]</code>
// or <code>[*T****FF*]</code>
// or <code>[***T**FF*]</code>
// or <code>[****T*FF*]</code>
func (im *IntersectionMatrix) IsCovers() bool {
	hasPointInCommon :=
		isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) ||
			isTrue(im.matrix[calc.INTERIOR][calc.BOUNDARY]) ||
			isTrue(im.matrix[calc.BOUNDARY][calc.INTERIOR]) ||
			isTrue(im.matrix[calc.BOUNDARY][calc.BOUNDARY])

	return hasPointInCommon &&
		im.matrix[calc.EXTERIOR][calc.INTERIOR] == calc.FALSE &&
		im.matrix[calc.EXTERIOR][calc.BOUNDARY] == calc.FALSE
}

// IsCoveredBy Tests if this matrix matches
//  <code>[T*F**F***]</code>
// or <code>[*TF**F***]</code>
// or <code>[**FT*F***]</code>
// or <code>[**F*TF***]</code>
func (im *IntersectionMatrix) IsCoveredBy() bool {
	hasPointInCommon := isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) ||
		isTrue(im.matrix[calc.INTERIOR][calc.BOUNDARY]) ||
		isTrue(im.matrix[calc.BOUNDARY][calc.INTERIOR]) ||
		isTrue(im.matrix[calc.BOUNDARY][calc.BOUNDARY])

	return hasPointInCommon &&
		im.matrix[calc.INTERIOR][calc.EXTERIOR] == calc.FALSE &&
		im.matrix[calc.BOUNDARY][calc.EXTERIOR] == calc.FALSE
}

// IsEquals Tests whether the argument dimensions are equal and
// this matrix matches the pattern <tt>[T*F**FFF*]</tt>.
func (im *IntersectionMatrix) IsEquals(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if dimensionOfGeometryA != dimensionOfGeometryB {
		return false
	}
	return isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) &&
		im.matrix[calc.INTERIOR][calc.EXTERIOR] == calc.FALSE &&
		im.matrix[calc.BOUNDARY][calc.EXTERIOR] == calc.FALSE &&
		im.matrix[calc.EXTERIOR][calc.INTERIOR] == calc.FALSE &&
		im.matrix[calc.EXTERIOR][calc.BOUNDARY] == calc.FALSE
}

// IsOverlaps Tests if this matrix matches
// <UL>
//    <LI><tt>[T*T***T**]</tt> (for two points or two surfaces)
//    <LI><tt>[1*T***T**]</tt> (for two curves)
// </UL>.
func (im *IntersectionMatrix) IsOverlaps(dimensionOfGeometryA, dimensionOfGeometryB int) bool {
	if (dimensionOfGeometryA == calc.P && dimensionOfGeometryB == calc.P) ||
		(dimensionOfGeometryA == calc.A && dimensionOfGeometryB == calc.A) {
		return isTrue(im.matrix[calc.INTERIOR][calc.INTERIOR]) &&
			isTrue(im.matrix[calc.INTERIOR][calc.EXTERIOR]) &&
			isTrue(im.matrix[calc.EXTERIOR][calc.INTERIOR])
	}
	if dimensionOfGeometryA == calc.L && dimensionOfGeometryB == calc.L {
		return im.matrix[calc.INTERIOR][calc.INTERIOR] == 1 &&
			isTrue(im.matrix[calc.INTERIOR][calc.EXTERIOR]) &&
			isTrue(im.matrix[calc.EXTERIOR][calc.INTERIOR])
	}
	return false
}

// Matches Tests whether this matrix matches the given matrix pattern.
func (im *IntersectionMatrix) Matches(pattern string) (bool, error) {
	if len(pattern) != 9 {
		return false, algoerr.ErrorShouldBeLength9(pattern)
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

// Set Changes the value of one of this IntersectionMatrixs  elements.
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

// SetAtLeast  Changes the specified element to <code>minimumDimensionValue</code> if the element is less.
func (im *IntersectionMatrix) SetAtLeast(row, column, minimumDimensionValue int) {
	if im.matrix[row][column] < minimumDimensionValue {
		im.matrix[row][column] = minimumDimensionValue
	}
}

// SetAtLeastIfValid If row &gt;= 0 and column &gt;= 0, changes the specified element to minimumDimensionValue
// if the element is less. Does nothing if row &lt;0 or column &lt; 0.
func (im *IntersectionMatrix) SetAtLeastIfValid(row, column, minimumDimensionValue int) {
	if row >= 0 && column >= 0 {
		im.SetAtLeast(row, column, minimumDimensionValue)
	}
}

// SetAtLeastString  For each element in this <code>IntersectionMatrix</code>, changes the
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
	case calc.FALSE:
		return calc.SYMFALSE, nil
	case calc.TRUE:
		return calc.SYMTRUE, nil
	case calc.DONTCARE:
		return calc.SYMDONTCARE, nil
	case calc.P:
		return calc.SYMP, nil
	case calc.L:
		return calc.SYML, nil
	case calc.A:
		return calc.SYMA, nil
	default:
		return byte('_'), algoerr.ErrorUnknownDimension(dimensionValue)
	}
}

// toDimensionValue Converts the dimension symbol to a dimension value, for example, '*' = DONTCARE
func toDimensionValue(dimensionSymbol byte) (int, error) {
	switch []byte(bytes.ToUpper([]byte{dimensionSymbol}))[0] {
	case calc.SYMFALSE:
		return calc.FALSE, nil
	case calc.SYMTRUE:
		return calc.TRUE, nil
	case calc.SYMDONTCARE:
		return calc.DONTCARE, nil
	case calc.SYMP:
		return calc.P, nil
	case calc.SYML:
		return calc.L, nil
	case calc.SYMA:
		return calc.A, nil
	default:
		return -1, algoerr.ErrorUnknownDimension(dimensionSymbol)
	}
}

// matches Tests if the dimension value satisfies the dimension symbol.
func matches(actualDimensionValue int, requiredDimensionSymbol byte) bool {
	if requiredDimensionSymbol == calc.SYMDONTCARE {
		return true
	}
	if requiredDimensionSymbol == calc.SYMTRUE && (actualDimensionValue >= 0 || actualDimensionValue == calc.TRUE) {
		return true
	}
	if requiredDimensionSymbol == calc.SYMFALSE && actualDimensionValue == calc.FALSE {
		return true
	}
	if requiredDimensionSymbol == calc.SYMP && actualDimensionValue == calc.P {
		return true
	}
	if requiredDimensionSymbol == calc.SYML && actualDimensionValue == calc.L {
		return true
	}
	if requiredDimensionSymbol == calc.SYMA && actualDimensionValue == calc.A {
		return true
	}
	return false
}

func isTrue(actualDimensionValue int) bool {
	if actualDimensionValue >= 0 || actualDimensionValue == calc.TRUE {
		return true
	}
	return false
}
