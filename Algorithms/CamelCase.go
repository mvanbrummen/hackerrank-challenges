package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)
	n := 1
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			n++
		}
	}
	fmt.Println(n)
}
