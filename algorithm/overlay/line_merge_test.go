package overlay

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestLineMerge(t *testing.T) {
	type args struct {
		ml matrix.Collection
	}
	tests := []struct {
		name string
		args args
		want matrix.Collection
	}{
		{name: "line line", args: args{matrix.Collection{matrix.LineMatrix{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}},
			matrix.LineMatrix{{-45, -33}, {-46, -32}}}},
			want: matrix.Collection{matrix.LineMatrix{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}, {-46, -32}}}},
		{name: "line line1", args: args{matrix.Collection{matrix.LineMatrix{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}},
			matrix.LineMatrix{{-45.2, -33.2}, {-46, -32}}}},
			want: matrix.Collection{matrix.LineMatrix{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}}, matrix.LineMatrix{{-45.2, -33.2}, {-46, -32}}}},
		{name: "line line2", args: args{matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{30, 150}, {80, 150}}}},
			want: matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{30, 150}, {80, 150}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LineMerge(tt.args.ml); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
