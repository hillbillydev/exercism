// Package leap gives you a package to determine leap years.
package leap

// IsLeapYear determines if the year you pass in is a LeapYear.
func IsLeapYear(year int) bool {
	// 	on every year that is evenly divisible by 4
	//  except every year that is evenly divisible by 100
	if year/100%2 == 0 {
		//  unless the year is also evenly divisible by 400
		return (year/400)%2 == 0
	}
	return year/4%2 == 0
}
