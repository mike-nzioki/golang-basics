package main
import ("fmt")

// declare global variable before main function
var product int

func main () {
	/*
		* Varible scope
		- scope is a specified region where we can access a variable
		- go variables can be decleared in 2 different ways
		- local scope
		- global scope
	*/

	  // cannot access sum out of its local scope
	//   fmt.Println("Sum is", sum)

	// sumVariable is a local veriable
	sumVariable := addNumbers()
	fmt.Println("sumVariable: ", sumVariable)


	// you can access varible product since its globally available
	product = 2 * 2
	fmt.Println("product: ", product)
}


func addNumbers() int {
	// local variables
	sum := 5 + 5
	return sum
}