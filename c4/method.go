package c4

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

type myflaot float64

func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (m myflaot) abs() float64 {
	if m < 0 {
		return -float64(m)
	}
	return float64(m)
}

func Methods() {
	v := vertex{3, 4}
	m := myflaot(-5)
	fmt.Println(v.abs())
	fmt.Println(m.abs())
}

/*
	Keep in mind method are just a function that have special
	receiver function for handle struct/type of function.
	receiver can be applied in type and struct

	----- Note -----
	You can't define a method to receiver type is present in another package
	like (int, which is built in types)
*/
