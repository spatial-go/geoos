package envelope

import (
	"math"
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestEnvelope_Area(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"area", fields{200, 100, 200, 100}, 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Area(); got != tt.want {
				t.Errorf("Envelope.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Diameter(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Diameter", fields{200, 100, 200, 100}, math.Sqrt(20000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Diameter(); got != tt.want {
				t.Errorf("Envelope.Diameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_HashCode(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"HashCode", fields{1, 1, 2, 2}, 31912835},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.HashCode(); got != tt.want {
				t.Errorf("Envelope.HashCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIntersects(t *testing.T) {
	type args struct {
		p1 matrix.Matrix
		p2 matrix.Matrix
		q  matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"IsIntersects", args{matrix.Matrix{1, 1}, matrix.Matrix{2, 2}, matrix.Matrix{1.5, 1.5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntersects(tt.args.p1, tt.args.p2, tt.args.q); got != tt.want {
				t.Errorf("IsIntersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIntersectsTwo(t *testing.T) {
	type args struct {
		p1 matrix.Matrix
		p2 matrix.Matrix
		q1 matrix.Matrix
		q2 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"IsIntersectsTwo", args{matrix.Matrix{1, 1}, matrix.Matrix{2, 2}, matrix.Matrix{1.5, 1.5}, matrix.Matrix{2.5, 2.5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntersectsTwo(tt.args.p1, tt.args.p2, tt.args.q1, tt.args.q2); got != tt.want {
				t.Errorf("IsIntersectsTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		name string
		want *Envelope
	}{
		{"Empty", &Envelope{-1, 0, -1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFourFloat(t *testing.T) {
	type args struct {
		x1 float64
		x2 float64
		y1 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want *Envelope
	}{
		{"FourFloat", args{1, 2, 1, 2}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourFloat(tt.args.x1, tt.args.x2, tt.args.y1, tt.args.y2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoMatrix(t *testing.T) {
	type args struct {
		p1 matrix.Matrix
		p2 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want *Envelope
	}{
		{"TwoMatrix", args{matrix.Matrix{1, 1}, matrix.Matrix{2, 2}}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoMatrix(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixList(t *testing.T) {
	type args struct {
		ps []matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want *Envelope
	}{
		{"MatrixList", args{[]matrix.Matrix{{1, 1}, {2, 2}}}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixList(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygonMatrixList(t *testing.T) {
	type args struct {
		ps []matrix.PolygonMatrix
	}
	tests := []struct {
		name string
		args args
		want *Envelope
	}{
		{"PolygonMatrixList", args{[]matrix.PolygonMatrix{{{{1, 1}}}, {{{2, 2}}}}}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PolygonMatrixList(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PolygonMatrixList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix(t *testing.T) {
	type args struct {
		p matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want *Envelope
	}{
		{"Matrix", args{matrix.Matrix{1, 1}}, &Envelope{1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Matrix(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBound(t *testing.T) {
	type args struct {
		b []matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want *Envelope
	}{
		{"Bound", args{[]matrix.Matrix{{1, 1}, {2, 2}}}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bound(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnv(t *testing.T) {
	type args struct {
		env *Envelope
	}
	tests := []struct {
		name string
		args args
		want *Envelope
	}{
		{"Env", args{&Envelope{2, 1, 2, 1}}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Env(tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Env() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Copy(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Envelope
	}{
		{"Env", fields{2, 1, 2, 1}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Envelope.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_IsNil(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"IsNil", fields{1, 2, 2, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.IsNil(); got != tt.want {
				t.Errorf("Envelope.IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Width(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Width", fields{2, 1, 2, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Width(); got != tt.want {
				t.Errorf("Envelope.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Height(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Width", fields{2, 1, 2, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Height(); got != tt.want {
				t.Errorf("Envelope.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_MinExtent(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"MinExtent", fields{2, 1, 2, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.MinExtent(); got != tt.want {
				t.Errorf("Envelope.MinExtent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_MaxExtent(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"MinExtent", fields{2, 1, 2, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.MaxExtent(); got != tt.want {
				t.Errorf("Envelope.MaxExtent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_ExpandToIncludeMatrix(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		p matrix.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Envelope
	}{
		{"ExpandToIncludeMatrix", fields{2, 1, 2, 1}, args{matrix.Matrix{3, 1}}, &Envelope{3, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			e.ExpandToIncludeMatrix(tt.args.p)
			if !e.Equals(tt.want) {
				t.Errorf("Envelope.ExpandToIncludeMatrix() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEnvelope_ExpandBy(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		distance float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Envelope
	}{
		{"ExpandBy", fields{2, 1, 2, 1}, args{1}, &Envelope{3, 0, 3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			e.ExpandBy(tt.args.distance)
			if !e.Equals(tt.want) {
				t.Errorf("Envelope.ExpandBy() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEnvelope_ExpandByXY(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		deltaX float64
		deltaY float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Envelope
	}{
		{"ExpandByXY", fields{2, 1, 2, 1}, args{1, 1}, &Envelope{3, 0, 3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			e.ExpandByXY(tt.args.deltaX, tt.args.deltaY)
			if !e.Equals(tt.want) {
				t.Errorf("Envelope.ExpandByXY() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEnvelope_ExpandToInclude(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Envelope
	}{
		{"ExpandToInclude", fields{2, 1, 2, 1}, args{3, 3}, &Envelope{3, 1, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			e.ExpandToInclude(tt.args.x, tt.args.y)
			if !e.Equals(tt.want) {
				t.Errorf("Envelope.ExpandToInclude() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEnvelope_ExpandToIncludeEnv(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Envelope
	}{
		{"ExpandToIncludeEnv", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, &Envelope{3, 1, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			e.ExpandToIncludeEnv(tt.args.other)
			if !e.Equals(tt.want) {
				t.Errorf("Envelope.ExpandToIncludeEnv() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEnvelope_Translate(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		transX float64
		transY float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Envelope
	}{
		{"Translate", fields{2, 1, 2, 1}, args{3, 3}, &Envelope{5, 4, 5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			e.Translate(tt.args.transX, tt.args.transY)
			if !e.Equals(tt.want) {
				t.Errorf("Envelope.Translate() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEnvelope_Centre(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   matrix.Matrix
	}{
		{"Centre", fields{2, 0, 2, 0}, matrix.Matrix{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Centre(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Envelope.Centre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Intersection(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		env *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Envelope
	}{
		{"Intersection", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, &Envelope{2, 1, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Intersection(tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Envelope.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_IsIntersects(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"IsIntersects", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.IsIntersects(tt.args.other); got != tt.want {
				t.Errorf("Envelope.IsIntersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Disjoint(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Disjoint", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Disjoint(tt.args.other); got != tt.want {
				t.Errorf("Envelope.Disjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Overlaps(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Overlaps", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Overlaps(tt.args.other); got != tt.want {
				t.Errorf("Envelope.Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Contains(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Contains", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Contains(tt.args.other); got != tt.want {
				t.Errorf("Envelope.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Covers(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Covers", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Covers(tt.args.other); got != tt.want {
				t.Errorf("Envelope.Covers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Distance(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		env *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"Distance", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Distance(tt.args.env); got != tt.want {
				t.Errorf("Envelope.Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Equals(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Equals", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Equals(tt.args.other); got != tt.want {
				t.Errorf("Envelope.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_ToString(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"ToString", fields{2, 1, 2, 1}, "Env[1 : 2, 1 : 2]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.ToString(); got != tt.want {
				t.Errorf("Envelope.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_CompareTo(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	type args struct {
		other *Envelope
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"CompareTo", fields{2, 1, 2, 1}, args{&Envelope{3, 1, 3, 1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.CompareTo(tt.args.other); got != tt.want {
				t.Errorf("Envelope.CompareTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_ToMatrix(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *matrix.PolygonMatrix
	}{
		{"CompareTo", fields{2, 1, 2, 1}, &matrix.PolygonMatrix{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.ToMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Envelope.ToMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
