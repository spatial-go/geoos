// Package matrix Define spatial matrix base.
package matrix

import (
	"reflect"
	"testing"
)

func TestCollection_Dimensions(t *testing.T) {
	tests := []struct {
		name string
		c    Collection
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Dimensions(); got != tt.want {
				t.Errorf("Collection.Dimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_BoundaryDimensions(t *testing.T) {
	tests := []struct {
		name string
		c    Collection
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.BoundaryDimensions(); got != tt.want {
				t.Errorf("Collection.BoundaryDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Boundary(t *testing.T) {
	tests := []struct {
		name    string
		c       Collection
		want    Steric
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Boundary()
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Boundary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collection.Boundary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Nums(t *testing.T) {
	tests := []struct {
		name string
		c    Collection
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Nums(); got != tt.want {
				t.Errorf("Collection.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		c    Collection
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsEmpty(); got != tt.want {
				t.Errorf("Collection.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Bound(t *testing.T) {
	tests := []struct {
		name string
		c    Collection
		want Bound
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Bound(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collection.Bound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Equals(t *testing.T) {
	type args struct {
		ms Steric
	}
	tests := []struct {
		name string
		c    Collection
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Equals(tt.args.ms); got != tt.want {
				t.Errorf("Collection.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Proximity(t *testing.T) {
	type args struct {
		ms Steric
	}
	tests := []struct {
		name string
		c    Collection
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Proximity(tt.args.ms); got != tt.want {
				t.Errorf("Collection.Proximity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_EqualsExact(t *testing.T) {
	type args struct {
		ms        Steric
		tolerance float64
	}
	tests := []struct {
		name string
		c    Collection
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.EqualsExact(tt.args.ms, tt.args.tolerance); got != tt.want {
				t.Errorf("Collection.EqualsExact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Filter(t *testing.T) {
	type args struct {
		f Filter
	}
	tests := []struct {
		name string
		c    Collection
		args args
		want Steric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Filter(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collection.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
