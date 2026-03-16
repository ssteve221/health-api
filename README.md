# Go System Health API 🖥️

A lightweight, zero-dependency REST API built in Go that exposes real-time server metrics, memory allocation, and OS-level details. 

This project demonstrates an understanding of Go's `net/http` standard library, Go 1.22+ routing syntax, and the `runtime` package to interface directly with the underlying system and Go scheduler.

## Features
* **Real-Time Metrics:** Fetches active goroutines, memory allocation, and system uptime on every request.
* **System-Level Insights:** Identifies the host operating system and architecture dynamically.
* **Zero Dependencies:** Built entirely with the Go standard library—no external web frameworks (like Gin or Fiber) required.
* **Standardized JSON Responses:** Uses strict Go structs to marshal system data into clean, readable JSON.

## Tech Stack
* **Language:** Go (Golang)
* **Libraries:** `net/http`, `encoding/json`, `runtime`, `time`

## Prerequisites
* [Go](https://go.dev/dl/) installed on your machine (Version 1.22+ recommended for the new routing syntax).

## Installation & Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/ssteve221/health-api.git
   ```
   ```bash
   cd health-api
   ```
2. Run server
      ```bash
      go run main.go
      ```
3. The server will start on port 8080.
4. API Endpoints
Get System Health
GET /health

Returns a JSON object containing the current state of the server.

Example Response:

```bash
{
  "status": "Healthy",
  "os": "linux",
  "architecture": "amd64",
  "active_goroutines": 2,
  "memory_allocated_mb": "0.15",
  "total_system_memory_mb": "2.85",
  "server_uptime": "14m32.5s"
}
```
What I Learned:

This project reinforced my ability to build robust REST APIs without relying on heavy external frameworks. By utilizing Go's runtime package, I gained a deeper understanding of memory management, garbage collection tracking, and how to safely expose internal server states to external clients.
