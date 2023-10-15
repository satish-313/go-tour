package chapter1

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func TypeOf() {
	fmt.Printf("Type : %T value : %v \n", ToBe, ToBe)
	fmt.Printf("Type : %T value : %v \n", MaxInt, MaxInt)
	fmt.Printf("Type : %T value : %v \n", z, z)
}

/*
	In golang their is 7 basic or primative type
	bool,string,int(int8,int16,int32,int64,uint,uintprt),
	byte - uint8
	rune - uint32
	float - float32,float64
	complex number - complex64,complex128
*/
