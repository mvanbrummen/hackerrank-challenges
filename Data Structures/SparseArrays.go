package main

import "fmt"

func main() {
	var N, Q int
	fmt.Scanf("%d", &N)
	strings := make([]string, N)
	for i := range strings {
		fmt.Scanf("%s", &strings[i])
	}
	fmt.Scanf("%d", &Q)
	queries := make([]string, Q)
	for i := range queries {
		fmt.Scanf("%s", &queries[i])
	}
	for _, query := range queries {
		count := 0
		for _, str := range strings {
			if str == query {
				count++
			}
		}
		fmt.Println(count)
	}
}
