package port

import "time"

type Reminder struct {
	UserID         string
	DayRemind      time.Weekday
	TimeRemind     time.Time
	TimezoneRemind time.Location
	Type           string
}
