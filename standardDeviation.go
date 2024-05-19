package main

// The standard deviation is a measure of the amount of variation of a random variable expected about its mean.
func StandardDiviation(x []float64) int {
	if len(x) == 0 {
		return 0
	}
	v := Variance(x)
	g := float64(v) / 2.0
	for i := 0; i < 10; i++ {
		g = (g + float64(v)/g) / 2.0
	}
	return int(g)
}
