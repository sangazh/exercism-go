package clock

import (
	"fmt"
)

// define the clock type
type Clock struct {
	hour, minute int
}

// a "constructor"
func New(hour, minute int) Clock {
	c := Clock{hour: hour, minute: minute}
	c.optimize()

	return c
}

func (c *Clock) optimize() {
	if c.minute >= 60 || c.minute < 0 {
		c.hour += c.minute / 60
		c.minute %= 60
	}
	if c.minute < 0 {
		c.hour -= 1
		c.minute += 60
	}
	if c.hour >= 24 || c.hour < 0 {
		c.hour %= 24
	}
	if c.hour < 0 {
		c.hour += 24
	}
}

// a "stringer"
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	c.optimize()
	return c
}

func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes
	c.optimize()
	return c
}
