package c4

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type t struct {
	s string
}

func (t *t) M() {
	fmt.Println(t)
}
func (f myflaot) M() {
	fmt.Println(f)
}

func InterfaceValue() {
	var i I

	i = &t{"hello"}

	descibe(i)
	i.M()

	i = myflaot(math.Pi)
	descibe(i)
	i.M()
}

func descibe(i I) {
	fmt.Printf("%v , %T\n", i, i)
}
