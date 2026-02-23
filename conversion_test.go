package pdate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGregorianToJalali(t *testing.T) {
	cases := []struct {
		gYear, gMonth, gDay int
		jYear, jMonth, jDay int
		label               string
	}{
		// Nowruz boundaries
		{2016, 3, 20, 1395, 1, 1, "Nowruz 1395"},
		{2017, 3, 20, 1395, 12, 30, "Last day of leap 1395"},
		{2017, 3, 21, 1396, 1, 1, "Nowruz 1396"},
		{2018, 3, 21, 1397, 1, 1, "Nowruz 1397"},
		{2019, 3, 21, 1398, 1, 1, "Nowruz 1398"},
		{2020, 3, 20, 1399, 1, 1, "Nowruz 1399"},
		{2021, 3, 20, 1399, 12, 30, "Last day of leap 1399"},
		{2021, 3, 21, 1400, 1, 1, "Nowruz 1400"},
		{2022, 3, 21, 1401, 1, 1, "Nowruz 1401"},
		{2023, 3, 21, 1402, 1, 1, "Nowruz 1402"},
		{2024, 3, 20, 1403, 1, 1, "Nowruz 1403"},
		{2025, 3, 20, 1403, 12, 30, "Last day of leap 1403"},
		{2025, 3, 21, 1404, 1, 1, "Nowruz 1404"},

		// End of non-leap years
		{2022, 3, 20, 1400, 12, 29, "End of non-leap 1400"},
		{2023, 3, 20, 1401, 12, 29, "End of non-leap 1401"},

		// Month boundaries: 31-day months (1–6)
		{2024, 4, 19, 1403, 1, 31, "Last day of Farvardin"},
		{2024, 4, 20, 1403, 2, 1, "First day of Ordibehesht"},
		{2024, 5, 20, 1403, 2, 31, "Last day of Ordibehesht"},
		{2024, 5, 21, 1403, 3, 1, "First day of Khordad"},
		{2024, 9, 21, 1403, 6, 31, "Last day of Shahrivar"},

		// 31→30 day transition (month 6→7)
		{2024, 9, 22, 1403, 7, 1, "First day of Mehr"},

		// 30-day months (7–11)
		{2024, 10, 21, 1403, 7, 30, "Last day of Mehr"},
		{2024, 10, 22, 1403, 8, 1, "First day of Aban"},
		{2025, 2, 18, 1403, 11, 30, "Last day of Bahman"},
		{2025, 2, 19, 1403, 12, 1, "First day of Esfand"},

		// Well-known dates
		{2024, 12, 31, 1403, 10, 11, "End of 2024"},
		{2000, 1, 1, 1378, 10, 11, "Y2K"},
		{1979, 2, 11, 1357, 11, 22, "Iranian Revolution"},
		{2001, 9, 11, 1380, 6, 20, "9/11"},

		// Older dates
		{1800, 1, 1, 1178, 10, 11, "Start of 1800"},
		{1600, 1, 1, 978, 10, 11, "Start of 1600"},
		{1599, 12, 31, 978, 10, 10, "End of 1599"},
		{1500, 3, 21, 879, 1, 1, "Nowruz 879"},
	}

	for _, tc := range cases {
		name := fmt.Sprintf(
			"%s_%d-%02d-%02d",
			tc.label,
			tc.gYear,
			tc.gMonth,
			tc.gDay,
		)
		t.Run(name, func(t *testing.T) {
			jy, jm, jd := gregorianToJalali(
				tc.gYear,
				tc.gMonth,
				tc.gDay,
			)
			assert.Equal(t, tc.jYear, jy, "year")
			assert.Equal(t, tc.jMonth, jm, "month")
			assert.Equal(t, tc.jDay, jd, "day")
		})
	}
}

func TestJalaliToGregorianRoundTrip(t *testing.T) {
	for gy := 700; gy <= 2025; gy += 25 {
		for gm := 1; gm <= 12; gm++ {
			for gd := 1; gd <= 28; gd++ {
				jy, jm, jd := gregorianToJalali(gy, gm, gd)
				gy2, gm2, gd2 := jalaliToGregorian(jy, jm, jd)
				if gy != gy2 || gm != gm2 || gd != gd2 {
					t.Fatalf(
						"Round-trip failed: G(%d-%02d-%02d) → J(%d-%02d-%02d) → G(%d-%02d-%02d)",
						gy,
						gm,
						gd,
						jy,
						jm,
						jd,
						gy2,
						gm2,
						gd2,
					)
				}
			}
		}
	}
}

func TestIsLeapYearConsistentWithConversion(t *testing.T) {
	// A year is leap if and only if Esfand has 30 days,
	// i.e., the day after Esfand 29 is still in the same year.
	for jy := 1; jy <= 2000; jy++ {
		// Convert Esfand 30 to Gregorian and back
		gy, gm, gd := jalaliToGregorian(jy, 12, 30)
		jy2, jm2, jd2 := gregorianToJalali(gy, gm, gd)

		isLeapByConversion := jy2 == jy && jm2 == 12 && jd2 == 30
		isLeapByFunc := IsLeapYear(jy)

		if isLeapByConversion != isLeapByFunc {
			t.Fatalf("Year %d: IsLeapYear()=%v but conversion says leap=%v",
				jy, isLeapByFunc, isLeapByConversion)
		}
	}
}

func TestLeapYears(t *testing.T) {
	leapYears := []int{
		1399,
		1403,
		1408,
		1370,
		1375,
		1379,
		1383,
		1387,
		1391,
		1395,
	}
	nonLeapYears := []int{1400, 1401, 1402, 1404, 1405, 1406, 1407}

	for _, y := range leapYears {
		assert.True(t, IsLeapYear(y), "Year %d should be leap", y)
	}
	for _, y := range nonLeapYears {
		assert.False(t, IsLeapYear(y), "Year %d should NOT be leap", y)
	}
}
