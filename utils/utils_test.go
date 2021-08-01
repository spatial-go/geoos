package utils

import "testing"

func TestGetStringEncoding(t *testing.T) {
	type args struct {
		dataStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{" string encoding", args{"way_id,pt_id,x,y"}, "UTF8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringEncoding(tt.args.dataStr); got != tt.want {
				t.Errorf("GetStringEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
