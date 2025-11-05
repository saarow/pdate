# pdate - Go Package for Persian (Jalali) Dates
### A comprehensive Go package for working with Persian (Jalali) dates. Convert between Gregorian and Jalali calendars, get localized Persian month and weekday names, and handle Persian date operations with ease.

## Features

- ✅ Convert Gregorian dates to Jalali (Persian) dates
- ✅ Get current Jalali date with Persian names
- ✅ Localized Persian month and weekday names
- ✅ Simple and intuitive API
- ✅ Thread-safe operations
- ✅ No external dependencies

## Installation

```bash
go get -u github.com/saarow/pdate
```
## Quick Start
``` go
package main

import (
    "fmt"
    "github.com/saarow/pdate"
)

func main() {
    // Get current Jalali date
    jalali, err := pdate.GetJalaliDate()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Today in Jalali: %d/%d/%d - %s\n", 
        jalali.Year, jalali.Month, jalali.Day, jalali.WeekdayName)
    fmt.Printf("Month: %s\n", jalali.MonthName)
    
    // Get current Gregorian date
    gregorian := pdate.GetGregorianDate()
    fmt.Printf("Today in Gregorian: %d/%d/%d - %s\n",
        gregorian.Year, gregorian.Month, gregorian.Day, gregorian.WeekdayName)
}
```
