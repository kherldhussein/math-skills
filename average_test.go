package main

import "testing"

func TestAverage(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{name: "Decimals", args: args{x: []float64{1.5, 2.5, 3.5}}, expected: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.x); got != tt.expected {
				t.Errorf("Average() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
