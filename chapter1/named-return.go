package chapter1

/*
 naked return
*/

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func NakedReturn() {
	fmt.Println(split(30))
}
