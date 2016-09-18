package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)
	var str string
	for i := 0; i < T; i++ {
		var count int
		fmt.Scanf("%s", &str)
		for j := 1; j <= len(str)-1; j++ {
			if str[j] == str[j-1] {
				count++
			}
		}
		fmt.Println(count)
	}
}
