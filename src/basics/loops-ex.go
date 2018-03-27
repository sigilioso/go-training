package main

import (
	"fmt"
	"math"
)

// Sqrt function
func Sqrt(x float64) float64 {
	z := 1.0
	for prev := 0.0; math.Abs(z-prev) > 0.01; {
		prev = z
		z -= (z*z - x) / 2 * z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
