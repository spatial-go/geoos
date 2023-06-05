// Package matrix Define spatial matrix base.
package matrix

import (
	"reflect"
	"testing"
)

func TestMultiPolygonMatrix_Dimensions(t *testing.T) {
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Dimensions(); got != tt.want {
				t.Errorf("MultiPolygonMatrix.Dimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_BoundaryDimensions(t *testing.T) {
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.BoundaryDimensions(); got != tt.want {
				t.Errorf("MultiPolygonMatrix.BoundaryDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_Boundary(t *testing.T) {
	tests := []struct {
		name    string
		m       MultiPolygonMatrix
		want    Steric
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Boundary()
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiPolygonMatrix.Boundary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiPolygonMatrix.Boundary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_Nums(t *testing.T) {
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Nums(); got != tt.want {
				t.Errorf("MultiPolygonMatrix.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsEmpty(); got != tt.want {
				t.Errorf("MultiPolygonMatrix.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_Bound(t *testing.T) {
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		want Bound
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Bound(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiPolygonMatrix.Bound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_Equals(t *testing.T) {
	type args struct {
		ms Steric
	}
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.args.ms); got != tt.want {
				t.Errorf("MultiPolygonMatrix.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_Proximity(t *testing.T) {
	type args struct {
		ms Steric
	}
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Proximity(tt.args.ms); got != tt.want {
				t.Errorf("MultiPolygonMatrix.Proximity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_EqualsExact(t *testing.T) {
	type args struct {
		ms        Steric
		tolerance float64
	}
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EqualsExact(tt.args.ms, tt.args.tolerance); got != tt.want {
				t.Errorf("MultiPolygonMatrix.EqualsExact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonMatrix_Filter(t *testing.T) {
	tests := []struct {
		name string
		m    MultiPolygonMatrix
		want Steric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Filter(CreateFilterMatrix()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiPolygonMatrix.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
