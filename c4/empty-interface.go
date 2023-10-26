package c4

import "fmt"

func EmptyInterface() {
	var i interface{}
	descibe2(i) // nil nil

	i = 43
	descibe2(i) //43 int

	i = "satish"
	descibe2(i) // satish string
}

func descibe2(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}

/*
	Empty interface is no or zero method in interface.
	Use of empty interface is that it can hold any value therefore
	it used for handle unknown type.

	It simillar to generic in other language
*/
