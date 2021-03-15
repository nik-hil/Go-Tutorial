package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string
	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

type Point struct {
	X, Y float64
}

func (p Point) offset(x, y float64) Point {
	return Point{p.X + x, p.Y + y}
}
func (p *Point) Move(x, y float64) {
	p.X += x
	p.Y += y
}

func main() {
	var v IntSlice = []int{1, 2, 4}
	var s fmt.Stringer = v

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}
	fmt.Printf("%T %[1]v\n", s)
	fmt.Printf("%T %[1]v\n", v)

	p := Point{
		X: 0,
		Y: 0,
	}
	q := p.offset(2.0, 2.0)
	fmt.Println(p, q)
	p.Move(3, 3)
	fmt.Println(p)

}
