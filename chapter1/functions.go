package chapter1

import (
	"fmt"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func devide(x, y int) (int, float64) {
	return x / y, float64(x) / float64(y)
}

func Mathlogic() int {
	r1 := rand.Intn(50)
	r2 := rand.Intn(50)

	fmt.Printf("values %d, %d \n", r1, r2)
	r := add(r1, r2)
	d1, d2 := devide(r1, r2)
	fmt.Printf("int devide %d, float devide %f", d1, d2)
	return r
}
