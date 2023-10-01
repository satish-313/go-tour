package chapter1

import (
	"fmt"
	"math"
	"math/rand"
)

func RandomNumber() int {
	r1 := rand.Intn(20) // this rand.Int take non-negative [0,n) 0 panic
	r2 := rand.Intn(30)
	fmt.Printf("first rand is %d and second %d \n", r1, r2)
	fmt.Println(math.Sqrt(float64(r1)) * math.Pi)
	return r1 + r2
}
