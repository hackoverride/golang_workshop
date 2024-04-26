package channels

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const workers = 10      // number of workers to process the data
var allFoodies []Foodie // slice to store all foodies

func AdvancedChannelsWithGoRoutines() {
	allFoodies = make([]Foodie, 0)
	fmt.Println("Advanced Channels starting up")
	start := time.Now()
	mu := sync.Mutex{}

	foodieChannel := make(chan Foodie, 100)
	var wg sync.WaitGroup

	// A wait group is just a way for us to make sure all goroutines finish before we exit the program.
	// We add the number of goroutines we want to wait for to the wait group.
	// Then we call wg.Done() in each goroutine when it finishes.

	/* The code underneath is the code without the Mutex */
	// go func() {
	// 	for f := range foodieChannel {
	// 		allFoodies = append(allFoodies, f)
	// 	}
	// }()
	// go func() {
	// 	for f := range foodieChannel {
	// 		allFoodies = append(allFoodies, f)
	// 	}
	// }()

	// This code will handle the data concurrently because we are using goroutines with mutexes.
	for i := 0; i < workers; i++ {
		go func() {
			for f := range foodieChannel {
				mu.Lock()
				allFoodies = append(allFoodies, f)
				mu.Unlock()
			}
		}()
	}
	// The code above will work with the Mutex

	csvFile, err := os.Open("onlinefoods.csv")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	timeUsedToReadTheFile := time.Since(start)
	// During testing the time used to read the file was ~ 11 ms

	dataLength := len(csvLines) - 1

	for i := 0; i < len(csvLines); i++ {
		wg.Add(1) // increment the WaitGroup counter for each goroutine
		go func(ch chan<- Foodie, line []string) {
			defer wg.Done()
			// the WaitGroup will be decremented when the goroutine exits
			// When the WaitGroup counter reaches zero, all goroutines have finished.
			// This of the WaitGroup as a counter and is similar to async/await.

			if line[0] == "Age" { // skip the header
				return
			}
			f, err := parseFoodie(line)
			if err != nil {
				fmt.Println("Error parsing foodie:", err)
			}
			// The arrows <- are used to send and receive messages from channels.
			// Channels are used to communicate between goroutines.
			ch <- f
		}(foodieChannel, csvLines[i])
	}

	wg.Wait()
	close(foodieChannel)
	fmt.Println("Processed", dataLength, "lines of data in", time.Since(start))
	fmt.Println("Time used to read the file:", timeUsedToReadTheFile)
	fmt.Println("Advanced Channels shutting down")
	fmt.Println("Goodbye!")

	fmt.Println("All foodies:", len(allFoodies))

	fmt.Println("Random foodie:", allFoodies[1490])

	// This code will be very quick, but we are not guaranteed to have all the foodies in the slice.
	// This is because the goroutines are running concurrently and we are not handling the data correctly.
}

func parseFoodie(line []string) (Foodie, error) {
	age, _ := strconv.Atoi(line[0])

	familySize, err := strconv.Atoi(line[6])
	if err != nil {
		return Foodie{}, err
	}
	latitude, err := strconv.ParseFloat(line[7], 64)
	if err != nil {
		return Foodie{}, err
	}
	longitude, err := strconv.ParseFloat(line[8], 64)
	if err != nil {
		return Foodie{}, err
	}
	pinCode, err := strconv.Atoi(line[9])
	if err != nil {
		return Foodie{}, err
	}

	return Foodie{
		Age:                       age,
		Gender:                    line[1],
		MaritalStatus:             line[2],
		Occupation:                line[3],
		MonthlyIncome:             line[4],
		EducationalQualifications: line[5],
		FamilySize:                familySize,
		Latitude:                  latitude,
		Longitude:                 longitude,
		PinCode:                   pinCode,
		Output:                    line[10],
		Feedback:                  line[11],
	}, nil
}
