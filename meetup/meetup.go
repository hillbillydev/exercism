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
		t      = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	)

	if w == Teenth {
		// When the WeekSchedule is Tenth is means we
		// want to find the first Monday during the days 10 to 19.
		t = t.AddDate(0, 0, 12)
	}

	for WeekSchedule(count) < w {
		if t.Weekday() == weekday {
			if w == Teenth {
				// Found the first monday between 10 and 19.
				return t.Day()
			}

			if t.Month() != month {
				// If we're entering a new month it means that the last
				// result is the last Weekday.
				return result
			}

			result = t.Day()
			count++
		}

		t = t.AddDate(0, 0, 1)
	}

	return result
}
