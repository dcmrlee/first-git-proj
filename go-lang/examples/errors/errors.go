package main

import (
    "fmt"
    "errors"
    "runtime"
)

func f1(arg int) (int, error) {
    if arg == 2 {
        return -1, errors.New("can't work with 2")
    }
    return arg + 1, nil
}

type argError struct {
    fname string
    lineno int
    arg int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("[%s: %d] %d - %s", e.fname, e.lineno, e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 2 {
        _, file, line, ok := runtime.Caller(0)
        if ok {
            return -1, &argError{file, line, arg, "can't work with it"}
        } else {
        }
        return -1, &argError{arg: arg, prob: "can't work with it"}
    }
    return arg + 1, nil
}

func main() {
    for _, i := range []int{7, 2} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed", e)
        } else {
            fmt.Println("f1 worked", r)
        }
    }

    for _, i := range []int{7, 2} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed", e)
        } else {
            fmt.Println("f2 worked", r)
        }
    }

    _, e := f2(2)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.fname)
        fmt.Println(ae.lineno)
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
