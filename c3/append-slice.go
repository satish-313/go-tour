package c3

import "fmt"

func AppendSlice() {
	var s []int
	fmt.Println(s, s == nil)

	s = append(s, 1)

	fmt.Println(s)

	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)

	pow := make([]int, 4)
	for i := range pow {
		pow[i] = i << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
