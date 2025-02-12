package clock

import "time"

type Clock interface {
	Now() time.Time
}

func NewSystemClock() Clock {
	return &systemClock{}
}

type systemClock struct {
}

func (c *systemClock) Now() time.Time {
	return time.Now()
}

var _ Clock = new(systemClock)

var defaultClock Clock

func init() {
	defaultClock = NewSystemClock()
}

func SetDefault(c Clock) {
	defaultClock = c
}

func Now() time.Time {
	return defaultClock.Now()
}
