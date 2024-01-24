package main

import "fmt"

func main() {
	fName := "Arnold"
	lName := "Adero"
	birthYear := 1988
	currentYear := 2024
	age := currentYear - birthYear

	// Create variable fName and lName

	fmt.Println(fName, lName)

	// Create variable date of birth

	fmt.Println(birthYear)

	// create variable current year

	fmt.Println(currentYear)

	// create variable age

	fmt.Println(age)

	currentYear = currentYear + 1
	age = currentYear - birthYear
	fmt.Println(currentYear)
	fmt.Println(age)
}
