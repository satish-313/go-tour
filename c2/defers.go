package c2

import (
	"fmt"
)

func ex1() {
	i := 0
	defer fmt.Println(i)
	i++
	/*
		output will 0, defer take previous value instead of letter but
		function execute latter then increament i
	*/
}

func exe2() (i int) {
	defer fmt.Println("before ", i)
	defer func() {
		i++
		fmt.Println("inside ", i)
	}()
	defer fmt.Println("after ", i)
	return 1
	/*
		Naked return take return and do function after it
	*/
}

func Defer() {
	defer fmt.Println("world")
	fmt.Println("Hello ")
	ex1()
	fmt.Println(exe2())
	fmt.Println("Counting")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

/*
when defer function calls, it didn't execute immediately push to stack.
after all surrounding function execute. It create goroutine to keep track.
In stack last in first out way execute all defer
*/
