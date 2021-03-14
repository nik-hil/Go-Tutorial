package main

import (
	"fmt"
	"os"
)

func do(a [3]int) [3]int {
	a[1] = 111
	return a
}

func do1(a []int) int {
	a[1] = 100
	return a[0]
}

func do2(a map[int]int) {
	a[2] = 5
	fmt.Println(a)
}
func do3(a *map[int]int) {
	(*a)[2] = 5
	fmt.Println(*a, a)
}

func main() {
	// defer
	f, err := os.Open("somefile.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "File not found, ", err)
	}
	// closes file at end of "func"
	// always close file. Os can handle only fix no of open file handlers
	defer f.Close()

	// array pass by value
	a := [3]int{1, 2, 3}
	u := do(a)
	fmt.Println(a, u)

	// slice pass by ref
	slice := []int{2, 3, 4, 6}
	fmt.Println(slice, do1(slice), slice)

	// map pass by value
	b := map[int]int{1: 2}
	do2(b)
	fmt.Println(b)

	// map pass by ref
	c := map[int]int{1: 2}
	do3(&c)
	fmt.Println(c)

	fmt.Println("Pass by ref is pass by value, value of descriptor defining slice or map")

}
