package gigasecond

import "time"

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
