package c2

import (
	"fmt"
	"runtime"
	"time"
)

func myOs() {
	fmt.Println("Golang run on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux ")
	default:
		fmt.Printf("%s. \n", os)
	}
}

func weeks() {
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	fmt.Printf("%v = type of time.Now %T\n", today, today)

	fmt.Printf("%T \n", time.Saturday)
	t := time.Now()
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour())

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("Two day letter")
	default:
		fmt.Println("Too far")
	}
}

func Switch() {
	myOs()
	weeks()
}
