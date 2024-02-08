package main
import ("fmt")

func main () {
	/*
		* EMPTY INTERFACE
		- the nterface are used to store a set of methods without implementation
		- we can also create an interface without any methods
		interface {}
	*/

	var a interface {}
	fmt.Println("value: ", a)
	// value is nil

	// - EMPTY INTERFACE TO PASS FUNCTION ARGUMENT OF ANY TYPE
	//  - 
	b := "Welcome to web with go"
	c := 20
	d := true
	// pass string to function param
	displayName(b)
	// pass int
	displayName(c)
	// pass bool
	displayName(d)

	// display value
	displayValue(d, b)
	displayValue(c)
}

func displayName(i interface {} ) {
	fmt.Println(i)
}


// any number of argument
func displayValue(i... interface {} ) {
	fmt.Println(i)
}