package main
import "fmt"
func main() {

/*
	a variable is  a container that is used to store data
	variables have fixed data type associated with it
	variable are decleared using var keyword, the variable name and takes a data type at the end i.e
*/

// String
var name string = "Mike"
fmt.Println(name)

// Int
var price1 int = 10
fmt.Println(price1)

	//  or

var price2 = 10 
fmt.Println(price2)

/*
	in the obove example we are not specifying the data type of the variable
	so in this case the compiler automatically detects the type depending on the value
	assigned to it.

	thus since the value of 10 is an intenger the data type is int

	so it recommended to use this shorthand notation := to assgn values i.e
*/

// declear : var price string
// assign price = 20

// declear and assign: price:=20



	price := 10
	fmt.Println(price)

	price = 100
	fmt.Printf("changed value: %d", price)



	// creating multiple variables at thesame time

	var name1, age = "Mike", 26
	fmt.Printf("name: %s", name1, "age: %s", age)

	car, color := "volvo", "blue"
	fmt.Printf("car: %s", car , "color: %s", color)



	/*
		creating variable rules:
		 - variable name consist of alphabets. digits and an underscore
		 - variable can not have other simples like (@,#,$)
		 - varible names cannot be reserved words like int, type, for, etc.
	*/





	// CONSTANTS IN GO
	
		/*
			constants are fixed values that cannot be changed once decleared 
			- in go we use const key word to create a constant  ariables i.e
		*/

		const artist = "caligraph"
		fmt.Println(artist)

		// BRINGS ERROR (shorthand := is not used to create constants)
		// const lightSpeed := 299792458 
		// fmt.Println(lightSpeed)
}
