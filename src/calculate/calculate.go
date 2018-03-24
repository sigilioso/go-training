package calculate

import (
	"fmt"
	"strconv"
)

// Add sum two floa(ts
func Add(a, b float64) float64 {
	return a + b
}

// Substract subtract b from a
func Substract(a, b float64) float64 {
	return a - b
}

// Multiply multiply a and b
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide divide a and b
func Divide(a, b float64) float64 {
	return a / b
}

// round floats to 0.xx precision
func round(f float64) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
}
