package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	clouds := make([]int, n)
	for i := range clouds {
		fmt.Scanf("%d", &clouds[i])
	}
	e := 100
	for i := 0; i < n; i += k {
		if clouds[i%n] == 1 {
			e -= 2
		}
		e--
	}
	fmt.Println(e)
}
