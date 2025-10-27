package pdate

import (
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int

	WeekdayName string
	MonthName   string
}

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
