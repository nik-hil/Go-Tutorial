package main

// import is specific to each file.
// In same package, import is specific to each file.
// Don't import packages which wont be used.
import "fmt"

// Variable can't use short notation outside the func
// a := 2 is not valid
var a int = 2

// Func starting with capital letter is exported.
func main() {

	myArray := []int{1, 2, 4, 8, 16}

	for i := range myArray {
		fmt.Println(i, myArray[i])
	}
	fmt.Println("Prints copy of each elem of myArray")
	for i, arr := range myArray {
		fmt.Println(i, arr)
	}

	fmt.Println("Dict")

	myMap := map[int]int{1: 2, 2: 3, 3: 4}
	for i := range myMap {
		fmt.Println(i, myMap[i])
	}
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	fmt.Println("Switch")

	switch a := 5; a {
	case 0, 1, 2, 3, 4:
		fmt.Println("less")
	case 5, 6:
		fmt.Println("more")
	default:
		fmt.Println("default")
	}
	fmt.Println("Another Switch")

	a := 8
	switch {
	case a%2 == 0:
		fmt.Println("Even")
	case a%2 == 1:
		fmt.Println("Odd")
	default:
		fmt.Println("Not a Even or Odd number")
	}

}

func init() {
	fmt.Println("Only runtime can call init. It gets called before main.")
	var a int
	fmt.Println("Print a: ", a)
}
