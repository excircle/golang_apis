# "Site Latency API"

This folder contains source code for a Go-based API which parses a list of Debian Linux mirrors and returns the URL which can be visited with lowest latency.

## API Anatomy

### <a href="https://github.com/excircle/golang_apis/blob/master/web_latency_API/main.go">main.go</a>

2 files. One for the main, one for the data.

The main is relatively small but utilizes highly useful Goroutines to act in parallel.

1 custom struct (response).
1 custom Function utilizing a Goroutine (findFastest).

A simple web server is started in the main. It also loads a slice of web URLs from the package "mirrors" into the function findFastest.

The findFastest function utilizes a Goroutine to gather latency on the URLs. Once rathered, this API returns a string bearing the URL which was contacted fastest.

```go
package main

import (
	"encoding/json"
	"fmt"
	"mirrors"
	"log"
	"net/http"
	"time"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return response{<-urlChan, <-latencyChan}
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})
	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
```

### <a href="https://github.com/excircle/golang_apis/blob/master/web_latency_API/data.go">data.go</a>

Again, this is just a simple slice which is sourced into the main/findFastest function.
```go
package mirrors

// MirrorList is list of Debian mirror sites
var MirrorList = []string {
	"http://ftp.am.debian.org/debian/",
	"http://ftp.au.debian.org/debian/",
	"http://ftp.at.debian.org/debian/",
	"http://ftp.by.debian.org/debian/",
	"http://ftp.be.debian.org/debian/",
}
```
