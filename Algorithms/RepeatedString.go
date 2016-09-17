package main

import "fmt"

func main() {
	var (
		str string
		n   int
	)
	fmt.Scanf("%s", &str)
	fmt.Scanf("%d", &n)
	count := 0
	for _, c := range str {
		if c == 'a' {
			count++
		}
	}
	total := n / len(str) * count
	remainder := n % len(str)
	for i, c := range str {
		if i == remainder {
			break
		}
		if c == 'a' {
			total++
		}
	}
	fmt.Println(total)
}
