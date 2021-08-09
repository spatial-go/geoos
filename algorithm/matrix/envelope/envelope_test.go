package envelope

import (
	"math"
	"testing"
)

func TestEnvelope_Area(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"area", fields{200, 100, 200, 100}, 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Area(); got != tt.want {
				t.Errorf("Envelope.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvelope_Diameter(t *testing.T) {
	type fields struct {
		MaxX float64
		MinX float64
		MaxY float64
		MinY float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Diameter", fields{200, 100, 200, 100}, math.Sqrt(20000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Envelope{
				MaxX: tt.fields.MaxX,
				MinX: tt.fields.MinX,
				MaxY: tt.fields.MaxY,
				MinY: tt.fields.MinY,
			}
			if got := e.Diameter(); got != tt.want {
				t.Errorf("Envelope.Diameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
