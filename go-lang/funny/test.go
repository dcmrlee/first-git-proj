package main

import "fmt"

func main() {
	var s string = "abcdef"
	var b []byte
	b = []byte(s)
	fmt.Printf("%v\n", b)
	s1 := s[2:]
	fmt.Printf("%v\n", s1)
	fmt.Printf("%d\n", len(s1))
	fmt.Printf("%v\n", s[1:4])
	fmt.Printf("%v\n", s[1:2])
}
