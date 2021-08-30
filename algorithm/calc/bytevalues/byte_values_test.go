// Package bytevalues ead and write primitive datatypes from/to byte
package bytevalues

import "testing"

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
