package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	letters := make(map[rune]int, 26)
	for _, i := range str {
		if i == ' ' || i == '\n' {
			continue
		}
		i = unicode.ToLower(i)
		letters[i]++
	}
	if len(letters) == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
