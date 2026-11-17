package main

import (
	"fmt"
	"time"
)

const (
	USERS    = "/users"
	POSTS    = "/posts?userId="
	COMMENTS = "/comments?postId="
	TODOS    = "/todos?userId=1"
)

func main2() {

	t := time.Now()
	var users []User
	if err := fetch(baseURL+USERS, &users); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Fetched users length: %+v\n", len(users))
	var posts []Post
	for _, v := range users {
		var userPosts []Post
		if err := fetch(baseURL+POSTS+fmt.Sprint(v.ID), &userPosts); err != nil {
			fmt.Println(err)
			return
		}
		posts = append(posts, userPosts...)
	}

	var comments []Comment
	for _, v := range posts {
		var postComments []Comment
		if err := fetch(baseURL+COMMENTS+fmt.Sprint(v.ID), &postComments); err != nil {
			fmt.Println(err)
			return
		}
		comments = append(comments, postComments...)

	}

	fmt.Println("Fetched posts length:", len(posts))
	fmt.Println("Fetched comments length:", len(comments))
	fmt.Printf("Time taken: %v\n", time.Since(t)) // 27 seconds

}
