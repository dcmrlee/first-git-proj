package main

import "log"
import "time"
import "net/http"
import _ "net/http/pprof"

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
