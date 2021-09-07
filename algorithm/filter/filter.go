// Package filter Define  data filter function.
package filter

import "reflect"

// Filter  An interface  which use the values of the entity in a  entities.
type Filter interface {
	// Filter  Performs an operation with the provided .
	Filter(entity interface{})

	// Matrixes ...
	Entities() interface{}
}

// UniqueArrayFilter  A Filter that extracts a unique array.
type UniqueArrayFilter struct {
	entities []interface{}
}

// Entities  Returns the gathered Matrixes.
func (u *UniqueArrayFilter) Entities() interface{} {
	return u.entities
}

// Filter Performs an operation with the provided .
func (u *UniqueArrayFilter) Filter(entity interface{}) {
	u.add(entity)
}

func (u *UniqueArrayFilter) add(entity interface{}) {
	hasMatrix := false
	for _, v := range u.entities {
		if reflect.DeepEqual(v, entity) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.entities = append(u.entities, entity)
	}
}

// compile time checks
var (
	_ Filter = &UniqueArrayFilter{}
)
