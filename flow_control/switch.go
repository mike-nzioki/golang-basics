package main
import "fmt"

func main(){
	/*
		# SWITCH 
		- switch statement allows us to execute one code block among many alernatives
		- we can use fallthrough keyword to excute oher cases after the matching case
		- go switch terminates after the first match thus unlike other loanguages breaks is not needed
	*/

	dayOfTheWeek := 3

	switch dayOfTheWeek {
	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")
		fallthrough

	case 4:
		fmt.Println("Thurday")

	case 5:
		fmt.Println("Friday")

	case 6:
		fmt.Println("Satday")

	case 7:
		fmt.Println("Sunday")
		
	default:
		fmt.Println("Not even a day")
	}

	/*
		:GO SWITCH WITH MULTIPLE CASES
	*/

	weekDay := "Sunday"
	switch weekDay {
		case "Saturday", "Sunday":
			fmt.Println("WEEKEND")
		
		case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
			fmt.Println("WEEKDAY")
		
		default:
			fmt.Println("Not even a day")
	}


	/*
		: SWITCH WITHOUT EXPRESSION
		- the expression in switch is optional 
		- if we dont use a statement the swtichcase is true by default
		-
	*/

	numberOfDays := 28
	switch {
		case 28 == numberOfDays:
			fmt.Println("It's February")
		
		default: 
		fmt.Println("Not February")
	}


	/*
		# switch with optional statement
		- sitch also use an otional statement along with expression
		- the statement and expression are saperated by semicolons
	*/
	
	switch day := 4; day {
	case 1:
		fmt.Println("Sunday")
  
	  case 2:
		fmt.Println("Monday")
  
	  case 3:
		fmt.Println("Tuesday")
  
	  case 4:
		fmt.Println("Wednesday")
  
	  case 5:
		fmt.Println("Thursday")
  
	  case 6:
		fmt.Println("Friday") 
	 
	  case 7:
		fmt.Println("Saturday")
  
	  default:
		fmt.Println("Invalid Day!")

	}

}