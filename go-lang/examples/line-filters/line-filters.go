package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
