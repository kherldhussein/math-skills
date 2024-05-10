package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadData(path string) []float64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var data []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatalf("error parsing data %v", err)
		}
		data = append(data, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
