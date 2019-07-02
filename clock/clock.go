package clock

import (
	"fmt"
)

type Clock struct {
	m int
}

func New(hours, minutes int) Clock {
	return Clock{
		m: setupTime(hours*60 + minutes),
	}
}

func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	return c
}

func (c Clock) Subtract(minutes int) Clock {
	c.m -= minutes
	return c
}

func (c Clock) String() string {
	minutes := setupTime(c.m)
	return fmt.Sprintf("%02d:%02d", minutes/60, minutes%60)

}
func (c Clock) Equal(to Clock) bool {
	return setupTime(to.m) == setupTime(c.m)
}

func setupTime(minutes int) int {
	minutes = minutes % (24 * 60)
	if minutes < 0 {
		minutes += (24 * 60)
	}
	return minutes
}
