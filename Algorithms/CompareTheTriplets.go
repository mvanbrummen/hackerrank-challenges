package main

import "fmt"

func compareTriplets(a, b []int) (resultA, resultB int) {
	for i, _ := range a {
		if a[i] == b[i] {
			continue
		} else if a[i] > b[i] {
			resultA++
		} else {
			resultB++
		}
	}
	return resultA, resultB
}

func main() {
	a := make([]int, 3)
	b := make([]int, 3)
	for i, _ := range a {
		fmt.Scanf("%d", &a[i])
	}
	for j, _ := range b {
		fmt.Scanf("%d", &b[j])
	}
	resultA, resultB := compareTriplets(a, b)
	fmt.Printf("%d %d", resultA, resultB)
}
