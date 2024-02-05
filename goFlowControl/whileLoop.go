package main
import "fmt"

func main(){
	/*
		# While loop is used to execute a block of code until a certain condition is met
		 - unlike other languages go doesnot have a keword for a while loop
		 - in this case a for loop is used to perfor while loops i.e

		 for condition {
			// code
		 }
	*/

		 number := 1
		 for number <= 5 {
			fmt.Println(number)
			number++
		 }

	// # CREATE A MULTIPLICATION TABLE OF 5 
	multiplier := 1
	for multiplier <= 10 {
		product := 5 * multiplier
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier ++
	}


	/*
		# do...while loop
		- for loop can be use tp to perform do while loop
	*/
	num1 := 1
	for {
		if num1 > 5 {
			break;
		}
		fmt.Printf("%d\n", num1)
		num1 ++
	}
}