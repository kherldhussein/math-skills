package main

func Average(x []float64) int {
	if len(x) == 0 {
		return 0
	}
	var total float64
	for _, x1 := range x {
		total += x1
	}
	return int(total / float64(len(x)))
}
