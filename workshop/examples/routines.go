package examples

import (
	"fmt"
	"time"
)

func ExampleRoutine() {
	fmt.Println("Routine Example starting up")
	start := time.Now()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Routine finished after", time.Since(start))
	}()

	fmt.Println("Main finished after", time.Since(start))
	time.Sleep(3 * time.Second) // Wait for the routine to finish
	fmt.Println("ExampleRoutine finished after", time.Since(start))
}
