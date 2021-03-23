package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextId = make(chan int)

func handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "Received %v\n", <-nextId)
}
func counter() {
	for i := 0; ; i++ {
		nextId <- i
	}
}

func main() {
	go counter()
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
