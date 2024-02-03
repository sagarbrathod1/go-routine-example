package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const apiKey = "test_key"

// Added new function parameters: 1) a channel, 2) a pointer to a WaitGroup
func fetchStockPrice(symbol string, ch chan<-string, wg *sync.WaitGroup) string {
	defer wg.Done() // Ensure that the WaitGroup is closed when the function exits
	
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	ch <- fmt.Sprintf("%s: %s", symbol, body) // Send the result to the channel

	return string(body)
}

func main() {
	start := time.Now()

	companies := []string{"META", "AMZN", "AAPL"}

	ch := make(chan string, len(companies)) // Create a buffered channel
	var wg sync.WaitGroup // Create a WaitGroup

	// Launch a goroutine for each company
	for _, company := range companies {
		wg.Add(1) // Increment the WaitGroup counter
		go fetchStockPrice(company, ch, &wg) // Pass the function into a goroutine
	}

	// Wait for all the goroutines to finish and close the channel
	go func() {
		wg.Wait()
		close(ch) 
	}()

	// Read the results from the channel
	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("Time taken:", time.Since(start))
}