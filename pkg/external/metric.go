package external

import (
	"errors"
	"math/rand"
	"time"
)

type MetricTracker struct{}

func (m MetricTracker) Track(key string, element interface{}) error {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	if rand.Float32() < 0.5 {
		return nil
	} else {
		return errors.New("couldn't track")
	}
}

func NewMetricTracker() MetricTracker {
	return MetricTracker{}
}
