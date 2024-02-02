package main
import "fmt"

func main(){
	/*
		- go operators are simples that perform operations on variables or values
		- go has multiple operators that are categorized into;
		  		- Arithmatic
				- Assignment
				- Relational
				- Logical
	*/


	/*
		# Arithmatic
		 - Addition (a + b)
		 - Subtraction (a - b)
		 - Multiplication (a * b)
		 - Divition (a/b)
		 - Modulo Divition (a % b)
	*/
		num1 := 2
		num2 := 5
		num3 := 1

		var sum int
		var subtruct int
		var product int
		var modulo int
		var divitions int

		fmt.Printf("%d + %d = %d\n", num1, num2, sum)
		fmt.Printf("%d - %d = %d\n", num1, num2, subtruct)
		fmt.Printf("%d * %d = %d\n", num1, num2, product)
		fmt.Printf("%d % %d = %d\n", num2, num1, modulo)
		fmt.Printf("%d / %d = %d\n", num1, num3, divitions)

	/*
		# Assignment
		 - are used to assign values to a variable
		 - we also have compound assignment operators like:
		 	+= (a+=b)
			-+ (a-=b)
			*= (a*=b)
			/= (a/=b)
			%= (a%=b)
	*/

		name := "Mike"
		fmt.Println(name)

		// compound
		number := 2
		number += 7
		fmt.Println(number)

	/*
		# Relational
		  - used to compare 2 values or variables
	*/
	fmt.Println(5 == 6)
	// returns false 

	/*
		# Logical
			- used to perform logical operations and 
			- returns true or false depending upon the conditions
			    - &&  (Logical AND) - returns true of both expressions are true
				- ||  (Logical OR) - return true if one of the expression is true
				- !  (Logical Not) - returns true if expression is false and returns false if expression is true.
	*/


	
}