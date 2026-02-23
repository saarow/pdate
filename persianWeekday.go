package pdate

// PersianWeekday represents a day of the week in Persian calendar.
// Persian week starts on Shanbe (Saturday).
type PersianWeekday int

// Persian weekdays
const (
	Shanbe       PersianWeekday = 0 // شنبه - Saturday
	Yekshanbe    PersianWeekday = 1 // یکشنبه - Sunday
	Doshanbe     PersianWeekday = 2 // دوشنبه - Monday
	Seshanbe     PersianWeekday = 3 // سه‌شنبه - Tuesday
	Chaharshanbe PersianWeekday = 4 // چهارشنبه - Wednesday
	Panjshanbe   PersianWeekday = 5 // پنجشنبه - Thursday
	Jome         PersianWeekday = 6 // جمعه - Friday
)

var (
	persianWeekdayNames = [7]string{
		"شنبه",
		"یکشنبه",
		"دوشنبه",
		"سه\u200cشنبه",
		"چهارشنبه",
		"پنجشنبه",
		"جمعه",
	}

	englishWeekdayNames = [7]string{
		"Shanbe",
		"Yekshanbe",
		"Doshanbe",
		"Seshanbe",
		"Chaharshanbe",
		"Panjshanbe",
		"Jome",
	}
)

// String returns the Persian name of the weekday (e.g., "شنبه").
func (w PersianWeekday) String() string {
	if w < Shanbe || w > Jome {
		return ""
	}
	return persianWeekdayNames[w]
}

// EnglishName returns the English transliteration (e.g., "Shanbe").
func (w PersianWeekday) EnglishName() string {
	if w < Shanbe || w > Jome {
		return ""
	}
	return englishWeekdayNames[w]
}
