package c6

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func say1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(101 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func Go() {
	go say1("world")
	say("hello")
}
