package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go was launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	fmt.Printf("The month is %s\n", t.Month())
	fmt.Printf("The day is %d\n", t.Day())
	fmt.Printf("The weekday is %s\n", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, Jan 2, 2006"
	fmt.Printf("Tomorrow is %s\n", tomorrow.Format(longFormat))
	shortFormat := "01/02/2006"
	fmt.Printf("Tomorrow is %s\n", tomorrow.Format(shortFormat))

}
