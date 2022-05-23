package matrix

// Filter  An interface  which use the values of the matrix in a  Steric.
//  matrix filters can be used to implement centroid and
//  envelope computation, and many other functions.
type Filter interface {
	// IsChanged  Returns the true when need change.
	IsChanged() bool

	// FilterMatrixes  Performs an operation with the provided .
	FilterMatrixes(matrixes []Matrix)

	// Filter  Performs an operation with the provided .
	Filter(matrix Matrix)

	// Matrixes Returns matrixes.
	Matrixes() []Matrix

	// Clear  clear matrixes.
	Clear()
}

// UniqueArrayFilter  A Filter that extracts a unique array.
type UniqueArrayFilter struct {
	matrixes    []Matrix
	IsNotChange bool
}

// IsChanged  Returns the true when need change.
func (u *UniqueArrayFilter) IsChanged() bool {
	return !u.IsNotChange
}

// Matrixes  Returns the gathered Matrixes.
func (u *UniqueArrayFilter) Matrixes() []Matrix {
	return u.matrixes
}

// Clear  Returns the gathered Matrixes.
func (u *UniqueArrayFilter) Clear() {
	u.matrixes = []Matrix{}
}

// Filter Performs an operation with the provided .
func (u *UniqueArrayFilter) Filter(matrix Matrix) {
	u.add(matrix)
}

// FilterMatrixes Performs an operation with the provided .
func (u *UniqueArrayFilter) FilterMatrixes(matrixes []Matrix) {
	for _, v := range matrixes {
		u.Filter(v)
	}

}

func (u *UniqueArrayFilter) add(matrix Matrix) {
	hasMatrix := false
	for _, v := range u.matrixes {
		if v.Equals(matrix) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.matrixes = append(u.matrixes, matrix)
	}
}

// compile time checks
var (
	_ Filter = &UniqueArrayFilter{}
)
