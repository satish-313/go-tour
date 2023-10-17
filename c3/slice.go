package c3

import "fmt"

func Slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:4]
	var p []int = s[1:3]
	fmt.Println(&primes[0], &s[0]) // both have same address
	p[0] = 400
	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(primes)

	k := []int{1, 2, 4} // slice literal
	fmt.Println(k, len(k), cap(k))
}
