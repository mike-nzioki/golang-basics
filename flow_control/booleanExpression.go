package main
import "fmt"

func main(){
	/*
		- relational and logical operators
		- the boolean data type represents logical entities, it can have tow possible values, true or false
		- we create boolean type by usning bool keyword
		-
		
	*/
	isValid := false
	fmt.Println(isValid)


	// RELATION OPERATIONS
		num1 := 3
		num2 := 5
		result := num1 > num2
		fmt.Println(result)
	
	
	// Different types of relational operations
	/*
		== (equal to) - i.e a == b
		!= (not eqaul to) -i.e a != b
		> (greater than) -i.e a > b
		< (less than)	-i.e  a < b
		>= (greater than or equal to ) -i.e a >= b
		<= (less than or equal to) -i.e a <= b
	*/


	number1 := 12
  	number2 := 20
  	var results bool
	
  	// equal to operator
  	result = (number1 == number2)

  	fmt.Printf("%d == %d returns %t \n", number1, number2, results)

  	// not equal to operator
  	result = (number1 != number2)

  	fmt.Printf("%d != %d returns %t \n", number1, number2, results)

  	// greater than operator
  	result = (number1 > number2)

  	fmt.Printf("%d > %d returns %t \n", number1, number2, results)

  	// less than operator
  	result = (number1 < number2)

  	fmt.Printf("%d < %d returns %t \n", number1, number2, results)

}