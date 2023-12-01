package gcommon

import (
	"testing"
)

func TestTernary(t *testing.T) {
	tests := []struct {
		name          string
		condition     bool
		trueVal       interface{}
		falseVal      interface{}
		expectedValue interface{}
	}{
		{
			name:          "Condition true, positive integer values",
			condition:     true,
			trueVal:       5,
			falseVal:      10,
			expectedValue: 5,
		},
		{
			name:          "Condition false, positive integer values",
			condition:     false,
			trueVal:       5,
			falseVal:      10,
			expectedValue: 10,
		},
		{
			name:          "Condition true, string values",
			condition:     true,
			trueVal:       "foo",
			falseVal:      "bar",
			expectedValue: "foo",
		},
		{
			name:          "Condition false, string values",
			condition:     false,
			trueVal:       "foo",
			falseVal:      "bar",
			expectedValue: "bar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Ternary(test.condition, test.trueVal, test.falseVal)
			if result != test.expectedValue {
				t.Errorf("Expected %v, but got %v", test.expectedValue, result)
			}
		})
	}
}
