package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	rocks := make([]string, N)
	for i := range rocks {
		fmt.Scanf("%s", &rocks[i])
	}
	var gems string
	for _, c := range rocks[0] {
		if strings.ContainsRune(gems, c) {
			continue
		}
		contains := true
		for _, rock := range rocks {
			if !strings.ContainsRune(rock, c) {
				contains = false
			}
		}
		if contains {
			gems += string(c)
		}
	}
	fmt.Println(len(gems))
}
