package external

import (
	"math/rand"
	"time"
)

type Order struct {
	ID uint
}

type OrderProvider struct{}

func (o OrderProvider) Get(orderID uint) Order {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return Order{
		ID: orderID,
	}
}

func NewOrderProvider() OrderProvider {
	return OrderProvider{}
}
