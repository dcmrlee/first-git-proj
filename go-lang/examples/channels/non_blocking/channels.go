package main

import "fmt"

func main() {
    // unbuffer chan need first set receive edge, then send something to it
    // buffered chan can send something to it firstly
    // so, if the channel 'messages' changed to buffered chan as: messages := make(chan string, 1)
    // the output should like:
    // 'no message received'
    // 'sent message hi'
    // 'received message hi'
    messages := make(chan string)
    signals := make(chan bool)
    

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
