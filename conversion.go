package pdate

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
