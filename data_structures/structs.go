package main
import "fmt"

func main () {
	/*
		* Struct
		- struct is used to store variables of different data types
		- the syntax
		type StructureName struct [
			// struct difinition
		]
	*/

	type Person struct {
		name string
		age int
	}

	//  # STRUCT INSTANCES
	student := Person{"Luke", 25}
	fmt.Println(student)


	// # accesing a struct in go
	studentName := student.name
	fmt.Println("Name: ", studentName)

	studentAge := student.age
	fmt.Println("Age: ", studentAge)


	// # FUNCTIONS INSIDE A STRUCT IN GO
	/*
		* go allows us to create a function inside a struct
		- treats function as a field of struct
		- eg
	*/

	// initialize the function
	type Area func (int, int) int

	// create the structure
	type Rectangular struct {
		length int
		width int
		color string

		// function s a field of struct (assign field area to an istance of function Area)
		area Area
	}

	// ASSiGN VALUES TO STRUCT VARIABLES
	result := Rectangular{
		length: 20,
		width: 30,
		color: "blue",
		area: func(length int, width int) int {
			return length * width
		},
	}

	fmt.Println("Color of Rectangular: ", result.color)
	fmt.Println("Area of  Rectangular shape: ", result.area(result.length, result.width))
}