package sharedpaths

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestSharedPaths(t *testing.T) {
	type args struct {
		g1 matrix.Steric
		g2 matrix.Steric
	}
	tests := []struct {
		name        string
		args        args
		wantForwDir matrix.Collection
		wantBackDir matrix.Collection
		wantErr     bool
	}{
		{name: "share paths", args: args{
			g1: matrix.Collection{
				matrix.LineMatrix{{26, 125}, {26, 200}, {126, 200}, {126, 125}, {26, 125}},
				matrix.LineMatrix{{51, 150}, {101, 150}, {76, 175}, {51, 150}},
			},
			g2: matrix.LineMatrix{{151, 100}, {126, 156.25}, {126, 125}, {90, 161}, {76, 175}},
		},
			wantForwDir: matrix.Collection{
				matrix.LineMatrix{{126, 156.25}, {126, 125}},
				matrix.LineMatrix{{101, 150}, {90, 161}},
			},
			wantBackDir: nil,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotForwDir, gotBackDir, err := SharedPaths(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("SharedPaths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotForwDir, tt.wantForwDir) {
				t.Errorf("SharedPaths() gotForwDir = %v, want %v", gotForwDir, tt.wantForwDir)
			}
			if !reflect.DeepEqual(gotBackDir, tt.wantBackDir) {
				t.Errorf("SharedPaths() gotBackDir = %v, want %v", gotBackDir, tt.wantBackDir)
			}
		})
	}
}
