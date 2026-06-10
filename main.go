package main

import (
	"fmt"

	"github.com/Kair97/CICD/calc"
)

func main() {
	fmt.Println("2 + 3 =", calc.Add(2, 3))
	fmt.Println("10 / 2 =", mustDiv(10, 2))
}

func mustDiv(a, b float64) float64 {
	res, err := calc.Div(a, b)
	if err != nil {
		panic(err)
	}
	return res
}
