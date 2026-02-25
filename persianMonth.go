package pdate

// PersianMonth represents a month in the Persian calendar
// (Farvardin = 1, ...).
type PersianMonth int

// Persian months
const (
	Farvardin PersianMonth = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
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

// String returns the English transliteration of the month
// (e.g., "Farvardin", "Ordibehesht", ...).
func (m PersianMonth) String() string {
	return englishMonthNames[m]
}

// PersianName returns the month name in Persian
// (e.g., "فروردین", "اردیبهشت", ...).
func (m PersianMonth) PersianName() string {
	return persianMonthNames[m]
}
