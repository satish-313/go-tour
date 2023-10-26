package c6

import (
	"fmt"
)

func sum(s []int, c chan int, str string) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println(str, sum)
	// if str == "down" {
	// 	time.Sleep(2000 * time.Millisecond)
	// }
	c <- sum
}

func Ch() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 5) // here 5 is buffer length if more than buffer it block
	// n := len(s) / 2
	// fmt.Println(s[:n], n)
	// fmt.Println(s[n:])
	fmt.Println(cap(c))
	go sum(s[:len(s)/2], c, "top")
	go sum(s[len(s)/2:], c, "down")

	x, y := <-c, <-c // receive from c
	fmt.Println(x, y)
}

/*
	ch <- v // sending to channel
	v := <-ch // receive channel
*/
