package cache

import (
	"sync"
	"task_1/model"
)

// UserCache - Bir xil muallifni qayta-qayta so'ramaslik uchun xotiradagi xavfsiz kesh type UserCache struct
type UserCache struct {
	mu    sync.RWMutex
	users map[int]model.User
}

func NewUserCache() *UserCache {
	return &UserCache{
		users: make(map[int]model.User),
	}
}

func (u *UserCache) Get(id int) (model.User, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	user, ok := u.users[id]
	return user, ok

}

func (u *UserCache) Set(id int, user model.User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.users[id] = user
}
