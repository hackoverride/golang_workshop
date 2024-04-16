package channels

import (
	"fmt"

	"gogame/workshop/tools"
)

// Channels demonstrates how to use channels in Go.
func Channels() {
	fmt.Println("Channels:")

	// Create a new channel of type int
	myChannel := make(chan int)

	// Start a goroutine to send data to the channel
	go func() {
		myChannel <- 42
	}()

	// A Go routine is a lightweight thread managed by the Go runtime
	// It is a function that runs concurrently with other functions
	// The 'go' keyword is used to start a goroutine

	// Read from the channel
	value := <-myChannel // The '<-' operator is used to send and receive data from a channel
	fmt.Println("Value from channel:", value)
	tools.Pause()
}

func init() {
	// It is used to initialize the package
	// Normally the init function is used to initialize variables, setup configurations, etc.
	// It is called only once when the package is imported
	fmt.Println("Channels package initialized.")
	// The init function is called before the main function
	// It is used to initialize the package
}
