package loadgen

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// Global variables to track success/failure counts and response times
var successCount, failureCount int64
var totalResponseTime time.Duration
var responseTimes []time.Duration
var mu sync.Mutex

// SendRequest sends an HTTP request and tracks metrics
func SendRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	mu.Lock()
	responseTimes = append(responseTimes, duration)
	mu.Unlock()

	if err != nil || resp.StatusCode >= 400 {
		atomic.AddInt64(&failureCount, 1)
		fmt.Printf("Failed Request. Error: %v\n", err)
		return
	}

	atomic.AddInt64(&successCount, 1)
	atomic.AddInt64((*int64)(&totalResponseTime), int64(duration))

	fmt.Printf("Response Code: %d, Response Time: %v\n", resp.StatusCode, duration)
}

// StartRequests manages concurrent request sending
func StartRequests(url string, concurrency int, totalReq int, wg *sync.WaitGroup) {
	for i := 0; i < totalReq; i++ {
		wg.Add(1)
		go SendRequest(url, wg)

		if (i+1)%concurrency == 0 {
			wg.Wait() // Wait for the current batch to finish
		}
	}
}
