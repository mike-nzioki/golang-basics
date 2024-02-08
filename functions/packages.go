package main

import (
	"fmt"
	"calculator"
)

func main() {
	/*
		* packages
		- package is a container that conains various functions to perform a specific task
		- for instance math package includes the sqrt() function to perform a squareroot of a number
		- we start each project with package main  to tell the compiler that it is an executable code
		- we use import keyword to import packges eg import( "fmt" "math" "string")
	*/

	//  # FMT PACKAGE
	/*
		* FMT
		Print()	prints the text to output screen
		Println()	prints the text to output with a new line character at the end
		Printf()	prints the formatted string to the output screen
		Scan()	get input values from the user
		Scanf()	get input values using the format specifier
		Scanln()	get input values until the new line is detected
	*/

	// # MATH
	/*
		* math
		Sqrt()	returns the square root of the number
		Cbrt()	returns the cube root of the number
		Max()	returns the larger number between two
		Min()	returns the smaller number between two
		Mod()	computes the remainder after division

	*/

	//  # STRINGS
	/*
		* strings

		Functions	Descriptions
		Compare()	checks if two strings are equal
		Contains()	checks if the string contains a substring
		Count()	counts the number of times a substring present in the string
		Join()	creates a new string by concatenating elements of a string array
		ToLower()	converts the string to lowercase
		ToUpper()	converts the string to uppercase
	*/

	// CREATE CUSTOME PACKAGE
	//  - create new file and declear package
	//  - import the package at the top of your file (calculator package)
	// - use package

	number := 3
	number1 := 5

	fmt.Println("Add:", calculator.Add(number, number1))
	fmt.Println("Subtract:", calculator.Subtract(number1, number))
}
