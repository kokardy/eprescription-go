package epres

import (
	"fmt"
	"time"
)

type Date time.Time
type Duration time.Duration

func (d Date) NewDate(year int, month time.Month, day int, location string) Date {
	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Errorf("err: %s location: '%s' is unavailable", err, loc)
	}
	return Date(time.Date(year, month, day, 0, 0, 0, 0, loc))
}

func (d Date) Distance(d2 Date) Duration {
	return Duration(time.Time(d).Sub(time.Time(d2)))
}
