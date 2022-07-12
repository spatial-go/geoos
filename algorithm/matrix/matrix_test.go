package matrix

import (
	"reflect"
	"testing"
)

func TestBound_Equals(t *testing.T) {
	type args struct {
		g Bound
	}
	tests := []struct {
		name string
		b    Bound
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Equals(tt.args.g); got != tt.want {
				t.Errorf("Bound.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBound_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		b    Bound
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.IsEmpty(); got != tt.want {
				t.Errorf("Bound.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBound_ToRing(t *testing.T) {
	tests := []struct {
		name string
		b    Bound
		want LineMatrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ToRing(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bound.ToRing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBound_ToPolygon(t *testing.T) {
	tests := []struct {
		name string
		b    Bound
		want PolygonMatrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ToPolygon(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bound.ToPolygon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBound_Contains(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		b    Bound
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Contains(tt.args.m); got != tt.want {
				t.Errorf("Bound.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBound_ContainsBound(t *testing.T) {
	type args struct {
		bound Bound
	}
	tests := []struct {
		name string
		b    Bound
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ContainsBound(tt.args.bound); got != tt.want {
				t.Errorf("Bound.ContainsBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBound_IntersectsBound(t *testing.T) {
	type args struct {
		other Bound
	}
	tests := []struct {
		name string
		b    Bound
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.IntersectsBound(tt.args.other); got != tt.want {
				t.Errorf("Bound.IntersectsBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Compare(t *testing.T) {
	type args struct {
		m1 Matrix
	}
	tests := []struct {
		name    string
		m       Matrix
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Compare(tt.args.m1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Matrix.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransMatrixes(t *testing.T) {
	type args struct {
		inputGeom Steric
	}
	tests := []struct {
		name string
		args args
		want []Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransMatrixes(tt.args.inputGeom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransMatrixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineArray(t *testing.T) {
	type args struct {
		l LineMatrix
	}
	tests := []struct {
		name      string
		args      args
		wantLines []*LineSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLines := LineArray(tt.args.l); !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("LineArray() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}
