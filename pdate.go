package pdate

import (
	"fmt"
	"time"
)

// Pdate represents an immutable date in the Persian (Jalali/Solar Hijri) calendar.
// Pdate is designed to be similar to time.Time.
type Pdate struct {
	year    int
	month   PersianMonth
	day     int
	weekday PersianWeekday
	t       time.Time
}

var tehranLocation *time.Location

func init() {
	var err error
	tehranLocation, err = time.LoadLocation("Asia/Tehran")
	if err != nil {
		tehranLocation = time.FixedZone("IRST", 12600)
	}
}

// Now returns the current Persian date in Tehran timezone.
func Now() Pdate {
	return FromTime(time.Now().In(tehranLocation))
}

// Date create and returns a Pdate from Persian year, month, and day.
func Date(year int, month PersianMonth, day int) Pdate {
	gy, gm, gd := jalaliToGregorian(year, int(month), day)
	t := time.Date(gy, time.Month(gm), gd, 12, 0, 0, 0, tehranLocation)

	return Pdate{
		year:    year,
		month:   month,
		day:     day,
		weekday: PersianWeekdayFromGo(t.Weekday()),
		t:       t,
	}
}

// FromTime create and returns a Pdate from any time.Time value.
// The conversion uses the date portion in t's timezone.
func FromTime(t time.Time) Pdate {
	gy, gm, gd := t.Date()
	jy, jm, jd := gregorianToJalali(gy, int(gm), gd)

	return Pdate{
		year:    jy,
		month:   PersianMonth(jm),
		day:     jd,
		weekday: PersianWeekdayFromGo(t.Weekday()),
		t:       t,
	}
}

// Year returns the Persian year (e.g., 1404).
func (p Pdate) Year() int { return p.year }

// Month returns the Persian month as PersianMonth.
func (p Pdate) Month() PersianMonth { return p.month }

// Day returns the day of the month (1-31).
func (p Pdate) Day() int { return p.day }

// Weekday returns the day of the week as PersianWeekday.
func (p Pdate) Weekday() PersianWeekday { return p.weekday }

// Time returns the underlying time.Time value.
func (p Pdate) Time() time.Time { return p.t }

// String returns the date formatted using the format string
// "YYYY-MM-DD".
func (p Pdate) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", p.year, int(p.month), p.day)
}
