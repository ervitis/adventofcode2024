package main

import (
	"fmt"
	"sort"

	"github.com/ervitis/adventofcode/utils"
)

func main() {
	lines := utils.ReaderInputToPairInts("./input")
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i, line := range lines {
		left[i] = line[0]
		right[i] = line[1]
	}
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		d := left[i] - right[i]
		if d < 0 {
			d *= -1
		}
		sum += d
	}
	fmt.Println(sum)

	setRight := make(map[int]int)
	for _, v := range right {
		if _, ok := setRight[v]; !ok {
			setRight[v] = 1
		} else {
			setRight[v] += 1
		}
	}
	sum2 := 0
	for _, v := range left {
		if k, ok := setRight[v]; ok {
			sum2 += (v * k)
		}
	}
	fmt.Println(sum2)
}
