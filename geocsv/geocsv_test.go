package geocsv

import (
	"testing"

	"github.com/spatial-go/geoos/space"
)

func TestGeoCSV_Test1(t *testing.T) {
	type args struct {
		filePath string
		options  Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				filePath: "./test1.csv",
				options: Options{
					XField: "x",
					YField: "y",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gc, err := Read(tt.args.filePath, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("GeoCSV.Read() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if len(gc.headers) != 4 {
					t.Error("length of headers is wrong")
				}
				if len(gc.rows) != 4 {
					t.Error("length of rows is wrong")
				}
				features := gc.ToGeoJSON()
				if len(features.Features) != 4 {
					t.Error("length of features is wrong")
				}
				point := features.Features[0].Geometry.Coordinates.(space.Point)
				if point[0] != 2 || point[1] != 49 {
					t.Error("Coordinates is wrong")
				}
			}
		})
	}
}

func TestGeoCSV_Test2(t *testing.T) {
	type args struct {
		filePath string
		options  Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test2",
			args: args{
				filePath: "./test2.csv",
				options: Options{
					WKTField: "wkt",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gc, err := Read(tt.args.filePath, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("GeoCSV.Read() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if len(gc.headers) != 3 {
					t.Error("length of headers is wrong")
				}
				if len(gc.rows) != 4 {
					t.Error("length of rows is wrong")
				}
				features := gc.ToGeoJSON()
				if len(features.Features) != 4 {
					t.Error("length of features is wrong")
				}
				point := features.Features[1].Geometry.Coordinates.(space.Point)
				if point[0] != 3 || point[1] != 50 {
					t.Error("Coordinates is wrong")
				}
			}
		})
	}
}
