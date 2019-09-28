package time

import "time"

type TimePoint struct {
	Point       time.Time     // the point in time when the Action should be triggered
	WaitTime    time.Duration // the effective waiting time
	IsRepetable func() bool   // should determine if the action should be repeated
	Action      func()        // the logic that should be done when the point in time is triggered
}
