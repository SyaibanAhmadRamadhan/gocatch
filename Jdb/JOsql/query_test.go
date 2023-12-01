package JOsql

import (
	"testing"
)

func TestFilterToQuery(t *testing.T) {
	testCases := QFilterNamedArgs{
		{
			Column:      "age",
			Value:       "rama",
			Comparasion: IsNotNull,
			Logical:     Or,
		},
		{
			Column:      "umur",
			Value:       "rama",
			NamedArg:    "test",
			Comparasion: Equals,
			Logical:     And,
		},
		{
			Column:      "umurs",
			Value:       "rama",
			Comparasion: Equals,
			Logical:     And,
		},
	}

	t.Log(testCases.ToQuery(true, Pgx))
}
