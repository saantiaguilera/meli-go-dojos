package usecase

import "github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/domain"

type userProvider interface {
	GetUser(userID uint) (domain.User, error)
}

type User struct {
	provider userProvider
}

func (u User) GetUser(userID uint) (domain.User, error) {
	return u.provider.GetUser(userID)
}

func NewUser(p userProvider) User {
	return User{provider: p}
}
