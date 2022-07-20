package matrix

import (
	"reflect"
	"testing"
)

func TestIntersectionMatrixDefault(t *testing.T) {
	tests := []struct {
		name string
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{-1, -1, -1},
				{-1, -1, -1},
				{-1, -1, -1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectionMatrixDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectionMatrixDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_SetAll(t *testing.T) {
	type args struct {
		dimensionValue int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{-1, -1, -1},
				{-1, -1, -1},
				{-1, -1, -1}}}, args{1},
			&IntersectionMatrix{
				[][]int{{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.im.SetAll(tt.args.dimensionValue)
			if !reflect.DeepEqual(tt.im, tt.want) {
				t.Errorf("IntersectionMatrixDefault() = %v, want %v", tt.im, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsDisjoint(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{-1, -1, -1},
				{-1, -1, -1},
				{-1, -1, -1}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsDisjoint(); got != tt.want {
				t.Errorf("IntersectionMatrix.IsDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsIntersects(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, -1},
				{-1, -1, -1},
				{-1, -1, -1}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsIntersects(); got != tt.want {
				t.Errorf("IntersectionMatrix.IsIntersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsTouches(t *testing.T) {
	type args struct {
		dimensionOfGeometryA int
		dimensionOfGeometryB int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{-1, -1, -1},
				{1, -1, -1},
				{-1, -1, -1}}},
			args{1, 1}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsTouches(tt.args.dimensionOfGeometryA, tt.args.dimensionOfGeometryB); got != tt.want {
				t.Errorf("IntersectionMatrix.IsTouches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsCrosses(t *testing.T) {
	type args struct {
		dimensionOfGeometryA int
		dimensionOfGeometryB int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}},
			args{2, 1}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsCrosses(tt.args.dimensionOfGeometryA, tt.args.dimensionOfGeometryB); got != tt.want {
				t.Errorf("IntersectionMatrix.IsCrosses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsWithin(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, -1},
				{1, -1, -1},
				{1, -1, -1}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsWithin(); got != tt.want {
				t.Errorf("IntersectionMatrix.IsWithin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsContains(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, -1},
				{1, -1, -1},
				{-1, -1, -1}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsContains(); got != tt.want {
				t.Errorf("IntersectionMatrix.IsContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsCovers(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, -1},
				{1, -1, -1},
				{-1, -1, -1}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsCovers(); got != tt.want {
				t.Errorf("IntersectionMatrix.IsCovers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsCoveredBy(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, -1},
				{1, -1, -1},
				{-1, -1, -1}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsCoveredBy(); got != tt.want {
				t.Errorf("IntersectionMatrix.IsCoveredBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsEquals(t *testing.T) {
	type args struct {
		dimensionOfGeometryA int
		dimensionOfGeometryB int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, -1},
				{1, -1, -1},
				{-1, -1, -1}}}, args{2, 2}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsEquals(tt.args.dimensionOfGeometryA, tt.args.dimensionOfGeometryB); got != tt.want {
				t.Errorf("IntersectionMatrix.IsEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_IsOverlaps(t *testing.T) {
	type args struct {
		dimensionOfGeometryA int
		dimensionOfGeometryB int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, args{2, 2}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.IsOverlaps(tt.args.dimensionOfGeometryA, tt.args.dimensionOfGeometryB); got != tt.want {
				t.Errorf("IntersectionMatrix.IsOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_Matches(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		im      *IntersectionMatrix
		args    args
		want    bool
		wantErr bool
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, args{"TFTTFFTFF"}, true, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.im.Matches(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntersectionMatrix.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntersectionMatrix.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_ToString(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want string
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, "1F11FF1FF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.ToString(); got != tt.want {
				t.Errorf("IntersectionMatrix.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_Set(t *testing.T) {
	type args struct {
		row            int
		column         int
		dimensionValue int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, args{0, 1, 1},
			&IntersectionMatrix{
				[][]int{{1, 1, 1},
					{1, -1, -1},
					{1, -1, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.im.Set(tt.args.row, tt.args.column, tt.args.dimensionValue)
			if !reflect.DeepEqual(tt.want, tt.im) {
				t.Errorf("IntersectionMatrix.Set() = %v, want %v", tt.im, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_SetString(t *testing.T) {
	type args struct {
		dimensionSymbols string
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, args{"1111FF1FF"},
			&IntersectionMatrix{
				[][]int{{1, 1, 1},
					{1, -1, -1},
					{1, -1, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.im.SetString(tt.args.dimensionSymbols)
			if !reflect.DeepEqual(tt.want, tt.im) {
				t.Errorf("IntersectionMatrix.SetString() = %v, want %v", tt.im, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_SetAtLeast(t *testing.T) {
	type args struct {
		row                   int
		column                int
		minimumDimensionValue int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, args{0, 1, 1},
			&IntersectionMatrix{
				[][]int{{1, 1, 1},
					{1, -1, -1},
					{1, -1, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.im.SetAtLeast(tt.args.row, tt.args.column, tt.args.minimumDimensionValue)
			if !reflect.DeepEqual(tt.want, tt.im) {
				t.Errorf("IntersectionMatrix.SetAtLeast() = %v, want %v", tt.im, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_SetAtLeastIfValid(t *testing.T) {
	type args struct {
		row                   int
		column                int
		minimumDimensionValue int
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, args{0, 1, 1},
			&IntersectionMatrix{
				[][]int{{1, 1, 1},
					{1, -1, -1},
					{1, -1, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.im.SetAtLeastIfValid(tt.args.row, tt.args.column, tt.args.minimumDimensionValue)
			if !reflect.DeepEqual(tt.want, tt.im) {
				t.Errorf("IntersectionMatrix.SetAtLeastIfValid() = %v, want %v", tt.im, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_SetAtLeastString(t *testing.T) {
	type args struct {
		minimumDimensionSymbols string
	}
	tests := []struct {
		name string
		im   *IntersectionMatrix
		args args
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}}, args{"1111FF1FF"},
			&IntersectionMatrix{
				[][]int{{1, 1, 1},
					{1, -1, -1},
					{1, -1, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.im.SetAtLeastString(tt.args.minimumDimensionSymbols)
			if !reflect.DeepEqual(tt.want, tt.im) {
				t.Errorf("IntersectionMatrix.SetAtLeastString() = %v, want %v", tt.im, tt.want)
			}
		})
	}
}

func TestIntersectionMatrix_Transpose(t *testing.T) {
	tests := []struct {
		name string
		im   *IntersectionMatrix
		want *IntersectionMatrix
	}{
		{"case1", &IntersectionMatrix{
			[][]int{{1, -1, 1},
				{1, -1, -1},
				{1, -1, -1}}},
			&IntersectionMatrix{
				[][]int{{1, 1, 1},
					{-1, -1, -1},
					{1, -1, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.im.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectionMatrix.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDimensionSymbol(t *testing.T) {
	type args struct {
		dimensionValue int
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{"case 1", args{2}, '2', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toDimensionSymbol(tt.args.dimensionValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("toDimensionSymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toDimensionSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDimensionValue(t *testing.T) {
	type args struct {
		dimensionSymbol byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"case 1", args{'1'}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toDimensionValue(tt.args.dimensionSymbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("toDimensionValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toDimensionValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matches(t *testing.T) {
	type args struct {
		actualDimensionValue    int
		requiredDimensionSymbol byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{1, 'T'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matches(tt.args.actualDimensionValue, tt.args.requiredDimensionSymbol); got != tt.want {
				t.Errorf("matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTrue(t *testing.T) {
	type args struct {
		actualDimensionValue int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTrue(tt.args.actualDimensionValue); got != tt.want {
				t.Errorf("isTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}
