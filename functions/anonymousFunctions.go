package main
import ("fmt")

func main () {
	/*
		* we can create a function without a function name
	*/


	anonymous := func (){
		fmt.Println("Function without a name")
	}
	anonymous()



	// unonymous function with parameters and returns a value
	area := func(length, width int) int {
		return length * width
	}
	fmt.Println("The area of a ractangle is: ", area(3, 6))


	// unonymous function as an argument to other functions
	sum := func(number2 int, number3 int) int {
		return number2 * number3
	}

	result := findSquare(sum(4, 7))
	fmt.Println("Result is: ", result)


	// return anonymous function in go
	a := displayNumber()
	fmt.Println(a())

}

func findSquare( num3 int) int {
	square := num3 * num3
	return square
}


// function that return an anonymous function
func displayNumber() func() int {
	numbers := 10
	return func() int{
		numbers ++
		return numbers
	}
}