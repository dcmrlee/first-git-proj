package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("timestamp"))
	checkError(err)
	fmt.Println("write success...")
	resp := make([]byte, 128)
	respLen, err := conn.Read(resp)
	checkError(err)
	fmt.Println("read response: ", string(resp[:respLen]))
	conn.Close()
	fmt.Println("close connection")
	os.Exit(0)
}
