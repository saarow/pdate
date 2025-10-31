// Package pdate provides utilities for working with Persian (Jalali) dates.
package pdate

import (
	"fmt"
	"time"
)

// Date represents a complete date with year, month, day, and localized names.
// It can represent both Gregorian and Jalali dates with appropriate naming.
type Date struct {
	Year  int
	Month int
	Day   int

	WeekdayName string
	MonthName   string
}

// GetGregorianDate returns the current date in Gregorian format with English names.
func GetGregorianDate() Date {
	now := time.Now()
	year, month, day := now.Date()

	return Date{
		Year:  year,
		Month: int(month),
		Day:   day,

		WeekdayName: now.Format("Monday"),
		MonthName:   now.Month().String(),
	}
}

// GetJalaliDate returns the current date in Jalali (Persian) format with Persian names.
func GetJalaliDate() (Date, error) {
	gDate := GetGregorianDate()
	jDate := GregorianToJalali(gDate)

	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		return Date{}, fmt.Errorf(
			"failed to load Tehran timezone location: %w",
			err,
		)
	}
	tehranTime := time.Now().In(loc)
	weekday := tehranTime.Weekday()

	jDate.WeekdayName = PersianWeekdayName(weekday)
	jDate.MonthName = PersianMonthName(jDate.Month)

	return jDate, nil
}
