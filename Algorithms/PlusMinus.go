package main

import "fmt"

func plusMinus(floats []float32) (float32, float32, float32) {
	n := float32(len(floats))
	var pTotal, nTotal, zTotal float32
	for _, i := range floats {
		if i > 0 {
			pTotal++
		} else if i < 0 {
			nTotal++
		} else {
			zTotal++
		}
	}
	return pTotal / n, nTotal / n, zTotal / n
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	floats := make([]float32, n)
	for i := range floats {
		fmt.Scanf("%f", &floats[i])
	}
	positive, negative, zero := plusMinus(floats)
	fmt.Printf("%f\n%f\n%f\n", positive, negative, zero)
}
