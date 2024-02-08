package main
import ("fmt")

func main () {
	/*
		* pointers to structs
		- we can create a pointer variable of struct type
	*/


	type Person struct {
		name string
		age int
	}

	student := Person{"Mike", 25}
	var studentPointer *Person
	studentPointer = &student
	fmt.Println("student Pointer: ", studentPointer)




	// ACCESS STRUCT USING POINTERS
	// - you can access an individual member of struct using the pointer

	fmt.Println("Name: ", studentPointer.name)
	fmt.Println("Age: ", studentPointer.age)


	// - CHANGE THE STRUCT MEMBER IN GO
	studentPointer.name = "Mitch"
	fmt.Println("New Name: ", studentPointer.name)
}