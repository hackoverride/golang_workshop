package shines

// Java and .Net handles this with threads and locks and async await.
// Go handles this with goroutines and channels - It is in my opinion - less verbose, takes less code, easier to read and understand.
// Go takes less resources on the cpu and memory compared to Java and .Net.

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const workers int = 20

var apiPaths = [20]string{
	"https://www.vg.no",
	"https://www.dagbladet.no",
	"https://www.aftenposten.no",
	"https://www.nrk.no",
	"https://www.tv2.no",
	"https://www.nettavisen.no",
	"https://www.sb.no",
	"https://www.kode24.no",
	"https://www.bt.no",
	"https://www.op.no",
	"https://www.seher.no/",
	"https://dinside.dagbladet.no/",
	"https://www.prisjakt.no/",
	"https://www.finn.no/",
	"https://www.jernia.no/",
	"https://www.capgemini.com/",
	"https://www.power.no/",
	"https://www.bohus.no/",
	"https://www.ikea.no/",
	"https://www.skeidar.no/",
}

func fetchURL(url string, ch chan<- string) {
	startFetch := time.Now()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to fetch %s: %v", url, err)
		ch <- ""
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Non-OK HTTP status: %d from %s", resp.StatusCode, url)
		ch <- ""
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body from %s: %v", url, err)
		ch <- ""
		return
	}
	fmt.Printf("Fetched %s in %v\n", url, time.Since(startFetch))

	ch <- string(body)
}

func Shines() {
	fmt.Println("Shines starting up")
	start := time.Now()
	var wg sync.WaitGroup
	dataCh := make(chan string, workers)

	for _, path := range apiPaths {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fetchURL(url, dataCh)
		}(path)
	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	allTheData := make([]string, 0)
	for data := range dataCh {
		if data != "" {
			allTheData = append(allTheData, data)
		}
	}

	fmt.Println("Done in:", time.Since(start))
	fmt.Println("Data length:", len(allTheData))

	// Save the data to a file
	file, err := os.Create("shines.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	for _, data := range allTheData {
		_, err := file.WriteString(data)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
	}

	fmt.Println("Data saved to shines.txt")
}
