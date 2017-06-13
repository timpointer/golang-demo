package time

import "time"
import "fmt"

// FirstDayOfISOWeek get first day of the week
func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

// GetListMonth get a array of time string ex[201402,201403,201404] between start and end time
func GetListMonth(start, end time.Time) []string {
	list := []string{}
	includeEnd := end.AddDate(0, 1, 0)
	for time := start; time.Before(includeEnd); time = time.AddDate(0, 1, 0) {
		list = append(list, fmt.Sprintf("%04d%02d", time.Year(), time.Month()))
	}
	return list
}
