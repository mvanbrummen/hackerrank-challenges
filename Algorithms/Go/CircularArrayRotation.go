package main

import "fmt"

func rotate(ints []int, k int) {
	l := len(ints)
	k = k % l
	if k == 0 {
		return
	}
	tmp := make([]int, l)
	for i := 0; i < l; i++ {
		for i+k >= l {
			k -= l
		}
		tmp[i+k] = ints[i]
	}
	copy(ints, tmp)
}

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
	rotate(ints, k)
	for _, query := range queries {
		fmt.Printf("%d\n", ints[query])
	}
}
