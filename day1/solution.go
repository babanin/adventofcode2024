package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve(file *os.File) {
	var a []int
	var b []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, "   ")

		ai, _ := strconv.Atoi(res[0])
		a = append(a, ai)

		bi, _ := strconv.Atoi(res[1])
		b = append(b, bi)
	}

	sort.Ints(a)
	sort.Ints(b)

	res := uint64(0)
	for i := 0; i < len(a); i++ {
		diff := a[i] - b[i]
		if diff < 0 {
			diff = -diff
		}

		res = res + uint64(diff)
	}

	fmt.Println(res)
}
