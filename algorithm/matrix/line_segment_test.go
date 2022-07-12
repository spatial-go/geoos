package matrix

import (
	"reflect"
	"testing"
)

func TestLineSegment_PointAlong(t *testing.T) {
	type args struct {
		segmentLengthFraction float64
	}
	tests := []struct {
		name string
		l    *LineSegment
		args args
		want Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PointAlong(tt.args.segmentLengthFraction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineSegment.PointAlong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineSegment_PointAlongOffset(t *testing.T) {
	type args struct {
		segmentLengthFraction float64
		offsetDistance        float64
	}
	tests := []struct {
		name    string
		l       *LineSegment
		args    args
		want    Matrix
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.PointAlongOffset(tt.args.segmentLengthFraction, tt.args.offsetDistance)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineSegment.PointAlongOffset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineSegment.PointAlongOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
