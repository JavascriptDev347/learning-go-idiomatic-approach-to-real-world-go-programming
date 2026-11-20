package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const URL = "https://jsonplaceholder.typicode.com/posts/9999"

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(URL)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error: status code", resp.StatusCode)
		return
	}

	var post Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}

	fmt.Printf("Post ID: %d\nUser ID: %d\nTitle: %s\nBody: %s\n", post.ID, post.UserID, post.Title, post.Body)
}
