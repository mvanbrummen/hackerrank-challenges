package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	str := make([]rune, n)
	for i := range str {
		fmt.Scanf("%c", &str[i])
	}
	move := 0
	for i := 0; i < n-2; i++ {
		if str[i] == '0' && str[i+1] == '1' && str[i+2] == '0' {
			str[i+2] = '1'
			move++
		}
	}
	fmt.Println(move)
}
