package main

import (
	"fmt"
	"time"
)

func main() {
	var t string
	fmt.Scanf("%s", &t)
	time, _ := time.Parse("3:04:05PM", t)
	fmt.Print(time.Format("15:04:05"))
}
