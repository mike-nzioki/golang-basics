package main
import (
	"fmt"
	"time"
)

func main () {
	/*
		* GOROUTINES
		- goroutines are used to create concurrent programs
		- concurrent programs are able to run multiple processes at thesame time
		- to convert a normal function to goroutine you have to call the function with go keyword..example

		func display() {
  			// code inside the function
		}

		// start goroutine
		go display()
	*/

	// call goroutine
	go display("Process 1")
	display("Process 2")
	// the process one is not displayed because the goroutine is running independently and concurrent with main ()


	// RUN TWO FUNCTIONS CONCURRENTLY
	/*
		* The main() is always a default goroutine. other goroutines can be executed if the main is not executing
		- in order to make sure that all the goroutines are executed before the main function ends, we 
		-> ..sleep the process so that other proccesses get a chance
	*/

	go display1("Process 1")
	display1("Process 2")
}


func display(message string) {
	fmt.Println(message)
}


func display1(message string){
	// infinite loop
	for {
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}
}