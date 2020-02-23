package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./1_fromString.txt")
	checkError(err)

	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with first file of %v characters\n", ln)

	bytes := []byte(content)

	ioutil.WriteFile("./1_fromByte.txt", bytes, 0644)
	fmt.Println("All done with second file")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
