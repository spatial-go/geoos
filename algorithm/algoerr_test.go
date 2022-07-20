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

func TestErrorUnknownDimension(t *testing.T) {
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
			if err := ErrorUnknownDimension(tt.args.obj...); (err != nil) != tt.wantErr {
				t.Errorf("ErrorUnknownDimension() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		str string
		obj []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test error", args{"%v", []interface{}{"I am unknown type"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Error(tt.args.str, tt.args.obj...); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrorDimension(t *testing.T) {
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
			if err := ErrorDimension(tt.args.obj...); (err != nil) != tt.wantErr {
				t.Errorf("ErrorDimension() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrorShouldBeLength9(t *testing.T) {
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
			if err := ErrorShouldBeLength9(tt.args.obj...); (err != nil) != tt.wantErr {
				t.Errorf("ErrorShouldBeLength9() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
