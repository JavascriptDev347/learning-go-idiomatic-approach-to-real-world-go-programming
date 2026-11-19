package model

import "time"

type EnrichedPost struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
}

type StatsResponse struct {
	TotalProcessed int           `json:"total_processed"`
	Duration       string        `json:"duration"`
	DurationNano   time.Duration `json:"duration_ns"`
}
