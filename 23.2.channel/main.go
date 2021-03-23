package main

import (
	"fmt"
	"log"
	"net/http"
)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "Received %v\n", <-ch)
}
func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	var nextId nextCh = make(chan int)
	go counter(nextId)
	http.HandleFunc("/", nextId.handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// Run as seq 1 200 | xargs -I $ -n1 -P10  curl "http://localhost:8000/"
