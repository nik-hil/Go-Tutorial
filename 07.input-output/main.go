package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("On std out")
	fmt.Fprintln(os.Stderr, "On std err")

	s := 10
	a := [3]rune{'a', 'b', 'c'}
	m := map[int]string{1: "1", 2: "2"}

	fmt.Printf("%T \n%v \n%#v \n", s, s, s)
	fmt.Printf("%[1]T \n%[1]v \n%#[1]v \n%[1]q \n", a)
	fmt.Printf("%T \n%v \n%#v \n", m, m, m)

	// execute using $ go run . main.go
	fname := os.Args[1]
	if f, err := os.Open(fname); err != nil {
		fmt.Fprintln(os.Stderr, "bad file", err)
	} else if d, err := ioutil.ReadAll(f); err != nil {
		fmt.Fprintln(os.Stderr, "can't read", err)
	} else {
		fmt.Printf("The file has %d bytes", len(d))
	}
}
