package external

import (
	"math/rand"
	"time"
)

type (
	Driver struct {
		ID   uint
		Name string
	}

	DriverService struct{}
)

func (d DriverService) GetDriverInfo(driverID uint) (Driver, error) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return Driver{
		ID:   driverID,
		Name: "test-driver",
	}, nil
}

func NewDriverService() DriverService {
	return DriverService{}
}
