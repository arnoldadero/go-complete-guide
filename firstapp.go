package main

import "fmt"

func main() {
	// create var greetingText
	var greetingText string = "Hello World!"

	greetingText2 := "Hello Planet Earth!"

	// create var Int leap year days
	leapYearDays := 366

	// create var Int non leap year days
	nonLeapYearDays := leapYearDays - 1

	// create var float64  nonLeapYeaDays
	nonLeapYearDaysFloat := float64(nonLeapYearDays) / 3
	fmt.Println(nonLeapYearDaysFloat)

	// create var float64 leapYearDays
	leapYearDaysFloat := float64(leapYearDays) / 3
	fmt.Println(leapYearDaysFloat)


	// print greetingText2
	fmt.Println(greetingText2)

	// print "Hello World!"
	fmt.Println("Hello World!")
	// print nonLeapYearDays
	fmt.Println(nonLeapYearDays)

	// print leapYearDays
	fmt.Println(leapYearDays)

	// print greetingText
	fmt.Println(greetingText)

}
