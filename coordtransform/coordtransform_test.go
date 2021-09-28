package coordtransform

import (
	"github.com/spatial-go/geoos/space"
	"testing"
)

func TestTransformer_TransformLatLng(t1 *testing.T) {
	type fields struct {
		CoordType string
	}
	type args struct {
		lng float64
		lat float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
		want1  float64
		tolerance float64
	}{
		{
			name: "lnglat to mercator", fields: fields{CoordType: LLTOMERCATOR},
			args: args{lng: 110,lat:40}, want: 12245143.99, want1: 4865942.28, tolerance: 0.1,
		},
		{
			name: "mercator to lnglat", fields: fields{CoordType: MERCATORTOLL},
			args: args{lng: 12245143,lat:4865942}, want: 109.9999911, want1: 39.9999981, tolerance: 0.0000001,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := GetInstance()
			t.CoordType = tt.fields.CoordType
			got, got1 := t.TransformLatLng(tt.args.lng, tt.args.lat)
			gotPoint := space.Point{got, got1}
			wantPoint := space.Point{tt.want, tt.want1}
			if !gotPoint.EqualsExact(wantPoint, tt.tolerance) {
				t1.Errorf("TransformLatLng() got = %v %v, want %v %v", got, got1, tt.want, tt.want1)
			}
		})
	}
}
