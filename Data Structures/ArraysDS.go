package main

import "fmt"

func reverseArray(ints []int) []int {
	stni := make([]int, len(ints))
	j := len(ints) - 1
	for i := range ints {
		stni[j] = ints[i]
		j--
	}
	return stni
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ints := make([]int, n)
	for i := range ints {
		fmt.Scanf("%d", &ints[i])
	}
	for _, i := range reverseArray(ints) {
		fmt.Printf("%d ", i)
	}

}
