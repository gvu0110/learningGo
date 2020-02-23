package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d *Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poolde := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poolde)
	poolde.Speak()

	poolde.Sound = "Arf"
	poolde.Speak()

	poolde.SpeakThreeTimes()
	poolde.SpeakThreeTimes()
}
