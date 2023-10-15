package chapter1

import (
	"fmt"
)

var c, python, java bool
var i int = 20

func Var() {
	fmt.Println(i, c, python, java)
}

func ShortVar() {
	name, age, married := "satish", 25, false

	fmt.Printf("name %s of age %d and married %t \n", name, age, married)
}

/*
 In golang their is no undefined as javascript, because it is fail safe or no panic
 Inner var take precedence over outer or global var
*/
