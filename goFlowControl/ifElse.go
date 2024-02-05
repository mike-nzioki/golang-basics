package main
import "fmt"

func main(){
	// IF ELSE
	/*
		- we use if statements to run a block of code only when a criteria is met
		- in go we use if keyword follwed by test condition then curly braces
		-
	*/
	age := 20
	if age > 5 {
		// code
		fmt.Println(age)
	}


	// create a program that displays a number if its positive

	number1 := 16
	if number1 > 0{
		fmt.Printf("%d is a positive number\n", number1)
	}
	fmt.Println("Out of the loop")


	// IF ...ELSE Statements
	/*
		- if statement may have an optional else block
		- the syntax of if ..else stament is
		if test_condition{
			// code
		} else {
			// code
		}
	*/

	number2 := -18
	if number2 > 0 {
		fmt.Printf("%d is a positive number\n", number2)
	} else {
		fmt.Printf("%d is a negative number\n", number2)
	}



	// ELSE..IF..ELSE Ladder
	/*
		- the statement is used to execute a block of code among two alternatives
		- but if you want to make a choice among more than 2 alternatives the we use else if
		- 
	*/

	number3 := 3
	if number2 == number3 {
		fmt.Printf("Result: %/d == %d", number2, number3)
	} else if number2 > number3 {
		fmt.Printf("Result: %d > %d", number2, number3)
	} else {
		fmt.Printf("Result: %d < %d", number2, number3)
	}

	/*
		# NESTED IF STATEMENTS
		- you can have a nested if statment in go
	*/
}