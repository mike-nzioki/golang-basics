package main
import "fmt"

func main(){
	/*
		# GO TypeCasting
			- this is the process of converting the value if one data type to the other
			eg convert int to float
	*/

	floatVal := 7.9
	intVal := int( floatVal )
	fmt.Println(intVal)
	// returns 7

	/*
		# explicite typecasting
		- go provides predifined functions like int(), float32(),float64(), string()
	*/


	/*
		# Implicite Tyoe Casting
		- in this case one data type is automatically converted to the other 
		- go does not support implicite type casting
	*/


	// var number int = 67.8
	// fmt.Println(number)

	// return the error below
	// cannot use 67.8 (untyped float constant) as int value in variable declaration



	// ADD INT TO FLOAT
	number1 := 20
	var number2 float32 = 6.8
	var sum float32

	sum = float32(number1) + number2
	fmt.Printf("Sum is %g", sum)
}