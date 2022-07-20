// Package clipping the spatial geometric operation and reconstruction between entities is realized.
// a method for spatial geometry operation by update geometric relation graph.
package clipping

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
		name    string
		args    args
		want    []matrix.LineMatrix
		wantErr bool
	}{
		{name: "line line", args: args{matrix.Collection{matrix.LineMatrix{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}},
			matrix.LineMatrix{{-45, -33}, {-46, -32}}}},
			want: []matrix.LineMatrix{{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}, {-46, -32}}}},
		{name: "line line1", args: args{matrix.Collection{matrix.LineMatrix{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}},
			matrix.LineMatrix{{-45.2, -33.2}, {-46, -32}}}},
			want: []matrix.LineMatrix{{{-29, -27}, {-30, -29.7}, {-36, -31}, {-45, -33}}, {{-45.2, -33.2}, {-46, -32}}}},
		{name: "line line2", args: args{matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{30, 150}, {80, 150}}}},
			want: []matrix.LineMatrix{{{50, 100}, {50, 200}}, {{30, 150}, {80, 150}}}},
		{name: "line overlay", args: args{matrix.Collection{matrix.LineMatrix{{50, 80}, {50, 100}, {50, 120}, {120, 120}}, matrix.LineMatrix{{50, 100}, {50, 120}, {50, 150}}}},
			want: []matrix.LineMatrix{{{50, 80}, {50, 100}, {50, 120}, {120, 120}}, {{50, 120}, {50, 150}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LineMerge(tt.args.ml)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineMerge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
