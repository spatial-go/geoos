// Package matrix Define spatial matrix base.
package matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLineMatrix_Reverse(t *testing.T) {
	tests := []struct {
		name string
		l    LineMatrix
		want LineMatrix
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, LineMatrix{{2, 2}, {1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.l, tt.want)
			if got := tt.l.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineMatrix.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
