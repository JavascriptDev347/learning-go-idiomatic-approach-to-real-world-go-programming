package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	var users []User
	if err := fetch(baseURL+USERS, &users); err != nil {
		println(err.Error())
		return
	}

	postsCh := make(chan PostsResult)
	for _, v := range users {
		go func() {
			var posts []Post
			if err := fetch(baseURL+POSTS+fmt.Sprint(v.ID), &posts); err != nil {
				postsCh <- PostsResult{Err: err}
				return
			}
			postsCh <- PostsResult{Posts: posts}
		}()
	}

	var posts []Post
	for i := 0; i < len(users); i++ {
		r := <-postsCh
		if r.Err != nil {
			fmt.Println("xato:", r.Err)
			continue
		}
		posts = append(posts, r.Posts...)
	}

	commentsCh := make(chan CommentsResult, len(posts))
	for _, v := range posts {
		go func() {
			var comments []Comment
			if err := fetch(baseURL+COMMENTS+fmt.Sprint(v.ID), &comments); err != nil {
				commentsCh <- CommentsResult{Err: err}
				return
			}
			commentsCh <- CommentsResult{Comments: comments}
		}()
	}

	var comments []Comment
	for i := 0; i < len(posts); i++ {
		r := <-commentsCh
		if r.Err != nil {
			fmt.Println("xato:", r.Err)
			continue
		}
		comments = append(comments, r.Comments...)
	}

	fmt.Println("Fetched users length:", len(users))
	fmt.Println("Total posts:", len(posts))
	fmt.Println("Total comments:", len(comments))
	fmt.Println("Time taken:", time.Since(t))
}
