package service

import "time"

type ReminderAdapter struct {
	UserID         string
	DayRemind      time.Weekday
	TimeRemind     time.Time
	TimezoneRemind time.Location
	Type           string
}
