package main
import "fmt"

func main () {
	/*
		# Range is used with forloop to iterate thorough the elements of an array, string or map
		- to lean range you should have knowledge og loops
	*/

	numberArray := [6] int{23, 20, 17, 14, 11, 8}
	for index, item := range numberArray {
		fmt.Printf("numberArray[%d] = %d \n", index, item)
	}


	/*
		# RANGE WITH FO MAP
		- we can also use for range keyword with map to access key pairs
		- exampe
	*/

	subMarks := map[string]float32{"Java": 80, "Python": 90, "Go": 99}
	for subject, marks := range subMarks {
		fmt.Println(subject, ":" , marks)
	}


	/*
		# ACCESSING KEYS of a MAP USING GO
	*/

	fmt.Println("Subjects:")
	for subject := range subMarks {
		fmt.Println( subject )
	}
}