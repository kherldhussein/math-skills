package main

import (
	"fmt"
	"log"
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
	variance := Variance(data)
	sd := StandardDiviation(data)

	fmt.Printf("Average: %v\n", mean)
	fmt.Printf("Median: %v\n", median)
	fmt.Printf("Variance: %v\n", variance)
	fmt.Printf("Standard Deviation: %v\n", sd)
}
