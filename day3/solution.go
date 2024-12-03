package day3

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func Solve(reader io.Reader, writer io.Writer) {
	sum := 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
		matches := r.FindAllStringSubmatch(line, -1)
		for _, v := range matches {
			a, _ := strconv.Atoi(v[1])
			b, _ := strconv.Atoi(v[2])
			sum += a * b
		}
	}

	fmt.Fprint(writer, sum)
}
