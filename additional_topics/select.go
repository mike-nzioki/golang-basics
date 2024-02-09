package main
import (
	"fmt"
	"time"
)

func main () {
	/*
		* Go select
		- select statment in go allows us to execute a channell amonng many alternatives
		- syntax
		select {

  			case firstChannel:
			
  			case secondChannel:
			
  			case thirdChannel:

		}
	*/


	// Select
	number := make(chan int)
	message := make(chan string)
	
	// function call with goroutine
	go channelNumber(number)
	go channelMessage(message)

	select {
		case firstChannel := <- number:
			fmt.Println("Channel Data:", firstChannel)
		case secondChannel := <- message:
			fmt.Println("Channel Data:", secondChannel)
		 // default case 
		// default:
		// 	fmt.Println("Wait!! Channels are not ready for execution")
	}

	/*
		* Select one channel
		- in the above examples channels are ready for execution 
		- the select statment executes the channel randomly
		- if only one of the channel is ready for execution, it executes the channel
		- use time.sleep() to make the message channel unavailble eg
		- time.Sleep(2 * time.Second)
		- that way the number channel is executed
	*/


	/*
		* Go select to block channels
		- the statement blocks channels if they are not ready for execution
		- use time.sleep() to make both message channel unavailble eg
		- time.Sleep(2 * time.Second)
		- the channels will be blocked in the first 2 seconds then executed
	*/
}

// goroutines
func channelNumber(number chan int) {
	// time.Sleep(2 * time.Second)
	number <- 17
}

func channelMessage(message chan string) {
	time.Sleep(2 * time.Second)
	message <- "Learning go select"
}