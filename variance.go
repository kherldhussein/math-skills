package main

func Variance(x []float64) int {
	m := Average(x)
	t := 0
	if len(x) == 0 {
		return 0
	}
	for _, x1 := range x {
		y := int(x1) - m
		t += y * y
	}
	return t / len(x)
}
