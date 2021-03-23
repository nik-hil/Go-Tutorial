package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		lapse := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, lapse}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result)
	list := []string{}
	for i := 0; i < 2400; i++ {
		url := fmt.Sprintf("http://xkcd.com/%d/", i)
		list = append(list, url)
	}
	fmt.Println("Total urls are ", len(list))
	for _, url := range list {
		go get(url, results)
	}
	for range list {
		r := <-results
		if r.err != nil {
			log.Printf("%-20s %s \n", r.url, r.err)
		} else {
			log.Printf("%-20s %s \n", r.url, r.latency)
		}
	}
}
