package c4

import (
	"fmt"
)

type nilI interface {
	M()
}

type strT struct {
	s string
}

func (t *strT) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t)
}

func NilInterface() {
	var i nilI
	descibe(i)
	var t *strT

	i = t
	descibe(i)
	i.M()

	i = &strT{"satish"}
	descibe(i)
	i.M()
}
