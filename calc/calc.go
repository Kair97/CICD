package calc

import "errors"

// ErrDivByZero is returned when dividing by zero.
var ErrDivByZero = errors.New("division by zero")

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// Div divides a by b, returning an error if b is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}
