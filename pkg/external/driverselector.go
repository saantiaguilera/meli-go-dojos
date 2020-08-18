package external

import (
	"errors"
	"math/rand"
	"time"
)

type DriverSelector struct{}

func (d DriverSelector) Select(order Order) ([]uint, error) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	if rand.Float32() < 0.5 {
		var arr []uint
		l := rand.Intn(25)
		for i := 0; i < l; i++ {
			arr = append(arr, uint(i))
		}
		return arr, nil
	} else {
		return nil, errors.New("error")
	}
}

func NewDriverSelector() DriverSelector {
	return DriverSelector{}
}
