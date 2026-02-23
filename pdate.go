package pdate

import "time"

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

// Now returns the current Jalali date in Tehran timezone.
func Now() Pdate {
	return FromTime(time.Now().In(tehranLocation))
}

// New creates a Pdate from Jalali year, month, and day.
func New(year int, month PersianMonth, day int) Pdate {
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

// FromTime creates a Pdate from any time.Time value.
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

// Year returns the Jalali year (e.g., 1404).
func (p Pdate) Year() int { return p.year }

// Month returns the Jalali month as PersianMonth.
func (p Pdate) Month() PersianMonth { return p.month }

// MonthInt returns the month as an int (1-12).
func (p Pdate) MonthInt() int { return int(p.month) }

// MonthName returns the Persian name of the month (e.g., "خرداد").
func (p Pdate) MonthName() string { return p.month.String() }

// Day returns the day of the month (1-31).
func (p Pdate) Day() int { return p.day }

// Weekday returns the day of the week as PersianWeekday.
func (p Pdate) Weekday() PersianWeekday { return p.weekday }

// WeekdayName returns the Persian name of the weekday (e.g., "پنجشنبه").
func (p Pdate) WeekdayName() string { return p.weekday.String() }

// Time returns the underlying time.Time value.
func (p Pdate) Time() time.Time { return p.t }

// Location returns the timezone of the underlying time.
func (p Pdate) Location() *time.Location { return p.t.Location() }

// Unix returns the Unix timestamp (seconds since Unix epoch).
func (p Pdate) Unix() int64 { return p.t.Unix() }

// IsLeapYear reports whether this date's year is a leap year.
func (p Pdate) IsLeapYear() bool {
	return IsLeapYear(p.year)
}
