package main
import "fmt"

func main(){
	/*
		- Go uses data types to determin the type of data associeted with a variable
		- 
	*/

	//  DAta TYPES in GO

	/*
		-	int - (intenger numbers) - 123
		-	float - (decimal point numbers)- 1.30
		-	string - 'mike'
		-	complex (complex numbers) - 2+4Ii, -9.7+18.3i
		-	bool (true, false)
		-	byte - (a byte (8bits) of non negative intengers) - 2 , 115 , 97
		-	rune -( used for charactors internally used as 32-bit intengers ) - 'a', '7', '<'
	*/


	//  Intenger Data type
		var age int
		age = 25
		fmt.Printf("Intenger: %d",age)

	// Float Data Type
		var balance float32
		balance = 29.67
		fmt.Printf("float32: %g", balance)

		var price float64
		price = 787878.07
		fmt.Printf("float64: %g", price)

		/*
			float have 2 sizes which include
				- float32
				- float64
		*/

	// String Data Type
		var car string
		car = "BMW"
		fmt.Printf("string: %s", car)


	// Boolean Data Type
		var isValid bool
		isValid = false
		fmt.Printf("bool: %t", isValid)
}