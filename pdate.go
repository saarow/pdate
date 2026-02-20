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
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		return Date{}, fmt.Errorf(
			"failed to load Tehran timezone location: %w",
			err,
		)
	}

	now := time.Now().In(loc)
	year, month, day := now.Date()

	gDate := Date{
		Year:  year,
		Month: int(month),
		Day:   day,
	}

	jDate := GregorianToJalali(gDate)
	jDate.WeekdayName = PersianWeekdayName(now.Weekday())
	jDate.MonthName = PersianMonthName(gDate.Month)

	return jDate, nil
}
