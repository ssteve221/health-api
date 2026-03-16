package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// define json to expose
type HealthStats struct {
	Status       string `json:"status"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	Goroutines   int    `json:active_goroutines"`
	MemoryAlloc  string `json:"memory_allocated_mb"`
	TotalMemory  string `json:"total_system_memory_mb"`
	Uptime       string `json:"server_uptime"`
}

// variable to track when server started
var startTime time.Time

func init() {
	startTime = time.Now()
}

// handler function for endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	//populate struct
	stats := HealthStats{
		Status:       "healthy",
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		Goroutines:   runtime.NumGoroutine(),
		//convert bytes to megabytes
		MemoryAlloc: fmt.Sprintf("%.2f", float64(m.Alloc)/1024/1024),
		TotalMemory: fmt.Sprintf("%.2f", float64(m.Sys)/1024/1024),
		Uptime:      time.Since(startTime).String(),
	}
	//format response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//encode struct into HTTP response
	json.NewEncoder(w).Encode(stats)
}

func main() {
	//define route
	http.HandleFunc("GET /health", healthHandler)
	port := ":8080"
	fmt.Printf("System Health MOnitor running on http://localhost%s/health\n", port)
	//start server
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server failed to start :%v\n", err)
	}
}
