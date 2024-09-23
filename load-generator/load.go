package main

import (
	"flag"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// Global variables to track success/failure counts and response times
var successCount, failureCount int64
var totalResponseTime time.Duration
var responseTimes []time.Duration
var mu sync.Mutex

// Function to send an HTTP request
func sendRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the goroutine completes

	start := time.Now()

	// Build request
	resp, err := http.Get(url)
	duration := time.Since(start)

	mu.Lock()
	responseTimes = append(responseTimes, duration)
	mu.Unlock()

	// Err handling for fail or status code
	if err != nil || resp.StatusCode >= 400 {
		atomic.AddInt64(&failureCount, 1)
		fmt.Printf("Failed Request. Error: %v\n", err)
		return
	}

	// Success rate and response time
	atomic.AddInt64(&successCount, 1)
	atomic.AddInt64((*int64)(&totalResponseTime), int64(duration))

	fmt.Printf("Response Code: %d, Response Time: %v\n", resp.StatusCode, duration)
}

func main() {
	// Define command-line flags
	url := flag.String("url", "http://google.com", "The URL to load test") // default: google.com
	concurrency := flag.Int("c", 10, "Number of concurrent requests")      // default: 10
	totalReq := flag.Int("r", 100, "Total number of requests")             // default: 100

	// Parse command-line flags
	flag.Parse()

	// Initialize WaitGroup for synchronizing goroutines
	var wg sync.WaitGroup

	// Fire off the requests
	for i := 0; i < *totalReq; i++ {
		wg.Add(1) // Increment for each goroutine

		// Start a goroutine to send the request
		go sendRequest(*url, &wg)

 		// Maintain concurrency level
		if (i+1)%*concurrency == 0 {
			wg.Wait() // Wait for the current batch to finish before starting more
		}
	}

	// Wait for all remaining requests
	wg.Wait()

	// Report the final metrics
	fmt.Printf("\n===== Metrics =====\n")
	fmt.Printf("Total Requests: %d\n", *totalReq)
	fmt.Printf("Successful Requests: %d\n", successCount)
	fmt.Printf("Failed Requests: %d\n", failureCount)

	// Calculate average response time
	avgResponseTime := time.Duration(totalResponseTime.Nanoseconds() / int64(*totalReq))
	fmt.Printf("Average Response Time: %v\n", avgResponseTime)

	// Sort response times to calculate percentiles
	sort.Slice(responseTimes, func(i, j int) bool {
		return responseTimes[i] < responseTimes[j]
	})

	// Calculate and print the 95th percentile response time
	percentileIndex := int(float64(len(responseTimes)) * 0.95)
	p95ResponseTime := responseTimes[percentileIndex]
	fmt.Printf("95th Percentile Response Time: %v\n", p95ResponseTime)

	fmt.Println("All requests completed.")
}
