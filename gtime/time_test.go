package gtime

import (
	"testing"
	"time"
)

// TestFormatDuration tests FormatDuration function with multiple scenarios
func TestFormatDuration(t *testing.T) {
	expectations := []struct {
		inputDuration      time.Duration
		expectedTimeFormat string
	}{
		{time.Nanosecond, "1 ns"},  // Test case for nanoseconds
		{time.Microsecond, "1 µs"}, // Test case for microseconds
		{time.Second, "1.00 s"},    // Test case for seconds
		{time.Minute, "1.00 m"},    // Test case for minutes
		{time.Hour, "1.00 h"},      // Test case for hours
		{48 * time.Hour, "2 d"},    // Test case for 1 day
	}
	for _, tt := range expectations {
		if got := FormatDuration(tt.inputDuration); got != tt.expectedTimeFormat {
			t.Errorf("Expected '%s', but got '%s'", tt.expectedTimeFormat, got)
		}
	}
}
