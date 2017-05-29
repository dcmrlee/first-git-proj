package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	defer conn.Close()
	request := make([]byte, 128)
	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println("connection closed by client")
			fmt.Println(err.Error())
			break
		}
		if readLen == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			fmt.Println("got content[request for timestamp]: ", request)
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
			fmt.Println("write timestamp success...")
		} else {
			fmt.Println("got content: ", request)
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
			fmt.Println("write daytime success...")
		}

		request = make([]byte, 128)
	}
}

func main() {
	service := ":8111"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("got request: ", conn.RemoteAddr().String())
		go handleClient(conn)
	}
}
