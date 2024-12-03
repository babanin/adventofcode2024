package main

import (
	"adventofcode/day3"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	day3.Solve(file, os.Stdout)
}
