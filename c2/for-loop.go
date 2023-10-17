package c2

import (
	"fmt"
)

func FoorLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	wSum := 1

	// we omit intial & termination condition it become while loop similar to
	// another language

	for wSum < 1000 {
		wSum += wSum
	}

	/*
		if no condition it become infinity solution
		kind of couter intutive golang alway give false or zero value in
		this case it give true case

		for {

		}
	*/

	fmt.Println(sum, wSum)
}
