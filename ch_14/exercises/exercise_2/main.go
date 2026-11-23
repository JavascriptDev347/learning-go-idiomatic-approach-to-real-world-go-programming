package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errFoundTarget = errors.New("1234 soni generatsiya qilindi")

func main() {

	ctx, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout()

	ctx, cancelCause := context.WithCancelCause(ctx)
	defer cancelCause(nil)

	var (
		sum        int64 = 0
		iterations int64 = 0
	)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Hisoblash boshlandi (maksimal 2 soniya)...")

loop:
	for {
		if err := context.Cause(ctx); err != nil {
			break loop
		}
		iterations++

		num := r.Intn(100_000_000)
		if num == 1234 {
			cancelCause(errFoundTarget)
			break loop
		}

		sum += int64(num)
	}
	fmt.Println("-------------------------------------------")
	fmt.Printf("Yakuniy yig'indi (Sum): %d\n", sum)
	fmt.Printf("Iteratsiyalar soni (Count): %d\n", iterations)
	cause := context.Cause(ctx)
	if errors.Is(cause, errFoundTarget) {
		fmt.Printf("To'xtash sababi: Maqsadga erishildi (%v)\n", cause)
	} else if errors.Is(cause, context.DeadlineExceeded) {
		fmt.Printf("To'xtash sababi: Vaqt chegarasi tugadi (%v)\n", cause)
	} else {
		fmt.Printf("To'xtash sababi: Noma'lum sabab (%v)\n", cause)
	}
	fmt.Println("-------------------------------------------")
}
