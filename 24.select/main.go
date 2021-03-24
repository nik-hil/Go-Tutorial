package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{make(chan int), make(chan int)}
	stopper := time.After(5 * time.Second)
	for i := range chans {
		go func(i int, ch chan int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}
	for i := 0; i < 20; i++ {
		select {
		case m0 := <-chans[0]:
			log.Println("Received ", m0)
		case m1 := <-chans[1]:
			log.Println("Received ", m1)
		case <-stopper:
			log.Fatal("timeout")
		}
	}
}

/*
run the program.
It will run till time out occurs.
Select allows us to read from multiple channel in a go.
*/
