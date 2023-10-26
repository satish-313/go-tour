package c4

import (
	"fmt"
	"math"
)

type abser interface {
	abs() float64
}

func Interface() {
	var a abser

	f := myflaot(-math.Sqrt2)
	v := vertex{3, 4}
	p := point{6, 8}

	a = f
	fmt.Println(a.abs())

	a = v
	fmt.Println(a.abs())

	a = &p
	fmt.Println(a.abs())
}

/*
	Interface always value assing not pointer assing
	to get value pointer we have to & infront of pointer

	Interface are implicit type
*/
