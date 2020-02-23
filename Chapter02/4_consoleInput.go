package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	str, _ := reader.ReadString('\n') // the sigle quotes designated as byte, rather than as a string
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // convert a string to a float

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of the input number:", number)
	}
}
