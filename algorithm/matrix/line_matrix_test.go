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

func TestSteric_Dimensions(t *testing.T) {
	tests := []struct {
		name string
		l    Steric
		want int
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, 1},
		{"case 2", Matrix{1, 1}, 0},
		{"case 3", PolygonMatrix{{{1, 1}, {2, 2}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Dimensions(); got != tt.want {
				t.Errorf("LineMatrix.Dimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineMatrix_BoundaryDimensions(t *testing.T) {
	tests := []struct {
		name string
		l    LineMatrix
		want int
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, 0},
		{"case 2", LineMatrix{{1, 1}, {2, 2}, {2, 1}, {1, 1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.BoundaryDimensions(); got != tt.want {
				t.Errorf("LineMatrix.BoundaryDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineMatrix_Boundary(t *testing.T) {
	tests := []struct {
		name    string
		l       LineMatrix
		want    Steric
		wantErr bool
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, Collection{Matrix{1, 1}, Matrix{2, 2}}, false},
		{"case 2", LineMatrix{{1, 1}, {2, 2}, {2, 1}, {1, 1}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Boundary()
			if (err != nil) != tt.wantErr {
				t.Errorf("LineMatrix.Boundary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineMatrix.Boundary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineMatrix_IsClosed(t *testing.T) {
	tests := []struct {
		name string
		l    LineMatrix
		want bool
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, false},
		{"case 2", LineMatrix{{1, 1}, {2, 2}, {2, 1}, {1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsClosed(); got != tt.want {
				t.Errorf("LineMatrix.IsClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteric_Nums(t *testing.T) {
	tests := []struct {
		name string
		l    Steric
		want int
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, 1},
		{"case 2", Matrix{1, 1}, 1},
		{"case 3", PolygonMatrix{{{1, 1}, {2, 2}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Nums(); got != tt.want {
				t.Errorf("Steric.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteric_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		l    Steric
		want bool
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, false},
		{"case 2", Matrix{1, 1}, false},
		{"case 3", PolygonMatrix{{{1, 1}, {2, 2}}}, false},
		{"case 4", Matrix{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsEmpty(); got != tt.want {
				t.Errorf("Steric.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteric_Bound(t *testing.T) {
	tests := []struct {
		name string
		l    Steric
		want []Matrix
	}{
		{"line bound", LineMatrix{
			{0, 0}, {0, 5}, {5, 5}, {5, 0}, {0, 0},
		}, []Matrix{{0, 0}, {5, 5}},
		},
		{"case 2", Matrix{1, 1}, []Matrix{{1, 1}, {1, 1}}},
		{"case 3", PolygonMatrix{
			{{0, 0}, {0, 5}, {5, 5}, {5, 0}, {0, 0}}}, []Matrix{{0, 0}, {5, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Bound(); !got.Equals(tt.want) {
				t.Errorf("Steric.Bound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineMatrix_ToLineArray(t *testing.T) {
	tests := []struct {
		name      string
		l         LineMatrix
		wantLines []*LineSegment
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}}, []*LineSegment{{Matrix{1, 1}, Matrix{2, 2}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLines := tt.l.ToLineArray(); !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("LineMatrix.ToLineArray() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}

func TestSteric_Equals(t *testing.T) {
	type args struct {
		ms Steric
	}
	tests := []struct {
		name string
		l    Steric
		args args
		want bool
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}},
			args{LineMatrix{{1, 1}, {2, 2}}}, true},
		{"case 2", Matrix{1, 1},
			args{LineMatrix{{1, 1}, {2, 2}}}, false},
		{"case 3", PolygonMatrix{{{1, 1}, {2, 2}}},
			args{LineMatrix{{1, 1}, {2, 2}}}, false},
		{"case 4", Matrix{},
			args{LineMatrix{{1, 1}, {2, 2}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Equals(tt.args.ms); got != tt.want {
				t.Errorf("Steric.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteric_Proximity(t *testing.T) {
	type args struct {
		ms Steric
	}
	tests := []struct {
		name string
		l    Steric
		args args
		want bool
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}},
			args{LineMatrix{{1, 1}, {2, 2}}}, true},
		{"case 2", LineMatrix{{1, 1}, {2, 2}},
			args{LineMatrix{{1, 1}, {2, 2.00000000001}}}, true},
		{"case 3", PolygonMatrix{{{1, 1}, {2, 2}}},
			args{LineMatrix{{1, 1}, {2, 2}}}, false},
		{"case 4", Matrix{},
			args{LineMatrix{{1, 1}, {2, 2}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.Proximity(tt.args.ms)
			if got != tt.want {
				t.Errorf("Steric.Proximity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteric_EqualsExact(t *testing.T) {
	type args struct {
		ms        Steric
		tolerance float64
	}
	tests := []struct {
		name string
		l    Steric
		args args
		want bool
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}},
			args{LineMatrix{{1, 1}, {2, 2}}, 0.01}, true},
		{"case 2", LineMatrix{{1, 1}, {2, 2}},
			args{LineMatrix{{1, 1}, {2, 2.00000000001}}, 0.01}, true},
		{"case 3", PolygonMatrix{{{1, 1}, {2, 2}}},
			args{LineMatrix{{1, 1}, {2, 2}}, 0.01}, false},
		{"case 4", Matrix{},
			args{LineMatrix{{1, 1}, {2, 2}}, 0.01}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.EqualsExact(tt.args.ms, tt.args.tolerance); got != tt.want {
				t.Errorf("Steric.EqualsExact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteric_Filter(t *testing.T) {
	type args struct {
		f Filter
	}
	tests := []struct {
		name string
		l    Steric
		args args
		want Steric
	}{
		{"case 1", LineMatrix{{1, 1}, {2, 2}},
			args{&UniqueArrayFilter{}}, LineMatrix{{1, 1}, {2, 2}}},
		{"case 2", LineMatrix{{1, 1}, {2, 2}, {2, 2}},
			args{&UniqueArrayFilter{}}, LineMatrix{{1, 1}, {2, 2}}},
		{"case 3", PolygonMatrix{{{1, 1}, {2, 2}}},
			args{&UniqueArrayFilter{}}, PolygonMatrix{{{1, 1}, {2, 2}, {1, 1}}}},
		{"case 4", Matrix{},
			args{&UniqueArrayFilter{}}, Matrix{}},

		{"poly ", PolygonMatrix{{
			{0, 0}, {0, 5}, {5, 5}, {5, 5}, {5, 0}, {0, 0},
		}}, args{&UniqueArrayFilter{}},
			PolygonMatrix{{{0, 0}, {0, 5}, {5, 5}, {5, 0}, {0, 0}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Filter(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Steric.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
