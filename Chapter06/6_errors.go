package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("My error string")
	fmt.Printf("%T\n", myError)
	fmt.Println(myError)
	attendance := map[string]bool{
		"Ann":  true,
		"Mike": true}
	attend, ok := attendance["Mike"]
	if ok {
		fmt.Println("Mike attended?", attend)
	} else {
		fmt.Println("No info for Mike")
	}
}
