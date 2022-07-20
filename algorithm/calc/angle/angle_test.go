package angle

import (
	"math"
	"testing"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestAngle(t *testing.T) {
	type args struct {
		p0 matrix.Matrix
		p1 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"angle", args{matrix.Matrix{100, 100}, matrix.Matrix{100, 200}}, math.Pi / 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Angle(tt.args.p0, tt.args.p1); got != tt.want {
				t.Errorf("Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDegrees(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"ToDegrees", args{math.Pi / 2}, 90.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDegrees(tt.args.radians); got != tt.want {
				t.Errorf("ToDegrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToRadians(t *testing.T) {
	type args struct {
		angleDegrees float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"ToRadians", args{90}, math.Pi / 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRadians(tt.args.angleDegrees); got != tt.want {
				t.Errorf("ToRadians() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixAngle(t *testing.T) {
	type args struct {
		p matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"MatrixAngle", args{matrix.Matrix{100, 100}}, PiOver4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixAngle(tt.args.p); got != tt.want {
				t.Errorf("MatrixAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAcute(t *testing.T) {
	type args struct {
		p0 matrix.Matrix
		p1 matrix.Matrix
		p2 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"IsAcute", args{matrix.Matrix{100, 100}, matrix.Matrix{50, 50}, matrix.Matrix{100, 30}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAcute(tt.args.p0, tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("IsAcute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsObtuse(t *testing.T) {
	type args struct {
		p0 matrix.Matrix
		p1 matrix.Matrix
		p2 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"IsObtuse", args{matrix.Matrix{100, 100}, matrix.Matrix{50, 50}, matrix.Matrix{100, 30}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsObtuse(tt.args.p0, tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("IsObtuse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		tip1 matrix.Matrix
		tail matrix.Matrix
		tip2 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Between", args{matrix.Matrix{100, 100}, matrix.Matrix{50, 50}, matrix.Matrix{100, 50}}, PiOver4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Between(tt.args.tip1, tt.args.tail, tt.args.tip2); got != tt.want {
				t.Errorf("Between() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetweenOriented(t *testing.T) {
	type args struct {
		tip1 matrix.Matrix
		tail matrix.Matrix
		tip2 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"BetweenOriented", args{matrix.Matrix{100, 100}, matrix.Matrix{50, 50}, matrix.Matrix{100, 50}}, -1 * PiOver4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BetweenOriented(tt.args.tip1, tt.args.tail, tt.args.tip2); got != tt.want {
				t.Errorf("BetweenOriented() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteriorAngle(t *testing.T) {
	type args struct {
		p0 matrix.Matrix
		p1 matrix.Matrix
		p2 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"InteriorAngle", args{matrix.Matrix{100, 100}, matrix.Matrix{50, 50}, matrix.Matrix{50, 100}}, PiOver4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InteriorAngle(tt.args.p0, tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("InteriorAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTurn(t *testing.T) {
	type args struct {
		ang1 float64
		ang2 float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Turn", args{90, 80}, calc.CounterClockWise},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Turn(tt.args.ang1, tt.args.ang2); got != tt.want {
				t.Errorf("Turn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Normalize", args{math.Pi*2 + 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Normalize(tt.args.angle); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizePositive(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"NormalizePositive", args{math.Pi*2 + 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizePositive(tt.args.angle); got != tt.want {
				t.Errorf("NormalizePositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args struct {
		ang1 float64
		ang2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Diff", args{math.Pi*2 + 1, -1}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.ang1, tt.args.ang2); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
