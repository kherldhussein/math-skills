package main

// Variance is the expected value of the squared deviation from the mean of a random variable.
func Variance(x []float64, m float64) float64 {
	t := 0.0
	if len(x) == 0 {
		return 0
	}
	for _, x1 := range x {
		y := x1 - m
		t += y * y
	}
	return t / float64(len(x))
}
