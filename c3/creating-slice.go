package c3

import (
	"fmt"
	"strconv"
)

func showSlice[T comparable](s string, x []T) {
	fmt.Printf("%v len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func MakeSlice() {
	a := make([]int, 5)
	showSlice("a", a)

	b := make([]int, 1, 5)
	showSlice("b", b)

	c := b[:2]
	showSlice("c", c)

	d := c[2:5]
	showSlice("d", d)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "X", "_"},
		{"_", "_", "_"},
	}

	for i, val := range board {
		showSlice(strconv.Itoa(i), val)
	}
	fmt.Println(board)
}

/*
	slice can create using make keyword
	make(type,lenght,cap)
*/
