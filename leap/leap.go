// Package leap defines if a given year is leap or not
package leap

// IsLeapYear returns true if a given year is leap
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
