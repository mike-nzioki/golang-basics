package main
import ("fmt")

func main () {
	/*
		* Recursion
		-  a recursive function calls itself
		- syntax
		func recurse(){
			// code
			recurse()
		}
	*/

	// call recurseve function
	countDown(3)

	// call sum
	result := sum(5)
	fmt.Println("suResult: ", result)

	// call factorial
	factorialResults := factorial(4)
	fmt.Println("factorial: ", factorialResults)

}

func countDown(num int) {

	if num > 0 {
		fmt.Println("CountDown Starts")
		fmt.Println(num)
		countDown( num - 1)
	} else {
		fmt.Println("CountDown Stops")
	}

}


// return sum of positive numbers
func sum (num1 int) int {
	if num1 == 0 {
		return 0
	} else {
		return num1 + sum(num1 - 1)
	}
}


// factorial of number using go recursion
func factorial (num2 int) int {
	if num2 == 0 {
		return 1
	} else {
		return num2 * factorial(num2 - 1)
	}
}