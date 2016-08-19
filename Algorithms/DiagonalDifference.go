package main

import "fmt"

func diagonalDifference(matrix [][]int) int {
	p, s := 0, 0
	n := len(matrix) - 1
	j := 0
	for i := range matrix {
		p += matrix[i][j]
		s += matrix[i][n]
		j++
		n--
	}
	if p > s {
		return p - s
	} else {
		return s - p
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	rows := make([][]int, n)
	for i := range rows {
		rows[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &rows[i][j])
		}
	}
	fmt.Print(diagonalDifference(rows))
}
