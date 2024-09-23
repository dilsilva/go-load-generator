package loadgen

import (
	"fmt"
	"sort"
	"time"
)

// ReportMetrics outputs the collected metrics
func ReportMetrics(totalReq int) {
	fmt.Printf("\n===== Metrics =====\n")
	fmt.Printf("Total Requests: %d\n", totalReq)
	fmt.Printf("Successful Requests: %d\n", successCount)
	fmt.Printf("Failed Requests: %d\n", failureCount)

	// Calculate average response time
	avgResponseTime := time.Duration(totalResponseTime.Nanoseconds() / int64(totalReq))
	fmt.Printf("Average Response Time: %v\n", avgResponseTime)

	// Sort response times to calculate percentiles
	sort.Slice(responseTimes, func(i, j int) bool {
		return responseTimes[i] < responseTimes[j]
	})

	// Calculate and print the 95th percentile response time
	if len(responseTimes) > 0 {
		percentileIndex := int(float64(len(responseTimes)) * 0.95)
		p95ResponseTime := responseTimes[percentileIndex]
		fmt.Printf("95th Percentile Response Time: %v\n", p95ResponseTime)
	} else {
		fmt.Println("No response times recorded.")
	}

	fmt.Println("All requests completed.")
}
