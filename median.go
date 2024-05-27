package main

// The median of a set of numbers is the value separating the higher half from the lower half of a data sample, a population, or a probability distribution.
func Median(x []float64) float64 {
	l := len(x)
	y := true
	if l == 0 {
		return 0
	}
	// Bubble Sort Algorithm
	for y {
		y = false
		for i := 0; i < l-1; i++ {
			if x[i] > x[i+1] {
				x[i], x[i+1] = x[i+1], x[i]
				y = true
			}
		}
	}
	idx := l / 2
	if l%2 == 0 {
		return (x[idx-1] + x[idx]) / 2.0
	} else {
		return x[idx]
	}
}
