package channels

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Foodie struct {
	Age                       int
	Gender                    string
	MaritalStatus             string
	Occupation                string
	MonthlyIncome             string
	EducationalQualifications string
	FamilySize                int
	Latitude                  float64
	Longitude                 float64
	PinCode                   int
	Output                    string
	Feedback                  string
}

// AdvancedChannels demonstrates how to use channels in Go.

func HandleFoodie(f Foodie, c chan Foodie) {
	// f.Output = "Processed"
	c <- f
}

func AdvancedChannels() {
	fmt.Println("Advanced Channels starting up")
	start := time.Now()

	// Start a gorutine to process the data with a buffered channel
	foodieChannel := make(chan Foodie, 100)
	go func() {
		for {
			f := <-foodieChannel
			allFoodies = append(allFoodies, f)
		}
	}()

	csvFile, err := os.Open("onlinefoods.csv")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	fmt.Println("File opened successfully")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	counter := 0
	for _, line := range csvLines {
		// Skip the header
		if counter == 0 {
			counter++
			continue
		}
		// Age,Gender,Marital Status,Occupation,Monthly Income,Educational Qualifications,Family size,latitude,longitude,Pin code,Output,Feedback,
		// 23,Male,Single,Student,No Income,Post Graduate,5,12.8988,77.5764,560078,Yes,Positive,Yes
		age, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("Failed to convert to integer:", err)
			return
		}
		gender := line[1]
		maritalStatus := line[2]
		occupation := line[3]
		monthlyIncome := line[4]
		educationalQualifications := line[5]
		familySize, err := strconv.Atoi(line[6])
		if err != nil {
			fmt.Println("Failed to convert to integer:", err)
			return
		}
		latitude, err := strconv.ParseFloat(line[7], 64)
		if err != nil {
			fmt.Println("Failed to convert to float:", err)
			return
		}
		longitude, err := strconv.ParseFloat(line[8], 64)
		if err != nil {
			fmt.Println("Failed to convert to float:", err)
			return
		}
		pinCode, err := strconv.Atoi(line[9])
		if err != nil {
			fmt.Println("Failed to convert to integer:", err)
			return
		}
		output := line[10]
		feedback := line[11]

		f := Foodie{
			Age:                       age,
			Gender:                    gender,
			MaritalStatus:             maritalStatus,
			Occupation:                occupation,
			MonthlyIncome:             monthlyIncome,
			EducationalQualifications: educationalQualifications,
			FamilySize:                familySize,
			Latitude:                  latitude,
			Longitude:                 longitude,
			PinCode:                   pinCode,
			Output:                    output,
			Feedback:                  feedback,
		}

		// Send the data to the channel
		foodieChannel <- f
		counter++
	}

	fmt.Println("Processed", counter, "lines of code in", time.Since(start))
	fmt.Println("Advanced Channels shutting down")
	fmt.Println("Goodbye!")
	fmt.Println("All foodies:", len(allFoodies))

	fmt.Println("Random foodie:", allFoodies[1490])
}
