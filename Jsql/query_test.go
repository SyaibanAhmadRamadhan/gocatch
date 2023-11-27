package Jsql

import (
	"testing"
)

func TestFilterToQuery(t *testing.T) {
	testCases := QFilterNamed{
		"age": FilterNamedQuery{
			Value:       "rama",
			Comparasion: IsNotNull,
			Logical:     Or,
		},
		"umur": GenerateQFilterNamed("rama", "test", Equals, And),
		"name": GenerateQFilterNamedArgByColumn("rama", Equals, And),
	}

	t.Log(testCases.ToQuery(true, Pgx))
}
