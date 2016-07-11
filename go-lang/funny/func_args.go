package main

import "fmt"

func main() {
	callback(1, myFunc)
}

func myFunc(a int) int {
	return (2*a + 1)
}

func callback(a int, f func(int) int) {
	fmt.Printf("a = %d, 2a+1 = %d\n", a, f(a))
}
