package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	if m < 0 {
		hDelta := -1 * m/60
		mDelta := m%60 * -1
		if hDelta >= 0 && mDelta > 0 {
			hDelta++
		}
		h, m = h - hDelta, 60 - mDelta
	} else if m >= 60 {	
		h, m = h + m/60, m%60
	}

	if h < 0 {
		h = h * -1
		h = 24 - h%24
	} else if h >= 24 {
		h = h%24
	}

	if m == 60 {
		m = 0
	}
	if h == 24 {
		h = 0
	}

	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

/*
just minutes 
type Clock int

const (
	dayMinutes  = 1440
	hourMinutes = 60
)

func normalize(minutes int) Clock {
	if minutes < 0 {
		return Clock(dayMinutes + (minutes % dayMinutes))
	} else if minutes >= dayMinutes {
		return Clock(minutes % dayMinutes)
	}
	return Clock(minutes)
}

func New(hour, minutes int) Clock {
	return normalize(hour*hourMinutes + minutes)
}

func (c Clock) Add(minutes int) Clock {
	return normalize(int(c) + minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return normalize(int(c) - minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/hourMinutes, c%hourMinutes)
}
*/