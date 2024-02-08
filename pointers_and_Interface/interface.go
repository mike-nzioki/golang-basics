package main
import ("fmt")

func main() {
	/*
		* Interface 
		- we use interface to store a set o methods without implementation
		- method of interface will not have method body
		- below interface has 2 mehtods without implementation
	*/


	/*
		* Implementation of a Go interface
		- to use interface you first have to implement it by a type ( struct )
		- to implement a struct should provide implementations for all methods of an interface
	*/


	  // assigns value to struct members
	  rect := Rectangle{7, 4}

	  // call calculate() with struct variable rect
	  calculate(rect)	

}
	// SHAPE INTERFACE
	type Shape interface {
		area() float32
		// perimeter() float32
	}


	// struct implementation
	type Rectangle struct {
		length, width float32
	}

	// use struct to implement area of interface
	func (r Rectangle) area() float32 {
		return r.length * r.width
	}


	// access method of the interface
	func calculate(s Shape) {
		fmt.Println("Area: ", s.area())
	}