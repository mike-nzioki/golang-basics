package main
import (
	"fmt"
	"math"
)

func main () {
	/*
		* FUNCTIONS
		- functions are used to divide code to to smaller chunks to make code look clean and easy to understand
		- a block of code that performs a specific task.
		- syntax;
		func functionName(){
			// code
		}
	*/
	//  # calling function greeting
	greeting()

	// #calling function addNumbers
	addNumbers()

	// # call divideNumbers function
	divideNumbers(21, 7)

	// # call multiplyNumbers
	productResult := multiplyNumbers(23, 4)
	fmt.Println("productResult: ", productResult)


	// c# call areaAndLength
	area, diagonal := areaAndLength(12, 5)
	fmt.Println("area: ", area, "Diagonal: ", diagonal)
}

// greeting function
func greeting(){
	fmt.Println("Good Morning")
}

// addNumbers function
func addNumbers() {
	n1 := 12
	n2 := 5
	sum := n1 + n2
	fmt.Println("Sum: ", sum)
}


// divide numbers
func divideNumbers(num1 int, num2 int){
	result := num1 / num2
	fmt.Println("result: ", result)
}


// return value
func multiplyNumbers(number1 int, number2 int) int {
	product := number1 * number2
	return product
}


// return multiple values
func areaAndLength(hight int, width int) (int, float64) {
	area := hight * width
	diagonal := math.Sqrt(float64((hight * hight) + (width * width)))
	return area, diagonal
}
