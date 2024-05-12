package main

func Median(x []float64) int {
	l := len(x)
	y := true
	if l == 0 {
		return 0
	}
	for y {
		y = false
		for i := 0; i < l-1; i++ {
			if x[i] > x[i+1] {
				tmp := x[i]
				x[i] = x[i+1]
				x[i+1] = tmp
				y = true
			}
		}
	}
	idx := l / 2
	if l%2 == 0 {
		return int((x[idx-1] + x[idx]) / 2.0)
	} else {
		return int(x[idx])
	}
}
