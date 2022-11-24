package random

import "testing"

func TestRandZN(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test rand zh", args: args{min: 5, max: 10}},
		{name: "test rand zh2", args: args{min: 2, max: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandZN(tt.args.min, tt.args.max); len(got)/3 > tt.args.max || len(got)/3 < tt.args.min {
				t.Errorf("RandZN() = %v; len: %d", got, len(got)/3)
			} else {
				t.Logf("RandZN() = %v", got)
			}
		})
	}
}
