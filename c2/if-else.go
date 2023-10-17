package c2

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// we can use v in else or else if condition
		fmt.Printf("%g >= %g \n", v, lim)
	}

	return lim
}

func IfElse() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
