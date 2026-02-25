package pdate

import "time"

// PersianWeekday represents a day of the week in Persian calendar
// (Shanbe = 0, ...).
type PersianWeekday int

// Persian weekdays
const (
	Shanbe PersianWeekday = iota
	Yekshanbe
	Doshanbe
	Seshanbe
	Chaharshanbe
	Panjshanbe
	Jome
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

// String returns the English transliteration of the weekday
// (e.g., "Shanbe", "Yekshanbe", ...).
func (w PersianWeekday) String() string {
	return englishWeekdayNames[w]
}

// PersianName returns the Persian name of the weekday
// (e.g., "شنبه", "یکشنبه", ...).
func (w PersianWeekday) PersianName() string {
	return persianWeekdayNames[w]
}

// GoWeekday converts PersianWeekday to Go's time.Weekday.
func (w PersianWeekday) GoWeekday() time.Weekday {
	// Shanbe(0) → Saturday(6), Yekshanbe(1) → Sunday(0), etc.
	return time.Weekday((int(w) + 6) % 7)
}

// PersianWeekdayFromGo converts Go's time.Weekday to PersianWeekday.
func PersianWeekdayFromGo(wd time.Weekday) PersianWeekday {
	// Saturday(6) → Shanbe(0), Sunday(0) → Yekshanbe(1), etc.
	return PersianWeekday((int(wd) + 1) % 7)
}
