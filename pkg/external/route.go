package external

import (
	"math/rand"
	"time"
)

type Route struct {
	ID    uint
	Order Order
}

type RouteCreator struct{}

func (r RouteCreator) Create(order Order) Route {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return Route{
		ID:    1,
		Order: order,
	}
}

func NewRouteCreator() RouteCreator {
	return RouteCreator{}
}
