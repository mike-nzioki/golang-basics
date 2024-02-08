package main
import ("fmt")

func main () {
	/*
		* Type assertion
		- type assertion allows us to access the data and the data type of value stored by interface
		-


	*/

	var a interface {}
	// store intenger to an empty interface
	a = 10
	// type assertion
	interfaceValue := a.(int)
	fmt.Println(interfaceValue)
}