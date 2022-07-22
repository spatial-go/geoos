package operation

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestValidOP_IsSimple(t *testing.T) {
	type fields struct {
		Steric matrix.Steric
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Valid isSimple polygon", fields: fields{matrix.PolygonMatrix{
			{{1, 2}, {3, 4}, {5, 6}, {5, 3}, {1, 2}},
		}},
			want: true},
		{name: "Valid isSimple line", fields: fields{matrix.LineMatrix{
			{1, 1}, {2, 2}, {2, 3.5}, {1, 3}, {1, 2}, {2, 1},
		}},
			want: false},
		{name: "Valid isSimple ring", fields: fields{matrix.LineMatrix{
			{1, 2}, {3, 4}, {5, 6}, {5, 3}, {1, 2}},
		},
			want: true},

		{name: "Valid isSimple ring1", fields: fields{matrix.LineMatrix{
			{113.00084732367,
				22.5083135499}, {113.00079141659,
				22.50835737586}, {113.0007903827,
				22.5083583067}, {113.00084732367,
				22.5083135499}},
		},
			want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := &ValidOP{
				Steric: tt.fields.Steric,
			}
			if got := el.IsSimple(); got != tt.want {
				t.Errorf("ValidOP.IsSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrectPolygonMatrixSelfIntersect(t *testing.T) {
	type args struct {
		ms matrix.Steric
	}
	tests := []struct {
		name string
		args args
		want matrix.Steric
	}{
		{"self intersect poly", args{matrix.PolygonMatrix{{{1, 1}, {1, 2}, {2, 2}, {2, 3}, {3, 3}, {3, 2}, {2, 2}, {2, 1}, {1, 1}}}},
			matrix.Collection{matrix.PolygonMatrix{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}},
				matrix.PolygonMatrix{{{2, 2}, {2, 3}, {3, 3}, {3, 2}, {2, 2}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorrectPolygonMatrixSelfIntersect(tt.args.ms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CorrectPolygonMatrixSelfIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrectRingSelfIntersect(t *testing.T) {
	type args struct {
		shell matrix.LineMatrix
	}
	tests := []struct {
		name  string
		args  args
		want  []matrix.LineMatrix
		want1 bool
	}{
		{"self intersect poly", args{matrix.LineMatrix{{1, 1}, {1, 2}, {2, 2}, {2, 3}, {3, 3}, {3, 2}, {2, 2}, {2, 1}, {1, 1}}},
			[]matrix.LineMatrix{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}},
				{{2, 2}, {2, 3}, {3, 3}, {3, 2}, {2, 2}}}, true,
		},
		{"self intersect poly", args{matrix.LineMatrix{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}},
			nil, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CorrectRingSelfIntersect(tt.args.shell)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CorrectRingSelfIntersect() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CorrectRingSelfIntersect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
