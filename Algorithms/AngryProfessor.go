package main

import "fmt"

func main() {
	var n, k, t, a int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &k)
		attendees := 0
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &a)
			if a <= 0 {
				attendees++
			}
		}
		if attendees < k {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
