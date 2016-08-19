package main

import "fmt"

func sumArray(ints []uint64) (sum uint64) {
	for _, i := range ints {
		sum += i
	}
	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ints := make([]uint64, n)
	for i, _ := range ints {
		fmt.Scanf("%d", &ints[i])
	}
	fmt.Print(sumArray(ints))
}
