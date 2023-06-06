// Package filter Define  data filter function.
package filter

// ShieldFunc Define a filter function.
type ShieldFunc func(param1, param2 any) bool

// Filter  An interface  which use the values of the entity in a  entities.
type Filter[T any] interface {

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
	ShieldFunc
}

// Entities  Returns the gathered Matrixes.
func (u *UniqueArrayFilter[T]) Entities() []T {
	return u.entities
}

// Filter Performs an operation with the provided .
func (u *UniqueArrayFilter[T]) Filter(entity T) bool {
	return u.add(entity)
}

// FilterEntities Performs an operation with the provided .
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
		if u.ShieldFunc(v, entity) {
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
