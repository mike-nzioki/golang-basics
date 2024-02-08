package main
import ("fmt")

func main() {
	/*
		* Closure
		- closure is a nested function that allows us to access variables of the outer function even after that outer fucntion is closed
		- 
	*/

	//  NESTED FUNCTIONS
	greet("Mike")



	// RETURNING A FUNCTION IN GO
	response := greetResponse()
	response()


	// CLOSURES
	theRepoonse := casualResponse()
	fmt.Println(theRepoonse())


	// PRINT ODD NUMBERS USING CLOSURE
	odd := calculate()
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())


	/*
		* Closure helps with data isolation
		- form the examples, a new closure is returned whenever we call the outer function
		- each closure is independent from each other and the change on one will not affect the other
		- thus we are able to work with multiple data in isolation from one another
		- examples
	*/

	// return closure
	odd1 := calculate()
	fmt.Println(odd1())
	fmt.Println(odd1())

	// returns new closure
	odd2 := calculate()
	fmt.Println(odd2())
	fmt.Println(odd2())
}

// outter function
func greet(name string) {
	// inner function
	displayName := func() {
		fmt.Println("Hello", name)
	}

	// call inner functionx
	displayName()
}

func greetResponse() func() {
	return func() {
		fmt.Println("Hello, how are You?")
	}
}


// closure
func casualResponse () func() string {
	respo := "I'm good,"
	return func() string {
		fullResponse := respo +" "+"How about you?"
		return fullResponse
	}
}

// ODD NUMBERS
func calculate() func() int {
	num := 1

	return func () int {
		num += 2
		return num
	}
}


