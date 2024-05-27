package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("usage: go run . data.txt")
		return
	}

	file := os.Args[1]
	data := ReadData(file)

	mean := Average(data)
	median := Median(data)
	variance := Variance(data, mean)
	sd := math.Sqrt(variance)

	fmt.Printf("Average: %v\n", math.Round(mean))
	fmt.Printf("Median: %v\n", math.Round(median))
	fmt.Printf("Variance: %v\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %v\n", math.Round(sd))
}
