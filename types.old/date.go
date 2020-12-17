package epres

import (
	"fmt"
	"time"
)

const LOCATION_JP = "Asia/Tokyo"

type Date struct {
	time.Time
}

func Today(location string) Date {
	now := time.Now()
	return NewDate(now.Year(), now.Month(), now.Day(), location)
}

func (d Date) Stringer() string {
	return d.Format("2006/01/02")
}

type Duration time.Duration

func (d Date) NewDate(year int, month time.Month, day int, location string) Date {
	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Errorf("err: %s location: '%s' is unavailable", err, loc)
	}
	return Date{time.Date(year, month, day, 0, 0, 0, 0, loc)}
}

func (d Date) Distance(d2 Date) Duration {
	return Duration(d.Sub(d2.Time))
}

type Birthdayer interface {
	Birthday() Date
}

func Age(b Birthdayer) Duration {
	today := Today()
	birthday := b.Birthday()
	return birthday.Distance(today)
}
