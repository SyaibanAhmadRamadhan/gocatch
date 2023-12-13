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

func TestNormalizeTimeUnit(t *testing.T) {
	testCases := []struct {
		name          string
		inputTime     time.Time
		unit          TimeUnit
		expectedNsec  int
		expectedMonth time.Month
	}{
		{
			name:          "Nanoseconds",
			inputTime:     time.Date(2022, time.January, 1, 12, 0, 0, 0, time.UTC),
			unit:          Nanoseconds,
			expectedNsec:  0,
			expectedMonth: time.January,
		},
		{
			name:          "Milliseconds",
			inputTime:     time.Date(2022, time.February, 1, 12, 0, 0, 500*int(time.Millisecond), time.UTC),
			unit:          Milliseconds,
			expectedNsec:  int(500 * time.Millisecond.Nanoseconds()),
			expectedMonth: time.February,
		},
		{
			name:          "Microseconds",
			inputTime:     time.Date(2022, time.March, 1, 12, 0, 0, 500*int(time.Microsecond), time.UTC),
			unit:          Microseconds,
			expectedNsec:  500 * int(time.Microsecond),
			expectedMonth: time.March,
		},
		{
			name:          "Default",
			inputTime:     time.Date(2022, time.April, 1, 12, 0, 0, 500*int(time.Nanosecond), time.UTC),
			unit:          42, // Invalid unit, uses default
			expectedNsec:  500 * int(time.Nanosecond),
			expectedMonth: time.April,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NormalizeTimeUnit(tc.inputTime, tc.unit)

			if result.Nanosecond() != tc.expectedNsec {
				t.Errorf("Expected Nanoseconds: %d, Got: %d", tc.expectedNsec, result.Nanosecond())
			}

			if result.Month() != tc.expectedMonth {
				t.Errorf("Expected Month: %s, Got: %s", tc.expectedMonth, result.Month())
			}
		})
	}
}
