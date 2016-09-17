package main

import "fmt"

// f(n) = f(n-1) + f(n-2)
// f(0) = 0
// f(1) = 1
func fibonacci() func() int {
	x1 := 1 // f(n-1)
	x2 := 0 // f(n-2)
	return func() int {
		ret := x2
		x1, x2 = x1+x2, x1
		return ret
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
