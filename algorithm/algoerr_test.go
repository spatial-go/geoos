package algorithm

import (
	"testing"
)

func TestErrUnknownType(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test error", args{[]interface{}{"I am unknown type"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrUnknownType(tt.args.obj...); (err != nil) != tt.wantErr {
				t.Errorf("ErrUnknownType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
