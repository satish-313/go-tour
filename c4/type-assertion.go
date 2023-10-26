package c4

import "fmt"

func Assert() {
	var i interface{} = "hello"

	s, ok := i.(string)
	fmt.Println(s, ok)

	num, ok := i.(float64)
	fmt.Println(num, ok) // nil type of float64

	n, _ := i.(float64)
	fmt.Println(n)

	// n := i.(float64)
	// fmt.Println(n) panic will occourse
}

/*
	Type assertion in interface is check it has that type or not
	if present it will give two varible value , ok
	if ok is true value present or it will panic if try to access

*/
