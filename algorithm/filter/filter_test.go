package filter

import (
	"reflect"
	"testing"
)

func TestUniqueArrayFilter_FilterSteric(t *testing.T) {

	type args struct {
		matr []int
	}
	tests := []struct {
		name string
		want interface{}
		args args
	}{
		{name: "unique array filter",
			args: args{[]int{1, 2, 3, 4, 5, 4, 5}},
			want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UniqueArrayFilter[int]{FilterFunc: func(param1, param2 any) bool {
				if reflect.DeepEqual(param1, param2) {
					return true
				}
				return false
			}}
			entities := tt.args.matr
			for _, v := range entities {
				u.Filter(v)
			}
			if got := u.Entities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique array filter = %v, want %v", got, tt.want)
			}
		})
	}
}
