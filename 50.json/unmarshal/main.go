package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	// "io"
	// "log"
	// "os"
	// "strings"
)

func main() {
	var jsonBlob = byte(`
	{"Id": 123, "Username": "nikhil1", "Password": "nikhil1", "Email": "email1@domain.com"},
	`)
	type Animal struct {
		Id       int64
		Username string
		Password string
		Email    string
	}
	var animals Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	// Output:
	// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
}
