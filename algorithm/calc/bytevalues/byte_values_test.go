// Package bytevalues ead and write primitive datatypes from/to byte
package bytevalues

import (
	"testing"
)

func TestGetInt32(t *testing.T) {
	type args struct {
		buf       []byte
		byteOrder int
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"Getint32", args{[]byte{1, 0, 0, 32}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt32(tt.args.buf, tt.args.byteOrder); (got&0xffff)%1000 != tt.want {
				t.Errorf("GetInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutInt32(t *testing.T) {
	type args struct {
		buf       []byte
		intValue  int32
		byteOrder int
	}
	tests := []struct {
		name string
		args args
	}{
		{"PutInt32", args{make([]byte, 4), 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PutInt32(tt.args.buf, tt.args.intValue, tt.args.byteOrder)
			t.Log("succeed.")
		})
	}
}

func TestGetInt64(t *testing.T) {
	type args struct {
		buf       []byte
		byteOrder int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"GetInt64", args{[]byte{1, 0, 0, 0, 0, 0, 0, 0}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt64(tt.args.buf, tt.args.byteOrder); got != tt.want {
				t.Errorf("GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutInt64(t *testing.T) {
	type args struct {
		buf       []byte
		intValue  int64
		byteOrder int
	}
	tests := []struct {
		name string
		args args
	}{
		{"PutInt32", args{make([]byte, 8), 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PutInt64(tt.args.buf, tt.args.intValue, tt.args.byteOrder)
			t.Logf("succeed.  %v", tt.args.buf)
		})
	}
}

func TestGetFloat32(t *testing.T) {
	type args struct {
		buf       []byte
		byteOrder int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"GetFloat32", args{[]byte{0, 0, 128, 63}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloat32(tt.args.buf, tt.args.byteOrder); got != tt.want {
				t.Errorf("GetFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutFloat32(t *testing.T) {
	type args struct {
		buf        []byte
		floatValue float32
		byteOrder  int
	}
	tests := []struct {
		name string
		args args
	}{
		{"PutFloat32", args{make([]byte, 4), 1.0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PutFloat32(tt.args.buf, tt.args.floatValue, tt.args.byteOrder)
			t.Logf("succeed.  %v", tt.args.buf)
		})
	}
}

func TestGetFloat64(t *testing.T) {
	type args struct {
		buf       []byte
		byteOrder int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"GetFloat64", args{[]byte{0, 0, 0, 0, 0, 0, 240, 63}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloat64(tt.args.buf, tt.args.byteOrder); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutFloat64(t *testing.T) {
	type args struct {
		buf        []byte
		floatValue float64
		byteOrder  int
	}
	tests := []struct {
		name string
		args args
	}{
		{"PutFloat32", args{make([]byte, 8), 1.0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PutFloat64(tt.args.buf, tt.args.floatValue, tt.args.byteOrder)
			t.Logf("succeed.  %v", tt.args.buf)
		})
	}
}
