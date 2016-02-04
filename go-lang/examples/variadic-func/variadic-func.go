package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func strSum(n int, strs ...string) {
    fmt.Println(n)
    for _, str := range strs {
        fmt.Println(str)
    }
}

func main() {
    sum(1, 2)
    sum(1, 2, 3)
    
    nums := []int{1, 2, 3, 4}
    sum(nums...)

    strSum(5, "li", "da", "chao")
}
