package main

import "fmt"

func main() {
	var n, k, q int
	fmt.Scanf("%d %d %d", &n, &k, &q)
	ints := make([]int, n)
	for i := range ints {
		fmt.Scanf("%d", &ints[i])
	}
	queries := make([]int, q)
	for i := range queries {
		fmt.Scanf("%d", &queries[i])
	}
	for i := 0; i < k; i++ {
		ints = append(ints[n-1:], ints[:n-1]...)
	}
	for _, query := range queries {
		fmt.Printf("%d\n", ints[query])
	}
}
