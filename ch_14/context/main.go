package main

import (
	"context"
	"fmt"
)

const (
	userKey = "user"
)

func main() {

	ctx := context.Background()

	if myVal, err := ctx.Value("myKey").(string); err {
		fmt.Println("myVal:", myVal)
	} else {
		fmt.Println("myVal not found")
	}
}

func contextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func userFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userKey).(string)
	return user, ok
}
