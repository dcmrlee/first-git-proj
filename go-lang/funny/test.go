package main

import "fmt"
import structPack "github.com/dcmrlee/first-git-proj/go-lang/funny/structPack"

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
	t([]int{})
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.
	fmt.Printf("%v\n", struct1)
}

func t(n []int) int {
	a := len(n)
	fmt.Printf("%v\n", a)
	return a
}
