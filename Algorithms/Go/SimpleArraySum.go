package main

import "fmt"

func sumIntSlice(ints []int) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ints := make([]int, n)
	for i, _ := range ints {
		fmt.Scanf("%d", &ints[i])
	}
	fmt.Printf("%d", sumIntSlice(ints))
}
