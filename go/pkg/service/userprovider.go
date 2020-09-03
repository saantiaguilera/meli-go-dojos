package service

import "github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg/domain"

type UserProvider struct {
	RetrieveUser getUser
}

type getUser func(id uint) (userService, error)

func NewUserProvider(g getUser) UserProvider {
	return UserProvider{RetrieveUser: g}
}

func (u UserProvider) GetUser(id uint) (domain.User, error) {
	uc, err := u.RetrieveUser(id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:      uc.ID,
		Name:    uc.Name,
		Surname: uc.Surname,
		Age:     uc.Age,
	}, nil
}
