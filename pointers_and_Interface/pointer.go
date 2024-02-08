package main
import ("fmt")

func main() {
	/*
		* Pointer
		- go pointers allows to us to work directly with memory addresses
		- when a variable is created a memory address is allocated for the variable to store the values
		- we can access the memory using &
	*/

	num := 7
	// pointer variable 
	ptr := &num
	fmt.Println("Varible value: ", num)
	fmt.Println("Memory Address: ", ptr)



	// get variable pointed by pointers
	// we use * operator to access the valu presented in a mpmory address pointed by pointers
	name := "Mike"
	var namePointer *string
	namePointer = &name
	fmt.Println(*namePointer)



	// WORKING OF POINTERS
	// - declear variables
	var num1 int
	var pointer1 *int

	// - assign value to the variable 
	num1 = 22
	pointer1 = &num1

	fmt.Println("Varible value: ", num1)
	fmt.Println("Pointer Address: ", pointer1)


	// - change the value of a variable
	num1 = 11
	fmt.Println("Varible value: ", num1)


	// - change the value using a pointer
	*pointer1 = 2
	fmt.Println("Varible value: ", num1)
}
