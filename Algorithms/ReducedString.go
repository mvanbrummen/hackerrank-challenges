package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			str = str[:i-1] + str[i+1:]
			i = 0
		}
	}
	if len(str) == 0 {
		str = "Empty String"
	}
	fmt.Println(str)
}
