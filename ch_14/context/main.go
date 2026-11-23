package main

import (
	"context"
	"fmt"
	"time"
)

const (
	userKey = "user"
)

func main() {
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	defer cancel2()
	start := time.Now()
	<-child.Done()
	end := time.Now()
	fmt.Println(end.Sub(start).Truncate(time.Second))

	// ctx := context.Background()

	// if myVal, err := ctx.Value("myKey").(string); err {
	// 	fmt.Println("myVal:", myVal)
	// } else {
	// 	fmt.Println("myVal not found")
	// }
}

func contextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func userFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userKey).(string)
	return user, ok
}
