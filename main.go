package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func makeRequest(client *http.Client, url string, wg *sync.WaitGroup) {

	defer wg.Done()

	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start)

	log.Printf("Request to %s completed in %s with status code %d\n", url, duration, resp.StatusCode)

}

func main() {

	url := "http://localhost:3000"

	numRequest := 10

	client := &http.Client{}

	var wg sync.WaitGroup
	wg.Add(numRequest)

	startTime := time.Now()

	for i := 0; i < numRequest; i++ {
		go makeRequest(client, url, &wg)
	}

	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("Stress test completed in %s\n", duration)

}
