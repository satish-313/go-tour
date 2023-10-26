package c5

import "fmt"

type list[T any] struct {
	next *list[T]
	val  T
}

func Gen() {
	si := []int{10, 20, 30, -20, 15}
	fmt.Println(index(si, 15))

	ss := []string{"foo", "bar", "art"}
	fmt.Println(index(ss, "bar"))
}

func index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}
