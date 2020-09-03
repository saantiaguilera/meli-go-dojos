package service

import (
	"errors"
	"sync"

	"github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/domain"
)

type userService struct {
	ID      uint
	Name    string
	Surname string
	Age     uint
}
type UserCache struct {
	user      map[uint]userService
	lockMutex *sync.RWMutex
}

func NewUserCache() UserCache {
	u := make(map[uint]userService)
	mu := new(sync.RWMutex)
	return UserCache{user: u, lockMutex: mu}
}

func (c *UserCache) GetUser(id uint) (userService, error) {
	c.lockMutex.RLock()
	user, exist := c.user[id]
	c.lockMutex.RUnlock()
	if exist {
		return user, nil
	}
	return userService{}, errors.New("no exist")
}

// For testing, this should be non-domain
func (c *UserCache) PutUser(user domain.User) (userService, error) {
	us := userService{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Age:     user.Age,
	}
	c.lockMutex.Lock()
	c.user[user.ID] = us
	c.lockMutex.Unlock()
	return us, nil
}
