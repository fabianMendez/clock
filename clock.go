package clock

import (
	"time"
)

// CurrentTimeFunc is a function that returns the current time
type CurrentTimeFunc func() time.Time

// Clock is an interface to allow mocking the time in unit tests
type Clock interface {
	// Now returns the current time
	Now() time.Time
}

type funcClock struct {
	fn CurrentTimeFunc
}

func (u *funcClock) Now() time.Time {
	return u.fn()
}

// Func creates a Clock which returns the time returned but the given CurrentTimeFunc
func Func(fn CurrentTimeFunc) Clock {
	return &funcClock{fn}
}

// Real returns the system's Clock
func Real() Clock {
	return Func(time.Now)
}

// Fixed creates a Clock which always return the same time.Time
func Fixed(fixedTime time.Time) Clock {
	return Func(func() time.Time { return fixedTime })
}
