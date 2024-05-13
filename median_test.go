package main

import "testing"

func TestMedian(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{name: "Even elements, decimals", args: args{x: []float64{1.5, 2.5}}, expected: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.x); got != tt.expected {
				t.Errorf("Median() = %v, want %v", got, tt.expected)
			}
		})
	}
}
