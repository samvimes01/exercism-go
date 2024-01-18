// Contains a method for calculating time after a gigasecond.
package gigasecond

import "time"

// AddGigasecond calculating time after a gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1000000000))
}
