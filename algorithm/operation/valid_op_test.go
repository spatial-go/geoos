package operation

import (
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
