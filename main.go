package main

import (
	"adventofcode/day1"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	day1.Solve(file)
}
