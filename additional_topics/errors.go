package main
import (
	"fmt"
	"errors"
)

func main () {
	/*
		* Errors
		- to catch errors go uses errors.new() or ErrorF()
	*/


	message := "Hello"
	// create error using New() function
	myError := errors.New("WRONG MESSAGE")
  
  	if message != "Programiz" {
		fmt.Println(myError)
  	}

	//	return a value of error type
	//   checkName(name string) error {...}



	// USING ErrorF() IN GO
	// - we can also handle go errors using Errorf() function
	//  - unlike New() we can formart the error message wtih Errorf()
	// the errorf() is preseent in fmt package
	age := -14

  	// create an error using Efforf()
  	error := fmt.Errorf("%d is negative\nAge can't be negative", age)
	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Println("Age: %d", age);
	}
}