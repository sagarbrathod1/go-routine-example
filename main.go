package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiKey = "test_key"

func fetchStockPrice(symbol string) string {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)
	
	// Make a GET request to the Alpha Vantage API
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Return the response body as a string
	return string(body)
}

func main() {
	// Record the start time
	start := time.Now()

	companies := []string{"META", "AMZN", "AAPL"}

	// Fetch the stock price for each company
	for _, company := range companies {
		price := fetchStockPrice(company)
		fmt.Println(company, price)
	}

  // Log the total time taken to execute the API calls
	fmt.Println("Time taken:", time.Since(start))
}