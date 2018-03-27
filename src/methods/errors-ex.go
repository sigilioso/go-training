package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt complex values not allowed
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %v", float64(e))
}

// Sqrt function
func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for prev := 0.0; math.Abs(z-prev) > 0.01; {
		prev = z
		z -= (z*z - x) / 2 * z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-1))
}
