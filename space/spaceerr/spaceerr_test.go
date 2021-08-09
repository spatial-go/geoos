package spaceerr

import "testing"

func TestErrorUsageFunc(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test error", args{[]interface{}{"I am ErrorUsageFunc("}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrorUsageFunc(tt.args.obj...); (err != nil) != tt.wantErr {
				t.Errorf("ErrorUsageFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
