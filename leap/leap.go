// Holds a methid to calculate if it is a leap year.
package leap

// Given a year, report if it is a leap year.
func IsLeapYear(year int) bool {
	if year % 4 != 0 {
		return false
	}
	if year % 100 != 0 {
		return true
	} else if year % 400 != 0 {
		return false
	}
	return true
}
