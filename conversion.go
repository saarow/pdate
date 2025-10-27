package pdate

import "time"

var (
	persianWeekdays = [7]string{
		"یکشنبه",       // Sunday (0)
		"دوشنبه",       // Monday (1)
		"سه\u200cشنبه", // Tuesday (2)
		"چهارشنبه",     // Wednesday (3)
		"پنجشنبه",      // Thursday (4)
		"جمعه",         // Friday (5)
		"شنبه",         // Saturday (6)
	}

	persianMonths = map[int]string{
		1:  "فروردین",
		2:  "اردیبهشت",
		3:  "خرداد",
		4:  "تیر",
		5:  "مرداد",
		6:  "شهریور",
		7:  "مهر",
		8:  "آبان",
		9:  "آذر",
		10: "دی",
		11: "بهمن",
		12: "اسفند",
	}
)

func GregorianToJalali(gregorian Date) Date {
	result := Date{}
	array := [13]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

	year := gregorian.Year
	if year <= 1600 {
		year = year - 621
		result.Year = 0
	} else {
		year = year - 1600
		result.Year = 979
	}

	var temp int
	if year > 2 {
		temp = year + 1
	} else {
		temp = year
	}

	days := ((temp + 3) / 4) + (365 * year) - ((temp + 99) / 100) - 80 +
		array[gregorian.Month-1] + ((temp + 399) / 400) + gregorian.Day

	result.Year += 33 * (days / 12053)
	days = days % 12053
	result.Year += 4 * (days / 1461)
	days = days % 1461

	if days > 365 {
		result.Year += (days - 1) / 365
		days = (days - 1) % 365
	}

	if days < 186 {
		result.Month = 1 + (days / 31)
	} else {
		result.Month = 7 + (days-186)/30
	}

	if days < 186 {
		result.Day = 1 + (days % 31)
	} else {
		result.Day = 1 + (days-186)%30
	}

	return result
}

func PersianMonthName(month int) string {
	return persianMonths[month]
}

func PersianWeekdayName(weekday time.Weekday) string {
	return persianWeekdays[weekday]
}
