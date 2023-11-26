package Jsql

import (
	"testing"
)

type user struct {
	Name    string `db:"name"`
	Age     string `db:"age"`
	Address string `db:"address"`
	Nested  nested `db:"ignore"`
}

type nested struct {
	Name    string  `db:"name"`
	Nested2 nested2 `db:"nested2"`
}

type nested2 struct {
	Name string `db:"name"`
}

func TestSelectColumn_GetFromStruct(t *testing.T) {
	user := user{}
	selectColumn := GetQColumnFromStruct(&user, "user.", "db")
	t.Log(selectColumn.ToQuery() + "ra")
}

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
