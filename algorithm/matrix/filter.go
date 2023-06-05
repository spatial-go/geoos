package matrix

import "github.com/spatial-go/geoos/algorithm/filter"

func CreateFilterMatrix() filter.Filter[Matrix] {
	f := &filter.UniqueArrayFilter[Matrix]{FilterFunc: func(param1, param2 any) bool {
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
func CreateFilterMatrixNotChanged() filter.Filter[Matrix] {
	f := &filter.UniqueArrayFilter[Matrix]{IsNotChange: true, FilterFunc: func(param1, param2 any) bool {
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
