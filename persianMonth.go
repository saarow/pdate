package pdate

// PersianMonth represents a month in the Jalali (Persian) calendar.
type PersianMonth int

// Persian months
const (
	Farvardin   PersianMonth = 1
	Ordibehesht PersianMonth = 2
	Khordad     PersianMonth = 3
	Tir         PersianMonth = 4
	Mordad      PersianMonth = 5
	Shahrivar   PersianMonth = 6
	Mehr        PersianMonth = 7
	Aban        PersianMonth = 8
	Azar        PersianMonth = 9
	Dey         PersianMonth = 10
	Bahman      PersianMonth = 11
	Esfand      PersianMonth = 12
)

var (
	persianMonthNames = [13]string{
		0:  "",
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

	englishMonthNames = [13]string{
		0:  "",
		1:  "Farvardin",
		2:  "Ordibehesht",
		3:  "Khordad",
		4:  "Tir",
		5:  "Mordad",
		6:  "Shahrivar",
		7:  "Mehr",
		8:  "Aban",
		9:  "Azar",
		10: "Dey",
		11: "Bahman",
		12: "Esfand",
	}
)

// Valid reports whether m is a valid month (1-12).
func (m PersianMonth) Valid() bool {
	return m >= Farvardin && m <= Esfand
}

// String returns the Persian name of the month (e.g., "فروردین").
func (m PersianMonth) String() string {
	if !m.Valid() {
		return ""
	}
	return persianMonthNames[m]
}

// EnglishName returns the English transliteration (e.g., "Farvardin").
func (m PersianMonth) EnglishString() string {
	if !m.Valid() {
		return ""
	}
	return englishMonthNames[m]
}
