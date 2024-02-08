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
	// - 

}