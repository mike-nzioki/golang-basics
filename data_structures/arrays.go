package main
import "fmt"

func main () {
	/*
		* Go Arrays
		- an array is a collection of similar types of data 
		- below is a suntax of delearing an array in go

		var array_variable = [size]datatype{elements of array}
		- size represents the length of an arrays
		- the list of elements of the array goes inside {}
	*/

	arrayOfIntengers := [5] int {1, 2, 3, 4, 5}
	fmt.Println(arrayOfIntengers)

	/*
		* Declear an array of undifined size
	*/

	greetings := [...]string{"Hello", "Hi", "Hey", "Yoh"}
	fmt.Println(greetings)


	/*
		* shorthand notation to declear an array
	*/
	programmers := [...]string{"Mike", "Mitch", "John"}
	fmt.Println(programmers)



	/*
		* ACCESSING ARRAY ELEMENTS IN GO
		- 
	*/

	languages := [...] string{"Go", "Javascript", "Php", "Python"}
	fmt.Println(languages[0])
	fmt.Println(languages[2])


	/*
		* INITIALIZE AN ARRAY IN GO
	*/
	var arrayOfIntInitialized[3] int
	arrayOfIntInitialized[0] = 1
	arrayOfIntInitialized[1] = 2
	arrayOfIntInitialized[2] = 3
	fmt.Println(arrayOfIntInitialized)


	/*
		* INITIALIZE SPECIFIC ELEMENTS IN AN ARRAY
	*/
		ages := [...] int{0: 7, 2: 5}
		fmt.Println(ages)

	/*
		* Changing the array element in go
		- in the ages array ages[1] is currently 0
		- we change it to 9
	*/
		ages[1] = 9
		fmt.Println(ages)

	/*
		* Finding the length of an array
	*/

		length := len(ages)
		fmt.Println(length)


	/*
		* Looping through an array
	*/

	for i := 0; i < len(ages); i ++ {
		fmt.Println(ages[i])
	}


	/*
		* multidimentional array in go
		- this is an array of arrays
	*/

	multidimentionalArrayOfInts := [2][2] int{ {2, 4}, {3, 6} }
	fmt.Println(multidimentionalArrayOfInts)

	for l := 0; l < 2; l ++ {
		for n := 0; n < 2; n ++ {
			fmt.Println(multidimentionalArrayOfInts[l][n])
		}
	}
}