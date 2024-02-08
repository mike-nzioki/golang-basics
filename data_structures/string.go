package main
import ( 

	"fmt" 
	"strings" 
)


func main () {
	/*
		* String
		- a string is a sequence of characters 
		- we use double quotes to represent a string in go
	*/

	name := "Mike"
	fmt.Println(name)

	//  in go we can also represent a string suing backticks ``

	greeting := `Good morning`
	fmt.Println(greeting)


	// # Acess character in a string 
	characterName := "programmer"
	fmt.Printf("%c\n", characterName[0])
	fmt.Printf("%c\n", characterName[1])
	fmt.Printf("%c\n", characterName[2])
	fmt.Printf("%c\n", characterName[3])


	//  # Find string Length
	lenCharactor := len(characterName)
	fmt.Println(lenCharactor)


	// # Join Strings together

	sentance := name + " " + "is a " + " " + characterName
	fmt.Println(sentance)

	

	/*
		* STRING FUNCTIONS
		- 	// import strings at the top of the file to use the functions below

		Compare()  -   compares two strings
		Contains(  -   checks if a substring is present inside a string
		Replaces(  -   replaces a substring with another substring
		ToLower()  -   converts a string to lowercase
		ToUpper()  -   converts a string to uppercase
		Split()	   -   plits a string into multiple substrings

	*/
	


	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"

	//strings.Contains()
	result1 := strings.Contains(text, substring1)
	fmt.Println(result1)

	result2 := strings.Contains(text, substring2)
	fmt.Println(result2)


	// strings.Replace()
	// r - character to be replaced
	// t - new charactoer to replace the old
	// 1 - represents how many old characters to be replaced
	replacedText := strings.Replace(text, "r", "t", 1)
	fmt.Println("New String: ", replacedText)


	// string.ToUpper()
	// string.ToLower()

	upperChar := strings.ToUpper(text) 
	lowerChar := strings.ToLower(text)
	fmt.Println(upperChar)
	fmt.Println(lowerChar)

	
	// SPLIT CHARACTERS
	message := "I love GO"
	splittedText := strings.Split(message, " ")
	fmt.Println(splittedText)

}