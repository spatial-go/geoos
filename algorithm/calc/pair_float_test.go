package calc

import (
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
