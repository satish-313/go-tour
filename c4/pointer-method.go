package c4

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func (p *point) abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func PointerMethod() {
	v := vertex{6, 8}
	v.scale(2)
	fmt.Println(v, v.abs())
	p := point{3, 4}
	fmt.Println(p.abs())
}

/*
	Most of method are pointer based because pointer modify
	original value.
*/
