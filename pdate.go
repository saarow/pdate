package pdate

import "time"

type Date struct {
	Year  int
	Month int
	Day   int
}

func GetGregorianDate() Date {
	now := time.Now()
	year, month, day := now.Date()

	return Date{
		Year:  year,
		Month: int(month),
		Day:   day,
	}
}

func GetJalaliDate() Date {
	gDate := GetGregorianDate()
	jDate := GregorianToJalali(gDate)

	return jDate
}
