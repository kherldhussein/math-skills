package main

import (
	"math"
	"testing"
)

func TestVariance(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		mean     float64
		expected float64
		epsilon  float64
	}{
		{"Element", []float64{5}, 5, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Variance(test.input, test.mean)
			if math.Abs(result-test.expected) > test.epsilon {
				t.Errorf("Test case %s failed: expected %f, got %f", test.name, test.expected, result)
			}
		})
	}
}
