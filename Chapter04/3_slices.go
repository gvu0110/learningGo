package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Pink")
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 9
	numbers[1] = 3
	numbers[2] = 1
	numbers[3] = 5
	numbers[4] = 7
	fmt.Println(numbers)

	numbers = append(numbers, 13)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)
}
