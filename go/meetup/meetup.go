package meetup

import (
	"time"
)

type WeekSchedule int

const (
	Teenth WeekSchedule = iota
	First
	Second
	Third
	Fourth
	Last
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) (day int) {
	var start int
	switch week {
	case First:
		start = 1
	case Second:
		start = 8
	case Third:
		start = 15
	case Fourth:
		start = 22
	case Last:
		lastDayOfMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
		start = lastDayOfMonth.Day() - 6
	case Teenth:
		start = 13
	}

	startDate := time.Date(year, month, start, 0, 0, 0, 0, time.Local)
	days := int(weekday-startDate.Weekday()+7) % 7

	targetDate := startDate.AddDate(0, 0, days)
	day = targetDate.Day()

	return day
}
