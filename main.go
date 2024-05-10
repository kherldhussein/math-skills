package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Print("usage: go run . data.txt")
		return
	}
	file := os.Args[1]
	data := ReadData(file)
	fmt.Println(data)
}
