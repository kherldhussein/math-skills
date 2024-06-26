package main

// Average is a single number or value that best represents a set of data.
func Average(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	var total float64
	for _, x1 := range x {
		total += x1
	}
	return total / float64(len(x))
}
