// Package matrix Define spatial matrix base.
package matrix

import (
	"reflect"
	"testing"
)

func TestPolygonMatrix_Bound(t *testing.T) {
	tests := []struct {
		name string
		p    PolygonMatrix
		want Bound
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Bound(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PolygonMatrix.Bound() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPolygonMatrix_IsRectangle(t *testing.T) {
	tests := []struct {
		name string
		p    PolygonMatrix
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsRectangle(); got != tt.want {
				t.Errorf("PolygonMatrix.IsRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
