package clock

import (
	"fmt"
)

// Clock holds measurments of time.
type Clock struct {
	minutes int
}

// New sets up a new Clock.
func New(hours, minutes int) Clock {
	return Clock{
		minutes: setupTime(hours*60 + minutes),
	}
}

// Add adds hours to a clock.
func (c Clock) Add(minutes int) Clock {
	c.minutes = setupTime(c.minutes + minutes)
	return c
}

// Subtract minutes from a clock.
func (c Clock) Subtract(minutes int) Clock {
	c.minutes = setupTime(c.minutes - minutes)
	return c
}

// String implmenets the Stringer interface.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)

}

func setupTime(m int) int {
	m %= (24 * 60)
	if m < 0 {
		m += (24 * 60)
	}
	return m
}
