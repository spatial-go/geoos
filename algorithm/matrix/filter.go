package matrix

import "github.com/spatial-go/geoos/algorithm/filter"

// CreateFilterMatrix create a filter by matrix.
func CreateFilterMatrix() filter.Filter[Matrix] {
	f := &filter.UniqueArrayFilter[Matrix]{ShieldFunc: func(param1, param2 any) bool {
		if v1, ok := param1.(Matrix); ok {
			if v2, ok := param2.(Matrix); ok {
				if v1.Equals(v2) {
					return true
				}
			}
		}
		return false
	}}
	return f
}
