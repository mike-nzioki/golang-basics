package main
import "fmt"

func main(){
	/*
		# a loop is used to repeat a block of code
		- this is the for loop syntax
		for initialization; condition; update {
  			statement(s)
		}
	*/

	for index := 0; index <= 5; index++ {
		fmt.Println(index)
	}
	// PRINTS 0 TO 5


	n, sum := 10, 0
	for i := 1; i <= n; i++ {
		sum += i
	} 
	fmt.Println("Sum = ", sum)

	// sum = 55
}