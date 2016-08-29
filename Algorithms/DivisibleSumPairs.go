package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	ints := make([]int, n)
	for i := range ints {
		fmt.Scanf("%d", &ints[i])
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (ints[i]+ints[j])%k == 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
