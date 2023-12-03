package gtime

import (
	"fmt"
	"time"
)

// FormatDuration function takes a time.Duration parameter 'd'
// and returns a formatted string representing the duration.
// The output format varies depending on the magnitude of 'd'.
// This function categorizes duration into
// nanoseconds (ns), microseconds (µs), seconds (s), minutes (m), hours (h) and days (d).
func FormatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	} else if d < time.Second {
		return fmt.Sprintf("%d µs", d.Microseconds())
	} else if d < time.Minute {
		return fmt.Sprintf("%.2f s", d.Seconds())
	} else if d < time.Hour {
		return fmt.Sprintf("%.2f m", d.Minutes())
	} else if d < 24*time.Hour {
		return fmt.Sprintf("%.2f h", d.Hours())
	}

	return fmt.Sprintf("%.0f d", d.Round(24*time.Hour).Hours()/24)
}

// TimeTrack function accepts a starting time and
// returns the elapsed time since the start time until now.
// The elapsed time is formatted by the FormatDuration function.
func TimeTrack(start time.Time) string {
	elapsed := time.Since(start)

	return FormatDuration(elapsed)
}
