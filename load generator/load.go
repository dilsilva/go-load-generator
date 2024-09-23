package main

import (
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
	var wg sync.WaitGroup

	// Configuration
	url := "url"        // Target URL
	concurrentReq := 10 // Amount of concurrent requests
	totalReq := 100     // Total requests to send

	// Fire off requests in batches using Goroutines
	for i := 0; i < totalReq; i++ {
		wg.Add(1) // Increment for each goroutine
		go sendRequest(url, &wg)

		// Maintain concurrency level
		if (i+1)%concurrentReq == 0 {
			wg.Wait() // Wait for current batch to finish
		}
	}

	// Wait for all remaining requests
	wg.Wait()
	fmt.Println("All requests completed.")
}
