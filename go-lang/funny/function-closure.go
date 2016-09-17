package main

import (
	"fmt"
)

func closure_func() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	test_closure := closure_func()
	for i := 0; i < 10; i++ {
		fmt.Println(test_closure(i))
	}
}
