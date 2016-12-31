package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	str := ""
	for i := 0; i < n; i++ {
		str += "#"
		fmt.Printf(fmt.Sprintf("%%%ds\n", n), str)
	}
}
