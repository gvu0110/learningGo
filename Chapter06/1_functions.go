package main

import "fmt"

func main() {
	doSomething()

	sum := addValue(23, 54)
	fmt.Println(sum)

	sum = addAllValues(12, 54, 79)
	fmt.Println(sum)
}

func doSomething() {
	fmt.Println("Doing something!")
}

func addValue(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	fmt.Printf("%T\n", values)
	for i := range values {
		sum += values[i]
	}
	return sum
}
