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

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func main() {
	start := time.Now()

	companies := []string{"META", "AMZN", "AAPL"}

	for _, company := range companies {
		price := fetchStockPrice(company)
		fmt.Println(company, price)
	}

	fmt.Println("Time taken:", time.Since(start))
}