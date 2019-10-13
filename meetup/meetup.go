package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(w WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	var (
		result int
		count  int
		t      = time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
	)

	if w == Teenth {
		// When the WeekSchedule is Tenth is means we
		// want to find the first Monday during the days 10 to 19.
		t = t.AddDate(0, 0, 12)
	}

	for count < int(w) {
		t = t.AddDate(0, 0, 1)

		if t.Month() != month {
			// If we're entering a new month it means that the last
			// result is the last Weekday.
			break
		}

		if t.Weekday() != weekday {
			continue
		}

		if w == Teenth {
			// Found the first monday between 10 and 19.
			return t.Day()
		}

		result = t.Day()
		count++
	}

	return result
}
