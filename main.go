package main

import (
	"adventofcode/day2"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	day2.Solve(file, os.Stdout)
}
