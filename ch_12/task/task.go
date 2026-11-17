package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	t := time.Now()
	// fetch users
	var users []User
	if err := fetch(baseURL+USERS, &users); err != nil {
		fmt.Println(err)
		return
	}

	var posts []Post
	for _, v := range users {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var userPosts []Post
			if err := fetch(baseURL+POSTS+fmt.Sprint(v.ID), &userPosts); err != nil {
				fmt.Println(err)
				return
			}
			mu.Lock()
			posts = append(posts, userPosts...)
			mu.Unlock()
		}()
	}
	wg.Wait()

	var comments []Comment
	for _, v := range posts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var postComments []Comment
			if err := fetch(baseURL+COMMENTS+fmt.Sprint(v.ID), &postComments); err != nil {
				fmt.Println(err)
				return
			}
			mu.Lock()
			comments = append(comments, postComments...)
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Fetched users length:", len(users))
	fmt.Println("Total posts:", len(posts))
	fmt.Println("Total comments:", len(comments))
	fmt.Printf("Time taken: %v\n", time.Since(t))
}
