package order

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"task_context/internal/tracker"
	"time"
)

type Order struct {
	ID      string `json:"id"`
	Owner   string `json:"owner"`
	InStock int    `json:"in_stock"`
}

type Service struct {
	Client   *http.Client
	StockURL string
}

// Get — user context'dan emas, PARAMETR orqali keladi
func (s Service) Get(ctx context.Context, user, id string) (Order, error) {
	tracker.Log(ctx).Info("fetching order", "user", user, "order_id", id)

	if err := slowDB(ctx); err != nil {
		return Order{}, fmt.Errorf("db: %w", err)
	}

	stock, err := s.fetchStock(ctx, id)
	if err != nil {
		return Order{}, fmt.Errorf("stock: %w", err)
	}

	return Order{ID: id, Owner: user, InStock: stock}, nil
}

// slowDB — 3 soniya ishlaydigan "DB query".
// DIQQAT: bu select — 3-bo'limning (Cancellation) oldindan ko'rinishi.
// Context goroutine'ni o'chirmaydi: biz o'zimiz ctx.Done()'ni tinglab chiqib ketamiz.
func slowDB(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err() // context.Canceled
	}
}

// fetchStock — tashqi servisga chaqiruv. ctx so'rovga biriktiriladi (1-bo'lim).
func (s Service) fetchStock(ctx context.Context, id string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.StockURL+id, nil)
	if err != nil {
		return 0, err
	}
	req = tracker.Request(req) // request ID'ni keyingi servisga uzatamiz

	res, err := s.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status %d", res.StatusCode)
	}

	var body struct {
		Count int `json:"count"`
	}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return 0, err
	}
	return body.Count, nil
}
