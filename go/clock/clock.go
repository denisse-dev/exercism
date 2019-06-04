package clock

import "fmt"

// Clock stores an hour and minute integers to represent a clock.
type Clock struct {
	hour   int
	minute int
}

const maxMinutes = 60
const maxHours = 24

// New creates a clock that handles times without dates.
func New(hour int, minute int) Clock {
	var clock Clock
	clock.hour = hour
	clock.minute = minute

	clock = clock.Add(0)
	clock = clock.Subtract(0)

	return clock
}

// String formats a clock struct to show time correctly.
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add sums certain minutes to a previously created clock.
func (clock Clock) Add(minutes int) Clock {
	clock.minute += minutes

	clock.hour += int(clock.minute / maxMinutes)
	clock.hour %= maxHours
	clock.minute %= maxMinutes

	return clock
}

// Subtract subtracts certain minutes to a previously clock.
func (clock Clock) Subtract(minutes int) Clock {
	clock.minute -= minutes

	for clock.minute < 0 {
		clock.minute += maxMinutes
		clock.hour--
	}
	for clock.hour < 0 {
		clock.hour += maxHours
	}

	return clock
}
