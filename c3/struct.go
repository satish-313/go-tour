package c3

import (
	"fmt"
)

type vertex struct {
	x int
	y int
}

type job struct {
	name  string
	value string
}

func StructT() {
	fmt.Println(vertex{2, 3})
	name := "satish"
	value := "golang"

	val := job{
		name:  name,
		value: value,
	}

	fmt.Printf("name : %s and value %s\n", val.name, val.value)
	p := &val
	p.value = "golang developer"
	fmt.Printf("name : %s and value %s\n", val.name, val.value)
}

/*
	struct is collection of field (field can be var,function,other struct)
*/
