package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// define month first and last days to iterate through
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	// teenth means only 13th to 19th
	if wSched == Teenth {
		startOfMonth = startOfMonth.AddDate(0, 0, 12) // 13th
		endOfMonth = startOfMonth.AddDate(0, 0, 6)    // 19th
	}
	dt := startOfMonth.AddDate(0,0,0)

	dayCounter := 1
	// not the most efficient way, but quite readable.
	// just iterate through days of month, use label to break loop from inside switch
Loop:
	for d := startOfMonth; !d.After(endOfMonth); d = d.AddDate(0, 0, 1) {
		weekDay := d.Weekday()
		switch wSched {
			// last and teenth share the same logic - the latest will be correct
			// teenth day always appear last (once) since it's only 7 days
			case Last, Teenth:
				if weekDay == wDay {
					dt = d
				}
			// 1 to 4th weekday share same logic - count the weekday, break when found
			case First, Second, Third, Fourth:
				if weekDay == wDay && dayCounter == int(wSched) {
					dt = d
					break Loop
				}
				if weekDay == wDay {
					dayCounter++
				}
		}
	}
	return dt.Day()
}
