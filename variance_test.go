package main

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{name: "Multiple elements", args: args{[]float64{
			189,
			113,
			121,
			114,
			145,
			110,
		}}, expected: 784},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.x); got != tt.expected {
				t.Errorf("Variance() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
