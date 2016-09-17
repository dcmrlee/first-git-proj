package main

import (
	"fmt"
	"time"
)

type myError struct {
	when time.Time
	what string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &myError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
