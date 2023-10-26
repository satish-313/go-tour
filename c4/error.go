package c4

import (
	"fmt"
	"time"
)

type myError struct {
	when time.Time
	what string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s\n", e.when, e.what)
}

func run() error {
	return &myError{time.Now(), "it didn't work"}
}

func Err() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
