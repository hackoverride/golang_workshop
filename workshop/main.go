package main

import (
	"bufio"
	"fmt"
	"gogame/workshop/examples/shines"
	"os"
	"time"
)

func main() {

	// examples.Example()
	// examples.ExampleRoutine()

	// Regular AdvancedChannels handle 17_081 lines of data in the onlinefoods.csv file
	// normally around 1.65 seconds.
	// channels.AdvancedChannels()
	// channels.AdvancedChannelsWithGoRoutines()
	shines.Shines()
}

func init() {
	fmt.Println("Welcome to the Go workshop!")
}

func HelloCap() {
	fmt.Println("Hello, Cap Sandefjord!")
	fmt.Println("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read name:", err)
		return
	}

	// Print formatted string
	fmt.Printf("Hello, %sWelcome to the Go workshop!\n", name)

	fmt.Println("Current time:", time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"))
	// https://www.geeksforgeeks.org/time-formatting-in-golang/
	// Learn about time formatting in Go

	fmt.Println("Did you know? Go is a statically typed language. This means that the type of a variable is known at compile time.")
}
