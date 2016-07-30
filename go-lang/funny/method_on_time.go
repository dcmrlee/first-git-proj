package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time "anomynous field"
}

func (mt *myTime) first3Chars() string {
	return mt.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	fmt.Println("Full time now:", m.String())
	fmt.Println("First 3 chars:", m.first3Chars())
}
