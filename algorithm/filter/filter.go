// Package filter Define  data filter function.
package filter

type FilterFunc func(param1, param2 any) bool

// Filter  An interface  which use the values of the entity in a  entities.
type Filter[T any] interface {

	// IsChanged  Returns the true when need change.
	IsChanged() bool

	// Filter  Performs an operation with the provided .
	Filter(entity T) bool

	// FilterEntities  Performs an operation with the provided .
	FilterEntities(entities []T)

	// Entities ...
	Entities() []T

	// Clear  clear entities.
	Clear()
}

// UniqueArrayFilter  A Filter that extracts a unique array.
type UniqueArrayFilter[T any] struct {
	entities []T
	FilterFunc
	IsNotChange bool
}

// IsChanged  Returns the true when need change.
func (u *UniqueArrayFilter[T]) IsChanged() bool {
	return !u.IsNotChange
}

// Entities  Returns the gathered Matrixes.
func (u *UniqueArrayFilter[T]) Entities() []T {
	return u.entities
}

// Filter Performs an operation with the provided .
func (u *UniqueArrayFilter[T]) Filter(entity T) bool {
	return u.add(entity)
}

// FilterMatrixes Performs an operation with the provided .
func (u *UniqueArrayFilter[T]) FilterEntities(es []T) {
	for _, v := range es {
		u.Filter(v)
	}

}

// Clear  Returns the gathered Matrixes.
func (u *UniqueArrayFilter[T]) Clear() {
	u.entities = []T{}
}

func (u *UniqueArrayFilter[T]) add(entity T) bool {
	hasMatrix := false
	for _, v := range u.entities {
		if u.FilterFunc(v, entity) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.entities = append(u.entities, entity)
		return true
	}
	return false
}

// compile time checks
var (
	_ Filter[interface{}] = &UniqueArrayFilter[interface{}]{}
)
