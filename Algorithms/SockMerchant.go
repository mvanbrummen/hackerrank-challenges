package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	colours := make(map[int]int)
	for i := 0; i < n; i++ {
		var c int
		fmt.Scanf("%d", &c)
		colours[c]++
	}
	count := 0
	for c := range colours {
		count += colours[c] / 2
	}
	fmt.Println(count)
}
