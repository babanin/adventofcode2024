package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	Safe   = "safe"
	Unsafe = "unsafe"
)

func Solve(reader io.Reader, writer io.Writer) {
	safe := 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " ")

		if isSafe(res) {
			safe += 1
			fmt.Fprintln(writer, Safe)
		} else {
			fmt.Fprintln(writer, Unsafe)
		}
	}

	fmt.Fprint(writer, safe)
}

func isSafe(res []string) bool {
	inc := true
	nums := make([]int, len(res))
	for i := range res {
		nums[i], _ = strconv.Atoi(res[i])

		if i == 1 {
			if nums[i] > nums[i-1] {
				inc = true
			} else if nums[i] < nums[i-1] {
				inc = false
			} else {
				return false
			}
		}

		if i > 0 {
			diff := nums[i] - nums[i-1]

			if diff < 0 {
				if inc {
					return false
				}
			} else if diff > 0 {
				if !inc {
					return false
				}
			}

			if diff < 0 {
				diff = -diff
			}

			if diff < 1 || diff > 3 {
				return false
			}
		}
	}

	return true
}
