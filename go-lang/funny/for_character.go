package main

import "fmt"

func main() {
	var str string = "G"
	for i := 1; i <= 25; i++ {
		fmt.Printf("%s\n", str)
		str = str + "G"
	}
}
