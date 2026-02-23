package pdate

// gregorianToJalali converts Gregorian (gy, gm, gd) to Jalali (jy,jm,jd).
func gregorianToJalali(gy, gm, gd int) (jy, jm, jd int) {
	var gdm = [12]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

	var gy2 int
	if gm > 2 {
		gy2 = gy + 1
	} else {
		gy2 = gy
	}

	days := 355666 + (365 * gy) + ((gy2 + 3) / 4) - ((gy2 + 99) / 100) +
		((gy2 + 399) / 400) + gd + gdm[gm-1]

	jy = -1595 + (33 * (days / 12053))
	days %= 12053

	jy += 4 * (days / 1461)
	days %= 1461

	if days > 365 {
		jy += (days - 1) / 365
		days = (days - 1) % 365
	}

	if days < 186 {
		jm = 1 + days/31
		jd = 1 + days%31
	} else {
		jm = 7 + (days-186)/30
		jd = 1 + (days-186)%30
	}

	return jy, jm, jd
}

// jalaliToGregorian converts Jalali (jy, jm, jd) to Gregorian (gy, gm, gd).
func jalaliToGregorian(jy, jm, jd int) (gy, gm, gd int) {
	jy2 := jy + 1595

	days := -355668 + (365 * jy2) + ((jy2 / 33) * 8) + (((jy2 % 33) + 3) / 4) + jd

	if jm < 7 {
		days += (jm - 1) * 31
	} else {
		days += ((jm - 7) * 30) + 186
	}

	gy = 400 * (days / 146097)
	days %= 146097

	if days > 36524 {
		days--
		gy += 100 * (days / 36524)
		days %= 36524
		if days >= 365 {
			days++
		}
	}

	gy += 4 * (days / 1461)
	days %= 1461

	if days > 365 {
		gy += (days - 1) / 365
		days = (days - 1) % 365
	}

	gd = days + 1

	gdm := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (gy%4 == 0 && gy%100 != 0) || gy%400 == 0 {
		gdm[1] = 29
	}

	gm = 0
	for gm < 12 && gd > gdm[gm] {
		gd -= gdm[gm]
		gm++
	}
	gm++

	return gy, gm, gd
}

// IsLeapYear reports whether a Jalali year is a leap year.
func IsLeapYear(year int) bool {
	r := year % 33
	if r < 0 {
		r += 33
	}
	switch r {
	case 1, 5, 9, 13, 17, 22, 26, 30:
		return true
	}
	return false
}
