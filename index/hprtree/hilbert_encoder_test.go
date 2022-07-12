package hprtree

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

func TestNewHilbertEncoder(t *testing.T) {
	type args struct {
		level  int
		extent *envelope.Envelope
	}
	tests := []struct {
		name string
		args args
		want *HilbertEncoder
	}{
		{"case1", args{1, &envelope.Envelope{MaxX: 4, MinX: 1, MaxY: 5, MinY: 1}}, &HilbertEncoder{1, 1, 1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHilbertEncoder(tt.args.level, tt.args.extent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHilbertEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
