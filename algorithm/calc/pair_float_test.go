package calc

import (
	"math"
	"reflect"
	"testing"
)

func TestPairFloat_SelfAdd(t *testing.T) {
	type fields struct {
		Hi float64
		Lo float64
	}
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PairFloat
	}{
		{name: "pair float self add", fields: fields{1.0, 2.0}, args: args{0.5, 0.5},
			want: &PairFloat{4.0, 0.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &PairFloat{
				Hi: tt.fields.Hi,
				Lo: tt.fields.Lo,
			}
			if got := d.SelfAdd(tt.args.yhi, tt.args.ylo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.SelfAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Ge(t *testing.T) {
	type fields struct {
		Hi float64
		Lo float64
	}
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Ge", fields{
			Hi: 112.64007568359376,
			Lo: 0.0,
		},
			args{&PairFloat{
				Hi: 112.64007568359375,
				Lo: 0.0,
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &PairFloat{
				Hi: tt.fields.Hi,
				Lo: tt.fields.Lo,
			}
			if got := d.Ge(tt.args.y); got != tt.want {
				t.Errorf("PairFloat.Ge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfDividePair(t *testing.T) {
	type fields struct {
		Hi float64
		Lo float64
	}
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PairFloat
	}{
		{"divide", fields{
			Hi: 112.64007568359376,
			Lo: 0.0,
		},
			args{&PairFloat{
				Hi: 0.0,
				Lo: 0.0,
			},
			},
			&PairFloat{math.NaN(), math.NaN()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &PairFloat{
				Hi: tt.fields.Hi,
				Lo: tt.fields.Lo,
			}
			if got := d.SelfDividePair(tt.args.y); got.Equals(tt.want) {
				t.Errorf("PairFloat.SelfDividePair() = %v, want %v, <1 %v", got, tt.want, got.Value() < 1)
			}
		})
	}
}

func TestValueOf(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want *PairFloat
	}{
		{"valueOf", args{0.15}, &PairFloat{0.15, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValueOf(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_AddOne(t *testing.T) {
	type args struct {
		y float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"addOne", &PairFloat{0.15, 0}, args{0.15}, &PairFloat{0.3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddOne(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.AddOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_AddPair(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"add", &PairFloat{0.15, 0}, args{&PairFloat{0.15, 0}}, &PairFloat{0.3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddPair(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.AddPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Add(t *testing.T) {
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"add", &PairFloat{0.15, 0}, args{0.15, 0}, &PairFloat{0.3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Add(tt.args.yhi, tt.args.ylo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfAddOne(t *testing.T) {
	type args struct {
		y float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"add", &PairFloat{0.15, 0}, args{0.15}, &PairFloat{0.3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfAddOne(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.SelfAddOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfAddPair(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"add", &PairFloat{0.15, 0}, args{&PairFloat{0.15, 0}}, &PairFloat{0.3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfAddPair(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.SelfAddPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SubtractPair(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"sub", &PairFloat{0.15, 0}, args{&PairFloat{0.15, 0}}, &PairFloat{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SubtractPair(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.SubtractPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Subtract(t *testing.T) {
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"sub", &PairFloat{0.15, 0}, args{0.15, 0}, &PairFloat{0.0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Subtract(tt.args.yhi, tt.args.ylo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfSubtractPair(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"sub", &PairFloat{0.15, 0}, args{&PairFloat{0.15, 0}}, &PairFloat{0.0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfSubtractPair(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.SelfSubtractPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfSubtract(t *testing.T) {
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"sub", &PairFloat{0.15, 0}, args{0.15, 0}, &PairFloat{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfSubtract(tt.args.yhi, tt.args.ylo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairFloat.SelfSubtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_MultiplyPair(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"mul", &PairFloat{0.1, 0}, args{&PairFloat{0.1, 0}}, &PairFloat{0.01, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.MultiplyPair(tt.args.y); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("PairFloat.MultiplyPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Multiply(t *testing.T) {
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"mul", &PairFloat{0.1, 0}, args{0.1, 0}, &PairFloat{0.01, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Multiply(tt.args.yhi, tt.args.ylo); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("PairFloat.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfMultiplyPair(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"mul", &PairFloat{0.1, 0}, args{&PairFloat{0.1, 0}}, &PairFloat{0.01, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfMultiplyPair(tt.args.y); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("PairFloat.SelfMultiplyPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfMultiply(t *testing.T) {
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"mul", &PairFloat{0.1, 0}, args{0.1, 0}, &PairFloat{0.01, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfMultiply(tt.args.yhi, tt.args.ylo); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("PairFloat.SelfMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_DividePair(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"div", &PairFloat{0.01, 0}, args{&PairFloat{0.1, 0}}, &PairFloat{0.1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.DividePair(tt.args.y); math.Abs(got.Hi-tt.want.Hi) > 1.0e-16 {
				t.Errorf("PairFloat.DividePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Divide(t *testing.T) {
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"div", &PairFloat{0.01, 0}, args{0.1, 0}, &PairFloat{0.1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Divide(tt.args.yhi, tt.args.ylo); math.Abs(got.Hi-tt.want.Hi) > 1.0e-16 {
				t.Errorf("PairFloat.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfDivide(t *testing.T) {
	type args struct {
		yhi float64
		ylo float64
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want *PairFloat
	}{
		{"div", &PairFloat{0.01, 0}, args{0.1, 0}, &PairFloat{0.1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfDivide(tt.args.yhi, tt.args.ylo); math.Abs(got.Hi-tt.want.Hi) > 1.0e-16 {
				t.Errorf("PairFloat.SelfDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Signum(t *testing.T) {
	tests := []struct {
		name string
		d    *PairFloat
		want int
	}{
		{"signum", &PairFloat{0.01, 0}, 1},
		{"signum", &PairFloat{-0.01, 0}, -1},
		{"signum", &PairFloat{0.0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Signum(); got != tt.want {
				t.Errorf("PairFloat.Signum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Pow2(t *testing.T) {
	tests := []struct {
		name string
		d    *PairFloat
		want *PairFloat
	}{
		{"sqr", &PairFloat{4, 0}, &PairFloat{16, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Pow2(); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("PairFloat.Sqr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_SelfSqr(t *testing.T) {
	tests := []struct {
		name string
		d    *PairFloat
		want *PairFloat
	}{
		{"sqr", &PairFloat{4, 0}, &PairFloat{16, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SelfPow2(); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("PairFloat.SelfSqr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_IsZero(t *testing.T) {
	tests := []struct {
		name string
		d    *PairFloat
		want bool
	}{
		{"iszero", &PairFloat{4, 0}, false},
		{"iszero", &PairFloat{0, 1}, false},
		{"iszero", &PairFloat{0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsZero(); got != tt.want {
				t.Errorf("PairFloat.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Equals(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want bool
	}{
		{"equals", &PairFloat{0.01, 0}, args{&PairFloat{0.01, 0}}, true},
		{"equals", &PairFloat{0.01, 0}, args{&PairFloat{0.01, 0.0000000000001}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Equals(tt.args.y); got != tt.want {
				t.Errorf("PairFloat.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Gt(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want bool
	}{
		{"gt 0", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.0000000000001}}, true},
		{"gt 1", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.00001}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Gt(tt.args.y); got != tt.want {
				t.Errorf("PairFloat.Gt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Lt(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want bool
	}{
		{"lt 0", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.0000000000001}}, false},
		{"lt 1", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.00001}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Lt(tt.args.y); got != tt.want {
				t.Errorf("PairFloat.Lt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Le(t *testing.T) {
	type args struct {
		y *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want bool
	}{
		{"le 0", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.0000000000001}}, false},
		{"le 1", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.00001}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Le(tt.args.y); got != tt.want {
				t.Errorf("PairFloat.Le() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_CompareTo(t *testing.T) {
	type args struct {
		other *PairFloat
	}
	tests := []struct {
		name string
		d    *PairFloat
		args args
		want int
	}{
		{"CompareTo 0", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.0000000000001}}, 1},
		{"CompareTo 1", &PairFloat{0.01, 0.00000001}, args{&PairFloat{0.01, 0.00001}}, -1},
		{"CompareTo 2", &PairFloat{0.01, 0.00001}, args{&PairFloat{0.01, 0.00001}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.CompareTo(tt.args.other); got != tt.want {
				t.Errorf("PairFloat.CompareTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairFloat_Value(t *testing.T) {
	tests := []struct {
		name string
		d    *PairFloat
		want float64
	}{
		{"value 1", &PairFloat{0.01, 0.00000001}, 0.01000001},
		{"value 2", &PairFloat{0.01, 0.0}, 0.01},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Value(); got != tt.want {
				t.Errorf("PairFloat.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignum(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"signum", args{0.01}, 1},
		{"signum", args{-0.01}, -1},
		{"signum", args{0.0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Signum(tt.args.x); got != tt.want {
				t.Errorf("Signum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeterminant(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want *PairFloat
	}{
		{"determinant", args{2, 3, 4, 5}, &PairFloat{-2, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Determinant(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeterminantPair(t *testing.T) {
	type args struct {
		x1 *PairFloat
		y1 *PairFloat
		x2 *PairFloat
		y2 *PairFloat
	}
	tests := []struct {
		name string
		args args
		want *PairFloat
	}{
		{"determinant", args{&PairFloat{3, 0},
			&PairFloat{3, 0}, &PairFloat{4, 0}, &PairFloat{5, 0}}, &PairFloat{3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeterminantPair(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); math.Abs(got.Hi-tt.want.Hi) > 1.0e-17 {
				t.Errorf("DeterminantPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
