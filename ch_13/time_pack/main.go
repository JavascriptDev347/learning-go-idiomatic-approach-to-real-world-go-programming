package main

import (
	"fmt"
	"time"
)

func main() {

	// time.Parse
	t, err := time.Parse("2006-01-02 15:04:05 -0700", "2023-03-13 00:00:00 +0000")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
}
