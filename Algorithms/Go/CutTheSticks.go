package main

import (
	"fmt"
	"sort"
)

func cutSticks1(ints []int) {
	sort.Ints(ints)
	for len(ints) > 0 {
		smallest := ints[0]
		for i := 0; i < len(ints); i++ {
			ints[i] -= smallest
			if ints[i] == 0 {
				ints = ints[i+1:]
				i = -1
			}
		}
		if len(ints) != 0 {
			fmt.Println(len(ints))
		}
	}
}

func cutSticks2(ints []int) {
	sort.Ints(ints)
	for i := 1; i < len(ints); i++ {
		if ints[i] != ints[i-1] {
			fmt.Println(len(ints) - i)
		}
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(n)
	ints := make([]int, n)
	for i := range ints {
		fmt.Scanf("%d", &ints[i])
	}
	cutSticks2(ints)
}
