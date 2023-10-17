package c3

import (
	"fmt"
)

func change(arr *[2]string) {
	arr[0] = "ribal"
}

func Arr() {
	var a [2]string

	a[0] = "satish"
	a[1] = "pradhan"
	fmt.Println(a[0], a[1])
	change(&a)
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11}
	primes[5] = 13
	fmt.Println(primes)
}

/*
	similar to other language like c,c++
*/
