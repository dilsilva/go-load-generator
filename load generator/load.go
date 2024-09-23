package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sendRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the goroutine completes
	start := time.Now()

	// Build request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Response time
	duration := time.Since(start)
	fmt.Printf("Response Code: %d, Response Time: %v\n", resp.StatusCode, duration)
}

func main() {
	// Define command-line flags
	url := flag.String("url", "http://google.com", "The URL to load test") // default: google.com
	concurrency := flag.Int("c", 10, "Number of concurrent requests")      // default: 10
	totalReq := flag.Int("r", 100, "Total number of requests")             // default: 100

	// Parse the flags from the command line
	flag.Parse()

	// Initialize WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Fire off the requests
	for i := 0; i < *totalReq; i++ {
		wg.Add(1) // Increment for each goroutine

		// Start a goroutine to send the request
		go sendRequest(*url, &wg)

		// Maintain concurrency level
		if (i+1)%*concurrency == 0 {
			wg.Wait() // Wait for current batch to finish
		}
	}

	// Wait for all remaining requests
	wg.Wait()
	fmt.Println("All requests completed.")
}
