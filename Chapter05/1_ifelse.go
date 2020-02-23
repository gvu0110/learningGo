package main

import "fmt"

func main() {

	//var x float64 = 42
	var result string

	if x := 42; x < 0 {
		result = "Less than 0"
	} else if x == 0 {
		result = "Equal to 0"
	} else {
		result = "Greater than 0"
	}

	fmt.Println(result)
	//fmt.Println("Value of x:", x)

}
