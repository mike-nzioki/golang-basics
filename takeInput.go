package main
import "fmt"

func main (){
	/*
		go uses scan function from fmt to take inputs from user
		-	fmt.Scan()
		-	fmt.Scanln()
		-	fmt.Scanf()
	*/
	var name string
	var age int
	var role string

	fmt.Println("what is your name?")
	fmt.Scan(&name)


	fmt.Println("how old are you?")
	fmt.Scanln(&age)


	fmt.Println("what do you do?")
	fmt.Scanf("I am a : %s", &role)
}