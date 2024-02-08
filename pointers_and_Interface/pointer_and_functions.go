package main
import ("fmt")

func main(){
	/*
		* Pointers & functions
		- pointers can be used as function argument
		- pointers can be returned in functions
		-
	*/
	number := 55
	// call the function
	update(&number)
	fmt.Println("The number is: ", number)

	// Return Pointer From functions
	msg := display()
	fmt.Println("memory alocation: ", msg)
	fmt.Println("message: ", *msg)
}


// update function
func update(num *int) {
	*num = 30 
}


// Return Pointer From functions
func display() *string {
	message := "Mike is a programmer"
	return &message
}


/*
	* call by reference
	- while passing pointers to a function, we are actually passing a reffernce of a variable.
	- 
*/