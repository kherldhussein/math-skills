package main

import "testing"

func TestStandardDiviation(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{name: "Known SD", args: args{[]float64{
			189,
			113,
			121,
			114,
			145,
			110,
		}}, expected: 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDiviation(tt.args.x); got != tt.expected {
				t.Errorf("StandardDiviation() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
