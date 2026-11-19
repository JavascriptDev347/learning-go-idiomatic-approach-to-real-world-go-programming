package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"task_1/cache"
	"task_1/model"
	"time"
)

type AggregatorService struct {
	client    *http.Client
	userCache *cache.UserCache
	logger    *slog.Logger
}

func NewAggregatorService(logger *slog.Logger) *AggregatorService {
	return &AggregatorService{
		client:    &http.Client{Timeout: 10 * time.Second},
		userCache: cache.NewUserCache(),
		logger:    logger,
	}
}

func (s *AggregatorService) fetchPosts(limit int) ([]model.Post, error) {
	resp, err := s.client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, fmt.Errorf("postlarni yuklab olishda xatolik: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("kutilmagan status kodi: %d", resp.StatusCode)
	}

	var allPosts []model.Post
	if err := json.NewDecoder(resp.Body).Decode(&allPosts); err != nil {
		return nil, fmt.Errorf("postlarni JSON formatida dekodlashda xatolik: %w", err)
	}

	if len(allPosts) > limit {
		allPosts = allPosts[:limit]
	}
	return allPosts, nil
}

func (s *AggregatorService) fetchUsers(userID int) (model.User, error) {
	if user, ok := s.userCache.Get(userID); ok {
		return user, nil
	}

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userID)
	resp, err := s.client.Get(url)
	if err != nil {
		return model.User{}, fmt.Errorf("foydalanuvchini yuklab olishda xatolik: %w", err)
	}
	defer resp.Body.Close()

	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return model.User{}, fmt.Errorf("foydalanuvchini JSON formatida dekodlashda xatolik: %w", err)
	}

	s.userCache.Set(userID, user)
	return user, nil
}

func (s *AggregatorService) AggregatePosts(limit, numWorkers int) ([]model.EnrichedPost, time.Duration, error) {
	startTime := time.Now()

	posts, err := s.fetchPosts(limit)
	if err != nil {
		return nil, 0, fmt.Errorf("postlarni olishda xatolik: %w", err)
	}

	tasks := make(chan model.Post, len(posts))
	results := make(chan model.EnrichedPost, len(posts))

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()

			for post := range tasks {
				user, err := s.fetchUsers(post.UserID)
				if err != nil {
					s.logger.Error("Foydalanuvchini yuklashda xatolik", "userID", post.UserID, "err", err)
					user = model.User{Name: "Noma'lum", Email: "Noma'lum"}
				}

				results <- model.EnrichedPost{
					ID:          post.ID,
					Title:       post.Title,
					Body:        post.Body,
					AuthorName:  user.Name,
					AuthorEmail: user.Email,
				}
			}
		}(i)
	}

	for _, p := range posts {
		tasks <- p
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	var enriched []model.EnrichedPost
	for res := range results {
		enriched = append(enriched, res)
	}
	duration := time.Since(startTime)
	return enriched, duration, nil
}
