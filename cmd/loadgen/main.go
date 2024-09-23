package main

import (
	"flag"
	"sync"

	"load-generator/internal/loadgen"
)

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
	loadgen.StartRequests(*url, *concurrency, *totalReq, &wg)

	// Wait for all requests to complete
	wg.Wait()

	// Report final metrics
	loadgen.ReportMetrics(*totalReq)
}
