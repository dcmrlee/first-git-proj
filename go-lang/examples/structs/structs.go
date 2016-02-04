package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{"Bob", 20})
    
    fmt.Println(person{name: "Alice", age: 10})

    fmt.Println(person{name: "Fred"})

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    fmt.Println(sp)
}
