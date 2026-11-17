package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://jsonplaceholder.typicode.com"

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func fetch(url string, v any) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("so'rov xatosi %s: %w", url, err)
	}
	// Body yopilmasa, ulanish pool'ga qaytmaydi va resurs oqib ketadi.
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("kutilmagan status %s: %d", url, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("JSON decode xatosi %s: %w", url, err)
	}
	return nil
}
