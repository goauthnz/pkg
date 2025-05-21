package postgres

import "time"

// TimeNow returns the current time in UTC.
// It is a helper function used to unify the time format across the application.
func TimeNow() time.Time {
	return time.Now().UTC()
}
