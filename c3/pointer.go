package c3

import (
	"fmt"
)

func Pointer() {
	i, j := 23, 24

	p := &i

	fmt.Printf("%T \n", p)
	fmt.Println(p)  // addrerss is i = xx10
	fmt.Println(*p) // 23

	p = &j
	fmt.Println(p, *p) // address of var j = xx18
	*p = 78
	fmt.Println(j) // j = 78
}
