package main

import (
	"fmt"
	"learningGo/stringutil"
)

func main() {

	n1, l1 := stringutil.FullName("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of chars: %v\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v, number of chars: %v\n", n2, l2)

}
