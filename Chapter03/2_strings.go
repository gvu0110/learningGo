package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)
	str2 := "An explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str2))

	lvalue := "hello"
	rvalue := "HELLO"
	fmt.Println("Equal?", (lvalue == rvalue))
	fmt.Println("Equal non-case-sensitive?", strings.EqualFold(lvalue, rvalue))

	fmt.Println("Contain expression?", strings.Contains(str1, "exp"))
	fmt.Println("Contain expression?", strings.Contains(str2, "exp"))
}
