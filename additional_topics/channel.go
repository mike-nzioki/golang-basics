package main
import (
	"fmt"
)

func main () {
	/*
		* Go Channel
		- acts as medium for goroutiens to communicate with each other
		- to create channel in go we use make() function
		- syntax

		channelName := make(chan int)

		- channelName - name of the channel
		- (chan int) - indicates that the channel is of integer type
	*/

	number := make(chan int)
	message := make(chan string)

	// access type and value of channel
	fmt.Printf("Channel Type: %T\n", number)
	fmt.Printf("Channel Value: %v", number)

	fmt.Printf("Channel Type: %T", message)
	fmt.Printf("Channel Value: %v", message)
	/*
		- %T - to print the type of the channel
		- %v - to print the value of the channel
	*/


	// GO CHANNEL OPERATIONS
	// - create a channel and send data to it
	// - we use this syntax - channelName <- data 

	  // function call with goroutine
	  go channelData(number, message)

	  // retrieve channel data
	  fmt.Println("Channel Data:", <-number)
	  fmt.Println("Channel Data:", <-message)


	// # BLOCKING NATURE OF CHANNELS
	// - the channel automatically blocks the send and recive operations depending on the status of goroutine
	// - when a goroutine sends data into a channel, the operations is blocked until the data is recieved by another goroutine

	ch := make(chan string)
	// function call with goroutine
	go sendData(ch)
	// recieve channel data
	fmt.Println(<-ch)
}


func channelData(number chan int, message chan string) {

	// send data into channel
	number <- 15
	message <- "Learning Go channel"
  
}


// send data
func sendData(ch chan string) {
	// data send to the channel
	ch <- "received. send operation success"
	fmt.Println("No receiver! Send Operation Blocked")
}