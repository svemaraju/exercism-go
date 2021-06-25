package clock

import (
	"fmt"
)

// Clock struct represented by hour and minute values.
type Clock struct {
	hour int
	minute int
}


// New creates a new Clock object.
func New(hour, minute int) Clock {
	// calculate overall minutes.
	total_minutes := 60*hour + minute
	// Rollover the minutes until they become positive.
    for total_minutes<0 {
		total_minutes = 24*60 + total_minutes
	}
	// calculate hour and minute parts from total_minutes.
	hour = (total_minutes/60)%24
	minute = total_minutes%60
	return Clock{hour: hour, minute: minute}
}

// String function returns a string representation of Clock instance.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add function adds minutes to an existing Clock and returns a new Clock instance.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute + minutes)
}


// Subtract function subtracts minutes from an existing Clock and returns a new Clock instance.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute - minutes)
}

