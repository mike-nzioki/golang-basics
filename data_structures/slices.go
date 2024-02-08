package main
import ("fmt")

func main () {
	/*
		* Slices are collections of similar types of data, just like arrays
		- unlike arrays slices doesnot have fixed sizes

		- in below slice  numbers represent the name of the slice
		- int - data type
		{} - eleents of the slice goes here
	*/

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)


	/*
		* CREATE SLICE FROM AN ARRAY IN GO
	*/
	
	// an array 
	numbers1 := [8]int{1, 2, 3, 4, 5, 6, 7, 9}

	// slice
	sliceNumbers := numbers1[ 3:6 ]
	fmt.Println(sliceNumbers)


	/*
		* SLICE FUNCTIONS

		- append() - adds elements to slice
		- copy () - copy elements of one slice to another
		- Equal() - compares two slices
		- len() - finds lengths of a slice
	*/

	phones := []int{9, 5}
	phones = append(phones, 7)
	fmt.Println(phones)
	
	copy(phones, sliceNumbers)
	fmt.Println("phones:", phones)
}