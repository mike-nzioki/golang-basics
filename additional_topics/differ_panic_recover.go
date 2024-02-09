package main
import (
	"fmt"
)


func main() {
	/*
		* Go panic, deffer, and recover
		- we use differ, panic and recover statements to handle errors
		- differ is used to delay the execution of function that might cause an error
		- panic statement terminates the program immmedietely
		- recover is used to recover the message during panic
	*/

	// # differ in go
	// - used to prevent the execution of a function before until all other functions are executed

	defer fmt.Println("Theree")

	fmt.Println("One")
	fmt.Println("Two")
	// three is printed after one and two are executed
	// note that incase of multiple deffer statements the order of execution will be LIFO (last in first out)



	// # panic
	// - we use panic to immediately terminate execution of a program
	fmt.Println("Help! something is wrong")
	panic( "Ending the Program")
	fmt.Println("Waiting to be executed")

	division(4, 2)
	division(8, 0)
	division(2, 8)



	// # RECOVER IN Go
	//  - the panic statement terminates the program but it is important for program to complete its execution
	//  - we use recover statement to to hadle panic in go 
	//  - recover statement prevents termination of the program and recovers the program for panic


}


func division(num1, num2 int) {

	// execute handlePanic even befor panic occurs
	// defer handlePanic()
	// if num2 is 0, program is terminated due to panic

	if num2 == 0 {
	  panic("Cannot divide a number by zero")
	} else {	
	  result := num1 / num2
	  fmt.Println("Result: ", result)
	}
	
}


// recover
func handlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("RECOVER", a)
	}
}