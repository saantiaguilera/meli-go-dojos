package external

import (
	"errors"
	"math/rand"
	"time"
)

type RosteringService struct{}

func (r RosteringService) PerformRostering(route Route, drivers []Driver) (Driver, error) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	if rand.Float32() < 0.5 && len(drivers) > 0 {
		return drivers[0], nil
	} else {
		return Driver{}, errors.New("error")
	}
}

func NewRosteringService() RosteringService {
	return RosteringService{}
}
